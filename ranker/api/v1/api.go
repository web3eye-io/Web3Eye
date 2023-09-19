package v1

import (
	"context"

	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1"
	"github.com/web3eye-io/Web3Eye/ranker/api/v1/contract"
	"github.com/web3eye-io/Web3Eye/ranker/api/v1/endpoint"
	"github.com/web3eye-io/Web3Eye/ranker/api/v1/snapshot"
	"github.com/web3eye-io/Web3Eye/ranker/api/v1/synctask"
	"github.com/web3eye-io/Web3Eye/ranker/api/v1/token"
	"github.com/web3eye-io/Web3Eye/ranker/api/v1/transfer"

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
	snapshot.Register(server)
	endpoint.Register(server)
	synctask.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, end string, opts []grpc.DialOption) error {
	if err := npool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, end, opts); err != nil {
		return err
	}
	return nil
}
