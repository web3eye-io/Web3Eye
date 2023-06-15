//nolint:nolintlint,dupl
package snapshot

import (
	"context"

	converter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/snapshot"
	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/snapshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"

	"github.com/google/uuid"
)

func (s *Server) CreateSnapshot(ctx context.Context, in *npool.CreateSnapshotRequest) (*npool.CreateSnapshotResponse, error) {
	var err error

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateSnapshot", "error", err)
		return &npool.CreateSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSnapshotResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateSnapshots(ctx context.Context, in *npool.CreateSnapshotsRequest) (*npool.CreateSnapshotsResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateSnapshots", "error", "Infos is empty")
		return &npool.CreateSnapshotsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateSnapshots", "error", err)
		return &npool.CreateSnapshotsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSnapshotsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateSnapshot(ctx context.Context, in *npool.UpdateSnapshotRequest) (*npool.UpdateSnapshotResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateSnapshot", "ID", in.GetID(), "error", err)
		return &npool.UpdateSnapshotResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("UpdateSnapshot", "ID", in.GetID(), "error", err)
		return &npool.UpdateSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSnapshotResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSnapshot(ctx context.Context, in *npool.GetSnapshotRequest) (*npool.GetSnapshotResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetSnapshot", "ID", in.GetID(), "error", err)
		return &npool.GetSnapshotResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetSnapshot", "ID", in.GetID(), "error", err)
		return &npool.GetSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSnapshotResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSnapshotOnly(ctx context.Context, in *npool.GetSnapshotOnlyRequest) (*npool.GetSnapshotOnlyResponse, error) {
	var err error

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetSnapshotOnly", "error", err)
		return &npool.GetSnapshotOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSnapshotOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSnapshots(ctx context.Context, in *npool.GetSnapshotsRequest) (*npool.GetSnapshotsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetSnapshots", "error", err)
		return &npool.GetSnapshotsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSnapshotsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistSnapshot(ctx context.Context, in *npool.ExistSnapshotRequest) (*npool.ExistSnapshotResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistSnapshot", "ID", in.GetID(), "error", err)
		return &npool.ExistSnapshotResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistSnapshot", "ID", in.GetID(), "error", err)
		return &npool.ExistSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSnapshotResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSnapshotConds(ctx context.Context, in *npool.ExistSnapshotCondsRequest) (*npool.ExistSnapshotCondsResponse, error) {
	var err error

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistSnapshotConds", "error", err)
		return &npool.ExistSnapshotCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSnapshotCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountSnapshots(ctx context.Context, in *npool.CountSnapshotsRequest) (*npool.CountSnapshotsResponse, error) {
	var err error

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("Counts", "error", err)
		return &npool.CountSnapshotsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountSnapshotsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteSnapshot(ctx context.Context, in *npool.DeleteSnapshotRequest) (*npool.DeleteSnapshotResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteSnapshot", "ID", in.GetID(), "error", err)
		return &npool.DeleteSnapshotResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteSnapshot", "ID", in.GetID(), "error", err)
		return &npool.DeleteSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSnapshotResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
