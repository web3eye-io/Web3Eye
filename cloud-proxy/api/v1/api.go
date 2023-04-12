package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
