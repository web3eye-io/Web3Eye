package snapshot

import (
	"context"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/snapshot"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

func Create(ctx context.Context, in *npool.SnapshotReq) (*ent.Snapshot, error) {
	var info *ent.Snapshot
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Snapshot.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.SnapshotCreate, in *npool.SnapshotReq) *ent.SnapshotCreate {
	c.SetID(uuid.MustParse(in.GetID()))
	c.SetIndex(in.GetIndex())
	c.SetSnapshotCommP(in.GetSnapshotCommP())
	c.SetSnapshotRoot(in.GetSnapshotRoot())
	c.SetSnapshotURI(in.GetSnapshotURI())
	c.SetBackupState(in.GetBackupState())
	return c
}

func CreateBulk(ctx context.Context, in []*npool.SnapshotReq) ([]*ent.Snapshot, error) {
	var err error
	rows := []*ent.Snapshot{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.SnapshotCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Snapshot.Create(), info)
		}
		rows, err = tx.Snapshot.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.UpdateSnapshotRequest) (*ent.Snapshot, error) {
	var err error
	var info *ent.Snapshot

	if _, err := uuid.Parse(in.Info.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.Snapshot.UpdateOneID(uuid.MustParse(in.Info.GetID()))
		if in.Info.Index != nil {
			u.SetIndex(in.Info.GetIndex())
		}
		if in.Info.SnapshotCommP != nil {
			u.SetSnapshotCommP(in.Info.GetSnapshotCommP())
		}
		if in.Info.SnapshotRoot != nil {
			u.SetSnapshotRoot(in.Info.GetSnapshotRoot())
		}
		if in.Info.SnapshotURI != nil {
			u.SetSnapshotURI(in.Info.GetSnapshotURI())
		}
		if in.Info.BackupState != nil {
			u.SetBackupState(in.Info.GetBackupState())
		}

		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// func UpdateSet(u *ent.SnapshotUpdateOne, in *npool.SnapshotReq) *ent.SnapshotUpdateOne {
// 	if in.VectorID != nil {
// 		u.SetVectorID(in.GetVectorID())
// 	}
// 	if in.Remark != nil {
// 		u.SetRemark(in.GetRemark())
// 	}
// 	return u
// }

func Row(ctx context.Context, id uuid.UUID) (*ent.Snapshot, error) {
	var info *ent.Snapshot
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Snapshot.Query().Where(snapshot.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.SnapshotQuery, error) {
	stm := cli.Snapshot.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(snapshot.ID(id))
		default:
			return nil, fmt.Errorf("invalid snapshot field")
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
			stm.Where(snapshot.IDIn(ids...))
		}
	}

	if conds.Index != nil {
		switch conds.GetIndex().GetOp() {
		case cruder.EQ:
			stm.Where(snapshot.Index(conds.GetIndex().GetValue()))
		default:
			return nil, fmt.Errorf("invalid snapshot field")
		}
	}
	if conds.SnapshotCommP != nil {
		switch conds.GetSnapshotCommP().GetOp() {
		case cruder.EQ:
			stm.Where(snapshot.SnapshotCommP(conds.GetSnapshotCommP().GetValue()))
		default:
			return nil, fmt.Errorf("invalid snapshot field")
		}
	}
	if conds.SnapshotRoot != nil {
		switch conds.GetSnapshotRoot().GetOp() {
		case cruder.EQ:
			stm.Where(snapshot.SnapshotRoot(conds.GetSnapshotRoot().GetValue()))
		default:
			return nil, fmt.Errorf("invalid snapshot field")
		}
	}
	if conds.SnapshotURI != nil {
		switch conds.GetSnapshotURI().GetOp() {
		case cruder.EQ:
			stm.Where(snapshot.SnapshotURI(conds.GetSnapshotURI().GetValue()))
		default:
			return nil, fmt.Errorf("invalid snapshot field")
		}
	}
	if conds.BackupState != nil {
		switch conds.GetBackupState().GetOp() {
		case cruder.EQ:
			stm.Where(snapshot.BackupState(conds.GetBackupState().GetValue()))
		default:
			return nil, fmt.Errorf("invalid snapshot field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Snapshot, int, error) {
	var err error
	rows := []*ent.Snapshot{}
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
			Order(ent.Desc(snapshot.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.Snapshot, err error) {
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
		exist, err = cli.Snapshot.Query().Where(snapshot.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Snapshot, error) {
	var info *ent.Snapshot
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Snapshot.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
