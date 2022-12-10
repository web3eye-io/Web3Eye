package synctask

import (
	"context"
	"fmt"
	"time"

	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/synctask"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	npool "github.com/web3eye-io/cyber-tracer/message/cybertracer/nftmeta/v1/synctask"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"
)

func Create(ctx context.Context, in *npool.SyncTaskReq) (*ent.SyncTask, error) {
	var info *ent.SyncTask
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.SyncTask.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.SyncTaskCreate, in *npool.SyncTaskReq) *ent.SyncTaskCreate {
	if in.ID != nil {
		id, err := uuid.Parse(*in.ID)
		if err != nil {
			id = uuid.New()
		}
		c.SetID(id)
	}
	if in.ChainType != nil {
		c.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		c.SetChainID(in.GetChainID())
	}
	if in.Start != nil {
		c.SetStart(in.GetStart())
	}
	if in.End != nil {
		c.SetEnd(in.GetEnd())
	}
	if in.Current != nil {
		c.SetCurrent(in.GetCurrent())
	}
	if in.Topic != nil {
		c.SetTopic(in.GetTopic())
	}
	if in.Description != nil {
		c.SetDescription(in.GetDescription())
	}
	if in.SyncState != nil {
		c.SetSyncState(in.GetSyncState().String())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*npool.SyncTaskReq) ([]*ent.SyncTask, error) {
	var err error
	rows := []*ent.SyncTask{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.SyncTaskCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.SyncTask.Create(), info)
		}
		rows, err = tx.SyncTask.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.SyncTaskReq) (*ent.SyncTask, error) {
	var err error
	var info *ent.SyncTask

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.SyncTask.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.ChainType != nil {
			u.SetChainType(in.GetChainType().String())
		}
		if in.ChainID != nil {
			u.SetChainID(in.GetChainID())
		}
		if in.Start != nil {
			u.SetStart(in.GetStart())
		}
		if in.End != nil {
			u.SetEnd(in.GetEnd())
		}
		if in.Current != nil {
			u.SetCurrent(in.GetCurrent())
		}
		if in.Topic != nil {
			u.SetTopic(in.GetTopic())
		}
		if in.Description != nil {
			u.SetDescription(in.GetDescription())
		}
		if in.SyncState != nil {
			u.SetSyncState(in.GetSyncState().String())
		}
		if in.Remark != nil {
			u.SetRemark(in.GetRemark())
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// func UpdateSet(u *ent.SyncTaskUpdateOne, in *npool.SyncTaskReq) *ent.SyncTaskUpdateOne {
// 	if in.VectorID != nil {
// 		u.SetVectorID(in.GetVectorID())
// 	}
// 	if in.Remark != nil {
// 		u.SetRemark(in.GetRemark())
// 	}
// 	return u
// }

func Row(ctx context.Context, id uuid.UUID) (*ent.SyncTask, error) {
	var info *ent.SyncTask
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.SyncTask.Query().Where(synctask.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.SyncTaskQuery, error) {
	stm := cli.SyncTask.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.ID(id))
		default:
			return nil, fmt.Errorf("invalid synctask field")
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
			stm.Where(synctask.IDIn(ids...))
		}
	}
	if conds.ChainType != nil {
		switch conds.GetChainType().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.ChainType(conds.GetChainType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}
	if conds.ChainID != nil {
		switch conds.GetChainID().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.ChainID(conds.GetChainID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}
	if conds.Start != nil {
		switch conds.GetStart().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.Start(conds.GetStart().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}
	if conds.End != nil {
		switch conds.GetEnd().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.End(conds.GetEnd().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}
	if conds.Current != nil {
		switch conds.GetCurrent().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.Current(conds.GetCurrent().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}
	if conds.Topic != nil {
		switch conds.GetTopic().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.Topic(conds.GetTopic().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}
	if conds.Description != nil {
		switch conds.GetDescription().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.Description(conds.GetDescription().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}
	if conds.SyncState != nil {
		switch conds.GetSyncState().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.SyncState(conds.GetSyncState().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}
	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(synctask.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid synctask field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.SyncTask, int, error) {
	var err error
	rows := []*ent.SyncTask{}
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
			Order(ent.Desc(synctask.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.SyncTask, err error) {
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
		exist, err = cli.SyncTask.Query().Where(synctask.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.SyncTask, error) {
	var info *ent.SyncTask
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.SyncTask.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
