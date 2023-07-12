package block

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/block"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

func Create(ctx context.Context, in *npool.BlockReq) (*ent.Block, error) {
	var info *ent.Block
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Block.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Upsert(ctx context.Context, in *npool.BlockReq) (*ent.Block, error) {
	conds := &npool.Conds{
		ChainType: &web3eye.StringVal{
			Op:    "eq",
			Value: in.GetChainType().String(),
		},
		ChainID: &web3eye.StringVal{
			Op:    "eq",
			Value: in.GetChainID(),
		},
		BlockNumber: &web3eye.Uint64Val{
			Op:    "eq",
			Value: in.GetBlockNumber(),
		},
	}
	rows, total, err := Rows(ctx, conds, 0, 1)
	if err != nil {
		return nil, err
	}

	if total != 0 {
		id := rows[0].ID.String()
		in.ID = &id
		return Update(ctx, in)
	}

	return Create(ctx, in)
}

//nolint:gocyclo
func CreateSet(c *ent.BlockCreate, in *npool.BlockReq) *ent.BlockCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.ChainType != nil {
		c.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		c.SetChainID(in.GetChainID())
	}
	if in.BlockNumber != nil {
		c.SetBlockNumber(in.GetBlockNumber())
	}
	if in.BlockHash != nil {
		c.SetBlockHash(in.GetBlockHash())
	}
	if in.BlockTime != nil {
		c.SetBlockTime(in.GetBlockTime())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*npool.BlockReq) ([]*ent.Block, error) {
	var err error
	rows := []*ent.Block{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.BlockCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Block.Create(), info)
		}
		rows, err = tx.Block.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

//nolint:gocyclo
func Update(ctx context.Context, in *npool.BlockReq) (*ent.Block, error) {
	var err error
	var info *ent.Block

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.Block.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.ChainType != nil {
			u.SetChainType(in.GetChainType().String())
		}
		if in.ChainID != nil {
			u.SetChainID(in.GetChainID())
		}
		if in.BlockNumber != nil {
			u.SetBlockNumber(in.GetBlockNumber())
		}
		if in.BlockHash != nil {
			u.SetBlockHash(in.GetBlockHash())
		}
		if in.BlockTime != nil {
			u.SetBlockTime(in.GetBlockTime())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// func UpdateSet(u *ent.BlockUpdateOne, in *npool.BlockReq) *ent.BlockUpdateOne {
// 	if in.VectorID != nil {
// 		u.SetVectorID(in.GetVectorID())
// 	}
// 	if in.Remark != nil {
// 		u.SetRemark(in.GetRemark())
// 	}
// 	return u
// }

func Row(ctx context.Context, id uuid.UUID) (*ent.Block, error) {
	var info *ent.Block
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Block.Query().Where(block.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.BlockQuery, error) {
	stm := cli.Block.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(block.ID(id))
		default:
			return nil, fmt.Errorf("invalid block field")
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
			stm.Where(block.IDIn(ids...))
		}
	}
	if conds.ChainType != nil {
		switch conds.GetChainType().GetOp() {
		case cruder.EQ:
			stm.Where(block.ChainType(conds.GetChainType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid block field")
		}
	}
	if conds.ChainID != nil {
		switch conds.GetChainID().GetOp() {
		case cruder.EQ:
			stm.Where(block.ChainID(conds.GetChainID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid block field")
		}
	}
	if conds.BlockNumber != nil {
		switch conds.GetBlockNumber().GetOp() {
		case cruder.EQ:
			stm.Where(block.BlockNumber(conds.GetBlockNumber().GetValue()))
		default:
			return nil, fmt.Errorf("invalid block field")
		}
	}
	if conds.BlockHash != nil {
		switch conds.GetBlockHash().GetOp() {
		case cruder.EQ:
			stm.Where(block.BlockHash(conds.GetBlockHash().GetValue()))
		default:
			return nil, fmt.Errorf("invalid block field")
		}
	}
	if conds.BlockTime != nil {
		switch conds.GetBlockTime().GetOp() {
		case cruder.EQ:
			stm.Where(block.BlockTime(conds.GetBlockTime().GetValue()))
		default:
			return nil, fmt.Errorf("invalid block field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Block, int, error) {
	var err error
	rows := []*ent.Block{}
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
			Order(ent.Desc(block.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.Block, err error) {
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
		exist, err = cli.Block.Query().Where(block.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Block, error) {
	var info *ent.Block
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Block.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
