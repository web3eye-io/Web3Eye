package transfer

import (
	"context"
	"fmt"

	constant "github.com/web3eye-io/Web3Eye/nft-meta/pkg/const"
	transfercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/transfer"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	transferproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	ChainType   *basetype.ChainType
	ChainID     *string
	Contract    *string
	TokenType   *basetype.TokenType
	TokenID     *string
	From        *string
	To          *string
	Amount      *uint64
	BlockNumber *uint64
	TxHash      *string
	BlockHash   *string
	TxTime      *uint64
	Remark      *string
	LogIndex    *uint32

	Reqs   []*transfercrud.Req
	Conds  *transfercrud.Conds
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
func WithContract(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid contract")
			}
			return nil
		}
		h.Contract = u
		return nil
	}
}
func WithTokenType(u *basetype.TokenType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid tokentype")
			}
			return nil
		}
		if _, ok := basetype.TokenType_name[int32(*u)]; !ok {
			return fmt.Errorf("invalid tokentype field")
		}
		h.TokenType = u
		return nil
	}
}
func WithTokenID(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid tokenid")
			}
			return nil
		}
		h.TokenID = u
		return nil
	}
}
func WithFrom(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid from")
			}
			return nil
		}
		h.From = u
		return nil
	}
}
func WithTo(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid to")
			}
			return nil
		}
		h.To = u
		return nil
	}
}
func WithAmount(u *uint64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid amount")
			}
			return nil
		}
		h.Amount = u
		return nil
	}
}
func WithBlockNumber(u *uint64, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid blocknumber")
			}
			return nil
		}
		h.BlockNumber = u
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
func WithBlockHash(u *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid blockhash")
			}
			return nil
		}
		h.BlockHash = u
		return nil
	}
}
func WithTxTime(u *uint64, must bool) func(context.Context, *Handler) error {
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

func WithLogIndex(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid logindex")
			}
			return nil
		}
		h.LogIndex = u
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
func WithReqs(reqs []*transferproto.TransferReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*transfercrud.Req{}
		for _, req := range reqs {
			_req := &transfercrud.Req{}
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
			if req.Contract != nil {
				_req.Contract = req.Contract
			}
			if req.TokenType != nil {
				if _, ok := basetype.TokenType_name[int32(*req.TokenType)]; !ok {
					return fmt.Errorf("invalid tokentype field")
				}
				_req.TokenType = req.TokenType
			}
			if req.TokenID != nil {
				_req.TokenID = req.TokenID
			}
			if req.From != nil {
				_req.From = req.From
			}
			if req.To != nil {
				_req.To = req.To
			}
			if req.Amount != nil {
				_req.Amount = req.Amount
			}
			if req.BlockNumber != nil {
				_req.BlockNumber = req.BlockNumber
			}
			if req.TxHash != nil {
				_req.TxHash = req.TxHash
			}
			if req.BlockHash != nil {
				_req.BlockHash = req.BlockHash
			}
			if req.TxTime != nil {
				_req.TxTime = req.TxTime
			}
			if req.Remark != nil {
				_req.Remark = req.Remark
			}
			if req.LogIndex != nil {
				_req.LogIndex = req.LogIndex
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

//nolint:gocyclo,funlen
func WithConds(conds *transferproto.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &transfercrud.Conds{}
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
		if conds.Contract != nil {
			h.Conds.Contract = &cruder.Cond{
				Op:  conds.GetContract().GetOp(),
				Val: conds.GetContract().GetValue(),
			}
		}
		if conds.TokenType != nil {
			h.Conds.TokenType = &cruder.Cond{
				Op:  conds.GetTokenType().GetOp(),
				Val: basetype.TokenType(conds.GetTokenType().GetValue()),
			}
		}
		if conds.TokenID != nil {
			h.Conds.TokenID = &cruder.Cond{
				Op:  conds.GetTokenID().GetOp(),
				Val: conds.GetTokenID().GetValue(),
			}
		}
		if conds.From != nil {
			h.Conds.From = &cruder.Cond{
				Op:  conds.GetFrom().GetOp(),
				Val: conds.GetFrom().GetValue(),
			}
		}
		if conds.To != nil {
			h.Conds.To = &cruder.Cond{
				Op:  conds.GetTo().GetOp(),
				Val: conds.GetTo().GetValue(),
			}
		}
		if conds.Amount != nil {
			h.Conds.Amount = &cruder.Cond{
				Op:  conds.GetAmount().GetOp(),
				Val: conds.GetAmount().GetValue(),
			}
		}
		if conds.BlockNumber != nil {
			h.Conds.BlockNumber = &cruder.Cond{
				Op:  conds.GetBlockNumber().GetOp(),
				Val: conds.GetBlockNumber().GetValue(),
			}
		}
		if conds.TxHash != nil {
			h.Conds.TxHash = &cruder.Cond{
				Op:  conds.GetTxHash().GetOp(),
				Val: conds.GetTxHash().GetValue(),
			}
		}
		if conds.BlockHash != nil {
			h.Conds.BlockHash = &cruder.Cond{
				Op:  conds.GetBlockHash().GetOp(),
				Val: conds.GetBlockHash().GetValue(),
			}
		}
		if conds.TxTime != nil {
			h.Conds.TxTime = &cruder.Cond{
				Op:  conds.GetTxTime().GetOp(),
				Val: conds.GetTxTime().GetValue(),
			}
		}
		if conds.Remark != nil {
			h.Conds.Remark = &cruder.Cond{
				Op:  conds.GetRemark().GetOp(),
				Val: conds.GetRemark().GetValue(),
			}
		}
		if conds.LogIndex != nil {
			h.Conds.LogIndex = &cruder.Cond{
				Op:  conds.GetLogIndex().GetOp(),
				Val: conds.GetLogIndex().GetValue(),
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
