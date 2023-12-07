package order

import (
	"context"
	"fmt"

	constant "github.com/web3eye-io/Web3Eye/nft-meta/pkg/const"
	ordercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/order"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	orderproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	ChainType   *basetype.ChainType
	ChainID     *string
	TxHash      *string
	BlockNumber *uint64
	TxIndex     *uint32
	LogIndex    *uint32
	Recipient   *string
	TargetItems []*orderproto.OrderItem
	OfferItems  []*orderproto.OrderItem
	Remark      *string

	Reqs   []*ordercrud.Req
	Conds  *ordercrud.Conds
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

func WithTxHash(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		h.TxHash = u
		return nil
	}
}

func WithBlockNumber(u *uint64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		h.BlockNumber = u
		return nil
	}
}

func WithTxIndex(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		h.TxIndex = u
		return nil
	}
}

func WithLogIndex(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		h.LogIndex = u
		return nil
	}
}

func WithRecipient(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		h.Recipient = u
		return nil
	}
}

func WithTargetItems(u []*orderproto.OrderItem, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid targetitems")
			}
			return nil
		}
		h.TargetItems = u
		return nil
	}
}

func WithOfferItems(u []*orderproto.OrderItem, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid offeritems")
			}
			return nil
		}
		h.OfferItems = u
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

// nolint:gocyclo
func WithReqs(reqs []*orderproto.OrderReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*ordercrud.Req{}
		for _, req := range reqs {
			_req := &ordercrud.Req{}
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
			if req.TxHash != nil {
				_req.TxHash = req.TxHash
			}
			if req.BlockNumber != nil {
				_req.BlockNumber = req.BlockNumber
			}
			if req.TxIndex != nil {
				_req.TxIndex = req.TxIndex
			}
			if req.LogIndex != nil {
				_req.LogIndex = req.LogIndex
			}
			if req.Recipient != nil {
				_req.Recipient = req.Recipient
			}
			if req.TargetItems != nil {
				_req.TargetItems = req.TargetItems
			}
			if req.OfferItems != nil {
				_req.OfferItems = req.OfferItems
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
func WithConds(conds *orderproto.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &ordercrud.Conds{}
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
		if conds.TxHash != nil {
			h.Conds.TxHash = &cruder.Cond{
				Op:  conds.GetTxHash().GetOp(),
				Val: conds.GetTxHash().GetValue(),
			}
		}
		if conds.BlockNumber != nil {
			h.Conds.BlockNumber = &cruder.Cond{
				Op:  conds.GetBlockNumber().GetOp(),
				Val: conds.GetBlockNumber().GetValue(),
			}
		}
		if conds.TxIndex != nil {
			h.Conds.TxIndex = &cruder.Cond{
				Op:  conds.GetTxIndex().GetOp(),
				Val: conds.GetTxIndex().GetValue(),
			}
		}
		if conds.LogIndex != nil {
			h.Conds.LogIndex = &cruder.Cond{
				Op:  conds.GetLogIndex().GetOp(),
				Val: conds.GetLogIndex().GetValue(),
			}
		}
		if conds.Recipient != nil {
			h.Conds.Recipient = &cruder.Cond{
				Op:  conds.GetRecipient().GetOp(),
				Val: conds.GetRecipient().GetValue(),
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
