//nolint:nolintlint,dupl
package token

import (
	"context"

	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/token"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/imageconvert"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"

	"github.com/google/uuid"
)

func (s *Server) CreateToken(ctx context.Context, in *npool.CreateTokenRequest) (*npool.CreateTokenResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateToken", "error", err)
		return &npool.CreateTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	go func() {
		err = imageconvert.QueueDealVector(info)
		if err != nil {
			logger.Sugar().Error(err)
		}
	}()

	return &npool.CreateTokenResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateTokens(ctx context.Context, in *npool.CreateTokensRequest) (*npool.CreateTokensResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateTokens", "error", "Infos is empty")
		return &npool.CreateTokensResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateTokens", "error", err)
		return &npool.CreateTokensResponse{}, status.Error(codes.Internal, err.Error())
	}

	go func() {
		for i := 0; i < len(rows); i++ {
			err = imageconvert.QueueDealVector(rows[i])
			if err != nil {
				logger.Sugar().Error(err)
			}
		}
	}()

	return &npool.CreateTokensResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateToken(ctx context.Context, in *npool.UpdateTokenRequest) (*npool.UpdateTokenResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateToken", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateToken", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTokenResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetToken(ctx context.Context, in *npool.GetTokenRequest) (*npool.GetTokenResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetToken", "ID", in.GetID(), "error", err)
		return &npool.GetTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetToken", "ID", in.GetID(), "error", err)
		return &npool.GetTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTokenResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTokenOnly(ctx context.Context, in *npool.GetTokenOnlyRequest) (*npool.GetTokenOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetTokenOnly", "error", err)
		return &npool.GetTokenOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTokenOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTokens(ctx context.Context, in *npool.GetTokensRequest) (*npool.GetTokensResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetTokens", "error", err)
		return &npool.GetTokensResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTokensResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistToken(ctx context.Context, in *npool.ExistTokenRequest) (*npool.ExistTokenResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistToken", "ID", in.GetID(), "error", err)
		return &npool.ExistTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistToken", "ID", in.GetID(), "error", err)
		return &npool.ExistTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTokenResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTokenConds(ctx context.Context, in *npool.ExistTokenCondsRequest) (*npool.ExistTokenCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistTokenConds", "error", err)
		return &npool.ExistTokenCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTokenCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountTokens(ctx context.Context, in *npool.CountTokensRequest) (*npool.CountTokensResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountTokensResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTokensResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteToken(ctx context.Context, in *npool.DeleteTokenRequest) (*npool.DeleteTokenResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteToken", "ID", in.GetID(), "error", err)
		return &npool.DeleteTokenResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteToken", "ID", in.GetID(), "error", err)
		return &npool.DeleteTokenResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTokenResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
