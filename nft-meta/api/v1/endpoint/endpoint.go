//nolint:nolintlint,dupl
package endpoint

import (
	"context"

	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/endpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	endpointproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
)

func (s *Server) CreateEndpoint(ctx context.Context, in *endpointproto.CreateEndpointRequest) (*endpointproto.CreateEndpointResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"CreateEndpoint",
			"In", in,
		)
		return &endpointproto.CreateEndpointResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	in.GetInfo().ChainID = nil
	in.GetInfo().State = basetype.EndpointState_EndpointDefault.Enum()

	h, err := handler.NewHandler(ctx,
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithAddress(in.Info.Address, true),
		handler.WithState(in.Info.State, false),
		handler.WithRPS(in.Info.RPS, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoint", "error", err)
		return &endpointproto.CreateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.CreateEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoint", "error", err)
		return &endpointproto.CreateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to CreateEndpoint ,chaintype:%v chainid:%v rps:%v address:%v", info.ChainType, info.ChainID, info.RPS, info.Address)
	return &endpointproto.CreateEndpointResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateEndpoints(ctx context.Context, in *endpointproto.CreateEndpointsRequest) (*endpointproto.CreateEndpointsResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateEndpoints", "error", "Infos is empty")
		return &endpointproto.CreateEndpointsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoints", "error", err)
		return &endpointproto.CreateEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, err := h.CreateEndpoints(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoints", "error", err)
		return &endpointproto.CreateEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to CreateEndpoints ,have %v infos", len(infos))
	return &endpointproto.CreateEndpointsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateEndpoint(ctx context.Context, in *endpointproto.UpdateEndpointRequest) (*endpointproto.UpdateEndpointResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpdateEndpoint",
			"In", in,
		)
		return &endpointproto.UpdateEndpointResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithChainType(in.Info.ChainType, false),
		handler.WithChainID(in.Info.ChainID, false),
		handler.WithAddress(in.Info.Address, false),
		handler.WithState(in.Info.State, false),
		handler.WithRPS(in.Info.RPS, false),
		handler.WithRemark(in.Info.Remark, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "ID", in.GetInfo().GetID(), "error", err)
		return &endpointproto.UpdateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpdateEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "ID", in.GetInfo().GetID(), "error", err)
		return &endpointproto.UpdateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to UpdateEndpoint ,chaintype:%v chainid:%v rps:%v address:%v", info.ChainType, info.ChainID, info.RPS, info.Address)
	return &endpointproto.UpdateEndpointResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateEndpoints(ctx context.Context, in *endpointproto.UpdateEndpointsRequest) (*endpointproto.UpdateEndpointsResponse, error) {
	if len(in.Infos) == 0 {
		logger.Sugar().Errorw(
			"UpdateEndpoints",
			"In", in,
			"Msg", "have no input",
		)
		return &endpointproto.UpdateEndpointsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	failedInfos := []*endpointproto.FailedInfo{}
	for _, info := range in.Infos {
		h, err := handler.NewHandler(ctx,
			handler.WithID(info.ID, true),
			handler.WithChainType(info.ChainType, false),
			handler.WithChainID(info.ChainID, false),
			handler.WithAddress(info.Address, false),
			handler.WithState(info.State, false),
			handler.WithRPS(info.RPS, false),
			handler.WithRemark(info.Remark, false),
		)
		if err != nil {
			failedInfos = append(failedInfos, &endpointproto.FailedInfo{
				ID:  *info.ID,
				MSG: err.Error(),
			})
			logger.Sugar().Errorw("UpdateEndpoints", "ID", info.GetID(), "error", err)
			continue
		}

		_, err = h.UpdateEndpoint(ctx)
		if err != nil {
			failedInfos = append(failedInfos, &endpointproto.FailedInfo{
				ID:  *info.ID,
				MSG: err.Error(),
			})
			logger.Sugar().Errorw("UpdateEndpoints", "ID", info.GetID(), "error", err)
		}
	}

	logger.Sugar().Infof("success to CreateEndpoints ,all have %v infos, failed to update %v infos", len(in.Infos), len(failedInfos))
	return &endpointproto.UpdateEndpointsResponse{
		Infos: failedInfos,
	}, nil
}

func (s *Server) GetEndpoint(ctx context.Context, in *endpointproto.GetEndpointRequest) (*endpointproto.GetEndpointResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "ID", in.GetID(), "error", err)
		return &endpointproto.GetEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.GetEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "ID", in.GetID(), "error", err)
		return &endpointproto.GetEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to GetEndpoint ,chaintype:%v chainid:%v rps:%v address:%v", info.ChainType, info.ChainID, info.RPS, info.Address)
	return &endpointproto.GetEndpointResponse{
		Info: info,
	}, nil
}

func (s *Server) GetEndpointOnly(ctx context.Context, in *endpointproto.GetEndpointOnlyRequest) (*endpointproto.GetEndpointOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetEndpointOnly", "error", err)
		return &endpointproto.GetEndpointOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetEndpoints(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetEndpointOnly", "error", err)
		return &endpointproto.GetEndpointOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total != 1 {
		errMsg := "more than one result or have no result"
		return &endpointproto.GetEndpointOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}
	info := infos[0]

	logger.Sugar().Infof("success to GetEndpointOnly ,chaintype:%v chainid:%v rps:%v address:%v", info.ChainType, info.ChainID, info.RPS, info.Address)
	return &endpointproto.GetEndpointOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetEndpoints(ctx context.Context, in *endpointproto.GetEndpointsRequest) (*endpointproto.GetEndpointsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoints", "error", err)
		return &endpointproto.GetEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}
	infos, total, err := h.GetEndpoints(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoints", "error", err)
		return &endpointproto.GetEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to GetEndpoints ,all have %v infos", len(infos))
	return &endpointproto.GetEndpointsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) ExistEndpoint(ctx context.Context, in *endpointproto.ExistEndpointRequest) (*endpointproto.ExistEndpointResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpoint", "ID", in.GetID(), "error", err)
		return &endpointproto.ExistEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := h.ExistEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpoint", "ID", in.GetID(), "error", err)
		return &endpointproto.ExistEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &endpointproto.ExistEndpointResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistEndpointConds(ctx context.Context, in *endpointproto.ExistEndpointCondsRequest) (*endpointproto.ExistEndpointCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpointConds", "error", err)
		return &endpointproto.ExistEndpointCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := h.ExistEndpointConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpointConds", "error", err)
		return &endpointproto.ExistEndpointCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &endpointproto.ExistEndpointCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteEndpoint(ctx context.Context, in *endpointproto.DeleteEndpointRequest) (*endpointproto.DeleteEndpointResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "ID", in.GetID(), "error", err)
		return &endpointproto.DeleteEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := h.DeleteEndpoint(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "ID", in.GetID(), "error", err)
		return &endpointproto.DeleteEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to DeleteEndpoint ,chaintype:%v chainid:%v rps:%v address:%v", info.ChainType, info.ChainID, info.RPS, info.Address)
	return &endpointproto.DeleteEndpointResponse{
		Info: info,
	}, nil
}
