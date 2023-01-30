//nolint:nolintlint,dupl
package synctask

import (
	"context"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/common/ctkafka"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	"github.com/web3eye-io/Web3Eye/common/utils"
	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/synctask"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/synctask"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"

	"github.com/google/uuid"
)

const (
	MaxPutTaskNumOnce = 500
	ReportInterval    = 100
	RedisLockTimeout  = time.Second * 10
)

func (s *Server) CreateSyncTask(ctx context.Context, in *npool.CreateSyncTaskRequest) (*npool.CreateSyncTaskResponse, error) {
	var err error
	_info := in.GetInfo()
	if _info == nil {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "info is nil")
	}

	if _info.Start == nil || _info.End == nil || *_info.Start >= *_info.End {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "start >= end or set invalid")
	}

	if _info.ChainType == nil {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "chaintype not set")
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
		logger.Sugar().Errorw("CreateSyncTask", "error", err)
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := crud.Create(ctx, _info)
	if err != nil {
		logger.Sugar().Errorw("CreateSyncTask", "error", err)
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("CreateSyncTask success ,chaintype:%v chainid:%v start:%v end:%v", info.ChainType, info.ChainID, info.Start, info.End)

	return &npool.CreateSyncTaskResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) TriggerSyncTask(ctx context.Context, in *npool.TriggerSyncTaskRequest) (*npool.GetSyncTaskResponse, error) {
	conds := npool.Conds{
		Topic: &web3eye.StringVal{
			Value: in.Topic,
			Op:    "eq",
		},
	}
	info, err := crud.RowOnly(ctx, &conds)
	if err != nil {
		logger.Sugar().Errorw("TriggerSyncTask", "error", err)
		return &npool.GetSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	lockKey := "TriggerSyncTask_Lock"
	lockID, err := ctredis.TryLock(lockKey, RedisLockTimeout)
	if err != nil {
		logger.Sugar().Warn("TriggerSyncTask", "warning", err)
		return &npool.GetSyncTaskResponse{Info: converter.Ent2Grpc(info)}, nil
	}

	defer func() {
		err := ctredis.Unlock(lockKey, lockID)
		if err != nil {
			logger.Sugar().Warn("TriggerSyncTask", "warning", err)
		}
	}()

	if info.SyncState != cttype.SyncState_Start.String() {
		return &npool.GetSyncTaskResponse{
			Info: converter.Ent2Grpc(info),
		}, nil
	}

	syncEnd := info.End
	if info.End == 0 {
		syncEnd = eth.GetCurrentConfirmedHeight(ctx)
	}

	if info.Current >= syncEnd {
		info.SyncState = cttype.SyncState_Finsh.String()
	}
	info, err = crud.Update(ctx, converter.Ent2GrpcReq(info))
	if err != nil {
		logger.Sugar().Errorw("TriggerSyncTask", "error", err)
		return &npool.GetSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	lastNum := info.Current
	info.Current += MaxPutTaskNumOnce
	if info.Current > syncEnd {
		info.Current = syncEnd
	}

	p, err := ctkafka.NewCTProducer(in.Topic)
	if err != nil {
		logger.Sugar().Errorw("TriggerSyncTask", "error", err)
		return &npool.GetSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	for ; lastNum < info.Current; lastNum++ {
		b, err := utils.Uint642Bytes(lastNum)
		if err != nil {
			logger.Sugar().Error(err)
		}
		err = p.Produce(b, b)
		if err != nil {
			info.Current = lastNum
			logger.Sugar().Error(err)
			break
		}
	}

	info, err = crud.Update(ctx, converter.Ent2GrpcReq(info))

	if err != nil {
		logger.Sugar().Errorw("TriggerSyncTask", "error", err)
		return &npool.GetSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
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

	logger.Sugar().Infow("UpdateSyncTask", "ID", in.GetInfo().GetID())

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

	logger.Sugar().Infow("DeleteSyncTask", "ID", in.GetID())

	return &npool.DeleteSyncTaskResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
