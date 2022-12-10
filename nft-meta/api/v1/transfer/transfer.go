//nolint:nolintlint,dupl
package transfer

import (
	"context"

	converter "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/converter/v1/transfer"
	crud "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/crud/v1/transfer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/transfer"

	"github.com/google/uuid"
)

func (s *Server) CreateTransfer(ctx context.Context, in *npool.CreateTransferRequest) (*npool.CreateTransferResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateTransfer", "error", err)
		return &npool.CreateTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTransferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateTransfers(ctx context.Context, in *npool.CreateTransfersRequest) (*npool.CreateTransfersResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateTransfers", "error", "Infos is empty")
		return &npool.CreateTransfersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateTransfers", "error", err)
		return &npool.CreateTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTransfersResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateTransfer(ctx context.Context, in *npool.UpdateTransferRequest) (*npool.UpdateTransferResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateTransfer", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateTransfer", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTransferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTransfer(ctx context.Context, in *npool.GetTransferRequest) (*npool.GetTransferResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetTransfer", "ID", in.GetID(), "error", err)
		return &npool.GetTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetTransfer", "ID", in.GetID(), "error", err)
		return &npool.GetTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTransferOnly(ctx context.Context, in *npool.GetTransferOnlyRequest) (*npool.GetTransferOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetTransferOnly", "error", err)
		return &npool.GetTransferOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransferOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTransfers(ctx context.Context, in *npool.GetTransfersRequest) (*npool.GetTransfersResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetTransfers", "error", err)
		return &npool.GetTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransfersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistTransfer(ctx context.Context, in *npool.ExistTransferRequest) (*npool.ExistTransferResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistTransfer", "ID", in.GetID(), "error", err)
		return &npool.ExistTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistTransfer", "ID", in.GetID(), "error", err)
		return &npool.ExistTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTransferResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTransferConds(ctx context.Context, in *npool.ExistTransferCondsRequest) (*npool.ExistTransferCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistTransferConds", "error", err)
		return &npool.ExistTransferCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTransferCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountTransfers(ctx context.Context, in *npool.CountTransfersRequest) (*npool.CountTransfersResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountTransfersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTransfersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteTransfer(ctx context.Context, in *npool.DeleteTransferRequest) (*npool.DeleteTransferResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteTransfer", "ID", in.GetID(), "error", err)
		return &npool.DeleteTransferResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteTransfer", "ID", in.GetID(), "error", err)
		return &npool.DeleteTransferResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTransferResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
