package transfer

import (
	"context"

	entrancepool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/transfer"
	transferpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
	rankerpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/transfer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	entrancepool.UnimplementedManagerServer
}

func (s *Server) GetTransfer(ctx context.Context, in *transferpool.GetTransferRequest) (*transferpool.GetTransferResponse, error) {
	client.UseCloudProxyCC()
	return client.GetTransfer(ctx, in)
}

func (s *Server) GetTransferOnly(ctx context.Context, in *transferpool.GetTransferOnlyRequest) (*transferpool.GetTransferOnlyResponse, error) {
	client.UseCloudProxyCC()
	return client.GetTransferOnly(ctx, in)
}

func (s *Server) GetTransfers(ctx context.Context, in *rankerpool.GetTransfersRequest) (*transferpool.GetTransfersResponse, error) {
	client.UseCloudProxyCC()
	return client.GetTransfers(ctx, in)
}

func (s *Server) CountTransfers(ctx context.Context, in *rankerpool.CountTransfersRequest) (*transferpool.CountTransfersResponse, error) {
	client.UseCloudProxyCC()
	return client.CountTransfers(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancepool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancepool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
