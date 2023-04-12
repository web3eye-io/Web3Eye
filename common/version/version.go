package version

import (
	"fmt"

	"github.com/web3eye-io/Web3Eye/proto/web3eye"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"
)

func Version() (*web3eye.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		logger.Sugar().Errorf("get service version error: %+w", err)
		return nil, fmt.Errorf("get service version error: %w", err)
	}
	return &web3eye.VersionResponse{
		Info: info,
	}, nil
}
