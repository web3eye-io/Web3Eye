//nolint:nolintlint,dupl
package order

import (
	"context"

	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

func (s *Server) CreateOrder(ctx context.Context, in *npool.CreateOrderRequest) (*npool.CreateOrderResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"CreateOrder",
			"In", in,
		)
		return &npool.CreateOrderResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	h, err := handler.NewHandler(ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithTxHash(in.Info.TxHash, true),
		handler.WithBlockNumber(in.Info.BlockNumber, true),
		handler.WithTxIndex(in.Info.TxIndex, true),
		handler.WithLogIndex(in.Info.LogIndex, true),
		handler.WithRecipient(in.Info.Recipient, true),
		handler.WithTargetItems(in.Info.TargetItems, true),
		handler.WithOfferItems(in.Info.OfferItems, true),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateOrder", "error", err)
		return &npool.CreateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.CreateOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateOrder", "error", err)
		return &npool.CreateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateOrders(ctx context.Context, in *npool.CreateOrdersRequest) (*npool.CreateOrdersResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateOrders", "error", "Infos is empty")
		return &npool.CreateOrdersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateOrders", "error", err)
		return &npool.CreateOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, err := h.CreateOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateOrders", "error", err)
		return &npool.CreateOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrdersResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpsertOrder(ctx context.Context, in *npool.UpsertOrderRequest) (*npool.UpsertOrderResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpsertOrder",
			"In", in,
		)
		return &npool.UpsertOrderResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	h, err := handler.NewHandler(ctx,
		handler.WithEntID(in.Info.EntID, false),
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, true),
		handler.WithTxHash(in.Info.TxHash, true),
		handler.WithBlockNumber(in.Info.BlockNumber, true),
		handler.WithTxIndex(in.Info.TxIndex, true),
		handler.WithLogIndex(in.Info.LogIndex, true),
		handler.WithRecipient(in.Info.Recipient, true),
		handler.WithTargetItems(in.Info.TargetItems, true),
		handler.WithOfferItems(in.Info.OfferItems, true),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpsertOrder", "error", err)
		return &npool.UpsertOrderResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.UpdateOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpsertOrder", "error", err)
		return &npool.UpsertOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpsertOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) UpsertOrders(ctx context.Context, in *npool.UpsertOrdersRequest) (*npool.UpsertOrdersResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("UpsertOrders", "error", "Infos is empty")
		return &npool.UpsertOrdersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("UpsertOrders", "error", err)
		return &npool.UpsertOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, err := h.UpsertOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpsertOrders", "error", err)
		return &npool.UpsertOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpsertOrdersResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateOrder(ctx context.Context, in *npool.UpdateOrderRequest) (*npool.UpdateOrderResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpsertBlock",
			"In", in,
		)
		return &npool.UpdateOrderResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithChainType(in.Info.ChainType, false),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithTxHash(in.Info.TxHash, false),
		handler.WithBlockNumber(in.Info.BlockNumber, false),
		handler.WithTxIndex(in.Info.TxIndex, false),
		handler.WithLogIndex(in.Info.LogIndex, false),
		handler.WithRecipient(in.Info.Recipient, false),
		handler.WithTargetItems(in.Info.TargetItems, false),
		handler.WithOfferItems(in.Info.OfferItems, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateOrder", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpdateOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateOrder", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, in *npool.GetOrderRequest) (*npool.GetOrderResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetOrder", "error", err)
		return &npool.GetOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.GetOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetOrder", "ID", in.GetID(), "error", err)
		return &npool.GetOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOrderOnly(ctx context.Context, in *npool.GetOrderOnlyRequest) (*npool.GetOrderOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetOrderOnly", "error", err)
		return &npool.GetOrderOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetOrderOnly", "error", err)
		return &npool.GetOrderOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total != 1 {
		errMsg := "more than one result or have no result"
		return &npool.GetOrderOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &npool.GetOrderOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetOrders(ctx context.Context, in *npool.GetOrdersRequest) (*npool.GetOrdersResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetOrders", "error", err)
		return &npool.GetOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetOrders", "error", err)
		return &npool.GetOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) ExistOrder(ctx context.Context, in *npool.ExistOrderRequest) (*npool.ExistOrderResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistOrder", "ID", in.GetID(), "error", err)
		return &npool.ExistOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := h.ExistOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistOrder", "ID", in.GetID(), "error", err)
		return &npool.ExistOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistOrderConds(ctx context.Context, in *npool.ExistOrderCondsRequest) (*npool.ExistOrderCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistOrderConds", "error", err)
		return &npool.ExistOrderCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := h.ExistOrderConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistOrderConds", "error", err)
		return &npool.ExistOrderCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteOrder(ctx context.Context, in *npool.DeleteOrderRequest) (*npool.DeleteOrderResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("DeleteOrder", "error", err)
		return &npool.DeleteOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := h.DeleteOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteOrder", "ID", in.GetID(), "error", err)
		return &npool.DeleteOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteOrderResponse{
		Info: info,
	}, nil
}
