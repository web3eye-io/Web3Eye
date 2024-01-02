//nolint:nolintlint,dupl
package contract

import (
	"context"

	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/contract"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
)

func (s *Server) CreateContract(ctx context.Context, in *npool.CreateContractRequest) (*npool.CreateContractResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"CreateContract",
			"In", in,
		)
		return &npool.CreateContractResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(
		ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithAddress(in.Info.Address, true),
		handler.WithName(in.Info.Name, true),
		handler.WithSymbol(in.Info.Symbol, true),
		handler.WithDecimals(in.Info.Decimals, true),
		handler.WithCreator(in.Info.Creator, false),
		handler.WithBlockNum(in.Info.BlockNum, true),
		handler.WithTxHash(in.Info.TxHash, true),
		handler.WithTxTime(in.Info.TxTime, true),
		handler.WithProfileURL(in.Info.ProfileURL, false),
		handler.WithBaseURL(in.Info.BaseURL, false),
		handler.WithBannerURL(in.Info.BannerURL, false),
		handler.WithDescription(in.Info.Description, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateContract", "error", err)
		return &npool.CreateContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.CreateContract(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateContract", "error", err)
		return &npool.CreateContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateContractResponse{
		Info: info,
	}, nil
}

func (s *Server) UpsertContract(ctx context.Context, in *npool.UpsertContractRequest) (*npool.UpsertContractResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpsertContract",
			"In", in,
		)
		return &npool.UpsertContractResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(
		ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithAddress(in.Info.Address, true),
		handler.WithName(in.Info.Name, true),
		handler.WithSymbol(in.Info.Symbol, true),
		handler.WithDecimals(in.Info.Decimals, true),
		handler.WithCreator(in.Info.Creator, false),
		handler.WithBlockNum(in.Info.BlockNum, true),
		handler.WithTxHash(in.Info.TxHash, true),
		handler.WithTxTime(in.Info.TxTime, true),
		handler.WithProfileURL(in.Info.ProfileURL, false),
		handler.WithBaseURL(in.Info.BaseURL, false),
		handler.WithBannerURL(in.Info.BannerURL, false),
		handler.WithDescription(in.Info.Description, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpsertContract", "error", err)
		return &npool.UpsertContractResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.UpsertContract(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpsertContract", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpsertContractResponse{
		Info: info,
	}, err
}

func (s *Server) CreateContracts(ctx context.Context, in *npool.CreateContractsRequest) (*npool.CreateContractsResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateContracts", "error", "Infos is empty")
		return &npool.CreateContractsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateContracts", "error", err)
		return &npool.CreateContractsResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, err := h.CreateContracts(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateContracts", "error", err)
		return &npool.CreateContractsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateContractsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateContract(ctx context.Context, in *npool.UpdateContractRequest) (*npool.UpdateContractResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpdateContract",
			"In", in,
		)
		return &npool.UpdateContractResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(
		ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithChainType(in.Info.ChainType, false),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithAddress(in.Info.Address, false),
		handler.WithName(in.Info.Name, false),
		handler.WithSymbol(in.Info.Symbol, false),
		handler.WithDecimals(in.Info.Decimals, false),
		handler.WithCreator(in.Info.Creator, false),
		handler.WithBlockNum(in.Info.BlockNum, false),
		handler.WithTxHash(in.Info.TxHash, false),
		handler.WithTxTime(in.Info.TxTime, false),
		handler.WithProfileURL(in.Info.ProfileURL, false),
		handler.WithBaseURL(in.Info.BaseURL, false),
		handler.WithBannerURL(in.Info.BannerURL, false),
		handler.WithDescription(in.Info.Description, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateContract", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpdateContract(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateContract", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateContractResponse{
		Info: info,
	}, nil
}

func (s *Server) GetContract(ctx context.Context, in *npool.GetContractRequest) (*npool.GetContractResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetContract", "ID", in.GetID(), "error", err)
		return &npool.GetContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.GetContract(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetContract", "ID", in.GetID(), "error", err)
		return &npool.GetContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContractResponse{
		Info: info,
	}, nil
}

func (s *Server) GetContractOnly(ctx context.Context, in *npool.GetContractOnlyRequest) (*npool.GetContractOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetContractOnly", "error", err)
		return &npool.GetContractOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := h.GetContracts(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetContractOnly", "error", err)
		return &npool.GetContractOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total != 1 {
		errMsg := "more than one result or have no result"
		logger.Sugar().Errorw("GetContractOnly", "error", errMsg)
		return &npool.GetContractOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &npool.GetContractOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetContracts(ctx context.Context, in *npool.GetContractsRequest) (*npool.GetContractsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetContracts", "error", err)
		return &npool.GetContractsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := h.GetContracts(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetContracts", "error", err)
		return &npool.GetContractsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContractsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) ExistContract(ctx context.Context, in *npool.ExistContractRequest) (*npool.ExistContractResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistContract", "ID", in.GetID(), "error", err)
		return &npool.ExistContractResponse{}, status.Error(codes.Internal, err.Error())
	}
	exist, err := h.ExistContract(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistContract", "ID", in.GetID(), "error", err)
		return &npool.ExistContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistContractResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistContractConds(ctx context.Context, in *npool.ExistContractCondsRequest) (*npool.ExistContractCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistContractConds", "error", err)
		return &npool.ExistContractCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := h.ExistContractConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistContractConds", "error", err)
		return &npool.ExistContractCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistContractCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteContract(ctx context.Context, in *npool.DeleteContractRequest) (*npool.DeleteContractResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("DeleteContract", "ID", in.GetID(), "error", err)
		return &npool.DeleteContractResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.DeleteContract(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteContract", "ID", in.GetID(), "error", err)
		return &npool.DeleteContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteContractResponse{
		Info: info,
	}, nil
}
