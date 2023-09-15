package contract

import (
	"context"

	entrancepool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/contract"
	rankerpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/contract"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/contract"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	entrancepool.UnimplementedManagerServer
}

func (s *Server) GetContractAndTokens(ctx context.Context, in *rankerpool.GetContractAndTokensReq) (*rankerpool.GetContractAndTokensResp, error) {
	client.UseCloudProxyCC()
	return client.GetContractAndTokens(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancepool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancepool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
