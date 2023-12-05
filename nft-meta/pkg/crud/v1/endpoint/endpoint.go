package endpoint

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	entendpoint "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/endpoint"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	ChainType *basetype.ChainType
	ChainID   *string
	Address   *string
	State     *basetype.EndpointState
	Remark    *string
}

func CreateSet(c *ent.EndpointCreate, req *Req) *ent.EndpointCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ChainType != nil {
		c.SetChainType((*req.ChainType).String())
	}
	if req.ChainID != nil {
		c.SetChainID(*req.ChainID)
	}
	if req.Address != nil {
		c.SetAddress(*req.Address)
	}
	if req.State != nil {
		c.SetState((*req.State).String())
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.EndpointUpdateOne, req *Req) (*ent.EndpointUpdateOne, error) {
	if req.ChainType != nil {
		u.SetChainType(req.ChainType.String())
	}
	if req.ChainID != nil {
		u.SetChainID(*req.ChainID)
	}
	if req.Address != nil {
		u.SetAddress(*req.Address)
	}
	if req.State != nil {
		u.SetState((*req.State).String())
	}
	if req.Remark != nil {
		u.SetRemark(*req.Remark)
	}
	return u, nil
}

type Conds struct {
	EntID     *cruder.Cond
	ChainType *cruder.Cond
	ChainID   *cruder.Cond
	Address   *cruder.Cond
	State     *cruder.Cond
	Remark    *cruder.Cond
}

func SetQueryConds(q *ent.EndpointQuery, conds *Conds) (*ent.EndpointQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entendpoint.EntID(entid))
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
			q.Where(entendpoint.ChainType(chaintype.String()))
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
			q.Where(entendpoint.ChainID(chainid))
		default:
			return nil, fmt.Errorf("invalid chainid field")
		}
	}
	if conds.Address != nil {
		address, ok := conds.Address.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid address")
		}
		switch conds.Address.Op {
		case cruder.EQ:
			q.Where(entendpoint.Address(address))
		default:
			return nil, fmt.Errorf("invalid address field")
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid state")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entendpoint.State(state))
		default:
			return nil, fmt.Errorf("invalid state field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(entendpoint.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	return q, nil
}
