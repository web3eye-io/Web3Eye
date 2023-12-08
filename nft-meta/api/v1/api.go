package v1

import (
	"context"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/block"
	// "github.com/web3eye-io/Web3Eye/nft-meta/api/v1/token"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
	// token.Register(server)
	block.Register(server)
	// transfer.Register(server)
	// contract.Register(server)
	// synctask.Register(server)
	// snapshot.Register(server)
	// endpoint.Register(server)
	// order.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, e string, opts []grpc.DialOption) error {
	if err := npool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, e, opts); err != nil {
		return err
	}
	// if err := token.RegisterGateway(mux, e, opts); err != nil {
	// 	return err
	// }
	return nil
}
