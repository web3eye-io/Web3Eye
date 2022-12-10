package api

import (
	"context"

	npool "github.com/web3eye-io/cyber-tracer/message/cybertracer/nftmeta/v1"
	"github.com/web3eye-io/cyber-tracer/nft-meta/api/v1/contract"
	"github.com/web3eye-io/cyber-tracer/nft-meta/api/v1/synctask"
	"github.com/web3eye-io/cyber-tracer/nft-meta/api/v1/token"
	"github.com/web3eye-io/cyber-tracer/nft-meta/api/v1/transfer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
	token.Register(server)
	transfer.Register(server)
	contract.Register(server)
	synctask.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
