//nolint:nolintlint,dupl
package orderpair

import (
	"context"

	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/orderpair"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/orderpair"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/orderpair"

	"github.com/google/uuid"
)

func (s *Server) CreateOrderPair(ctx context.Context, in *npool.CreateOrderPairRequest) (*npool.CreateOrderPairResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateOrderPair", "error", err)
		return &npool.CreateOrderPairResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderPairResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateOrderPairs(ctx context.Context, in *npool.CreateOrderPairsRequest) (*npool.CreateOrderPairsResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateOrderPairs", "error", "Infos is empty")
		return &npool.CreateOrderPairsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateOrderPairs", "error", err)
		return &npool.CreateOrderPairsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderPairsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateOrderPair(ctx context.Context, in *npool.UpdateOrderPairRequest) (*npool.UpdateOrderPairResponse, error) {
	var err error

	if _, err := uuid.Parse(in.Info.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateOrderPair", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateOrderPairResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateOrderPair", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateOrderPairResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderPairResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderPair(ctx context.Context, in *npool.GetOrderPairRequest) (*npool.GetOrderPairResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetOrderPair", "ID", in.GetID(), "error", err)
		return &npool.GetOrderPairResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetOrderPair", "ID", in.GetID(), "error", err)
		return &npool.GetOrderPairResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderPairResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderPairOnly(ctx context.Context, in *npool.GetOrderPairOnlyRequest) (*npool.GetOrderPairOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetOrderPairOnly", "error", err)
		return &npool.GetOrderPairOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderPairOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOrderPairs(ctx context.Context, in *npool.GetOrderPairsRequest) (*npool.GetOrderPairsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetOrderPairs", "error", err)
		return &npool.GetOrderPairsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderPairsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistOrderPair(ctx context.Context, in *npool.ExistOrderPairRequest) (*npool.ExistOrderPairResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistOrderPair", "ID", in.GetID(), "error", err)
		return &npool.ExistOrderPairResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistOrderPair", "ID", in.GetID(), "error", err)
		return &npool.ExistOrderPairResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderPairResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistOrderPairConds(ctx context.Context, in *npool.ExistOrderPairCondsRequest) (*npool.ExistOrderPairCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistOrderPairConds", "error", err)
		return &npool.ExistOrderPairCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOrderPairCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) CountOrderPairs(ctx context.Context, in *npool.CountOrderPairsRequest) (*npool.CountOrderPairsResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountOrderPairsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountOrderPairsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteOrderPair(ctx context.Context, in *npool.DeleteOrderPairRequest) (*npool.DeleteOrderPairResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteOrderPair", "ID", in.GetID(), "error", err)
		return &npool.DeleteOrderPairResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteOrderPair", "ID", in.GetID(), "error", err)
		return &npool.DeleteOrderPairResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteOrderPairResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
