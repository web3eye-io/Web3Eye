//nolint:nolintlint,dupl
package order

import (
	"context"

	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/order"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"

	"github.com/google/uuid"
)

func (s *Server) CreateOrder(ctx context.Context, in *npool.CreateOrderRequest) (*npool.CreateOrderResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateOrder", "error", err)
		return &npool.CreateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateOrders(ctx context.Context, in *npool.CreateOrdersRequest) (*npool.CreateOrdersResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateOrders", "error", "Infos is empty")
		return &npool.CreateOrdersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateOrders", "error", err)
		return &npool.CreateOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrdersResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateOrder(ctx context.Context, in *npool.UpdateOrderRequest) (*npool.UpdateOrderResponse, error) {
	var err error

	if _, err := uuid.Parse(in.Info.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateOrder", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateOrder", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, in *npool.GetOrderRequest) (*npool.GetOrderResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetOrder", "ID", in.GetID(), "error", err)
		return &npool.GetOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetOrder", "ID", in.GetID(), "error", err)
		return &npool.GetOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderOnly(ctx context.Context, in *npool.GetOrderOnlyRequest) (*npool.GetOrderOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetOrderOnly", "error", err)
		return &npool.GetOrderOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrders(ctx context.Context, in *npool.GetOrdersRequest) (*npool.GetOrdersResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetOrders", "error", err)
		return &npool.GetOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrdersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistOrder(ctx context.Context, in *npool.ExistOrderRequest) (*npool.ExistOrderResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistOrder", "ID", in.GetID(), "error", err)
		return &npool.ExistOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistOrder", "ID", in.GetID(), "error", err)
		return &npool.ExistOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistOrderConds(ctx context.Context, in *npool.ExistOrderCondsRequest) (*npool.ExistOrderCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistOrderConds", "error", err)
		return &npool.ExistOrderCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) CountOrders(ctx context.Context, in *npool.CountOrdersRequest) (*npool.CountOrdersResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountOrdersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteOrder(ctx context.Context, in *npool.DeleteOrderRequest) (*npool.DeleteOrderResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteOrder", "ID", in.GetID(), "error", err)
		return &npool.DeleteOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteOrder", "ID", in.GetID(), "error", err)
		return &npool.DeleteOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteOrderResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
