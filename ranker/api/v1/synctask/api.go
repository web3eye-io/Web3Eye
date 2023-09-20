package synctask

import (
	"context"

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
		return nil, err
	}
	return &rankernpool.UpdateSyncTaskResponse{Info: resp.Info}, nil
}

func (s *Server) GetSyncTask(ctx context.Context, in *rankernpool.GetSyncTaskRequest) (*rankernpool.GetSyncTaskResponse, error) {
	resp, err := s.Server.GetSyncTask(ctx, &nftmetanpool.GetSyncTaskRequest{
		ID: in.ID,
	})
	if err != nil {
		return nil, err
	}
	return &rankernpool.GetSyncTaskResponse{Info: resp.Info}, nil
}

func (s *Server) GetSyncTasks(ctx context.Context, in *rankernpool.GetSyncTasksRequest) (*rankernpool.GetSyncTasksResponse, error) {
	conds := &nftmetanpool.Conds{
		ID:          &web3eye.StringVal{Op: "eq", Value: *in.ID},
		ChainType:   &web3eye.StringVal{Op: "eq", Value: in.ChainType.String()},
		ChainID:     &web3eye.StringVal{Op: "eq", Value: *in.ChainID},
		Start:       &web3eye.Uint64Val{Op: "eq", Value: *in.Start},
		End:         &web3eye.Uint64Val{Op: "eq", Value: *in.End},
		Current:     &web3eye.Uint64Val{Op: "eq", Value: *in.Current},
		Topic:       &web3eye.StringVal{Op: "eq", Value: *in.Topic},
		Description: &web3eye.StringVal{Op: "eq", Value: *in.Description},
		SyncState:   &web3eye.StringVal{Op: "eq", Value: in.SyncState.String()},
		Remark:      &web3eye.StringVal{Op: "eq", Value: *in.Remark},
	}
	resp, err := s.Server.GetSyncTasks(ctx,
		&nftmetanpool.GetSyncTasksRequest{
			Conds:  conds,
			Offset: in.Offset,
			Limit:  in.Limit,
		})
	if err != nil {
		return nil, err
	}
	return &rankernpool.GetSyncTasksResponse{Infos: resp.Infos}, nil
}

func (s *Server) DeleteSyncTask(ctx context.Context, in *rankernpool.DeleteSyncTaskRequest) (*rankernpool.DeleteSyncTaskResponse, error) {
	resp, err := s.Server.DeleteSyncTask(ctx, &nftmetanpool.DeleteSyncTaskRequest{
		ID: in.ID,
	})
	if err != nil {
		return nil, err
	}
	return &rankernpool.DeleteSyncTaskResponse{Info: resp.Info}, nil
}
