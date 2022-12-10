//nolint:nolintlint,dupl
package contract

import (
	"context"

	converter "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/converter/v1/contract"
	crud "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/crud/v1/contract"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/contract"

	"github.com/google/uuid"
)

func (s *Server) CreateContract(ctx context.Context, in *npool.CreateContractRequest) (*npool.CreateContractResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateContract", "error", err)
		return &npool.CreateContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateContractResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateContracts(ctx context.Context, in *npool.CreateContractsRequest) (*npool.CreateContractsResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateContracts", "error", "Infos is empty")
		return &npool.CreateContractsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateContracts", "error", err)
		return &npool.CreateContractsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateContractsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateContract(ctx context.Context, in *npool.UpdateContractRequest) (*npool.UpdateContractResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateContract", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateContractResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateContract", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateContractResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetContract(ctx context.Context, in *npool.GetContractRequest) (*npool.GetContractResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetContract", "ID", in.GetID(), "error", err)
		return &npool.GetContractResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetContract", "ID", in.GetID(), "error", err)
		return &npool.GetContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContractResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetContractOnly(ctx context.Context, in *npool.GetContractOnlyRequest) (*npool.GetContractOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetContractOnly", "error", err)
		return &npool.GetContractOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContractOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetContracts(ctx context.Context, in *npool.GetContractsRequest) (*npool.GetContractsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetContracts", "error", err)
		return &npool.GetContractsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContractsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistContract(ctx context.Context, in *npool.ExistContractRequest) (*npool.ExistContractResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistContract", "ID", in.GetID(), "error", err)
		return &npool.ExistContractResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistContract", "ID", in.GetID(), "error", err)
		return &npool.ExistContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistContractResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistContractConds(ctx context.Context, in *npool.ExistContractCondsRequest) (*npool.ExistContractCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistContractConds", "error", err)
		return &npool.ExistContractCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistContractCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountContracts(ctx context.Context, in *npool.CountContractsRequest) (*npool.CountContractsResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountContractsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountContractsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteContract(ctx context.Context, in *npool.DeleteContractRequest) (*npool.DeleteContractResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteContract", "ID", in.GetID(), "error", err)
		return &npool.DeleteContractResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteContract", "ID", in.GetID(), "error", err)
		return &npool.DeleteContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteContractResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
