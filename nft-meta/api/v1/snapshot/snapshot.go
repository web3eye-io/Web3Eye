//nolint:nolintlint,dupl
package snapshot

import (
	"context"

	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/snapshot"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

func (s *Server) CreateSnapshot(ctx context.Context, in *npool.CreateSnapshotRequest) (*npool.CreateSnapshotResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"CreateSnapshot",
			"In", in,
		)
		return &npool.CreateSnapshotResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithIndex(in.Info.Index, true),
		handler.WithSnapshotCommP(in.Info.SnapshotCommP, true),
		handler.WithSnapshotRoot(in.Info.SnapshotRoot, true),
		handler.WithSnapshotURI(in.Info.SnapshotURI, true),
		handler.WithBackupState(in.Info.BackupState, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateSnapshot", "error", err)
		return &npool.CreateSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}
	info, err := h.CreateSnapshot(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateSnapshot", "error", err)
		return &npool.CreateSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSnapshotResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateSnapshots(ctx context.Context, in *npool.CreateSnapshotsRequest) (*npool.CreateSnapshotsResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateSnapshots", "error", "Infos is empty")
		return &npool.CreateSnapshotsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	h, err := handler.NewHandler(ctx,
		handler.WithReqs(in.Infos, true),
	)
	if err != nil {
		logger.Sugar().Errorw("CreateSnapshots", "error", err)
		return &npool.CreateSnapshotsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := h.CreateSnapshots(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateSnapshots", "error", err)
		return &npool.CreateSnapshotsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSnapshotsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) UpdateSnapshot(ctx context.Context, in *npool.UpdateSnapshotRequest) (*npool.UpdateSnapshotResponse, error) {
	if req := in.GetInfo(); req == nil {
		logger.Sugar().Errorw(
			"UpdateSnapshot",
			"In", in,
		)
		return &npool.UpdateSnapshotResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	h, err := handler.NewHandler(ctx,
		handler.WithID(in.Info.ID, true),
		handler.WithIndex(in.Info.Index, false),
		handler.WithSnapshotCommP(in.Info.SnapshotCommP, false),
		handler.WithSnapshotRoot(in.Info.SnapshotRoot, false),
		handler.WithSnapshotURI(in.Info.SnapshotURI, false),
		handler.WithBackupState(in.Info.BackupState, false),
	)
	if err != nil {
		logger.Sugar().Errorw("UpdateSnapshot", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.UpdateSnapshot(ctx)
	if err != nil {
		logger.Sugar().Errorw("UpdateSnapshot", "ID", in.Info.GetID(), "error", err)
		return &npool.UpdateSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSnapshotResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSnapshot(ctx context.Context, in *npool.GetSnapshotRequest) (*npool.GetSnapshotResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("GetSnapshot", "ID", in.GetID(), "error", err)
		return &npool.GetSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.GetSnapshot(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetSnapshot", "ID", in.GetID(), "error", err)
		return &npool.GetSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSnapshotResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSnapshotOnly(ctx context.Context, in *npool.GetSnapshotOnlyRequest) (*npool.GetSnapshotOnlyResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	if err != nil {
		logger.Sugar().Errorw("GetSnapshotOnly", "error", err)
		return &npool.GetSnapshotOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := h.GetSnapshots(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetSnapshotOnly", "error", err)
		return &npool.GetSnapshotOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total != 1 {
		errMsg := "more than one result or have no result"
		logger.Sugar().Errorw("GetBlockOnly", "error", errMsg)
		return &npool.GetSnapshotOnlyResponse{}, status.Error(codes.Internal, errMsg)
	}

	return &npool.GetSnapshotOnlyResponse{
		Info: infos[0],
	}, nil
}

func (s *Server) GetSnapshots(ctx context.Context, in *npool.GetSnapshotsRequest) (*npool.GetSnapshotsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
		handler.WithOffset(in.Offset),
		handler.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw("GetSnapshots", "error", err)
		return &npool.GetSnapshotsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := h.GetSnapshots(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetSnapshots", "error", err)
		return &npool.GetSnapshotsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSnapshotsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) ExistSnapshot(ctx context.Context, in *npool.ExistSnapshotRequest) (*npool.ExistSnapshotResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistSnapshot", "ID", in.GetID(), "error", err)
		return &npool.ExistSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}
	exist, err := h.ExistSnapshot(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistSnapshot", "ID", in.GetID(), "error", err)
		return &npool.ExistSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSnapshotResponse{
		Exist: exist,
	}, nil
}

func (s *Server) ExistSnapshotConds(ctx context.Context, in *npool.ExistSnapshotCondsRequest) (*npool.ExistSnapshotCondsResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw("ExistSnapshotConds", "error", err)
		return &npool.ExistSnapshotCondsResponse{}, status.Error(codes.Internal, err.Error())
	}
	exist, err := h.ExistSnapshotConds(ctx)
	if err != nil {
		logger.Sugar().Errorw("ExistSnapshotConds", "error", err)
		return &npool.ExistSnapshotCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSnapshotCondsResponse{
		Exist: exist,
	}, nil
}

func (s *Server) DeleteSnapshot(ctx context.Context, in *npool.DeleteSnapshotRequest) (*npool.DeleteSnapshotResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw("DeleteSnapshot", "ID", in.GetID(), "error", err)
		return &npool.DeleteSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := h.DeleteSnapshot(ctx)
	if err != nil {
		logger.Sugar().Errorw("DeleteSnapshot", "ID", in.GetID(), "error", err)
		return &npool.DeleteSnapshotResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSnapshotResponse{
		Info: info,
	}, nil
}
