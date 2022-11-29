//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/cyber-tracer/message/cybertracer"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/version"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*cybertracer.VersionResponse, error) {
	resp, err := version.Version()
	if err != nil {
		logger.Sugar().Errorw("[Version] get service version error: %w", err)
		return &cybertracer.VersionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
