package retriever

import (
	"context"

	client "github.com/web3eye-io/Web3Eye/dealer/pkg/client/v1/retrieve"
	entrancepool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/retriever"
	retrieverpool "github.com/web3eye-io/Web3Eye/proto/web3eye/retriever/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	entrancepool.UnimplementedManagerServer
}

func (s *Server) StartRetrieve(ctx context.Context, in *retrieverpool.StartRetrieveRequest) (*retrieverpool.StartRetrieveResponse, error) {
	client.UseCloudProxyCC()
	return client.StartRetrieve(ctx, in)
}

func (s *Server) StatRetrieve(ctx context.Context, in *retrieverpool.StatRetrieveRequest) (*retrieverpool.StatRetrieveResponse, error) {
	client.UseCloudProxyCC()
	return client.StatRetrieve(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancepool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancepool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
