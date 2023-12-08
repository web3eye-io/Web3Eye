//nolint:nolintlint,dupl
package endpoint

import (
	"context"

	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/endpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
)

func (s *Server) CreateEndpoint(ctx context.Context, in *npool.CreateEndpointRequest) (*npool.CreateEndpointResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"CreateEndpoint",
			"In", in,
		)
		return &npool.CreateEndpointResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	in.GetInfo().ChainID = nil
	in.GetInfo().State = basetype.EndpointState_EndpointDefault.Enum()

	h, err := handler.NewHandler(ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithAddress(in.Info.Address, true),
		handler.WithState(in.Info.State, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoint", "error", err)
		return &npool.CreateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.CreateEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoint", "error", err)
		return &npool.CreateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEndpointResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateEndpoints(ctx context.Context, in *npool.CreateEndpointsRequest) (*npool.CreateEndpointsResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateEndpoints", "error", "Infos is empty")
		return &npool.CreateEndpointsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoints", "error", err)
		return &npool.CreateEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, err := h.CreateEndpoints(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoints", "error", err)
		return &npool.CreateEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEndpointsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateEndpoint(ctx context.Context, in *npool.UpdateEndpointRequest) (*npool.UpdateEndpointResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpdateEndpoint",
			"In", in,
		)
		return &npool.UpdateEndpointResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithChainType(in.Info.ChainType, false),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithAddress(in.Info.Address, false),
		handler.WithState(in.Info.State, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpdateEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateEndpointResponse{
		Info: info,
	}, nil
}

func (s *Server) GetEndpoint(ctx context.Context, in *npool.GetEndpointRequest) (*npool.GetEndpointResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "ID", in.GetID(), "error", err)
		return &npool.GetEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.GetEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "ID", in.GetID(), "error", err)
		return &npool.GetEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEndpointResponse{
		Info: info,
	}, nil
}

func (s *Server) GetEndpointOnly(ctx context.Context, in *npool.GetEndpointOnlyRequest) (*npool.GetEndpointOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetEndpointOnly", "error", err)
		return &npool.GetEndpointOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetEndpoints(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetEndpointOnly", "error", err)
		return &npool.GetEndpointOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total != 1 {
		errMsg := "more than one result or have no result"
		logger.Sugar().Errorw("GetEndpointOnly", "error", errMsg)
		return &npool.GetEndpointOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &npool.GetEndpointOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetEndpoints(ctx context.Context, in *npool.GetEndpointsRequest) (*npool.GetEndpointsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoints", "error", err)
		return &npool.GetEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetEndpoints(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoints", "error", err)
		return &npool.GetEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEndpointsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) ExistEndpoint(ctx context.Context, in *npool.ExistEndpointRequest) (*npool.ExistEndpointResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpoint", "ID", in.GetID(), "error", err)
		return &npool.ExistEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := h.ExistEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpoint", "ID", in.GetID(), "error", err)
		return &npool.ExistEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEndpointResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistEndpointConds(ctx context.Context, in *npool.ExistEndpointCondsRequest) (*npool.ExistEndpointCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpointConds", "error", err)
		return &npool.ExistEndpointCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := h.ExistEndpointConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpointConds", "error", err)
		return &npool.ExistEndpointCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEndpointCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteEndpoint(ctx context.Context, in *npool.DeleteEndpointRequest) (*npool.DeleteEndpointResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "ID", in.GetID(), "error", err)
		return &npool.DeleteEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := h.DeleteEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "ID", in.GetID(), "error", err)
		return &npool.DeleteEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteEndpointResponse{
		Info: info,
	}, nil
}
