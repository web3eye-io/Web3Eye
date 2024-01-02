package snapshot

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	dealerclient "github.com/web3eye-io/Web3Eye/dealer/pkg/client/v1"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/snapshot"

	dealernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
	entrancernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/snapshot"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
	"google.golang.org/grpc"
)

type Server struct {
	entrancernpool.UnimplementedManagerServer
}

func (s *Server) GetSnapshot(ctx context.Context, in *nftmetanpool.GetSnapshotRequest) (*nftmetanpool.GetSnapshotResponse, error) {
	client.UseCloudProxyCC()
	return client.GetSnapshot(ctx, in)
}

func (s *Server) GetSnapshotOnly(ctx context.Context, in *nftmetanpool.GetSnapshotOnlyRequest) (*nftmetanpool.GetSnapshotOnlyResponse, error) {
	client.UseCloudProxyCC()
	return client.GetSnapshotOnly(ctx, in)
}

func (s *Server) GetSnapshots(ctx context.Context, in *nftmetanpool.GetSnapshotsRequest) (*nftmetanpool.GetSnapshotsResponse, error) {
	client.UseCloudProxyCC()
	return client.GetSnapshots(ctx, in)
}

func (s *Server) CreateBackup(ctx context.Context, in *dealernpool.CreateBackupRequest) (*dealernpool.CreateBackupResponse, error) {
	dealerclient.UseCloudProxyCC()
	return dealerclient.CreateBackup(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancernpool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancernpool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
