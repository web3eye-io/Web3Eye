package v1

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	retrieve "github.com/web3eye-io/Web3Eye/dealer/api/v1/retrieve"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
	retrieve.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
