package endpoint

import (
	"context"
	"fmt"

	constant "github.com/web3eye-io/Web3Eye/nft-meta/pkg/const"
	endpointcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/endpoint"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	endpointproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uint32
	EntID     *uuid.UUID
	ChainType *basetype.ChainType
	ChainID   *string
	Address   *string
	State     *basetype.EndpointState
	RPS       *uint32
	Remark    *string

	Reqs   []*endpointcrud.Req
	Conds  *endpointcrud.Conds
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

func WithAddress(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid address")
			}
			return nil
		}
		h.Address = u
		return nil
	}
}

func WithState(u *basetype.EndpointState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid state")
			}
			return nil
		}
		if _, ok := basetype.EndpointState_name[int32(*u)]; !ok {
			return fmt.Errorf("invalid state field")
		}
		h.State = u
		return nil
	}
}

func WithRPS(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid rps")
			}
			return nil
		}
		h.RPS = u
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

func WithReqs(reqs []*endpointproto.EndpointReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*endpointcrud.Req{}
		for _, req := range reqs {
			_req := &endpointcrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.ChainType != nil {
				if _, ok := basetype.ChainType_name[int32(*req.ChainType)]; !ok {
					return fmt.Errorf("invalid chaintype field")
				}
				_req.ChainType = req.ChainType
			}
			if req.ChainID != nil {
				_req.ChainID = req.ChainID
			}
			if req.Address != nil {
				_req.Address = req.Address
			}
			if req.State != nil {
				if _, ok := basetype.EndpointState_name[int32(*req.State)]; !ok {
					return fmt.Errorf("invalid state field")
				}
				_req.State = req.State
			}
			if req.Remark != nil {
				_req.Remark = req.Remark
			}
			if req.RPS != nil {
				_req.RPS = req.RPS
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *endpointproto.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &endpointcrud.Conds{}
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
		if conds.Address != nil {
			h.Conds.Address = &cruder.Cond{
				Op:  conds.GetAddress().GetOp(),
				Val: conds.GetAddress().GetValue(),
			}
		}
		if conds.State != nil {
			h.Conds.State = &cruder.Cond{
				Op:  conds.GetState().GetOp(),
				Val: basetype.EndpointState(conds.GetState().GetValue()),
			}
		}
		if conds.RPS != nil {
			h.Conds.Rps = &cruder.Cond{
				Op:  conds.GetRPS().GetOp(),
				Val: conds.GetRPS().GetValue(),
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
