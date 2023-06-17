//go:build !codeanalysis
// +build !codeanalysis

package v1

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"github.com/mr-tron/base58"
	"github.com/web3eye-io/Web3Eye/common/oss"
	"github.com/web3eye-io/Web3Eye/config"
	dealer_client "github.com/web3eye-io/Web3Eye/dealer/pkg/client/v1"
	"github.com/web3eye-io/Web3Eye/gen-car/pkg/car"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	dealer_proto "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
	gencar_proto "github.com/web3eye-io/Web3Eye/proto/web3eye/gencar/v1"
)

const (
	ImageResType = "image"
)

var (
	dataDir = config.GetConfig().GenCar.DataDir
)

type TokenBaseInfo struct {
	ChainType string
	ChainID   string
	Contract  string
	TokenID   string
}

type TokenResInfo struct {
	BaseInfo TokenBaseInfo
	FileName string
	Size     int64
	S3Key    string
	ID       string
	ResType  string
}

func (s *Server) ReportFile(ctx context.Context, in *gencar_proto.ReportFileRequest) (*gencar_proto.ReportFileResponse, error) {
	logger.Sugar().Infof("start deal file, id %v, s3key %v", in.ID, in.S3Key)
	info, err := token.GetToken(ctx, in.ID)
	if err != nil {
		logger.Sugar().Infof("failed get token by id, err: %v", in.ID, err)
		return nil, err
	}
	objAttr, err := oss.GetObjectAttributes(ctx, in.S3Key)
	if err != nil {
		logger.Sugar().Infof("failed get token image by id, err: %v", in.ID, err)
		return nil, err
	}
	resInfo := &TokenResInfo{
		BaseInfo: TokenBaseInfo{
			ChainType: info.ChainType.String(),
			ChainID:   info.ChainID,
			Contract:  info.Contract,
			TokenID:   info.TokenID,
		},
		S3Key:   in.S3Key,
		ID:      in.ID,
		ResType: ImageResType,
		Size:    objAttr.ContentLength,
	}
	sha2sum, err := Sha256Name(resInfo.BaseInfo)
	if err != nil {
		return nil, err
	}

	resInfo.FileName = fmt.Sprintf("%v%v", sha2sum, path.Ext(in.S3Key))

	go PutTokenResInfo(resInfo)

	return &gencar_proto.ReportFileResponse{
		ID:    in.ID,
		S3Key: in.S3Key,
	}, nil
}

func Sha256Name(info TokenBaseInfo) (string, error) {
	_name := fmt.Sprintf("%v-%v-%v-%v", info.ChainType, info.ChainID, info.Contract, info.TokenID)
	_h := sha256.Sum256([]byte(_name))
	return base58.Encode(_h[:]), nil
}

func CheckSha256Name(info TokenBaseInfo, base58String string) (bool, error) {
	h, err := Sha256Name(info)
	if err != nil {
		return false, err
	}
	if h == base58String {
		return true, nil
	}
	return false, nil
}

type CarManager struct {
	dataDir       string
	downloadChan  chan *TokenResInfo
	downloadClose chan struct{}
	genCarChan    chan *TokenResInfo
	genCarClose   chan struct{}
}

const (
	DefaultDownloadChanLen  = 100
	DefaultDownloadParallel = 3
	// 17GB
	// maxUnTarSize = 18253611008
	// 100M
	maxUnTarSize = 10485760
)

var carManager *CarManager

func RunCarManager() {
	carManager = &CarManager{
		downloadChan:  make(chan *TokenResInfo, DefaultDownloadChanLen),
		downloadClose: make(chan struct{}),
		genCarChan:    make(chan *TokenResInfo, DefaultDownloadChanLen),
		genCarClose:   make(chan struct{}),
		dataDir:       filePath(""),
	}

	var fileMode os.FileMode = 0777
	err := os.Mkdir(carManager.dataDir, fileMode)
	if err != nil && !strings.Contains(err.Error(), "file exists") {
		logger.Sugar().Errorf("mkdir failed, path: %v", carManager.dataDir)
		panic(err)
	}

	logger.Sugar().Info("start run car manager")

	go carManager.checkAndGenCar(maxUnTarSize)
	carManager.runDownloadTask(DefaultDownloadParallel)
}

func PutTokenResInfo(tokenResInfo *TokenResInfo) {
	if carManager == nil {
		panic("not init CarManager")
	}
	carManager.downloadChan <- tokenResInfo
}

func (cm *CarManager) runDownloadTask(parallel int) {
	for i := 0; i < parallel; i++ {
		go func() {
			for {
				info := <-cm.downloadChan
				err := oss.DownloadFile(context.Background(), filePath(info.FileName), info.S3Key)
				if err != nil {
					logger.Sugar().Errorf("failed to download file from s3, err: %v", err)
				}
				cm.genCarChan <- info
			}
		}()
	}
	<-cm.downloadClose
}

func filePath(fileName string) string {
	return fmt.Sprintf("%v/%v", dataDir, fileName)
}

type CarFileInfo struct {
	TokenList []*TokenResInfo
	Size      int64
	CarName   string
	S3Bucket  string
	RootCID   string
}

func newCarFileInfo() *CarFileInfo {
	return &CarFileInfo{
		CarName: fmt.Sprintf("%v.car", uuid.NewString()),
	}
}

func (cm *CarManager) checkAndGenCar(maxUnTarSize int64) {
	carFI := newCarFileInfo()
	size := int64(0)
	for {
		select {
		case info := <-cm.genCarChan:
			size = carFI.Size + info.Size
			if size > maxUnTarSize {
				err := GenCarAndUpdate(context.Background(), carFI)
				if err != nil {
					logger.Sugar().Errorf("gen car failed, err: %v", err)
				}

				carFI = newCarFileInfo()
				size = info.Size
			}
			carFI.Size = size
			carFI.TokenList = append(carFI.TokenList, info)
		case <-cm.genCarClose:
			return
		}
	}
}

func GenCarAndUpdate(ctx context.Context, carFI *CarFileInfo) error {
	files := make([]string, len(carFI.TokenList))
	for i, token := range carFI.TokenList {
		files[i] = filePath(token.FileName)
	}

	carInfo, err := car.CreateCar(ctx, filePath(carFI.CarName), files, car.DefaultCarVersion)
	if err != nil {
		return err
	}
	logger.Sugar().Infof("gen car file: %v successfully, rootCID: %v", carFI.CarName, carInfo.RootCID)

	err = oss.UploadFile(ctx, filePath(carFI.CarName), carFI.CarName)
	if err != nil {
		return err
	}

	carFI.RootCID = carInfo.RootCID
	carFI.S3Bucket = oss.GetS3Bucket()
	logger.Sugar().Infof("update car file: %v to s3 successfully", carFI.CarName)

	cleanUpUsedCarFI(ctx, carFI)
	logger.Sugar().Infof("cleanup files related to car file: %v", carFI.CarName)

	// report to dealer
	items := make([]*dealer_proto.ContentItem, len(carFI.TokenList))
	for i, v := range carFI.TokenList {
		items[i] = &dealer_proto.ContentItem{
			ID:        v.ID,
			URI:       v.S3Key,
			ChainType: v.BaseInfo.ChainType,
			ChainID:   v.BaseInfo.ChainID,
			Contract:  v.BaseInfo.Contract,
			TokenID:   v.BaseInfo.TokenID,
			FileName:  v.FileName,
		}
	}

	snapshot, err := dealer_client.CreateSnapshot(
		ctx,
		&dealer_proto.CreateSnapshotRequest{
			SnapshotCommP: carInfo.RootCID,
			SnapshotRoot:  carInfo.RootCID,
			SnapshotURI:   carFI.CarName,
			Items:         items,
		},
	)
	logger.Sugar().Infof("report to dealer for create snapshot: %v, car: %v", snapshot.Info.ID, carFI.CarName)

	if err != nil {
		return err
	}

	return nil
}

func cleanUpUsedCarFI(ctx context.Context, carFI *CarFileInfo) {
	os.Remove(filePath(carFI.CarName))

	for _, v := range carFI.TokenList {
		os.Remove(filePath(v.FileName))
	}
}
