//nolint:nolintlint,dupl
package synctask

import (
	"context"
	"fmt"

	"github.com/web3eye-io/cyber-tracer/common/ctkafka"
	converter "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/converter/v1/synctask"
	crud "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/crud/v1/synctask"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/cyber-tracer/proto/cybertracer"
	"github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/cttype"
	npool "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/synctask"
	cttype "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/cttype"

	"github.com/google/uuid"
)

const (
	MaxPutTaskNumOnce = 500
)

func (s *Server) CreateSyncTask(ctx context.Context, in *npool.CreateSyncTaskRequest) (*npool.CreateSyncTaskResponse, error) {
	var err error
	_info := in.GetInfo()
	if *_info.Start >= *_info.End {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "start >= end")
	}

	if _info.ChainType == nil {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "chaintype not set")
	}

	if _info.SyncState != nil &&
		*_info.SyncState != cttype.SyncState_Start &&
		*_info.SyncState != cttype.SyncState_Default &&
		*_info.SyncState != cttype.SyncState_Failed {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "state must be in [start default failed]")
	}

	if _info.SyncState == nil {
		syncS := cttype.SyncState_Default
		_info.SyncState = &syncS
	}

	_info.Current = _info.Start
	id := uuid.New().String()
	_info.ID = &id
	topic := fmt.Sprintf("%v_%v_%v_%v_%s",
		_info.ChainType.String(),
		*_info.ChainID,
		*_info.Start,
		*_info.End,
		*_info.ID)
	_info.Topic = &topic

	if err := ctkafka.CreateTopic(ctx, *_info.Topic); err != nil {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := crud.Create(ctx, _info)
	if err != nil {
		logger.Sugar().Errorw("CreateSyncTask", "error", err)
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSyncTaskResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) TriggerSyncTask(ctx context.Context, in *npool.TriggerSyncTaskRequest) (*npool.GetSyncTaskResponse, error) {
	conds := npool.Conds{
		Topic: &cybertracer.StringVal{
			Value: in.Topic,
			Op:    "EQ",
		},
	}
	info, err := crud.RowOnly(ctx, &conds)
	if err != nil {
		return &npool.GetSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	
	if info.End<=info.Current{
		info.SyncState=
	}

	return &npool.GetSyncTaskResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) UpdateSyncTask(ctx context.Context, in *npool.UpdateSyncTaskRequest) (*npool.UpdateSyncTaskResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateSyncTask", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateSyncTask", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSyncTaskResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSyncTask(ctx context.Context, in *npool.GetSyncTaskRequest) (*npool.GetSyncTaskResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetSyncTask", "ID", in.GetID(), "error", err)
		return &npool.GetSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetSyncTask", "ID", in.GetID(), "error", err)
		return &npool.GetSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSyncTaskResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSyncTaskOnly(ctx context.Context, in *npool.GetSyncTaskOnlyRequest) (*npool.GetSyncTaskOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetSyncTaskOnly", "error", err)
		return &npool.GetSyncTaskOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSyncTaskOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSyncTasks(ctx context.Context, in *npool.GetSyncTasksRequest) (*npool.GetSyncTasksResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetSyncTasks", "error", err)
		return &npool.GetSyncTasksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSyncTasksResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistSyncTask(ctx context.Context, in *npool.ExistSyncTaskRequest) (*npool.ExistSyncTaskResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistSyncTask", "ID", in.GetID(), "error", err)
		return &npool.ExistSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistSyncTask", "ID", in.GetID(), "error", err)
		return &npool.ExistSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSyncTaskResponse{
		Info: exist,
	}, nil
}

//nolint:lll
func (s *Server) ExistSyncTaskConds(ctx context.Context, in *npool.ExistSyncTaskCondsRequest) (*npool.ExistSyncTaskCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistSyncTaskConds", "error", err)
		return &npool.ExistSyncTaskCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSyncTaskCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountSyncTasks(ctx context.Context, in *npool.CountSyncTasksRequest) (*npool.CountSyncTasksResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountSyncTasksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountSyncTasksResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteSyncTask(ctx context.Context, in *npool.DeleteSyncTaskRequest) (*npool.DeleteSyncTaskResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteSyncTask", "ID", in.GetID(), "error", err)
		return &npool.DeleteSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteSyncTask", "ID", in.GetID(), "error", err)
		return &npool.DeleteSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSyncTaskResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
