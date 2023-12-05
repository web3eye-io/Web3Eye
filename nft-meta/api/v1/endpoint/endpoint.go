//nolint:nolintlint,dupl
package endpoint

import (
	"context"

	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/endpoint"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/endpoint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"

	"github.com/google/uuid"
)

func (s *Server) CreateEndpoint(ctx context.Context, in *npool.CreateEndpointRequest) (*npool.CreateEndpointResponse, error) {
	var err error

	in.GetInfo().ChainID = nil
	in.GetInfo().State = basetype.EndpointState_EndpointDefault.Enum()
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoint", "error", err)
		return &npool.CreateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateEndpoints(ctx context.Context, in *npool.CreateEndpointsRequest) (*npool.CreateEndpointsResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateEndpoints", "error", "Infos is empty")
		return &npool.CreateEndpointsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoints", "error", err)
		return &npool.CreateEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEndpointsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateEndpoint(ctx context.Context, in *npool.UpdateEndpointRequest) (*npool.UpdateEndpointResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) UpdateEndpoints(ctx context.Context, in *npool.UpdateEndpointsRequest) (*npool.UpdateEndpointsResponse, error) {
	failedInfos := []*npool.FailedInfo{}
	for _, v := range in.GetInfos() {
		if _, err := uuid.Parse(v.GetID()); err != nil {
			failedInfos = append(failedInfos, &npool.FailedInfo{
				ID:  *v.ID,
				MSG: err.Error(),
			})
			continue
		}

		_, err := crud.Update(ctx, v)
		if err != nil {
			failedInfos = append(failedInfos, &npool.FailedInfo{
				ID:  *v.ID,
				MSG: err.Error(),
			})
		}
	}

	return &npool.UpdateEndpointsResponse{
		Infos: failedInfos,
	}, nil
}

func (s *Server) GetEndpoint(ctx context.Context, in *npool.GetEndpointRequest) (*npool.GetEndpointResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "ID", in.GetID(), "error", err)
		return &npool.GetEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "ID", in.GetID(), "error", err)
		return &npool.GetEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEndpointOnly(ctx context.Context, in *npool.GetEndpointOnlyRequest) (*npool.GetEndpointOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetEndpointOnly", "error", err)
		return &npool.GetEndpointOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEndpointOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEndpoints(ctx context.Context, in *npool.GetEndpointsRequest) (*npool.GetEndpointsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetEndpoints", "error", err)
		return &npool.GetEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEndpointsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistEndpoint(ctx context.Context, in *npool.ExistEndpointRequest) (*npool.ExistEndpointResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistEndpoint", "ID", in.GetID(), "error", err)
		return &npool.ExistEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistEndpoint", "ID", in.GetID(), "error", err)
		return &npool.ExistEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEndpointResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistEndpointConds(ctx context.Context, in *npool.ExistEndpointCondsRequest) (*npool.ExistEndpointCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistEndpointConds", "error", err)
		return &npool.ExistEndpointCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEndpointCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) CountEndpoints(ctx context.Context, in *npool.CountEndpointsRequest) (*npool.CountEndpointsResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountEndpointsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteEndpoint(ctx context.Context, in *npool.DeleteEndpointRequest) (*npool.DeleteEndpointResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "ID", in.GetID(), "error", err)
		return &npool.DeleteEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "ID", in.GetID(), "error", err)
		return &npool.DeleteEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
