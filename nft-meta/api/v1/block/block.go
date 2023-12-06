//nolint:nolintlint,dupl
package block

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/block"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

func (s *Server) CreateBlock(ctx context.Context, in *npool.CreateBlockRequest) (*npool.CreateBlockResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithBlockNumber(in.Info.BlockNumber, true),
		handler.WithBlockHash(in.Info.BlockHash, true),
		handler.WithBlockTime(in.Info.BlockTime, true),
		handler.WithParseState(in.Info.ParseState, true),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateBlock", "error", err)
		return &npool.CreateBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.CreateBlock(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateBlock", "error", err)
		return &npool.CreateBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBlockResponse{
		Info: info,
	}, nil
}

func (s *Server) UpsertBlock(ctx context.Context, in *npool.UpsertBlockRequest) (*npool.UpsertBlockResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithBlockNumber(in.Info.BlockNumber, true),
		handler.WithBlockHash(in.Info.BlockHash, false),
		handler.WithBlockTime(in.Info.BlockTime, false),
		handler.WithParseState(in.Info.ParseState, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpsertBlock", "error", err)
		return &npool.UpsertBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpsertBlock(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpsertBlock", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpsertBlockResponse{
		Info: info,
	}, err
}

func (s *Server) CreateBlocks(ctx context.Context, in *npool.CreateBlocksRequest) (*npool.CreateBlocksResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateBlocks", "error", "Infos is empty")
		return &npool.CreateBlocksResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateBlocks", "error", err)
		return &npool.CreateBlocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := h.CreateBlocks(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateBlocks", "error", err)
		return &npool.CreateBlocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateBlocksResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateBlock(ctx context.Context, in *npool.UpdateBlockRequest) (*npool.UpdateBlockResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithChainType(in.Info.ChainType, false),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithBlockNumber(in.Info.BlockNumber, false),
		handler.WithBlockHash(in.Info.BlockHash, false),
		handler.WithBlockTime(in.Info.BlockTime, false),
		handler.WithParseState(in.Info.ParseState, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateBlock", "error", err)
		return &npool.UpdateBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpdateBlock(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateBlock", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateBlockResponse{
		Info: info,
	}, nil
}

func (s *Server) GetBlock(ctx context.Context, in *npool.GetBlockRequest) (*npool.GetBlockResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetBlock", "error", err)
		return &npool.GetBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.GetBlock(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetBlock", "ID", in.GetID(), "error", err)
		return &npool.GetBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBlockResponse{
		Info: info,
	}, nil
}

func (s *Server) GetBlockOnly(ctx context.Context, in *npool.GetBlockOnlyRequest) (*npool.GetBlockOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetBlockOnly", "error", err)
		return &npool.GetBlockOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetBlocks(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetBlockOnly", "error", err)
		return &npool.GetBlockOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total > 1 {
		errMsg := "more than one result"
		logger.Sugar().Errorw("GetBlockOnly", "error", errMsg)
		return &npool.GetBlockOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &npool.GetBlockOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetBlocks(ctx context.Context, in *npool.GetBlocksRequest) (*npool.GetBlocksResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetBlocks", "error", err)
		return &npool.GetBlocksResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetBlocks(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetBlocks", "error", err)
		return &npool.GetBlocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBlocksResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistBlock(ctx context.Context, in *npool.ExistBlockRequest) (*npool.ExistBlockResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistBlock", "ID", in.GetID(), "error", err)
		return &npool.ExistBlockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := h.ExistBlock(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistBlock", "ID", in.GetID(), "error", err)
		return &npool.ExistBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBlockResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistBlockConds(ctx context.Context, in *npool.ExistBlockCondsRequest) (*npool.ExistBlockCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistBlockConds", "error", err)
		return &npool.ExistBlockCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := h.ExistBlockConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistBlockConds", "error", err)
		return &npool.ExistBlockCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBlockCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteBlock(ctx context.Context, in *npool.DeleteBlockRequest) (*npool.DeleteBlockResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistBlockConds", "error", err)
		return &npool.DeleteBlockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := h.DeleteBlock(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteBlock", "ID", in.GetID(), "error", err)
		return &npool.DeleteBlockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteBlockResponse{
		Info: info,
	}, nil
}
