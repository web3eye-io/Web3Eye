//go:build !codeanalysis
// +build !codeanalysis

package v1

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"github.com/mr-tron/base58/base58"
	"github.com/web3eye-io/Web3Eye/common/ctfile"
	"github.com/web3eye-io/Web3Eye/common/oss"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/gen-car/pkg/car"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/gencar/v1"
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

func (s *Server) ReportFile(ctx context.Context, in *v1.ReportFileRequest) (*v1.ReportFileResponse, error) {
	logger.Sugar().Infof("start deal file, id %v, s3key %v", in.ID, in.S3Key)
	info, err := token.GetToken(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	objAttr, err := oss.GetObjectAttributes(ctx, in.S3Key)
	if err != nil {
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
	sha2sum, err := AnySHA256Sum(resInfo.BaseInfo)
	if err != nil {
		return nil, err
	}

	resInfo.FileName = fmt.Sprintf("%v%v", sha2sum, path.Ext(in.S3Key))

	go PutTokenResInfo(resInfo)

	return &v1.ReportFileResponse{
		ID:    in.ID,
		S3Key: in.S3Key,
	}, nil
}

func AnySHA256Sum(obj any) (string, error) {
	rawObj, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	_h := sha256.Sum256(rawObj)
	h := base58.Encode(_h[:])
	return h, nil
}

func CheckAnySHA256Sum(obj any, base58str string) (bool, error) {
	h, err := AnySHA256Sum(obj)
	if err != nil {
		return false, err
	}
	if h == base58str {
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
	maxUnTarSize = 24253611
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
	fmt.Println(carManager.dataDir)
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
	TarGzName string
	CarName   string
	S3Bucket  string
	RootCID   string
}

func newCarFileInfo() *CarFileInfo {
	return &CarFileInfo{
		TarGzName: fmt.Sprintf("%v.tar.gz", uuid.NewString()),
		CarName:   fmt.Sprintf("%v.car", uuid.NewString()),
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
	err := ctfile.GenTarGZ(filePath(carFI.TarGzName), files)
	if err != nil {
		return err
	}
	logger.Sugar().Infof("gen tar.gz file:%v successully, car: %v, has %v files", carFI.TarGzName, carFI.CarName, len(carFI.TokenList))

	rootCID, err := car.CreateCar(ctx, filePath(carFI.CarName), []string{filePath(carFI.TarGzName)}, car.DefaultCarVersion)
	if err != nil {
		return err
	}
	logger.Sugar().Infof("gen car file: %v successully, tar: %v, rootCID: %v", carFI.CarName, carFI.TarGzName, rootCID)

	err = oss.UploadFile(ctx, filePath(carFI.CarName), carFI.CarName)
	if err != nil {
		return err
	}
	carFI.RootCID = rootCID
	carFI.S3Bucket = oss.GetS3Bucket()
	logger.Sugar().Infof("update car file: %v to s3 successully", carFI.CarName)

	cleanUpUsedCarFI(ctx, carFI)
	logger.Sugar().Infof("cleanup files related to car file: %v", carFI.CarName)

	return nil
}

func cleanUpUsedCarFI(ctx context.Context, carFI *CarFileInfo) {
	os.Remove(filePath(carFI.CarName))
	os.Remove(filePath(carFI.TarGzName))

	for _, v := range carFI.TokenList {
		os.Remove(filePath(v.FileName))
	}
}
