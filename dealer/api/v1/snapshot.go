//nolint:nolintlint,dupl
package v1

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	snapshot "github.com/web3eye-io/Web3Eye/dealer/pkg/snapshot"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSnapshot(ctx context.Context, in *npool.CreateSnapshotRequest) (*npool.CreateSnapshotResponse, error) {
	handler, err := snapshot.NewHandler(
		snapshot.WithSnapshotURI(in.GetSnapshotURI()),
		snapshot.WithItems(in.GetItems()),
	)
	if err != nil {
		logger.Sugar().Infow(
			"CreateSnapshot",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSnapshotResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := handler.CreateSnapshot(ctx); err != nil {
		logger.Sugar().Infow(
			"CreateSnapshot",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSnapshotResponse{}, nil
}
