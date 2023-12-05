package snapshot

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	entsnapshot "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/snapshot"

	"github.com/google/uuid"
)

type Req struct {
	ID            *uint32
	EntID         *uuid.UUID
	Index         *uint64
	SnapshotCommP *string
	SnapshotRoot  *string
	SnapshotURI   *string
	BackupState   *string
}

func CreateSet(c *ent.SnapshotCreate, req *Req) *ent.SnapshotCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	if req.SnapshotCommP != nil {
		c.SetSnapshotCommP(*req.SnapshotCommP)
	}
	if req.SnapshotRoot != nil {
		c.SetSnapshotRoot(*req.SnapshotRoot)
	}
	if req.SnapshotURI != nil {
		c.SetSnapshotURI(*req.SnapshotURI)
	}
	if req.BackupState != nil {
		c.SetBackupState(*req.BackupState)
	}
	return c
}

func UpdateSet(u *ent.SnapshotUpdateOne, req *Req) (*ent.SnapshotUpdateOne, error) {
	if req.Index != nil {
		u.SetIndex(*req.Index)
	}
	if req.SnapshotCommP != nil {
		u.SetSnapshotCommP(*req.SnapshotCommP)
	}
	if req.SnapshotRoot != nil {
		u.SetSnapshotRoot(*req.SnapshotRoot)
	}
	if req.SnapshotURI != nil {
		u.SetSnapshotURI(*req.SnapshotURI)
	}
	if req.BackupState != nil {
		u.SetBackupState(*req.BackupState)
	}
	return u, nil
}

type Conds struct {
	EntID         *cruder.Cond
	Index         *cruder.Cond
	SnapshotCommP *cruder.Cond
	SnapshotRoot  *cruder.Cond
	SnapshotURI   *cruder.Cond
	BackupState   *cruder.Cond
}

func SetQueryConds(q *ent.SnapshotQuery, conds *Conds) (*ent.SnapshotQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entsnapshot.EntID(entid))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.Index != nil {
		index, ok := conds.Index.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid index")
		}
		switch conds.Index.Op {
		case cruder.EQ:
			q.Where(entsnapshot.Index(index))
		default:
			return nil, fmt.Errorf("invalid index field")
		}
	}
	if conds.SnapshotCommP != nil {
		chainid, ok := conds.SnapshotCommP.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid chainid")
		}
		switch conds.SnapshotCommP.Op {
		case cruder.EQ:
			q.Where(entsnapshot.SnapshotCommP(chainid))
		default:
			return nil, fmt.Errorf("invalid chainid field")
		}
	}
	if conds.SnapshotRoot != nil {
		snapshotroot, ok := conds.SnapshotRoot.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid snapshotroot")
		}
		switch conds.SnapshotRoot.Op {
		case cruder.EQ:
			q.Where(entsnapshot.SnapshotRoot(snapshotroot))
		default:
			return nil, fmt.Errorf("invalid snapshotroot field")
		}
	}
	if conds.SnapshotURI != nil {
		snapshoturi, ok := conds.SnapshotURI.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid snapshoturi")
		}
		switch conds.SnapshotURI.Op {
		case cruder.EQ:
			q.Where(entsnapshot.SnapshotURI(snapshoturi))
		default:
			return nil, fmt.Errorf("invalid snapshoturi field")
		}
	}
	if conds.BackupState != nil {
		backupstate, ok := conds.BackupState.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid backupstate")
		}
		switch conds.BackupState.Op {
		case cruder.EQ:
			q.Where(entsnapshot.BackupState(backupstate))
		default:
			return nil, fmt.Errorf("invalid backupstate field")
		}
	}
	return q, nil
}
