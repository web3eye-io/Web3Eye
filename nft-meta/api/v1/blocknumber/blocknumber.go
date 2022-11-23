//nolint:nolintlint,dupl
package blocknumber

import (
	"context"

	converter "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/converter/v1/blocknumber"
	crud "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/crud/v1/blocknumber"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/imageconvert"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/nftmeta/v1/blocknumber"

	"github.com/google/uuid"
)

func (s *Server) CreateBlockNumber(ctx context.Context, in *npool.CreateBlockNumberRequest) (*npool.CreateBlockNumberResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateBlockNumber", "error", err)
		return &npool.CreateBlockNumberResponse{}, status.Error(codes.Internal, err.Error())
	}

	go func() {
		imageconvert.DealVectorState(context.Background(), info.ID)
	}()

	return &npool.CreateBlockNumberResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateBlockNumbers(ctx context.Context, in *npool.CreateBlockNumbersRequest) (*npool.CreateBlockNumbersResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateBlockNumbers", "error", "Infos is empty")
		return &npool.CreateBlockNumbersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateBlockNumbers", "error", err)
		return &npool.CreateBlockNumbersResponse{}, status.Error(codes.Internal, err.Error())
	}

	go func() {
		for _, info := range rows {
			imageconvert.DealVectorState(context.Background(), info.ID)
		}
	}()

	return &npool.CreateBlockNumbersResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateBlockNumber(ctx context.Context, in *npool.UpdateBlockNumberRequest) (*npool.UpdateBlockNumberResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateBlockNumber", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBlockNumberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateBlockNumber", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateBlockNumberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateBlockNumberResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBlockNumber(ctx context.Context, in *npool.GetBlockNumberRequest) (*npool.GetBlockNumberResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetBlockNumber", "ID", in.GetID(), "error", err)
		return &npool.GetBlockNumberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetBlockNumber", "ID", in.GetID(), "error", err)
		return &npool.GetBlockNumberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBlockNumberResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBlockNumberOnly(ctx context.Context, in *npool.GetBlockNumberOnlyRequest) (*npool.GetBlockNumberOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetBlockNumberOnly", "error", err)
		return &npool.GetBlockNumberOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBlockNumberOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetBlockNumbers(ctx context.Context, in *npool.GetBlockNumbersRequest) (*npool.GetBlockNumbersResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetBlockNumbers", "error", err)
		return &npool.GetBlockNumbersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetBlockNumbersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistBlockNumber(ctx context.Context, in *npool.ExistBlockNumberRequest) (*npool.ExistBlockNumberResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistBlockNumber", "ID", in.GetID(), "error", err)
		return &npool.ExistBlockNumberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistBlockNumber", "ID", in.GetID(), "error", err)
		return &npool.ExistBlockNumberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBlockNumberResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistBlockNumberConds(ctx context.Context, in *npool.ExistBlockNumberCondsRequest) (*npool.ExistBlockNumberCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistBlockNumberConds", "error", err)
		return &npool.ExistBlockNumberCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistBlockNumberCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountBlockNumbers(ctx context.Context, in *npool.CountBlockNumbersRequest) (*npool.CountBlockNumbersResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountBlockNumbersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountBlockNumbersResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteBlockNumber(ctx context.Context, in *npool.DeleteBlockNumberRequest) (*npool.DeleteBlockNumberResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteBlockNumber", "ID", in.GetID(), "error", err)
		return &npool.DeleteBlockNumberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteBlockNumber", "ID", in.GetID(), "error", err)
		return &npool.DeleteBlockNumberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteBlockNumberResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
