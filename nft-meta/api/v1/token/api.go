package token

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
