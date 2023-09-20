package synctask

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/synctask"

	entrancernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/synctask"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/synctask"
	"google.golang.org/grpc"
)

type Server struct {
	entrancernpool.UnimplementedManagerServer
}

func (s *Server) CreateSyncTask(ctx context.Context, in *rankernpool.CreateSyncTaskRequest) (*rankernpool.CreateSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.CreateSyncTask(ctx, in)
}

func (s *Server) UpdateSyncTask(ctx context.Context, in *rankernpool.UpdateSyncTaskRequest) (*rankernpool.UpdateSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.UpdateSyncTask(ctx, in)
}

func (s *Server) GetSyncTask(ctx context.Context, in *rankernpool.GetSyncTaskRequest) (*rankernpool.GetSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.GetSyncTask(ctx, in)
}

func (s *Server) GetSyncTasks(ctx context.Context, in *rankernpool.GetSyncTasksRequest) (*rankernpool.GetSyncTasksResponse, error) {
	client.UseCloudProxyCC()
	return client.GetSyncTasks(ctx, in)
}

func (s *Server) DeleteSyncTask(ctx context.Context, in *rankernpool.DeleteSyncTaskRequest) (*rankernpool.DeleteSyncTaskResponse, error) {
	client.UseCloudProxyCC()
	return client.DeleteSyncTask(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancernpool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, synctask string, opts []grpc.DialOption) error {
	return entrancernpool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, synctask, opts)
}
