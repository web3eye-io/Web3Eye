package endpoint

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/endpoint"

	entrancernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/endpoint"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
	"google.golang.org/grpc"
)

type Server struct {
	entrancernpool.UnimplementedManagerServer
}

func (s *Server) CreateEndpoint(ctx context.Context, in *nftmetanpool.CreateEndpointRequest) (*nftmetanpool.CreateEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.CreateEndpoint(ctx, in)
}

func (s *Server) CreateEndpoints(ctx context.Context, in *nftmetanpool.CreateEndpointsRequest) (*nftmetanpool.CreateEndpointsResponse, error) {
	client.UseCloudProxyCC()
	return client.CreateEndpoints(ctx, in)
}

func (s *Server) UpdateEndpoint(ctx context.Context, in *nftmetanpool.UpdateEndpointRequest) (*nftmetanpool.UpdateEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.UpdateEndpoint(ctx, in)
}

func (s *Server) UpdateEndpoints(ctx context.Context, in *nftmetanpool.UpdateEndpointsRequest) (*nftmetanpool.UpdateEndpointsResponse, error) {
	client.UseCloudProxyCC()
	return client.UpdateEndpoints(ctx, in)
}

func (s *Server) GetEndpoint(ctx context.Context, in *nftmetanpool.GetEndpointRequest) (*nftmetanpool.GetEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.GetEndpoint(ctx, in)
}

func (s *Server) GetEndpointOnly(ctx context.Context, in *nftmetanpool.GetEndpointOnlyRequest) (*nftmetanpool.GetEndpointOnlyResponse, error) {
	client.UseCloudProxyCC()
	return client.GetEndpointOnly(ctx, in)
}

func (s *Server) GetEndpoints(ctx context.Context, in *nftmetanpool.GetEndpointsRequest) (*nftmetanpool.GetEndpointsResponse, error) {
	client.UseCloudProxyCC()
	return client.GetEndpoints(ctx, in)
}

func (s *Server) ExistEndpoint(ctx context.Context, in *nftmetanpool.ExistEndpointRequest) (*nftmetanpool.ExistEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.ExistEndpoint(ctx, in)
}

func (s *Server) ExistEndpointConds(ctx context.Context, in *nftmetanpool.ExistEndpointCondsRequest) (*nftmetanpool.ExistEndpointCondsResponse, error) {
	client.UseCloudProxyCC()
	return client.ExistEndpointConds(ctx, in)
}

func (s *Server) CountEndpoints(ctx context.Context, in *nftmetanpool.CountEndpointsRequest) (*nftmetanpool.CountEndpointsResponse, error) {
	client.UseCloudProxyCC()
	return client.CountEndpoints(ctx, in)
}

func (s *Server) DeleteEndpoint(ctx context.Context, in *nftmetanpool.DeleteEndpointRequest) (*nftmetanpool.DeleteEndpointResponse, error) {
	client.UseCloudProxyCC()
	return client.DeleteEndpoint(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancernpool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancernpool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
