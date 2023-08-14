package v1

import (
	"context"
	"os"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/converter/pkg/filegetter"
	"github.com/web3eye-io/Web3Eye/converter/pkg/model"
	converternpool "github.com/web3eye-io/Web3Eye/proto/web3eye/converter/v1"
)

func (s *Server) UrlToVector(ctx context.Context, in *converternpool.UrlToVectorReq) (*converternpool.UrlToVectorResp, error) {
	dataDir := config.GetConfig().Converter.DataDir
	fileName := uuid.NewString()
	path, err := filegetter.GetFileFromURL(in.GetUrl(), dataDir, fileName)
	if err != nil {
		logger.Sugar().Errorf("failed to download file form url", err)
		return nil, err
	}
	defer os.Remove(*path)

	vector, err := model.ToImageVector(*path)
	if err != nil {
		logger.Sugar().Errorf("failed to converter url to vector", err)
		return nil, err
	}

	logger.Sugar().Infof("success to converter url(%.50s) to vector", in.Url)

	return &converternpool.UrlToVectorResp{
		Vector: vector,
		Msg:    "got it!",
	}, nil
}
