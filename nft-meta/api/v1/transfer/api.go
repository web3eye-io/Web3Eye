package transfer

import (
	npool "github.com/NpoolPlatform/message/npool/nftmeta/v1/transfer"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
}
