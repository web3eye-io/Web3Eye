package orderitem

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/orderitem"
)

func Create(ctx context.Context, in *npool.OrderItemReq) (*ent.OrderItem, error) {
	var info *ent.OrderItem
	var err error

	if in == nil {
		return nil, errors.New("input is nil")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.OrderItem.Create(), in)
		if err != nil {
			return err
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.OrderItemCreate, in *npool.OrderItemReq) *ent.OrderItemCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.Contract != nil {
		c.SetContract(in.GetContract())
	}
	if in.TokenType != nil {
		c.SetTokenType(in.GetTokenType())
	}
	if in.TokenID != nil {
		c.SetTokenID(in.GetTokenID())
	}
	if in.Amount != nil {
		c.SetAmount(in.GetAmount())
	}
	if in.PortionNum != nil {
		c.SetPortionNum(in.GetPortionNum())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func UpsertBulk(ctx context.Context, in []*npool.OrderItemReq) error {
	var err error
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.OrderItemCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.OrderItem.Create(), info)
		}
		err = tx.OrderItem.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
		return err
	})
	return err
}

func CreateBulk(ctx context.Context, in []*npool.OrderItemReq) ([]*ent.OrderItem, error) {
	var err error
	rows := []*ent.OrderItem{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.OrderItemCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.OrderItem.Create(), info)
		}
		rows, err = tx.OrderItem.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

//nolint:gocyclo
func Update(ctx context.Context, in *npool.OrderItemReq) (*ent.OrderItem, error) {
	var err error
	var info *ent.OrderItem

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.OrderItem.UpdateOneID(uuid.MustParse(in.GetID()))
		u = UpdateSet(u, in)
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.OrderItemUpdateOne, in *npool.OrderItemReq) *ent.OrderItemUpdateOne {
	if in.Contract != nil {
		u.SetContract(in.GetContract())
	}
	if in.TokenType != nil {
		u.SetTokenType(in.GetTokenType())
	}
	if in.TokenID != nil {
		u.SetTokenID(in.GetTokenID())
	}
	if in.Amount != nil {
		u.SetAmount(in.GetAmount())
	}
	if in.PortionNum != nil {
		u.SetPortionNum(in.GetPortionNum())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error) {
	var info *ent.OrderItem
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.OrderItem.Query().Where(orderitem.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.OrderItemQuery, error) {
	stm := cli.OrderItem.Query()
	if conds == nil {
		return stm, nil
	}

	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(orderitem.ID(id))
		default:
			return nil, fmt.Errorf("invalid orderitem field")
		}
	}

	if conds.IDs != nil {
		if conds.GetIDs().GetOp() == cruder.IN {
			var ids []uuid.UUID
			for _, val := range conds.GetIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(orderitem.IDIn(ids...))
		}
	}

	if conds.Contract != nil {
		switch conds.GetContract().GetOp() {
		case cruder.EQ:
			stm.Where(orderitem.Contract(conds.GetContract().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderitem field")
		}
	}
	if conds.TokenType != nil {
		switch conds.GetTokenType().GetOp() {
		case cruder.EQ:
			stm.Where(orderitem.TokenType(conds.GetTokenType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderitem field")
		}
	}
	if conds.TokenID != nil {
		switch conds.GetTokenID().GetOp() {
		case cruder.EQ:
			stm.Where(orderitem.TokenID(conds.GetTokenID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderitem field")
		}
	}
	if conds.TokenID != nil {
		switch conds.GetTokenID().GetOp() {
		case cruder.EQ:
			stm.Where(orderitem.TokenID(conds.GetTokenID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderitem field")
		}
	}
	if conds.Amount != nil {
		switch conds.GetAmount().GetOp() {
		case cruder.EQ:
			stm.Where(orderitem.Amount(conds.GetAmount().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderitem field")
		}
	}
	if conds.PortionNum != nil {
		switch conds.GetPortionNum().GetOp() {
		case cruder.EQ:
			stm.Where(orderitem.PortionNum(conds.GetPortionNum().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderitem field")
		}
	}
	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(orderitem.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderitem field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.OrderItem, int, error) {
	var err error
	rows := []*ent.OrderItem{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(orderitem.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.OrderItem, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error

	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.OrderItem.Query().Where(orderitem.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error

	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.OrderItem, error) {
	var info *ent.OrderItem
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.OrderItem.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
