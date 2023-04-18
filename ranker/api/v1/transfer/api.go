package transfer

import (
	"context"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/transfer"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
	transfer.Server
}

func (s *Server) GetTransfer(ctx context.Context, in *nftmetanpool.GetTransferRequest) (*nftmetanpool.GetTransferResponse, error) {
	return s.Server.GetTransfer(ctx, in)
}

func (s *Server) GetTransferOnly(ctx context.Context, in *nftmetanpool.GetTransferOnlyRequest) (*nftmetanpool.GetTransferOnlyResponse, error) {
	return s.Server.GetTransferOnly(ctx, in)
}

func (s *Server) GetTransfers(ctx context.Context, in *nftmetanpool.GetTransfersRequest) (*nftmetanpool.GetTransfersResponse, error) {
	return s.Server.GetTransfers(ctx, in)
}

func (s *Server) CountTransfers(ctx context.Context, in *nftmetanpool.CountTransfersRequest) (*nftmetanpool.CountTransfersResponse, error) {
	return s.Server.CountTransfers(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}
