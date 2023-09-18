package token

import (
	"context"

	entrancepool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/token"
	nftmetapool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/token"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	entrancepool.UnimplementedManagerServer
}

func (s *Server) GetToken(ctx context.Context, in *nftmetapool.GetTokenRequest) (*nftmetapool.GetTokenResponse, error) {
	client.UseCloudProxyCC()
	return client.GetToken(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancepool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancepool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
