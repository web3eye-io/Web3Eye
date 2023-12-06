package order

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	entorderitem "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	OrderID   *uuid.UUID
	Contract  *string
	TokenType *basetype.TokenType
	TokenID   *string
	Amount    *uint64
	Remark    *string
}

func CreateSet(c *ent.OrderItemCreate, req *Req) *ent.OrderItemCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
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
	if req.Amount != nil {
		amount := utils.Uint64ToDecStr(*req.Amount)
		c.SetAmount(amount)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.OrderUpdateOne, req *Req) (*ent.OrderUpdateOne, error) {

	if req.Remark != nil {
		u.SetRemark(*req.Remark)
	}
	return u, nil
}

type Conds struct {
	EntID     *cruder.Cond
	EntIDs    *cruder.Cond
	OrderID   *cruder.Cond
	Contract  *cruder.Cond
	TokenType *cruder.Cond
	TokenID   *cruder.Cond
	Amount    *cruder.Cond
	Remark    *cruder.Cond
}

func SetQueryConds(q *ent.OrderItemQuery, conds *Conds) (*ent.OrderItemQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entorderitem.EntID(entid))
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
			q.Where(entorderitem.EntIDIn(entids...))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.OrderID != nil {
		orderid, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entorderitem.OrderID(orderid))
		default:
			return nil, fmt.Errorf("invalid orderid field")
		}
	}
	if conds.Contract != nil {
		contract, ok := conds.Contract.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid contract")
		}
		switch conds.Contract.Op {
		case cruder.EQ:
			q.Where(entorderitem.Contract(contract))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.TokenType != nil {
		tokentype, ok := conds.TokenType.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid tokentype")
		}
		switch conds.TokenType.Op {
		case cruder.EQ:
			q.Where(entorderitem.TokenType(tokentype))
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
			q.Where(entorderitem.TokenID(tokenid))
		default:
			return nil, fmt.Errorf("invalid tokenid field")
		}
	}
	if conds.Amount != nil {
		amount, ok := conds.Amount.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid amount")
		}
		switch conds.Amount.Op {
		case cruder.EQ:
			q.Where(entorderitem.Amount(amount))
		default:
			return nil, fmt.Errorf("invalid amount field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(entorderitem.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	return q, nil
}
