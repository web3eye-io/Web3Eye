package endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/endpoint"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
)

func Create(ctx context.Context, in *npool.EndpointReq) (*ent.Endpoint, error) {
	var info *ent.Endpoint
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Endpoint.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:gocyclo
func CreateSet(c *ent.EndpointCreate, in *npool.EndpointReq) *ent.EndpointCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.ChainType != nil {
		c.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		c.SetChainID(in.GetChainID())
	}
	if in.Address != nil {
		c.SetAddress(in.GetAddress())
	}
	if in.State != nil {
		c.SetState(in.GetState().String())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*npool.EndpointReq) ([]*ent.Endpoint, error) {
	var err error
	rows := []*ent.Endpoint{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.EndpointCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Endpoint.Create(), info)
		}
		rows, err = tx.Endpoint.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

//nolint:gocyclo
func Update(ctx context.Context, in *npool.EndpointReq) (*ent.Endpoint, error) {
	var err error
	var info *ent.Endpoint

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.Endpoint.UpdateOneID(uuid.MustParse(in.GetID()))
		u = UpdateSet(u, in)
		info, err = u.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.EndpointUpdateOne, in *npool.EndpointReq) *ent.EndpointUpdateOne {
	if in.ChainType != nil {
		u.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		u.SetChainID(in.GetChainID())
	}
	if in.Address != nil {
		u.SetAddress(in.GetAddress())
	}
	if in.State != nil {
		u.SetState(in.GetState().String())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Endpoint, error) {
	var info *ent.Endpoint
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Endpoint.Query().Where(endpoint.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.EndpointQuery, error) {
	stm := cli.Endpoint.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.ID(id))
		default:
			return nil, fmt.Errorf("invalid endpoint field")
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
			stm.Where(endpoint.IDIn(ids...))
		}
	}
	if conds.ChainType != nil {
		switch conds.GetChainType().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.ChainType(conds.GetChainType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint field")
		}
	}
	if conds.ChainID != nil {
		switch conds.GetChainID().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.ChainID(conds.GetChainID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint field")
		}
	}
	if conds.Address != nil {
		switch conds.GetAddress().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.Address(conds.GetAddress().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint field")
		}
	}
	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.State(conds.GetState().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint field")
		}
	}
	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Endpoint, int, error) {
	var err error
	rows := []*ent.Endpoint{}
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
			Order(ent.Desc(endpoint.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.Endpoint, err error) {
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
		exist, err = cli.Endpoint.Query().Where(endpoint.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Endpoint, error) {
	var info *ent.Endpoint
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Endpoint.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
