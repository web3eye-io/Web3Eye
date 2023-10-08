package order

import (
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
}
