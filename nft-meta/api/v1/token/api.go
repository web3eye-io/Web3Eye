package token

import (
	npool "github.com/web3eye-io/cyber-tracer/message/cybertracer/nftmeta/v1/token"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
}
