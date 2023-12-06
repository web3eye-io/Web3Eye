package block

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	entblock "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/block"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	ChainType   *basetype.ChainType
	ChainID     *string
	BlockNumber *uint64
	BlockHash   *string
	BlockTime   *int64
	ParseState  *basetype.BlockParseState
	Remark      *string
}

func CreateSet(c *ent.BlockCreate, req *Req) *ent.BlockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ChainType != nil {
		c.SetChainType((*req.ChainType).String())
	}
	if req.ChainID != nil {
		c.SetChainID(*req.ChainID)
	}
	if req.BlockNumber != nil {
		c.SetBlockNumber(*req.BlockNumber)
	}
	if req.BlockHash != nil {
		c.SetBlockHash(*req.BlockHash)
	}
	if req.BlockTime != nil {
		c.SetBlockTime(*req.BlockTime)
	}
	if req.ParseState != nil {
		c.SetParseState((*req.ParseState).String())
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.BlockUpdateOne, req *Req) (*ent.BlockUpdateOne, error) {
	if req.ChainType != nil {
		u.SetChainType(req.ChainType.String())
	}
	if req.ChainID != nil {
		u.SetChainID(*req.ChainID)
	}
	if req.BlockNumber != nil {
		u.SetBlockNumber(*req.BlockNumber)
	}
	if req.BlockHash != nil {
		u.SetBlockHash(*req.BlockHash)
	}
	if req.BlockTime != nil {
		u.SetBlockTime(*req.BlockTime)
	}
	if req.ParseState != nil {
		u.SetParseState(req.ParseState.String())
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
	BlockNumber *cruder.Cond
	BlockHash   *cruder.Cond
	BlockTime   *cruder.Cond
	ParseState  *cruder.Cond
	Remark      *cruder.Cond
}

func SetQueryConds(q *ent.BlockQuery, conds *Conds) (*ent.BlockQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entblock.EntID(entid))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		entids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entblock.EntIDIn(entids...))
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
			q.Where(entblock.ChainType(chaintype.String()))
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
			q.Where(entblock.ChainID(chainid))
		default:
			return nil, fmt.Errorf("invalid chainid field")
		}
	}
	if conds.BlockNumber != nil {
		blocknumber, ok := conds.BlockNumber.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid blocknumber")
		}
		switch conds.BlockNumber.Op {
		case cruder.EQ:
			q.Where(entblock.BlockNumber(blocknumber))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}
	if conds.BlockHash != nil {
		blockhash, ok := conds.BlockHash.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid blockhash")
		}
		switch conds.BlockHash.Op {
		case cruder.EQ:
			q.Where(entblock.BlockHash(blockhash))
		default:
			return nil, fmt.Errorf("invalid blockhash field")
		}
	}
	if conds.BlockTime != nil {
		blocktime, ok := conds.BlockTime.Val.(int64)
		if !ok {
			return nil, fmt.Errorf("invalid blocktime")
		}
		switch conds.BlockTime.Op {
		case cruder.EQ:
			q.Where(entblock.BlockTime(blocktime))
		default:
			return nil, fmt.Errorf("invalid blocktime field")
		}
	}
	if conds.ParseState != nil {
		parsestate, ok := conds.ParseState.Val.(basetype.BlockParseState)
		if !ok {
			return nil, fmt.Errorf("invalid parsestate")
		}
		switch conds.ParseState.Op {
		case cruder.EQ:
			q.Where(entblock.ParseState(parsestate.String()))
		default:
			return nil, fmt.Errorf("invalid parsestate field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(entblock.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	return q, nil
}
