package v1

import (
	"context"

	_ "github.com/web3eye-io/Web3Eye/entrance/api/v1/search"
	"github.com/web3eye-io/Web3Eye/entrance/api/v1/snapshot"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
	snapshot.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
