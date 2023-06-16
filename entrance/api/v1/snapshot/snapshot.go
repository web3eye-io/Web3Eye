package snapshot

import (
	"context"

	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/snapshot"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/snapshot"
	entrancernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/snapshot"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
	"google.golang.org/grpc"
)

type Server struct {
	entrancernpool.UnimplementedManagerServer
	snapshot.Server
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

func (s *Server) CountSnapshots(ctx context.Context, in *nftmetanpool.CountSnapshotsRequest) (*nftmetanpool.CountSnapshotsResponse, error) {
	client.UseCloudProxyCC()
	return client.CountSnapshots(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancernpool.RegisterManagerServer(server, &Server{})
}
