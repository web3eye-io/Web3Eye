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
	"sort"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/mr-tron/base58"
	"github.com/web3eye-io/Web3Eye/common/oss"
	"github.com/web3eye-io/Web3Eye/config"
	dealer_client "github.com/web3eye-io/Web3Eye/dealer/pkg/client/v1"
	"github.com/web3eye-io/Web3Eye/gen-car/pkg/car"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	dealer_proto "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
	gencar_proto "github.com/web3eye-io/Web3Eye/proto/web3eye/gencar/v1"
	token_proto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

const (
	ImageResType = "image"
)

var (
	dataDir = config.GetConfig().GenCar.DataDir
)

const (
	DefaultDownloadChanLen  = 100
	DefaultDownloadParallel = 3
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
	ID       uint32
	ResType  string
}

func (s *Server) ReportFile(ctx context.Context, in *gencar_proto.ReportFileRequest) (*gencar_proto.ReportFileResponse, error) {
	logger.Sugar().Infof("start deal file, id %v, s3key %v", in.ID, in.S3Key)
	resp, err := token.GetToken(ctx, &token_proto.GetTokenRequest{ID: in.ID})
	if err != nil {
		logger.Sugar().Infof("failed get token by id, err: %v", in.ID, err)
		return nil, err
	}
	objAttr, err := oss.GetObjectAttributes(ctx, config.GetConfig().Minio.TokenImageBucket, in.S3Key)
	if err != nil {
		logger.Sugar().Infof("failed get token image by id, err: %v", in.ID, err)
		return nil, err
	}
	resInfo := &TokenResInfo{
		BaseInfo: TokenBaseInfo{
			ChainType: resp.Info.ChainType.String(),
			ChainID:   resp.Info.ChainID,
			Contract:  resp.Info.Contract,
			TokenID:   resp.Info.TokenID,
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

	go carManager.checkAndGenCar(int64(config.GetConfig().GenCar.MaxTarSize))
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
				err := oss.DownloadFile(context.Background(), filePath(info.FileName), config.GetConfig().Minio.TokenImageBucket, info.S3Key)
				if err != nil {
					logger.Sugar().Errorf("failed to download file from s3, err: %v", err)
					continue
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
			if info.Size > maxUnTarSize || info.BaseInfo.ChainID == "" {
				continue
			}
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

	defer cleanUpUsedCarFI(ctx, carFI)

	carInfo, err := car.CreateCar(ctx, filePath(carFI.CarName), files, car.DefaultCarVersion)
	if err != nil {
		return err
	}
	logger.Sugar().Infof("gen car file: %v successfully, rootCID: %v", carFI.CarName, carInfo.RootCID)

	err = oss.UploadFile(ctx, filePath(carFI.CarName), config.GetConfig().Minio.CarBucket, carFI.CarName)
	if err != nil {
		return err
	}
	return nil
}

func GenCarAndUpdate1(ctx context.Context, carFI *CarFileInfo) error {
	files := make([]string, len(carFI.TokenList))
	for i, token := range carFI.TokenList {
		files[i] = filePath(token.FileName)
	}

	carInfo, err := car.CreateCar(ctx, filePath(carFI.CarName), files, car.DefaultCarVersion)
	if err != nil {
		return err
	}
	logger.Sugar().Infof("gen car file: %v successfully, rootCID: %v", carFI.CarName, carInfo.RootCID)

	err = oss.UploadFile(ctx, filePath(carFI.CarName), config.GetConfig().Minio.CarBucket, carFI.CarName)
	if err != nil {
		return err
	}

	carFI.RootCID = carInfo.RootCID
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
			SnapshotCommP: carInfo.CommPCID,
			SnapshotRoot:  carInfo.RootCID,
			SnapshotURI:   carFI.CarName,
			Items:         items,
		},
	)
	if err != nil {
		return err
	}
	if snapshot != nil {
		logger.Sugar().Infof("report to dealer for create snapshot: %v, car: %v", snapshot.Info.ID, carFI.CarName)
	}

	return nil
}

func cleanUpUsedCarFI(ctx context.Context, carFI *CarFileInfo) {
	os.Remove(filePath(carFI.CarName))

	files := []string{}
	for _, v := range carFI.TokenList {
		//nolint
		// os.Remove(filePath(v.FileName))
		files = append(files, v.S3Key, v.FileName)
	}

	testFile, err := os.Create(fmt.Sprintf("%v/%v.list", config.GetConfig().GenCar.DataDir, carFI.CarName))
	if err != nil {
		logger.Sugar().Error(err)
		return
	}
	defer testFile.Close()
	filesByte, err := json.Marshal(files)
	if err != nil {
		logger.Sugar().Error(err)
		return
	}
	_, err = testFile.Write(filesByte)
	if err != nil {
		logger.Sugar().Error(err)
		return
	}

	//nolint
	// err = oss.DeleteFiles(ctx, config.GetConfig().Minio.TokenImageBucket, files)
	// if err != nil {
	// 	logger.Sugar().Error(err)
	// 	return
	// }

	err = deleteOverFiles(ctx, int(config.GetConfig().Minio.MaxCarNum), config.GetConfig().Minio.CarBucket)
	if err != nil {
		logger.Sugar().Error(err)
		return
	}

	err = deleteOverFiles(ctx, int(config.GetConfig().Minio.MaxTarNum), config.GetConfig().Minio.TarBucket)
	if err != nil {
		logger.Sugar().Error(err)
		return
	}
}

func deleteOverFiles(ctx context.Context, topN int, bucket string) error {
	out, err := oss.GetS3Client().ListObjects(ctx, &s3.ListObjectsInput{
		Bucket: &bucket,
	})
	if err != nil {
		return err
	}
	sort.Slice(out.Contents, func(i, j int) bool {
		return out.Contents[i].LastModified.After(*out.Contents[j].LastModified)
	})
	if len(out.Contents) < topN {
		return nil
	}
	out.Contents = out.Contents[topN:]
	files := []string{}
	for _, v := range out.Contents {
		files = append(files, *v.Key)
	}
	if len(files) == 0 {
		return nil
	}
	return oss.DeleteFiles(ctx, bucket, files)
}
