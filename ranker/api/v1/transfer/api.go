package transfer

import (
	"context"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/transfer"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
	transfer.Server
}

func (s *Server) GetTransfers(ctx context.Context, in *rankernpool.GetTransfersRequest) (*nftmetanpool.GetTransfersResponse, error) {
	_conds := &nftmetanpool.Conds{
		ChainType: &web3eye.StringVal{
			Op:    "eq",
			Value: in.ChainType.String(),
		},
		ChainID: &web3eye.StringVal{
			Op:    "eq",
			Value: in.GetChainID(),
		},
		Contract: &web3eye.StringVal{
			Op:    "eq",
			Value: in.GetContract(),
		},
		TokenID: &web3eye.StringVal{
			Op:    "eq",
			Value: in.GetTokenID(),
		},
	}
	return s.Server.GetTransfers(ctx, &nftmetanpool.GetTransfersRequest{Conds: _conds, Offset: in.Offset, Limit: in.Limit})
}

func (s *Server) CountTransfers(ctx context.Context, in *rankernpool.CountTransfersRequest) (*nftmetanpool.CountTransfersResponse, error) {
	_conds := &nftmetanpool.Conds{
		ChainType: &web3eye.StringVal{
			Op:    "eq",
			Value: in.ChainType.String(),
		},
		ChainID: &web3eye.StringVal{
			Op:    "eq",
			Value: in.GetChainID(),
		},
		Contract: &web3eye.StringVal{
			Op:    "eq",
			Value: in.GetContract(),
		},
		TokenID: &web3eye.StringVal{
			Op:    "eq",
			Value: in.GetTokenID(),
		},
	}
	return s.Server.CountTransfers(ctx, &nftmetanpool.CountTransfersRequest{Conds: _conds})
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}
