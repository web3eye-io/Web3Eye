package contract

import (
	"context"
	"fmt"

	constant "github.com/web3eye-io/Web3Eye/nft-meta/pkg/const"
	contractcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/contract"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	contractproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	ChainType   *basetype.ChainType
	ChainID     *string
	Address     *string
	Name        *string
	Symbol      *string
	Decimals    *uint32
	Creator     *string
	BlockNum    *uint64
	TxHash      *string
	TxTime      *uint32
	ProfileURL  *string
	BaseURL     *string
	BannerURL   *string
	Description *string
	Remark      *string
	Reqs        []*contractcrud.Req
	Conds       *contractcrud.Conds
	Offset      int32
	Limit       int32
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
func WithName(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		h.Name = u
		return nil
	}
}
func WithSymbol(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid symbol")
			}
			return nil
		}
		h.Symbol = u
		return nil
	}
}
func WithDecimals(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid decimals")
			}
			return nil
		}
		h.Decimals = u
		return nil
	}
}
func WithCreator(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid creator")
			}
			return nil
		}
		h.Creator = u
		return nil
	}
}
func WithBlockNum(u *uint64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid blocknum")
			}
			return nil
		}
		h.BlockNum = u
		return nil
	}
}
func WithTxHash(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid txhash")
			}
			return nil
		}
		h.TxHash = u
		return nil
	}
}
func WithTxTime(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid txtime")
			}
			return nil
		}
		h.TxTime = u
		return nil
	}
}
func WithProfileURL(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid profileurl")
			}
			return nil
		}
		h.ProfileURL = u
		return nil
	}
}
func WithBaseURL(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid baseurl")
			}
			return nil
		}
		h.BaseURL = u
		return nil
	}
}
func WithBannerURL(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid bannerurl")
			}
			return nil
		}
		h.BannerURL = u
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

//nolint:gocyclo
func WithReqs(reqs []*contractproto.ContractReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*contractcrud.Req{}
		for _, req := range reqs {
			_req := &contractcrud.Req{}
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
			if req.Address != nil {
				_req.Address = req.Address
			}
			if req.Name != nil {
				_req.Name = req.Name
			}
			if req.Symbol != nil {
				_req.Symbol = req.Symbol
			}
			if req.Decimals != nil {
				_req.Decimals = req.Decimals
			}
			if req.Creator != nil {
				_req.Creator = req.Creator
			}
			if req.BlockNum != nil {
				_req.BlockNum = req.BlockNum
			}
			if req.TxHash != nil {
				_req.TxHash = req.TxHash
			}
			if req.TxTime != nil {
				_req.TxTime = req.TxTime
			}
			if req.ProfileURL != nil {
				_req.ProfileURL = req.ProfileURL
			}
			if req.BaseURL != nil {
				_req.BaseURL = req.BaseURL
			}
			if req.BannerURL != nil {
				_req.BannerURL = req.BannerURL
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

//nolint:gocyclo,funlen
func WithConds(conds *contractproto.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &contractcrud.Conds{}
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
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		if conds.Symbol != nil {
			h.Conds.Symbol = &cruder.Cond{
				Op:  conds.GetSymbol().GetOp(),
				Val: conds.GetSymbol().GetValue(),
			}
		}
		if conds.Decimals != nil {
			h.Conds.Decimals = &cruder.Cond{
				Op:  conds.GetDecimals().GetOp(),
				Val: conds.GetDecimals().GetValue(),
			}
		}
		if conds.Creator != nil {
			h.Conds.Creator = &cruder.Cond{
				Op:  conds.GetCreator().GetOp(),
				Val: conds.GetCreator().GetValue(),
			}
		}
		if conds.BlockNum != nil {
			h.Conds.BlockNum = &cruder.Cond{
				Op:  conds.GetBlockNum().GetOp(),
				Val: conds.GetBlockNum().GetValue(),
			}
		}
		if conds.TxHash != nil {
			h.Conds.TxHash = &cruder.Cond{
				Op:  conds.GetTxHash().GetOp(),
				Val: conds.GetTxHash().GetValue(),
			}
		}
		if conds.TxTime != nil {
			h.Conds.TxTime = &cruder.Cond{
				Op:  conds.GetTxTime().GetOp(),
				Val: conds.GetTxTime().GetValue(),
			}
		}
		if conds.ProfileURL != nil {
			h.Conds.ProfileURL = &cruder.Cond{
				Op:  conds.GetProfileURL().GetOp(),
				Val: conds.GetProfileURL().GetValue(),
			}
		}
		if conds.BaseURL != nil {
			h.Conds.BaseURL = &cruder.Cond{
				Op:  conds.GetBaseURL().GetOp(),
				Val: conds.GetBaseURL().GetValue(),
			}
		}
		if conds.BannerURL != nil {
			h.Conds.BannerURL = &cruder.Cond{
				Op:  conds.GetBannerURL().GetOp(),
				Val: conds.GetBannerURL().GetValue(),
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
