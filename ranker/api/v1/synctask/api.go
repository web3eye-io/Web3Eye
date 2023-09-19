package synctask

import (
	"context"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/synctask"
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

func (s *Server) CreateSyncTask(ctx context.Context, in *nftmetanpool.CreateSyncTaskRequest) (*nftmetanpool.CreateSyncTaskResponse, error) {
	return s.Server.CreateSyncTask(ctx, in)
}

func (s *Server) UpdateSyncTask(ctx context.Context, in *nftmetanpool.UpdateSyncTaskRequest) (*nftmetanpool.UpdateSyncTaskResponse, error) {
	return s.Server.UpdateSyncTask(ctx, in)
}

func (s *Server) GetSyncTask(ctx context.Context, in *nftmetanpool.GetSyncTaskRequest) (*nftmetanpool.GetSyncTaskResponse, error) {
	return s.Server.GetSyncTask(ctx, in)
}

func (s *Server) GetSyncTaskOnly(ctx context.Context, in *nftmetanpool.GetSyncTaskOnlyRequest) (*nftmetanpool.GetSyncTaskOnlyResponse, error) {
	return s.Server.GetSyncTaskOnly(ctx, in)
}

func (s *Server) GetSyncTasks(ctx context.Context, in *nftmetanpool.GetSyncTasksRequest) (*nftmetanpool.GetSyncTasksResponse, error) {
	return s.Server.GetSyncTasks(ctx, in)
}

func (s *Server) ExistSyncTask(ctx context.Context, in *nftmetanpool.ExistSyncTaskRequest) (*nftmetanpool.ExistSyncTaskResponse, error) {
	return s.Server.ExistSyncTask(ctx, in)
}

func (s *Server) ExistSyncTaskConds(ctx context.Context, in *nftmetanpool.ExistSyncTaskCondsRequest) (*nftmetanpool.ExistSyncTaskCondsResponse, error) {
	return s.Server.ExistSyncTaskConds(ctx, in)
}

func (s *Server) CountSyncTasks(ctx context.Context, in *nftmetanpool.CountSyncTasksRequest) (*nftmetanpool.CountSyncTasksResponse, error) {
	return s.Server.CountSyncTasks(ctx, in)
}

func (s *Server) DeleteSyncTask(ctx context.Context, in *nftmetanpool.DeleteSyncTaskRequest) (*nftmetanpool.DeleteSyncTaskResponse, error) {
	return s.Server.DeleteSyncTask(ctx, in)
}
