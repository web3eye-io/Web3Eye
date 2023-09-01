package orderpair

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderpair"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/orderpair"
)

func Create(ctx context.Context, in *npool.OrderPairReq) (*ent.OrderPair, error) {
	var info *ent.OrderPair
	var err error

	if in == nil {
		return nil, errors.New("input is nil")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.OrderPair.Create(), in)
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

func CreateSet(c *ent.OrderPairCreate, in *npool.OrderPairReq) *ent.OrderPairCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.TxHash != nil {
		c.SetTxHash(in.GetTxHash())
	}
	if in.Recipient != nil {
		c.SetRecipient(in.GetRecipient())
	}
	if in.TargetID != nil {
		c.SetTargetID(in.GetTargetID())
	}
	if in.OfferID != nil {
		c.SetOfferID(in.GetOfferID())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func UpsertBulk(ctx context.Context, in []*npool.OrderPairReq) error {
	var err error
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.OrderPairCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.OrderPair.Create(), info)
		}
		err = tx.OrderPair.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
		return err
	})
	return err
}

func CreateBulk(ctx context.Context, in []*npool.OrderPairReq) ([]*ent.OrderPair, error) {
	var err error
	rows := []*ent.OrderPair{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.OrderPairCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.OrderPair.Create(), info)
		}
		rows, err = tx.OrderPair.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

//nolint:gocyclo
func Update(ctx context.Context, in *npool.OrderPairReq) (*ent.OrderPair, error) {
	var err error
	var info *ent.OrderPair

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.OrderPair.UpdateOneID(uuid.MustParse(in.GetID()))
		u = UpdateSet(u, in)
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.OrderPairUpdateOne, in *npool.OrderPairReq) *ent.OrderPairUpdateOne {
	if in.TxHash != nil {
		u.SetTxHash(in.GetTxHash())
	}
	if in.Recipient != nil {
		u.SetRecipient(in.GetRecipient())
	}
	if in.TargetID != nil {
		u.SetTargetID(in.GetTargetID())
	}
	if in.OfferID != nil {
		u.SetOfferID(in.GetOfferID())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.OrderPair, error) {
	var info *ent.OrderPair
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.OrderPair.Query().Where(orderpair.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.OrderPairQuery, error) {
	stm := cli.OrderPair.Query()
	if conds == nil {
		return stm, nil
	}

	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(orderpair.ID(id))
		default:
			return nil, fmt.Errorf("invalid orderpair field")
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
			stm.Where(orderpair.IDIn(ids...))
		}
	}

	if conds.TxHash != nil {
		switch conds.GetTxHash().GetOp() {
		case cruder.EQ:
			stm.Where(orderpair.TxHash(conds.GetTxHash().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderpair field")
		}
	}
	if conds.Recipient != nil {
		switch conds.GetRecipient().GetOp() {
		case cruder.EQ:
			stm.Where(orderpair.Recipient(conds.GetRecipient().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderpair field")
		}
	}
	if conds.TargetID != nil {
		switch conds.GetTargetID().GetOp() {
		case cruder.EQ:
			stm.Where(orderpair.TargetID(conds.GetTargetID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderpair field")
		}
	}
	if conds.TargetID != nil {
		switch conds.GetTargetID().GetOp() {
		case cruder.EQ:
			stm.Where(orderpair.TargetID(conds.GetTargetID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderpair field")
		}
	}
	if conds.OfferID != nil {
		switch conds.GetOfferID().GetOp() {
		case cruder.EQ:
			stm.Where(orderpair.OfferID(conds.GetOfferID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderpair field")
		}
	}
	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(orderpair.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid orderpair field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.OrderPair, int, error) {
	var err error
	rows := []*ent.OrderPair{}
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
			Order(ent.Desc(orderpair.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.OrderPair, err error) {
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
		exist, err = cli.OrderPair.Query().Where(orderpair.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.OrderPair, error) {
	var info *ent.OrderPair
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.OrderPair.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
