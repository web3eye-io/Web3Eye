package transfer

import (
	"context"

	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
	"github.com/web3eye-io/Web3Eye/ranker/pkg/crud/v1/transfer"
	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
}

func (s *Server) GetTransfers(ctx context.Context, in *rankernpool.GetTransfersRequest) (*rankernpool.GetTransfersResponse, error) {
	infos, total, err := transfer.Rows(ctx, in)
	if err != nil {
		return nil, err
	}
	return &rankernpool.GetTransfersResponse{Infos: infos, Total: uint32(total)}, nil
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}
