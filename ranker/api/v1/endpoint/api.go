package endpoint

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/endpoint"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
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

func (s *Server) CreateEndpoint(ctx context.Context, in *rankernpool.CreateEndpointRequest) (*rankernpool.CreateEndpointResponse, error) {
	resp, err := s.Server.CreateEndpoint(ctx, &nftmetanpool.CreateEndpointRequest{
		Info: &nftmetanpool.EndpointReq{
			ChainType: &in.ChainType,
			ChainID:   &in.ChainID,
			Address:   &in.Address,
			RPS:       &in.RPS,
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoint", "error", err)
		return nil, err
	}
	return &rankernpool.CreateEndpointResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) UpdateEndpoint(ctx context.Context, in *rankernpool.UpdateEndpointRequest) (*rankernpool.UpdateEndpointResponse, error) {
	resp, err := s.Server.UpdateEndpoint(ctx, &nftmetanpool.UpdateEndpointRequest{
		Info: &nftmetanpool.EndpointReq{
			ID:      &in.ID,
			Address: in.Address,
			State:   in.State,
			RPS:     in.RPS,
			Remark:  in.Remark,
		},
	})
	if err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "error", err)
		return nil, err
	}
	return &rankernpool.UpdateEndpointResponse{Info: resp.Info}, nil
}

func (s *Server) GetEndpoint(ctx context.Context, in *rankernpool.GetEndpointRequest) (*rankernpool.GetEndpointResponse, error) {
	resp, err := s.Server.GetEndpoint(ctx, &nftmetanpool.GetEndpointRequest{
		ID: in.ID,
	})
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "error", err)
		return nil, err
	}
	return &rankernpool.GetEndpointResponse{Info: resp.Info}, nil
}

func (s *Server) GetEndpoints(ctx context.Context, in *rankernpool.GetEndpointsRequest) (*rankernpool.GetEndpointsResponse, error) {
	conds := buildConds(in)
	resp, err := s.Server.GetEndpoints(ctx,
		&nftmetanpool.GetEndpointsRequest{
			Conds:  conds,
			Offset: in.Offset,
			Limit:  in.Limit,
		})
	if err != nil {
		logger.Sugar().Errorw("GetEndpoints", "error", err)
		return nil, err
	}
	return &rankernpool.GetEndpointsResponse{Infos: resp.Infos, Total: resp.Total}, nil
}

func buildConds(in *rankernpool.GetEndpointsRequest) *nftmetanpool.Conds {
	conds := &nftmetanpool.Conds{}
	if in.ID != nil {
		conds.ID = &web3eye.Uint32Val{Op: cruder.EQ, Value: *in.ID}
	}

	if in.ChainType != nil {
		conds.ChainType = &web3eye.Uint32Val{Op: cruder.EQ, Value: uint32(*in.ChainType.Enum())}
	}

	if in.ChainID != nil {
		conds.ChainID = &web3eye.StringVal{Op: cruder.EQ, Value: *in.ChainID}
	}

	if in.Address != nil {
		conds.Address = &web3eye.StringVal{Op: cruder.EQ, Value: *in.Address}
	}

	if in.State != nil {
		conds.State = &web3eye.Uint32Val{Op: cruder.EQ, Value: uint32(*in.State.Enum())}
	}

	if in.RPS != nil {
		conds.RPS = &web3eye.Uint32Val{Op: cruder.EQ, Value: *in.RPS}
	}

	if in.Remark != nil {
		conds.Remark = &web3eye.StringVal{Op: cruder.EQ, Value: *in.Remark}
	}
	return conds
}

func (s *Server) DeleteEndpoint(ctx context.Context, in *rankernpool.DeleteEndpointRequest) (*rankernpool.DeleteEndpointResponse, error) {
	resp, err := s.Server.DeleteEndpoint(ctx, &nftmetanpool.DeleteEndpointRequest{
		ID: in.ID,
	})
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "error", err)
		return nil, err
	}
	return &rankernpool.DeleteEndpointResponse{Info: resp.Info}, nil
}
