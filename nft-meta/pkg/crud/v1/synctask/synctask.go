package synctask

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	entsynctask "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/synctask"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	ChainType   *basetype.ChainType
	ChainID     *string
	Start       *uint64
	End         *uint64
	Current     *uint64
	Topic       *string
	Description *string
	SyncState   *basetype.SyncState
	Remark      *string
}

func CreateSet(c *ent.SyncTaskCreate, req *Req) *ent.SyncTaskCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ChainType != nil {
		c.SetChainType((*req.ChainType).String())
	}
	if req.ChainID != nil {
		c.SetChainID(*req.ChainID)
	}
	if req.Start != nil {
		c.SetStart(*req.Start)
	}
	if req.End != nil {
		c.SetEnd(*req.End)
	}
	if req.Current != nil {
		c.SetCurrent(*req.Current)
	}
	if req.Topic != nil {
		c.SetTopic(*req.Topic)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	if req.SyncState != nil {
		c.SetSyncState((*req.SyncState).String())
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	return c
}

func UpdateSet(u *ent.SyncTaskUpdateOne, req *Req) (*ent.SyncTaskUpdateOne, error) {
	if req.ChainType != nil {
		u.SetChainType(req.ChainType.String())
	}
	if req.ChainID != nil {
		u.SetChainID(*req.ChainID)
	}
	if req.Start != nil {
		u.SetStart(*req.Start)
	}
	if req.End != nil {
		u.SetEnd(*req.End)
	}
	if req.Current != nil {
		u.SetCurrent(*req.Current)
	}
	if req.Topic != nil {
		u.SetTopic(*req.Topic)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.SyncState != nil {
		u.SetSyncState(req.SyncState.String())
	}
	if req.Remark != nil {
		u.SetRemark(*req.Remark)
	}
	return u, nil
}

type Conds struct {
	EntID       *cruder.Cond
	ChainType   *cruder.Cond
	ChainID     *cruder.Cond
	Start       *cruder.Cond
	End         *cruder.Cond
	Current     *cruder.Cond
	Topic       *cruder.Cond
	Description *cruder.Cond
	SyncState   *cruder.Cond
	Remark      *cruder.Cond
}

func SetQueryConds(q *ent.SyncTaskQuery, conds *Conds) (*ent.SyncTaskQuery, error) { //nolint
	if conds.EntID != nil {
		entid, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entsynctask.EntID(entid))
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
			q.Where(entsynctask.ChainType(chaintype.String()))
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
			q.Where(entsynctask.ChainID(chainid))
		default:
			return nil, fmt.Errorf("invalid chainid field")
		}
	}
	if conds.Start != nil {
		start, ok := conds.Start.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid start")
		}
		switch conds.Start.Op {
		case cruder.EQ:
			q.Where(entsynctask.Start(start))
		default:
			return nil, fmt.Errorf("invalid start field")
		}
	}
	if conds.End != nil {
		end, ok := conds.End.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid end")
		}
		switch conds.End.Op {
		case cruder.EQ:
			q.Where(entsynctask.End(end))
		default:
			return nil, fmt.Errorf("invalid end field")
		}
	}
	if conds.Current != nil {
		current, ok := conds.Current.Val.(uint64)
		if !ok {
			return nil, fmt.Errorf("invalid current")
		}
		switch conds.Current.Op {
		case cruder.EQ:
			q.Where(entsynctask.Current(current))
		default:
			return nil, fmt.Errorf("invalid current field")
		}
	}
	if conds.Topic != nil {
		topic, ok := conds.Topic.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid topic")
		}
		switch conds.Topic.Op {
		case cruder.EQ:
			q.Where(entsynctask.Topic(topic))
		default:
			return nil, fmt.Errorf("invalid topic field")
		}
	}
	if conds.Description != nil {
		description, ok := conds.Description.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid description")
		}
		switch conds.Description.Op {
		case cruder.EQ:
			q.Where(entsynctask.Description(description))
		default:
			return nil, fmt.Errorf("invalid description field")
		}
	}
	if conds.SyncState != nil {
		syncstate, ok := conds.SyncState.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid syncstate")
		}
		switch conds.SyncState.Op {
		case cruder.EQ:
			q.Where(entsynctask.SyncState(syncstate))
		default:
			return nil, fmt.Errorf("invalid syncstate field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(entsynctask.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	return q, nil
}
