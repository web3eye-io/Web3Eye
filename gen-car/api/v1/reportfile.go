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
	"github.com/mr-tron/base58/base58"
	"github.com/web3eye-io/Web3Eye/common/oss"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/gencar/v1"
)

const (
	ImageResType = "image"
	ImageTempDir = "./img"
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

	PutTokenResInfo(resInfo)

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
}

const (
	DefaultDownloadChanLen  = 100
	DefaultDownloadParallel = 3
)

var carManager *CarManager

func RunCarManager() {
	carManager = &CarManager{
		downloadChan:  make(chan *TokenResInfo, DefaultDownloadChanLen),
		downloadClose: make(chan struct{}),
		dataDir:       config.GetConfig().GenCar.DataDir,
	}

	var fileMode os.FileMode = 0777
	err := os.Mkdir(carManager.dataDir, fileMode)
	if err != nil && !strings.Contains(err.Error(), "file exists") {
		panic(err)
	}

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
				err := oss.DownloadFile(context.Background(), fmt.Sprintf("%v/%v", cm.dataDir, info.FileName), info.S3Key)
				if err != nil {
					logger.Sugar().Errorf("failed to download file from s3, err: %v", err)
				}
			}
		}()
	}
	<-cm.downloadClose
}
