//nolint:nolintlint,dupl
package v1

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	backup "github.com/web3eye-io/Web3Eye/dealer/pkg/backup"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBackup(ctx context.Context, in *npool.CreateBackupRequest) (*npool.CreateBackupResponse, error) {
	handler, err := backup.NewHandler(
		backup.WithIndex(in.GetIndex()),
	)
	if err != nil {
		logger.Sugar().Infow(
			"CreateBackup",
			"In", in,
			"Error", err,
		)
		return &npool.CreateBackupResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateBackup(ctx)
	if err != nil {
		logger.Sugar().Infow(
			"CreateBackup",
			"In", in,
			"Error", err,
		)
		return &npool.CreateBackupResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBackupResponse{
		Info: info,
	}, nil
}
