package blocknumber

import (
	"context"
	"fmt"
	"time"

	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/blocknumber"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/nftmeta/v1/blocknumber"
	"github.com/google/uuid"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"
)

func Create(ctx context.Context, in *npool.BlockNumberReq) (*ent.BlockNumber, error) {
	var info *ent.BlockNumber
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.BlockNumber.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.BlockNumberCreate, in *npool.BlockNumberReq) *ent.BlockNumberCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.ChainType != nil {
		c.SetChainType(in.GetChainType())
	}
	if in.ChainID != nil {
		c.SetChainID(in.GetChainID())
	}
	if in.Identifier != nil {
		c.SetIdentifier(in.GetIdentifier())
	}
	if in.CurrentNum != nil {
		c.SetCurrentNum(in.GetCurrentNum())
	}
	if in.Topic != nil {
		c.SetTopic(in.GetTopic())
	}
	if in.Description != nil {
		c.SetDescription(in.GetDescription())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*npool.BlockNumberReq) ([]*ent.BlockNumber, error) {
	var err error
	rows := []*ent.BlockNumber{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.BlockNumberCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.BlockNumber.Create(), info)
		}
		rows, err = tx.BlockNumber.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.BlockNumberReq) (*ent.BlockNumber, error) {
	var err error
	var info *ent.BlockNumber

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.BlockNumber.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.ChainType != nil {
			u.SetChainType(in.GetChainType())
		}
		if in.ChainID != nil {
			u.SetChainID(in.GetChainID())
		}
		if in.Identifier != nil {
			u.SetIdentifier(in.GetIdentifier())
		}
		if in.CurrentNum != nil {
			u.SetCurrentNum(in.GetCurrentNum())
		}
		if in.Topic != nil {
			u.SetTopic(in.GetTopic())
		}
		if in.Description != nil {
			u.SetDescription(in.GetDescription())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// func UpdateSet(u *ent.BlockNumberUpdateOne, in *npool.BlockNumberReq) *ent.BlockNumberUpdateOne {
// 	if in.VectorID != nil {
// 		u.SetVectorID(in.GetVectorID())
// 	}
// 	if in.Remark != nil {
// 		u.SetRemark(in.GetRemark())
// 	}
// 	return u
// }

func Row(ctx context.Context, id uuid.UUID) (*ent.BlockNumber, error) {
	var info *ent.BlockNumber
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.BlockNumber.Query().Where(blocknumber.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.BlockNumberQuery, error) {
	stm := cli.BlockNumber.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(blocknumber.ID(id))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
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
			stm.Where(blocknumber.IDIn(ids...))
		}
	}
	if conds.ChainType != nil {
		switch conds.GetChainType().GetOp() {
		case cruder.EQ:
			stm.Where(blocknumber.ChainType(conds.GetChainType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}
	if conds.ChainID != nil {
		switch conds.GetChainID().GetOp() {
		case cruder.EQ:
			stm.Where(blocknumber.ChainID(conds.GetChainID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}
	if conds.Identifier != nil {
		switch conds.GetIdentifier().GetOp() {
		case cruder.EQ:
			stm.Where(blocknumber.Identifier(conds.GetIdentifier().GetValue()))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}
	if conds.CurrentNum != nil {
		switch conds.GetCurrentNum().GetOp() {
		case cruder.EQ:
			stm.Where(blocknumber.CurrentNum(conds.GetCurrentNum().GetValue()))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}
	if conds.Topic != nil {
		switch conds.GetTopic().GetOp() {
		case cruder.EQ:
			stm.Where(blocknumber.Topic(conds.GetTopic().GetValue()))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}
	if conds.Description != nil {
		switch conds.GetDescription().GetOp() {
		case cruder.EQ:
			stm.Where(blocknumber.Description(conds.GetDescription().GetValue()))
		default:
			return nil, fmt.Errorf("invalid blocknumber field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.BlockNumber, int, error) {
	var err error
	rows := []*ent.BlockNumber{}
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
			Order(ent.Desc(blocknumber.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.BlockNumber, err error) {
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
		exist, err = cli.BlockNumber.Query().Where(blocknumber.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.BlockNumber, error) {
	var info *ent.BlockNumber
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.BlockNumber.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
