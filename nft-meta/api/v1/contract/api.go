package contract

import (
	npool "github.com/NpoolPlatform/message/npool/nftmeta/v1/contract"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
}
