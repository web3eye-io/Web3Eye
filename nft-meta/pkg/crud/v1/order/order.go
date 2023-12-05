package order

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	entorder "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	ChainType   *basetype.ChainType
	ChainID     *string
	TxHash      *string
	BlockNumber *uint64
	TxIndex     *uint32
	LogIndex    *uint32
	Recipient   *string
	Remark      *string
}

func CreateSet(c *ent.OrderCreate, req *Req) *ent.OrderCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ChainType != nil {
		c.SetChainType((*req.ChainType).String())
	}
	if req.ChainID != nil {
		c.SetChainID(*req.ChainID)
	}
	if req.TxHash != nil {
		c.SetTxHash(*req.TxHash)
	}
	if req.BlockNumber != nil {
		c.SetBlockNumber(*req.BlockNumber)
	}
	if req.TxIndex != nil {
		c.SetTxIndex(*req.TxIndex)
	}
	if req.LogIndex != nil {
		c.SetLogIndex(*req.LogIndex)
	}
	if req.Recipient != nil {
		c.SetRecipient(*req.Recipient)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.OrderUpdateOne, req *Req) (*ent.OrderUpdateOne, error) {
	if req.ChainType != nil {
		u.SetChainType(req.ChainType.String())
	}
	if req.ChainID != nil {
		u.SetChainID(*req.ChainID)
	}
	if req.TxHash != nil {
		u.SetTxHash(*req.TxHash)
	}
	if req.BlockNumber != nil {
		u.SetBlockNumber(*req.BlockNumber)
	}
	if req.TxIndex != nil {
		u.SetTxIndex(*req.TxIndex)
	}
	if req.LogIndex != nil {
		u.SetLogIndex(*req.LogIndex)
	}
	if req.Recipient != nil {
		u.SetRecipient(*req.Recipient)
	}
	if req.Remark != nil {
		u.SetRemark(*req.Remark)
	}
	return u, nil
}

type Conds struct {
	EntID       *cruder.Cond
	ChainType   *cruder.Cond
	ChainID     *cruder.Cond
	TxHash      *cruder.Cond
	BlockNumber *cruder.Cond
	TxIndex     *cruder.Cond
	LogIndex    *cruder.Cond
	Recipient   *cruder.Cond
	Remark      *cruder.Cond
}

func SetQueryConds(q *ent.OrderQuery, conds *Conds) (*ent.OrderQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entorder.EntID(entid))
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
			q.Where(entorder.ChainType(chaintype.String()))
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
			q.Where(entorder.ChainID(chainid))
		default:
			return nil, fmt.Errorf("invalid chainid field")
		}
	}
	if conds.TxHash != nil {
		txhash, ok := conds.TxHash.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid txhash")
		}
		switch conds.TxHash.Op {
		case cruder.EQ:
			q.Where(entorder.TxHash(txhash))
		default:
			return nil, fmt.Errorf("invalid txhash field")
		}
	}
	if conds.BlockNumber != nil {
		blocknumber, ok := conds.BlockNumber.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid blocknumber")
		}
		switch conds.BlockNumber.Op {
		case cruder.EQ:
			q.Where(entorder.BlockNumber(blocknumber))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}
	if conds.TxIndex != nil {
		txindex, ok := conds.TxIndex.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid txindex")
		}
		switch conds.TxIndex.Op {
		case cruder.EQ:
			q.Where(entorder.TxIndex(txindex))
		default:
			return nil, fmt.Errorf("invalid txindex field")
		}
	}
	if conds.LogIndex != nil {
		logindex, ok := conds.LogIndex.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid logindex")
		}
		switch conds.LogIndex.Op {
		case cruder.EQ:
			q.Where(entorder.LogIndex(logindex))
		default:
			return nil, fmt.Errorf("invalid logindex field")
		}
	}
	if conds.Recipient != nil {
		recipient, ok := conds.Recipient.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid recipient")
		}
		switch conds.Recipient.Op {
		case cruder.EQ:
			q.Where(entorder.Recipient(recipient))
		default:
			return nil, fmt.Errorf("invalid recipient field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(entorder.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	return q, nil
}
