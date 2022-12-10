package contract

import (
	npool "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/contract"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
}
