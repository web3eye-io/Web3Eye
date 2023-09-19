package endpoint

import (
	"context"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/endpoint"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/endpoint"
	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
	endpoint.Server
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}

func (s *Server) CreateEndpoint(ctx context.Context, in *nftmetanpool.CreateEndpointRequest) (*nftmetanpool.CreateEndpointResponse, error) {
	return s.Server.CreateEndpoint(ctx, in)
}

func (s *Server) CreateEndpoints(ctx context.Context, in *nftmetanpool.CreateEndpointsRequest) (*nftmetanpool.CreateEndpointsResponse, error) {
	return s.Server.CreateEndpoints(ctx, in)
}

func (s *Server) UpdateEndpoint(ctx context.Context, in *nftmetanpool.UpdateEndpointRequest) (*nftmetanpool.UpdateEndpointResponse, error) {
	return s.Server.UpdateEndpoint(ctx, in)
}

func (s *Server) UpdateEndpoints(ctx context.Context, in *nftmetanpool.UpdateEndpointsRequest) (*nftmetanpool.UpdateEndpointsResponse, error) {
	return s.Server.UpdateEndpoints(ctx, in)
}

func (s *Server) GetEndpoint(ctx context.Context, in *nftmetanpool.GetEndpointRequest) (*nftmetanpool.GetEndpointResponse, error) {
	return s.Server.GetEndpoint(ctx, in)
}

func (s *Server) GetEndpointOnly(ctx context.Context, in *nftmetanpool.GetEndpointOnlyRequest) (*nftmetanpool.GetEndpointOnlyResponse, error) {
	return s.Server.GetEndpointOnly(ctx, in)
}

func (s *Server) GetEndpoints(ctx context.Context, in *nftmetanpool.GetEndpointsRequest) (*nftmetanpool.GetEndpointsResponse, error) {
	return s.Server.GetEndpoints(ctx, in)
}

func (s *Server) ExistEndpoint(ctx context.Context, in *nftmetanpool.ExistEndpointRequest) (*nftmetanpool.ExistEndpointResponse, error) {
	return s.Server.ExistEndpoint(ctx, in)
}

func (s *Server) ExistEndpointConds(ctx context.Context, in *nftmetanpool.ExistEndpointCondsRequest) (*nftmetanpool.ExistEndpointCondsResponse, error) {
	return s.Server.ExistEndpointConds(ctx, in)
}

func (s *Server) CountEndpoints(ctx context.Context, in *nftmetanpool.CountEndpointsRequest) (*nftmetanpool.CountEndpointsResponse, error) {
	return s.Server.CountEndpoints(ctx, in)
}

func (s *Server) DeleteEndpoint(ctx context.Context, in *nftmetanpool.DeleteEndpointRequest) (*nftmetanpool.DeleteEndpointResponse, error) {
	return s.Server.DeleteEndpoint(ctx, in)
}
