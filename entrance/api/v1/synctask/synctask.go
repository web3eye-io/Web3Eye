package synctask

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/synctask"

	entrancernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/synctask"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
	"google.golang.org/grpc"
)

type Server struct {
	entrancernpool.UnimplementedManagerServer
}

func (s *Server) CreateSyncTask(ctx context.Context, in *nftmetanpool.CreateSyncTaskRequest) (*nftmetanpool.CreateSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.CreateSyncTask(ctx, in)
}

func (s *Server) UpdateSyncTask(ctx context.Context, in *nftmetanpool.UpdateSyncTaskRequest) (*nftmetanpool.UpdateSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.UpdateSyncTask(ctx, in)
}

func (s *Server) GetSyncTask(ctx context.Context, in *nftmetanpool.GetSyncTaskRequest) (*nftmetanpool.GetSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.GetSyncTask(ctx, in)
}

func (s *Server) GetSyncTaskOnly(ctx context.Context, in *nftmetanpool.GetSyncTaskOnlyRequest) (*nftmetanpool.GetSyncTaskOnlyResponse, error) {
	client.UseCloudProxyCC()
	return client.GetSyncTaskOnly(ctx, in)
}

func (s *Server) GetSyncTasks(ctx context.Context, in *nftmetanpool.GetSyncTasksRequest) (*nftmetanpool.GetSyncTasksResponse, error) {
	client.UseCloudProxyCC()
	return client.GetSyncTasks(ctx, in)
}

func (s *Server) ExistSyncTask(ctx context.Context, in *nftmetanpool.ExistSyncTaskRequest) (*nftmetanpool.ExistSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.ExistSyncTask(ctx, in)
}

func (s *Server) ExistSyncTaskConds(ctx context.Context, in *nftmetanpool.ExistSyncTaskCondsRequest) (*nftmetanpool.ExistSyncTaskCondsResponse, error) {
	client.UseCloudProxyCC()
	return client.ExistSyncTaskConds(ctx, in)
}

func (s *Server) CountSyncTasks(ctx context.Context, in *nftmetanpool.CountSyncTasksRequest) (*nftmetanpool.CountSyncTasksResponse, error) {
	client.UseCloudProxyCC()
	return client.CountSyncTasks(ctx, in)
}

func (s *Server) DeleteSyncTask(ctx context.Context, in *nftmetanpool.DeleteSyncTaskRequest) (*nftmetanpool.DeleteSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.DeleteSyncTask(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancernpool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, synctask string, opts []grpc.DialOption) error {
	return entrancernpool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, synctask, opts)
}
