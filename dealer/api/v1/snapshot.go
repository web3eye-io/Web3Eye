//nolint:nolintlint,dupl
package v1

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

func (s *Server) CreateSnapshot(ctx context.Context, in *npool.CreateSnapshotRequest) (*npool.CreateSnapshotResponse, error) {
	logger.Sugar().Infow(
		"CreateSnapshot",
		"In", in,
	)
	return &npool.CreateSnapshotResponse{}, nil
}
