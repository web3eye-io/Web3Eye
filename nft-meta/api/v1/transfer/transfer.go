//nolint:nolintlint,dupl
package transfer

import (
	"context"

	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/transfer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func (s *Server) CreateTransfer(ctx context.Context, in *npool.CreateTransferRequest) (*npool.CreateTransferResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"CreateTransfer",
			"In", in,
		)
		return &npool.CreateTransferResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithContract(in.Info.Contract, true),
		handler.WithTokenType(in.Info.TokenType, true),
		handler.WithTokenID(in.Info.TokenID, true),
		handler.WithFrom(in.Info.From, true),
		handler.WithTo(in.Info.To, true),
		handler.WithAmount(in.Info.Amount, false),
		handler.WithBlockNumber(in.Info.BlockNumber, false),
		handler.WithTxHash(in.Info.TxHash, true),
		handler.WithBlockHash(in.Info.BlockHash, false),
		handler.WithTxTime(in.Info.TxTime, false),
		handler.WithRemark(in.Info.Remark, false),
		handler.WithLogIndex(in.Info.LogIndex, false),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateTransfer", "error", err)
		return &npool.CreateTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.GetTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateTransfer", "error", err)
		return &npool.CreateTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTransferResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateTransfers(ctx context.Context, in *npool.CreateTransfersRequest) (*npool.CreateTransfersResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Warnw("CreateTransfers", "error", "Infos is empty")
		return &npool.CreateTransfersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateTransfers", "error", err)
		return &npool.CreateTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := h.CreateTransfers(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateTransfers", "error", err)
		return &npool.CreateTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTransfersResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpsertTransfer(ctx context.Context, in *npool.UpsertTransferRequest) (*npool.UpsertTransferResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpsertTransfer",
			"In", in,
		)
		return &npool.UpsertTransferResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithContract(in.Info.Contract, true),
		handler.WithTokenType(in.Info.TokenType, true),
		handler.WithTokenID(in.Info.TokenID, true),
		handler.WithFrom(in.Info.From, true),
		handler.WithTo(in.Info.To, true),
		handler.WithAmount(in.Info.Amount, false),
		handler.WithBlockNumber(in.Info.BlockNumber, false),
		handler.WithTxHash(in.Info.TxHash, true),
		handler.WithBlockHash(in.Info.BlockHash, false),
		handler.WithTxTime(in.Info.TxTime, false),
		handler.WithRemark(in.Info.Remark, false),
		handler.WithLogIndex(in.Info.LogIndex, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpsertTransfer", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpsertTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpsertTransfer", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpsertTransferResponse{
		Info: info,
	}, err
}

func (s *Server) UpsertTransfers(ctx context.Context, in *npool.UpsertTransfersRequest) (*npool.UpsertTransfersResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Warnw("UpsertTransfers", "error", "Infos is empty")
		return &npool.UpsertTransfersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("UpsertTransfers", "error", err)
		return &npool.UpsertTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = h.UpsertTransfers(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpsertTransfers", "error", err)
		return &npool.UpsertTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpsertTransfersResponse{}, nil
}

func (s *Server) UpdateTransfer(ctx context.Context, in *npool.UpdateTransferRequest) (*npool.UpdateTransferResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpdateTransfer",
			"In", in,
		)
		return &npool.UpdateTransferResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithChainType(in.Info.ChainType, false),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithContract(in.Info.Contract, false),
		handler.WithTokenType(in.Info.TokenType, false),
		handler.WithTokenID(in.Info.TokenID, false),
		handler.WithFrom(in.Info.From, false),
		handler.WithTo(in.Info.To, false),
		handler.WithAmount(in.Info.Amount, false),
		handler.WithBlockNumber(in.Info.BlockNumber, false),
		handler.WithTxHash(in.Info.TxHash, false),
		handler.WithBlockHash(in.Info.BlockHash, false),
		handler.WithTxTime(in.Info.TxTime, false),
		handler.WithRemark(in.Info.Remark, false),
		handler.WithLogIndex(in.Info.LogIndex, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateTransfer", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpdateTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateTransfer", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTransferResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTransfer(ctx context.Context, in *npool.GetTransferRequest) (*npool.GetTransferResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetTransfer", "error", err)
		return &npool.GetTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.GetTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTransfer", "ID", in.GetID(), "error", err)
		return &npool.GetTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransferResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTransferOnly(ctx context.Context, in *npool.GetTransferOnlyRequest) (*npool.GetTransferOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetTransferOnly", "error", err)
		return &npool.GetTransferOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetTransfers(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTransferOnly", "error", err)
		return &npool.GetTransferOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total != 1 {
		errMsg := "more than one result or have no result"
		return &npool.GetTransferOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &npool.GetTransferOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetTransfers(ctx context.Context, in *npool.GetTransfersRequest) (*npool.GetTransfersResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetTransfers", "error", err)
		return &npool.GetTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetTransfers(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTransfers", "error", err)
		return &npool.GetTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransfersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) ExistTransfer(ctx context.Context, in *npool.ExistTransferRequest) (*npool.ExistTransferResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistTransfer", "ID", in.GetID(), "error", err)
		return &npool.ExistTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := h.ExistTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistTransfer", "ID", in.GetID(), "error", err)
		return &npool.ExistTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTransferResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistTransferConds(ctx context.Context, in *npool.ExistTransferCondsRequest) (*npool.ExistTransferCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistTransferConds", "error", err)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := h.ExistTransferConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistTransferConds", "error", err)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTransferCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteTransfer(ctx context.Context, in *npool.DeleteTransferRequest) (*npool.DeleteTransferResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("DeleteTransfer", "error", err)
		return &npool.DeleteTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := h.DeleteTransfer(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteTransfer", "ID", in.GetID(), "error", err)
		return &npool.DeleteTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTransferResponse{
		Info: info,
	}, nil
}
