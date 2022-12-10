package transfer

import (
	npool "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/transfer"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
}
