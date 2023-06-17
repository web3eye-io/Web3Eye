//nolint:nolintlint,dupl
package v1

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	retrieve "github.com/web3eye-io/Web3Eye/dealer/pkg/retrieve"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/retriever/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) StartRetrieve(ctx context.Context, in *npool.StartRetrieveRequest) (*npool.StartRetrieveResponse, error) {
	handler, err := retrieve.NewHandler(
		retrieve.WithChainType(in.GetChainType()),
		retrieve.WithChainID(in.GetChainID()),
		retrieve.WithContract(in.GetContract()),
		retrieve.WithTokenID(in.GetTokenID()),
	)
	if err != nil {
		logger.Sugar().Infow(
			"StartRetrieve",
			"In", in,
			"Error", err,
		)
		return &npool.StartRetrieveResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.StartRetrieve(ctx)
	if err != nil {
		logger.Sugar().Infow(
			"StartRetrieve",
			"In", in,
			"Error", err,
		)
		return &npool.StartRetrieveResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.StartRetrieveResponse{
		Info: info,
	}, nil
}

func (s *Server) StatRetrieve(ctx context.Context, in *npool.StatRetrieveRequest) (*npool.StatRetrieveResponse, error) {
	handler, err := retrieve.NewHandler(
		retrieve.WithChainType(in.GetChainType()),
		retrieve.WithChainID(in.GetChainID()),
		retrieve.WithContract(in.GetContract()),
		retrieve.WithTokenID(in.GetTokenID()),
	)
	if err != nil {
		logger.Sugar().Infow(
			"StatRetrieve",
			"In", in,
			"Error", err,
		)
		return &npool.StatRetrieveResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.StatRetrieve(ctx)
	if err != nil {
		logger.Sugar().Infow(
			"StatRetrieve",
			"In", in,
			"Error", err,
		)
		return &npool.StatRetrieveResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.StatRetrieveResponse{
		Info: info,
	}, nil
}
