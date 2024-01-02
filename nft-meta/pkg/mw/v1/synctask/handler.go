package synctask

import (
	"context"
	"fmt"

	constant "github.com/web3eye-io/Web3Eye/nft-meta/pkg/const"
	synctaskcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/synctask"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	synctaskproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	ChainType   *basetype.ChainType
	ChainID     *string
	Start       *uint64
	End         *uint64
	Current     *uint64
	Topic       *string
	SyncState   *basetype.SyncState
	Description *string
	Remark      *string

	Reqs   []*synctaskcrud.Req
	Conds  *synctaskcrud.Conds
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

func WithChainType(u *basetype.ChainType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid chaintype")
			}
			return nil
		}
		if _, ok := basetype.ChainType_name[int32(*u)]; !ok {
			return fmt.Errorf("invalid chaintype field")
		}
		h.ChainType = u
		return nil
	}
}
func WithChainID(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid chainid")
			}
			return nil
		}
		h.ChainID = u
		return nil
	}
}
func WithStart(u *uint64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid start")
			}
			return nil
		}
		h.Start = u
		return nil
	}
}
func WithEnd(u *uint64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid end")
			}
			return nil
		}
		h.End = u
		return nil
	}
}
func WithCurrent(u *uint64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid current")
			}
			return nil
		}
		h.Current = u
		return nil
	}
}
func WithTopic(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid topic")
			}
			return nil
		}
		h.Topic = u
		return nil
	}
}
func WithDescription(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid description")
			}
			return nil
		}
		h.Description = u
		return nil
	}
}
func WithSyncState(u *basetype.SyncState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid syncstate")
			}
			return nil
		}
		if _, ok := basetype.SyncState_name[int32(*u)]; !ok {
			return fmt.Errorf("invalid syncstate field")
		}
		h.SyncState = u
		return nil
	}
}
func WithRemark(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		h.Remark = u
		return nil
	}
}

func WithReqs(reqs []*synctaskproto.SyncTaskReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*synctaskcrud.Req{}
		for _, req := range reqs {
			_req := &synctaskcrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.ChainType != nil {
				if _, ok := basetype.ChainType_name[int32(*req.ChainType)]; !ok {
					return fmt.Errorf("invalid parsestate field")
				}
				_req.ChainType = req.ChainType
			}
			if req.ChainID != nil {
				_req.ChainID = req.ChainID
			}
			if req.Start != nil {
				_req.Start = req.Start
			}
			if req.End != nil {
				_req.End = req.End
			}
			if req.Current != nil {
				_req.Current = req.Current
			}
			if req.Topic != nil {
				_req.Topic = req.Topic
			}
			if req.SyncState != nil {
				if _, ok := basetype.SyncState_name[int32(*req.SyncState)]; !ok {
					return fmt.Errorf("invalid synctask field")
				}
				_req.SyncState = req.SyncState
			}
			if req.Description != nil {
				_req.Description = req.Description
			}
			if req.Remark != nil {
				_req.Remark = req.Remark
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *synctaskproto.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &synctaskcrud.Conds{}
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
		if conds.ChainType != nil {
			h.Conds.ChainType = &cruder.Cond{
				Op:  conds.GetChainType().GetOp(),
				Val: basetype.ChainType(conds.GetChainType().GetValue()),
			}
		}
		if conds.ChainID != nil {
			h.Conds.ChainID = &cruder.Cond{
				Op:  conds.GetChainID().GetOp(),
				Val: conds.GetChainID().GetValue(),
			}
		}
		if conds.Start != nil {
			h.Conds.Start = &cruder.Cond{
				Op:  conds.GetStart().GetOp(),
				Val: conds.GetStart().GetValue(),
			}
		}
		if conds.End != nil {
			h.Conds.End = &cruder.Cond{
				Op:  conds.GetEnd().GetOp(),
				Val: conds.GetEnd().GetValue(),
			}
		}
		if conds.Current != nil {
			h.Conds.Current = &cruder.Cond{
				Op:  conds.GetCurrent().GetOp(),
				Val: conds.GetCurrent().GetValue(),
			}
		}
		if conds.Topic != nil {
			h.Conds.Topic = &cruder.Cond{
				Op:  conds.GetTopic().GetOp(),
				Val: conds.GetTopic().GetValue(),
			}
		}
		if conds.SyncState != nil {
			h.Conds.SyncState = &cruder.Cond{
				Op:  conds.GetSyncState().GetOp(),
				Val: basetype.SyncState(conds.GetSyncState().GetValue()),
			}
		}
		if conds.Description != nil {
			h.Conds.Description = &cruder.Cond{
				Op:  conds.GetDescription().GetOp(),
				Val: conds.GetDescription().GetValue(),
			}
		}
		if conds.Remark != nil {
			h.Conds.Remark = &cruder.Cond{
				Op:  conds.GetRemark().GetOp(),
				Val: conds.GetRemark().GetValue(),
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
