package v1

import (
	"context"
	"os"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/config"
	transformnpool "github.com/web3eye-io/Web3Eye/proto/web3eye/transform/v1"
	"github.com/web3eye-io/Web3Eye/transform/pkg/filegetter"
	"github.com/web3eye-io/Web3Eye/transform/pkg/model"
)

func (s *Server) URLToVector(ctx context.Context, in *transformnpool.UrlToVectorReq) (*transformnpool.UrlToVectorResp, error) {
	dataDir := config.GetConfig().Transform.DataDir
	fileName := uuid.NewString()
	path, _, err := filegetter.GetFileFromURL(ctx, in.GetUrl(), dataDir, fileName)
	if err != nil {
		logger.Sugar().Errorf("failed to download file form url", err)
		return nil, err
	}
	defer os.Remove(*path)

	vector, err := model.ToImageVector(ctx, *path)
	if err != nil {
		logger.Sugar().Errorf("failed to transform url to vector", err)
		return nil, err
	}

	logger.Sugar().Infof("success to transform url(%.50s) to vector", in.Url)

	return &transformnpool.UrlToVectorResp{
		Vector: vector,
		Msg:    "got it!",
	}, nil
}
