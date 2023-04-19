package contract

import (
	"context"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/contract"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/contract"
	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
	contract.Server
}

func (s *Server) GetContract(ctx context.Context, in *nftmetanpool.GetContractRequest) (*nftmetanpool.GetContractResponse, error) {
	return s.Server.GetContract(ctx, in)
}

func (s *Server) GetContractOnly(
	ctx context.Context,
	in *nftmetanpool.GetContractOnlyRequest) (*nftmetanpool.GetContractOnlyResponse, error) {
	return s.Server.GetContractOnly(ctx, in)
}

func (s *Server) GetContracts(ctx context.Context, in *nftmetanpool.GetContractsRequest) (*nftmetanpool.GetContractsResponse, error) {
	return s.Server.GetContracts(ctx, in)
}

func (s *Server) CountContracts(ctx context.Context, in *nftmetanpool.CountContractsRequest) (*nftmetanpool.CountContractsResponse, error) {
	return s.Server.CountContracts(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}
