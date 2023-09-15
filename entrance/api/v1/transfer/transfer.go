package transfer

import (
	"context"

	entrancepool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/transfer"
	rankerpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/transfer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	entrancepool.UnimplementedManagerServer
}

func (s *Server) GetTransfers(ctx context.Context, in *rankerpool.GetTransfersRequest) (*rankerpool.GetTransfersResponse, error) {
	client.UseCloudProxyCC()
	return client.GetTransfers(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancepool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancepool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
