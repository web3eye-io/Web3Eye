package snapshot

import (
	"context"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/snapshot"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/snapshot"
	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
	snapshot.Server
}

func (s *Server) GetSnapshot(ctx context.Context, in *nftmetanpool.GetSnapshotRequest) (*nftmetanpool.GetSnapshotResponse, error) {
	return s.Server.GetSnapshot(ctx, in)
}

func (s *Server) GetSnapshotOnly(ctx context.Context, in *nftmetanpool.GetSnapshotOnlyRequest) (*nftmetanpool.GetSnapshotOnlyResponse, error) {
	return s.Server.GetSnapshotOnly(ctx, in)
}

func (s *Server) GetSnapshots(ctx context.Context, in *nftmetanpool.GetSnapshotsRequest) (*nftmetanpool.GetSnapshotsResponse, error) {
	return s.Server.GetSnapshots(ctx, in)
}

func (s *Server) CountSnapshots(ctx context.Context, in *nftmetanpool.CountSnapshotsRequest) (*nftmetanpool.CountSnapshotsResponse, error) {
	return s.Server.CountSnapshots(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}
