//nolint:nolintlint,dupl
package block

import (
	"context"

	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/block"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/block"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"

	"github.com/google/uuid"
)

func (s *Server) CreateBlock(ctx context.Context, in *npool.CreateBlockRequest) (*npool.CreateBlockResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateBlock", "error", err)
		return &npool.CreateBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBlockResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) UpsertBlock(ctx context.Context, in *npool.UpsertBlockRequest) (*npool.UpsertBlockResponse, error) {
	var err error

	info, err := crud.Upsert(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateBlock", "error", err)
		return &npool.UpsertBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpsertBlockResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateBlocks(ctx context.Context, in *npool.CreateBlocksRequest) (*npool.CreateBlocksResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateBlocks", "error", "Infos is empty")
		return &npool.CreateBlocksResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateBlocks", "error", err)
		return &npool.CreateBlocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBlocksResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateBlock(ctx context.Context, in *npool.UpdateBlockRequest) (*npool.UpdateBlockResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateBlock", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBlockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateBlock", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateBlockResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBlock(ctx context.Context, in *npool.GetBlockRequest) (*npool.GetBlockResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetBlock", "ID", in.GetID(), "error", err)
		return &npool.GetBlockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetBlock", "ID", in.GetID(), "error", err)
		return &npool.GetBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBlockResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBlockOnly(ctx context.Context, in *npool.GetBlockOnlyRequest) (*npool.GetBlockOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetBlockOnly", "error", err)
		return &npool.GetBlockOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBlockOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBlocks(ctx context.Context, in *npool.GetBlocksRequest) (*npool.GetBlocksResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetBlocks", "error", err)
		return &npool.GetBlocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBlocksResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistBlock(ctx context.Context, in *npool.ExistBlockRequest) (*npool.ExistBlockResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistBlock", "ID", in.GetID(), "error", err)
		return &npool.ExistBlockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistBlock", "ID", in.GetID(), "error", err)
		return &npool.ExistBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBlockResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistBlockConds(ctx context.Context, in *npool.ExistBlockCondsRequest) (*npool.ExistBlockCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistBlockConds", "error", err)
		return &npool.ExistBlockCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBlockCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) CountBlocks(ctx context.Context, in *npool.CountBlocksRequest) (*npool.CountBlocksResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountBlocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountBlocksResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteBlock(ctx context.Context, in *npool.DeleteBlockRequest) (*npool.DeleteBlockResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteBlock", "ID", in.GetID(), "error", err)
		return &npool.DeleteBlockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteBlock", "ID", in.GetID(), "error", err)
		return &npool.DeleteBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteBlockResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
