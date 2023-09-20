package synctask

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/synctask"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/synctask"

	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
	synctask.Server
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}

func (s *Server) CreateSyncTask(ctx context.Context, in *rankernpool.CreateSyncTaskRequest) (*rankernpool.CreateSyncTaskResponse, error) {
	resp, err := s.Server.CreateSyncTask(ctx, &nftmetanpool.CreateSyncTaskRequest{
		Info: &nftmetanpool.SyncTaskReq{
			ChainType:   &in.ChainType,
			ChainID:     &in.ChainID,
			Start:       &in.Start,
			End:         &in.End,
			Current:     &in.Current,
			Description: in.Description,
			SyncState:   &in.SyncState,
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateSyncTask", "error", err)
		return nil, err
	}
	return &rankernpool.CreateSyncTaskResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateSyncTask(ctx context.Context, in *rankernpool.UpdateSyncTaskRequest) (*rankernpool.UpdateSyncTaskResponse, error) {
	resp, err := s.Server.UpdateSyncTask(ctx, &nftmetanpool.UpdateSyncTaskRequest{
		Info: &nftmetanpool.SyncTaskReq{
			ID:          &in.ID,
			ChainType:   in.ChainType,
			ChainID:     in.ChainID,
			Start:       in.Start,
			End:         in.End,
			Current:     in.Current,
			Topic:       in.Topic,
			Description: in.Description,
			SyncState:   in.SyncState,
		},
	})
	if err != nil {
		logger.Sugar().Errorw("UpdateSyncTask", "error", err)
		return nil, err
	}
	return &rankernpool.UpdateSyncTaskResponse{Info: resp.Info}, nil
}

func (s *Server) GetSyncTask(ctx context.Context, in *rankernpool.GetSyncTaskRequest) (*rankernpool.GetSyncTaskResponse, error) {
	resp, err := s.Server.GetSyncTask(ctx, &nftmetanpool.GetSyncTaskRequest{
		ID: in.ID,
	})
	if err != nil {
		logger.Sugar().Errorw("GetSyncTask", "error", err)
		return nil, err
	}
	return &rankernpool.GetSyncTaskResponse{Info: resp.Info}, nil
}

func (s *Server) GetSyncTasks(ctx context.Context, in *rankernpool.GetSyncTasksRequest) (*rankernpool.GetSyncTasksResponse, error) {
	conds := buildConds(in)
	resp, err := s.Server.GetSyncTasks(ctx,
		&nftmetanpool.GetSyncTasksRequest{
			Conds:  conds,
			Offset: in.Offset,
			Limit:  in.Limit,
		})
	if err != nil {
		logger.Sugar().Errorw("GetSyncTasks", "error", err)
		return nil, err
	}
	return &rankernpool.GetSyncTasksResponse{Infos: resp.Infos}, nil
}

func buildConds(in *rankernpool.GetSyncTasksRequest) *nftmetanpool.Conds {
	conds := &nftmetanpool.Conds{}
	if in.ID != nil {
		conds.ID = &web3eye.StringVal{Op: "eq", Value: *in.ID}
	}

	if in.ChainType != nil {
		conds.ChainType = &web3eye.StringVal{Op: "eq", Value: in.ChainType.String()}
	}

	if in.ChainID != nil {
		conds.ChainID = &web3eye.StringVal{Op: "eq", Value: *in.ChainID}
	}

	if in.Start != nil {
		conds.Start = &web3eye.Uint64Val{Op: "eq", Value: *in.Start}
	}

	if in.End != nil {
		conds.End = &web3eye.Uint64Val{Op: "eq", Value: *in.End}
	}

	if in.Current != nil {
		conds.Current = &web3eye.Uint64Val{Op: "eq", Value: *in.Current}
	}

	if in.Topic != nil {
		conds.Topic = &web3eye.StringVal{Op: "eq", Value: *in.Topic}
	}

	if in.Description != nil {
		conds.Description = &web3eye.StringVal{Op: "eq", Value: *in.Description}
	}

	if in.SyncState != nil {
		conds.SyncState = &web3eye.StringVal{Op: "eq", Value: in.SyncState.String()}
	}

	if in.Remark != nil {
		conds.Remark = &web3eye.StringVal{Op: "eq", Value: *in.Remark}
	}

	return conds
}

func (s *Server) DeleteSyncTask(ctx context.Context, in *rankernpool.DeleteSyncTaskRequest) (*rankernpool.DeleteSyncTaskResponse, error) {
	resp, err := s.Server.DeleteSyncTask(ctx, &nftmetanpool.DeleteSyncTaskRequest{
		ID: in.ID,
	})
	if err != nil {
		logger.Sugar().Errorw("DeleteSyncTask", "error", err)
		return nil, err
	}
	return &rankernpool.DeleteSyncTaskResponse{Info: resp.Info}, nil
}
