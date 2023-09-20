package endpoint

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/endpoint"

	entrancernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/endpoint"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/endpoint"
	"google.golang.org/grpc"
)

type Server struct {
	entrancernpool.UnimplementedManagerServer
}

func (s *Server) CreateEndpoint(ctx context.Context, in *rankernpool.CreateEndpointRequest) (*rankernpool.CreateEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.CreateEndpoint(ctx, in)
}

func (s *Server) UpdateEndpoint(ctx context.Context, in *rankernpool.UpdateEndpointRequest) (*rankernpool.UpdateEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.UpdateEndpoint(ctx, in)
}

func (s *Server) GetEndpoint(ctx context.Context, in *rankernpool.GetEndpointRequest) (*rankernpool.GetEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.GetEndpoint(ctx, in)
}

func (s *Server) GetEndpoints(ctx context.Context, in *rankernpool.GetEndpointsRequest) (*rankernpool.GetEndpointsResponse, error) {
	client.UseCloudProxyCC()
	return client.GetEndpoints(ctx, in)
}

func (s *Server) DeleteEndpoint(ctx context.Context, in *rankernpool.DeleteEndpointRequest) (*rankernpool.DeleteEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.DeleteEndpoint(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancernpool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancernpool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
