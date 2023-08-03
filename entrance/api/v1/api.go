package v1

import (
	"context"

	"github.com/web3eye-io/Web3Eye/entrance/api/v1/contract"
	"github.com/web3eye-io/Web3Eye/entrance/api/v1/retriever"
	"github.com/web3eye-io/Web3Eye/entrance/api/v1/snapshot"
	"github.com/web3eye-io/Web3Eye/entrance/api/v1/token"
	"github.com/web3eye-io/Web3Eye/entrance/api/v1/transfer"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterManagerServer(server, &Server{})
	snapshot.Register(server)
	retriever.Register(server)
	transfer.Register(server)
	token.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := snapshot.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := retriever.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := transfer.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := contract.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := token.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
