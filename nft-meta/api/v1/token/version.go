//go:build !codeanalysis
// +build !codeanalysis

package token

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/version"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*npool.VersionResponse, error) {
	resp, err := version.Version()
	if err != nil {
		logger.Sugar().Errorw("[Version] get service version error: %w", err)
		return &npool.VersionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
