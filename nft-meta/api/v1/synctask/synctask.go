//nolint:nolintlint,dupl
package synctask

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	"github.com/web3eye-io/Web3Eye/common/utils"
	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/synctask"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	blockhandler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/block"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	blockproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
)

const (
	MaxPutTaskNumOnce         = 100
	MaxPutBadBlockTaskNumOnce = 50
	ReportInterval            = 100
	RedisLockTimeout          = time.Second * 10
)

func (s *Server) CreateSyncTask(ctx context.Context, in *npool.CreateSyncTaskRequest) (*npool.CreateSyncTaskResponse, error) {
	var err error
	_info := in.GetInfo()
	if _info == nil {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "info is nil")
	}

	if _info.Start == nil || _info.End == nil || (*_info.Start >= *_info.End && *_info.End != 0) {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "start >= end or set invalid")
	}

	if _info.ChainType == nil {
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, "chaintype not set")
	}

	if _info.SyncState == nil {
		syncS := basetype.SyncState_Default
		_info.SyncState = &syncS
	}

	_info.Current = _info.Start
	entID := uuid.New().String()
	_info.EntID = &entID
	topic := fmt.Sprintf("%v_%v_%v_%v_%s",
		_info.ChainType.String(),
		*_info.ChainID,
		*_info.Start,
		*_info.End,
		*_info.EntID)
	_info.Topic = &topic

	h, err := handler.NewHandler(ctx,
		handler.WithEntID(_info.EntID, true),
		handler.WithChainType(_info.ChainType, true),
		handler.WithChainID(_info.ChainID, true),
		handler.WithStart(_info.Start, true),
		handler.WithEnd(_info.End, true),
		handler.WithCurrent(_info.Current, true),
		handler.WithTopic(_info.Topic, true),
		handler.WithSyncState(_info.SyncState, true),
		handler.WithDescription(_info.Description, false),
		handler.WithRemark(_info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateSyncTask", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.CreateSyncTask(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateSyncTask", "error", err)
		return &npool.CreateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("CreateSyncTask success ,chaintype:%v chainid:%v start:%v end:%v", info.ChainType, info.ChainID, info.Start, info.End)

	return &npool.CreateSyncTaskResponse{
		Info: info,
	}, nil
}

//nolint:funlen,gocyclo
func (s *Server) TriggerSyncTask(ctx context.Context, in *npool.TriggerSyncTaskRequest) (*npool.TriggerSyncTaskResponse, error) {
	// query synctask
	// lock
	lockKey := "TriggerSyncTask_Lock"
	lockID, err := ctredis.TryLock(lockKey, RedisLockTimeout)
	if err != nil {
		logger.Sugar().Warn("TriggerSyncTask", "warning", err)
		return &npool.TriggerSyncTaskResponse{}, nil
	}

	defer func() {
		err := ctredis.Unlock(lockKey, lockID)
		if err != nil {
			logger.Sugar().Warn("TriggerSyncTask", "warning", err)
		}
	}()

	conds := &npool.Conds{
		Topic: &web3eye.StringVal{
			Value: in.Topic,
			Op:    cruder.EQ,
		},
	}

	resp, err := s.GetSyncTaskOnly(ctx, &npool.GetSyncTaskOnlyRequest{Conds: conds})
	if err != nil {
		return &npool.TriggerSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info := resp.Info

	// check state
	if info.SyncState != basetype.SyncState_Start || in.CurrentBlockNum < info.Current {
		return &npool.TriggerSyncTaskResponse{
			Info: info,
		}, nil
	}

	// check sync state
	if info.End != 0 && info.Current >= info.End {
		info.SyncState = basetype.SyncState_Finish
		_, err := s.UpdateSyncTask(ctx, &npool.UpdateSyncTaskRequest{
			Info: &npool.SyncTaskReq{
				ID:        &info.ID,
				SyncState: &info.SyncState,
			},
		})

		if err != nil {
			logger.Sugar().Errorw("TriggerSyncTask", "error", err)
			return &npool.TriggerSyncTaskResponse{Info: info}, status.Error(codes.Internal, err.Error())
		}

		return &npool.TriggerSyncTaskResponse{Info: info}, nil
	}

	syncEnd := info.End
	if syncEnd == 0 || syncEnd > in.CurrentBlockNum {
		syncEnd = in.CurrentBlockNum
	}

	lastNum := info.Current
	info.Current += MaxPutTaskNumOnce
	if info.Current > syncEnd {
		info.Current = syncEnd
	}

	blocks, err := getBadBlocksNum(ctx, info.ChainType, info.ChainID)
	if err != nil {
		logger.Sugar().Errorw("TriggerSyncTask", "error", err)
		return &npool.TriggerSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = sendTasks(ctx, info.Topic, lastNum, info.Current, blocks)
	if err != nil {
		logger.Sugar().Errorw("TriggerSyncTask", "error", err)
		return &npool.TriggerSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	_, err = s.UpdateSyncTask(ctx, &npool.UpdateSyncTaskRequest{
		Info: &npool.SyncTaskReq{
			ID:      &info.ID,
			Current: &info.Current,
		},
	})

	if err != nil {
		logger.Sugar().Errorw("TriggerSyncTask", "error", err)
		return &npool.TriggerSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.TriggerSyncTaskResponse{
		Info: info,
	}, nil
}

// send to topic
func sendTasks(ctx context.Context, topic string, start, end uint64, blocks []uint64) error {
	pulsarCli, err := ctpulsar.Client()
	if err != nil {
		return err
	}
	defer pulsarCli.Close()

	producer, err := pulsarCli.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return err
	}
	defer producer.Close()

	sendTask := func(num uint64) error {
		payload, err := utils.Uint642Bytes(num)
		if err != nil {
			return err
		}

		_, err = producer.Send(ctx, &pulsar.ProducerMessage{Payload: payload})
		if err != nil {
			return err
		}
		return nil
	}

	for ; start < end; start++ {
		err = sendTask(start)
		if err != nil {
			return err
		}
	}

	for _, num := range blocks {
		err = sendTask(num)
		if err != nil {
			return err
		}
	}

	return nil
}

func getBadBlocksNum(ctx context.Context,
	chainType basetype.ChainType,
	chainID string,
) ([]uint64, error) {
	failedBlocks, err := getBlocksNum(ctx, chainType, chainID, basetype.BlockParseState_BlockTypeFailed, MaxPutBadBlockTaskNumOnce)
	if err != nil {
		return nil, err
	}
	notFinishBlocks, err := getBlocksNum(ctx, chainType, chainID, basetype.BlockParseState_BlockTypeStart, MaxPutBadBlockTaskNumOnce)
	if err != nil {
		return nil, err
	}
	return append(failedBlocks, notFinishBlocks...), nil
}

func getBlocksNum(
	ctx context.Context,
	chainType basetype.ChainType,
	chainID string,
	parseState basetype.BlockParseState,
	limit int32) ([]uint64, error) {
	conds := blockproto.Conds{
		ChainType: &web3eye.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(*chainType.Enum()),
		},
		ChainID: &web3eye.StringVal{
			Op:    cruder.EQ,
			Value: chainID,
		},
		ParseState: &web3eye.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(*parseState.Enum()),
		},
	}
	h, err := blockhandler.NewHandler(
		ctx,
		blockhandler.WithConds(&conds),
		blockhandler.WithOffset(0),
		blockhandler.WithLimit(limit),
	)
	if err != nil {
		return nil, err
	}

	blocks, total, err := h.GetBlocks(ctx)
	if err != nil || total == 0 {
		return nil, err
	}
	taskNums := make([]uint64, len(blocks))
	for i, item := range blocks {
		taskNums[i] = item.BlockNumber
	}
	logger.Sugar().Infof("find %v bloks in %v", total, parseState)
	return taskNums, err
}

func (s *Server) UpdateSyncTask(ctx context.Context, in *npool.UpdateSyncTaskRequest) (*npool.UpdateSyncTaskResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpdateSyncTask",
			"In", in,
		)
		return &npool.UpdateSyncTaskResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithChainType(in.Info.ChainType, false),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithStart(in.Info.Start, false),
		handler.WithEnd(in.Info.End, false),
		handler.WithCurrent(in.Info.Current, false),
		handler.WithTopic(in.Info.Topic, false),
		handler.WithSyncState(in.Info.SyncState, false),
		handler.WithDescription(in.Info.Description, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateSyncTask", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpdateSyncTask(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateSyncTask", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infow("UpdateSyncTask", "ID", in.GetInfo().GetID())

	return &npool.UpdateSyncTaskResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSyncTask(ctx context.Context, in *npool.GetSyncTaskRequest) (*npool.GetSyncTaskResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetSyncTask", "error", err)
		return &npool.GetSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.GetSyncTask(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetSyncTask", "ID", in.GetID(), "error", err)
		return &npool.GetSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSyncTaskResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSyncTaskOnly(ctx context.Context, in *npool.GetSyncTaskOnlyRequest) (*npool.GetSyncTaskOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetSyncTaskOnly", "error", err)
		return &npool.GetSyncTaskOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetSyncTasks(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetSyncTaskOnly", "error", err)
		return &npool.GetSyncTaskOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total != 1 {
		errMsg := "more than one result or have no result"
		return &npool.GetSyncTaskOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &npool.GetSyncTaskOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetSyncTasks(ctx context.Context, in *npool.GetSyncTasksRequest) (*npool.GetSyncTasksResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetSyncTasks", "error", err)
		return &npool.GetSyncTasksResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetSyncTasks(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetSyncTasks", "error", err)
		return &npool.GetSyncTasksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSyncTasksResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) ExistSyncTask(ctx context.Context, in *npool.ExistSyncTaskRequest) (*npool.ExistSyncTaskResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistSyncTask", "ID", in.GetID(), "error", err)
		return &npool.ExistSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := h.ExistSyncTask(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistSyncTask", "ID", in.GetID(), "error", err)
		return &npool.ExistSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSyncTaskResponse{
		Exist: exist,
	}, nil
}

//nolint:lll
func (s *Server) ExistSyncTaskConds(ctx context.Context, in *npool.ExistSyncTaskCondsRequest) (*npool.ExistSyncTaskCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistSyncTaskConds", "error", err)
		return &npool.ExistSyncTaskCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := h.ExistSyncTaskConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistSyncTaskConds", "error", err)
		return &npool.ExistSyncTaskCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSyncTaskCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteSyncTask(ctx context.Context, in *npool.DeleteSyncTaskRequest) (*npool.DeleteSyncTaskResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("DeleteSyncTask", "error", err)
		return &npool.DeleteSyncTaskResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := h.DeleteSyncTask(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteSyncTask", "ID", in.GetID(), "error", err)
		return &npool.DeleteSyncTaskResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSyncTaskResponse{
		Info: info,
	}, nil
}
