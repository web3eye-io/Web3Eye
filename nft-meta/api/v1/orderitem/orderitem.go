//nolint:nolintlint,dupl
package orderitem

import (
	"context"

	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/orderitem"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/orderitem"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/orderitem"

	"github.com/google/uuid"
)

func (s *Server) CreateOrderItem(ctx context.Context, in *npool.CreateOrderItemRequest) (*npool.CreateOrderItemResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateOrderItem", "error", err)
		return &npool.CreateOrderItemResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderItemResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateOrderItems(ctx context.Context, in *npool.CreateOrderItemsRequest) (*npool.CreateOrderItemsResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateOrderItems", "error", "Infos is empty")
		return &npool.CreateOrderItemsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateOrderItems", "error", err)
		return &npool.CreateOrderItemsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderItemsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateOrderItem(ctx context.Context, in *npool.UpdateOrderItemRequest) (*npool.UpdateOrderItemResponse, error) {
	var err error

	if _, err := uuid.Parse(in.Info.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateOrderItem", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateOrderItemResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateOrderItem", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateOrderItemResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderItemResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderItem(ctx context.Context, in *npool.GetOrderItemRequest) (*npool.GetOrderItemResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetOrderItem", "ID", in.GetID(), "error", err)
		return &npool.GetOrderItemResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetOrderItem", "ID", in.GetID(), "error", err)
		return &npool.GetOrderItemResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderItemResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderItemOnly(ctx context.Context, in *npool.GetOrderItemOnlyRequest) (*npool.GetOrderItemOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetOrderItemOnly", "error", err)
		return &npool.GetOrderItemOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderItemOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderItems(ctx context.Context, in *npool.GetOrderItemsRequest) (*npool.GetOrderItemsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetOrderItems", "error", err)
		return &npool.GetOrderItemsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderItemsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistOrderItem(ctx context.Context, in *npool.ExistOrderItemRequest) (*npool.ExistOrderItemResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistOrderItem", "ID", in.GetID(), "error", err)
		return &npool.ExistOrderItemResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistOrderItem", "ID", in.GetID(), "error", err)
		return &npool.ExistOrderItemResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderItemResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistOrderItemConds(ctx context.Context, in *npool.ExistOrderItemCondsRequest) (*npool.ExistOrderItemCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistOrderItemConds", "error", err)
		return &npool.ExistOrderItemCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderItemCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) CountOrderItems(ctx context.Context, in *npool.CountOrderItemsRequest) (*npool.CountOrderItemsResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountOrderItemsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountOrderItemsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteOrderItem(ctx context.Context, in *npool.DeleteOrderItemRequest) (*npool.DeleteOrderItemResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteOrderItem", "ID", in.GetID(), "error", err)
		return &npool.DeleteOrderItemResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteOrderItem", "ID", in.GetID(), "error", err)
		return &npool.DeleteOrderItemResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteOrderItemResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
