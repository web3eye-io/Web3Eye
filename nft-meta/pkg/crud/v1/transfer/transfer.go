package transfer

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	enttransfer "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/google/uuid"
)

type Req struct {
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
}

//nolint:gocyclo
func CreateSet(c *ent.TransferCreate, req *Req) *ent.TransferCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ChainType != nil {
		c.SetChainType(req.ChainType.String())
	}
	if req.ChainID != nil {
		c.SetChainID(*req.ChainID)
	}
	if req.Contract != nil {
		c.SetContract(*req.Contract)
	}
	if req.TokenType != nil {
		c.SetTokenType(req.TokenType.String())
	}
	if req.TokenID != nil {
		c.SetTokenID(*req.TokenID)
	}
	if req.From != nil {
		c.SetFrom(*req.From)
	}
	if req.To != nil {
		c.SetTo(*req.To)
	}
	if req.Amount != nil {
		amount := utils.Uint64ToDecStr(*req.Amount)
		c.SetAmount(amount)
	}
	if req.BlockNumber != nil {
		c.SetBlockNumber(*req.BlockNumber)
	}
	if req.TxHash != nil {
		c.SetTxHash(*req.TxHash)
	}
	if req.BlockHash != nil {
		c.SetBlockHash(*req.BlockHash)
	}
	if req.TxTime != nil {
		c.SetTxTime(*req.TxTime)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.TransferUpdateOne, req *Req) (*ent.TransferUpdateOne, error) {
	if req.ChainType != nil {
		u.SetChainType(req.ChainType.String())
	}
	if req.ChainID != nil {
		u.SetChainID(*req.ChainID)
	}
	if req.Contract != nil {
		u.SetContract(*req.Contract)
	}
	if req.TokenType != nil {
		u.SetTokenType(req.TokenType.String())
	}
	if req.TokenID != nil {
		u.SetTokenID(*req.TokenID)
	}
	if req.From != nil {
		u.SetFrom(*req.From)
	}
	if req.To != nil {
		u.SetTo(*req.To)
	}
	if req.Amount != nil {
		amount := utils.Uint64ToDecStr(*req.Amount)
		u.SetAmount(amount)
	}
	if req.BlockNumber != nil {
		u.SetBlockNumber(*req.BlockNumber)
	}
	if req.TxHash != nil {
		u.SetTxHash(*req.TxHash)
	}
	if req.BlockHash != nil {
		u.SetBlockHash(*req.BlockHash)
	}
	if req.TxTime != nil {
		u.SetTxTime(*req.TxTime)
	}
	if req.Remark != nil {
		u.SetRemark(*req.Remark)
	}
	return u, nil
}

type Conds struct {
	EntID       *cruder.Cond
	EntIDs      *cruder.Cond
	ChainType   *cruder.Cond
	ChainID     *cruder.Cond
	Contract    *cruder.Cond
	TokenType   *cruder.Cond
	TokenID     *cruder.Cond
	From        *cruder.Cond
	To          *cruder.Cond
	Amount      *cruder.Cond
	BlockNumber *cruder.Cond
	TxHash      *cruder.Cond
	BlockHash   *cruder.Cond
	TxTime      *cruder.Cond
	Remark      *cruder.Cond
}

func SetQueryConds(q *ent.TransferQuery, conds *Conds) (*ent.TransferQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttransfer.EntID(entid))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		entids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttransfer.EntIDIn(entids...))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.ChainType != nil {
		chaintype, ok := conds.ChainType.Val.(basetype.ChainType)
		if !ok {
			return nil, fmt.Errorf("invalid chaintype")
		}
		switch conds.ChainType.Op {
		case cruder.EQ:
			q.Where(enttransfer.ChainType(chaintype.String()))
		default:
			return nil, fmt.Errorf("invalid chaintype field")
		}
	}
	if conds.ChainID != nil {
		chainid, ok := conds.ChainID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid chainid")
		}
		switch conds.ChainID.Op {
		case cruder.EQ:
			q.Where(enttransfer.ChainID(chainid))
		default:
			return nil, fmt.Errorf("invalid chainid field")
		}
	}
	if conds.Contract != nil {
		contract, ok := conds.Contract.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid contract")
		}
		switch conds.Contract.Op {
		case cruder.EQ:
			q.Where(enttransfer.Contract(contract))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.TokenType != nil {
		tokentype, ok := conds.TokenType.Val.(basetype.TokenType)
		if !ok {
			return nil, fmt.Errorf("invalid tokentype")
		}
		switch conds.TokenType.Op {
		case cruder.EQ:
			q.Where(enttransfer.TokenType(tokentype.String()))
		default:
			return nil, fmt.Errorf("invalid tokentype field")
		}
	}
	if conds.TokenID != nil {
		tokenid, ok := conds.TokenID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid tokenid")
		}
		switch conds.TokenID.Op {
		case cruder.EQ:
			q.Where(enttransfer.TokenID(tokenid))
		default:
			return nil, fmt.Errorf("invalid tokenid field")
		}
	}
	if conds.From != nil {
		from, ok := conds.From.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid from")
		}
		switch conds.From.Op {
		case cruder.EQ:
			q.Where(enttransfer.From(from))
		default:
			return nil, fmt.Errorf("invalid from field")
		}
	}
	if conds.To != nil {
		to, ok := conds.To.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid to")
		}
		switch conds.To.Op {
		case cruder.EQ:
			q.Where(enttransfer.To(to))
		default:
			return nil, fmt.Errorf("invalid to field")
		}
	}
	if conds.Amount != nil {
		amount, ok := conds.Amount.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid amount")
		}
		switch conds.Amount.Op {
		case cruder.EQ:
			q.Where(enttransfer.Amount(amount))
		default:
			return nil, fmt.Errorf("invalid amount field")
		}
	}
	if conds.BlockNumber != nil {
		blocknumber, ok := conds.BlockNumber.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid blocknumber")
		}
		switch conds.BlockNumber.Op {
		case cruder.EQ:
			q.Where(enttransfer.BlockNumber(blocknumber))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}
	if conds.TxHash != nil {
		txhash, ok := conds.TxHash.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid txhash")
		}
		switch conds.TxHash.Op {
		case cruder.EQ:
			q.Where(enttransfer.TxHash(txhash))
		default:
			return nil, fmt.Errorf("invalid txhash field")
		}
	}
	if conds.BlockHash != nil {
		blockhash, ok := conds.BlockHash.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid blockhash")
		}
		switch conds.BlockHash.Op {
		case cruder.EQ:
			q.Where(enttransfer.BlockHash(blockhash))
		default:
			return nil, fmt.Errorf("invalid blockhash field")
		}
	}
	if conds.TxTime != nil {
		txtime, ok := conds.TxTime.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid txtime")
		}
		switch conds.TxTime.Op {
		case cruder.EQ:
			q.Where(enttransfer.TxTime(txtime))
		default:
			return nil, fmt.Errorf("invalid txtime field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(enttransfer.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	return q, nil
}
