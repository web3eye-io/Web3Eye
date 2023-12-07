package snapshot

import (
	"context"
	"fmt"

	constant "github.com/web3eye-io/Web3Eye/nft-meta/pkg/const"
	snapshotcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/snapshot"
	snapshotproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID            *uint32
	EntID         *uuid.UUID
	Index         *uint64
	SnapshotCommP *string
	SnapshotRoot  *string
	SnapshotURI   *string
	BackupState   *string

	Reqs   []*snapshotcrud.Req
	Conds  *snapshotcrud.Conds
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithIndex(u *uint64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid index")
			}
			return nil
		}
		h.Index = u
		return nil
	}
}
func WithSnapshotCommP(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid snapshotcommp")
			}
			return nil
		}
		h.SnapshotCommP = u
		return nil
	}
}
func WithSnapshotRoot(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid snapshotroot")
			}
			return nil
		}
		h.SnapshotRoot = u
		return nil
	}
}
func WithSnapshotURI(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid snapshoturi")
			}
			return nil
		}
		h.SnapshotURI = u
		return nil
	}
}
func WithBackupState(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid backupstate")
			}
			return nil
		}
		h.BackupState = u
		return nil
	}
}

// nolint:gocyclo
func WithReqs(reqs []*snapshotproto.SnapshotReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*snapshotcrud.Req{}
		for _, req := range reqs {
			_req := &snapshotcrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.Index != nil {
				_req.Index = req.Index
			}
			if req.SnapshotCommP != nil {
				_req.SnapshotCommP = req.SnapshotCommP
			}
			if req.SnapshotRoot != nil {
				_req.SnapshotRoot = req.SnapshotRoot
			}
			if req.SnapshotURI != nil {
				_req.SnapshotURI = req.SnapshotURI
			}
			if req.BackupState != nil {
				_req.BackupState = req.BackupState
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *snapshotproto.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &snapshotcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.Index != nil {
			h.Conds.Index = &cruder.Cond{
				Op:  conds.GetIndex().GetOp(),
				Val: conds.GetIndex().GetValue(),
			}
		}
		if conds.SnapshotCommP != nil {
			h.Conds.SnapshotCommP = &cruder.Cond{
				Op:  conds.GetSnapshotCommP().GetOp(),
				Val: conds.GetSnapshotCommP().GetValue(),
			}
		}
		if conds.SnapshotRoot != nil {
			h.Conds.SnapshotRoot = &cruder.Cond{
				Op:  conds.GetSnapshotRoot().GetOp(),
				Val: conds.GetSnapshotRoot().GetValue(),
			}
		}
		if conds.SnapshotURI != nil {
			h.Conds.SnapshotURI = &cruder.Cond{
				Op:  conds.GetSnapshotURI().GetOp(),
				Val: conds.GetSnapshotURI().GetValue(),
			}
		}
		if conds.BackupState != nil {
			h.Conds.BackupState = &cruder.Cond{
				Op:  conds.GetBackupState().GetOp(),
				Val: conds.GetBackupState().GetValue(),
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
