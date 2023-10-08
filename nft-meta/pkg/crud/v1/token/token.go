package token

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

func Create(ctx context.Context, in *npool.TokenReq) (*ent.Token, error) {
	var info *ent.Token
	var err error

	if in == nil {
		return nil, errors.New("input is nil")
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		c := tx.Token.Create()
		info, err = CreateSet(c, in).Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Upsert(ctx context.Context, in *npool.TokenReq) (*ent.Token, error) {
	if in == nil {
		return nil, errors.New("input is nil")
	}
	var info *ent.Token
	var err error
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		row, _ := tx.Token.Query().Where(
			token.Contract(in.GetContract()),
			token.TokenID(in.GetTokenID()),
		).Only(ctx)
		if row == nil {
			info, err = CreateSet(tx.Token.Create(), in).Save(ctx)
			return err
		}
		info, err = UpdateSet(tx.Token.UpdateOneID(row.ID), in).Save(ctx)
		return err
	})
	return info, err
}

//nolint:gocyclo
func CreateSet(c *ent.TokenCreate, in *npool.TokenReq) *ent.TokenCreate {
	id, err := uuid.Parse(*in.ID)
	if err != nil {
		c.SetID(uuid.New())
	} else {
		c.SetID(id)
	}
	if in.ChainType != nil {
		c.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		c.SetChainID(in.GetChainID())
	}
	if in.Contract != nil {
		c.SetContract(in.GetContract())
	}
	if in.TokenType != nil {
		c.SetTokenType(in.GetTokenType().String())
	}
	if in.TokenID != nil {
		c.SetTokenID(in.GetTokenID())
	}
	if in.Owner != nil {
		c.SetOwner(in.GetOwner())
	}
	if in.URI != nil {
		c.SetURI(in.GetURI())
	}
	if in.URIType != nil {
		c.SetURIType(in.GetURIType())
	}
	if in.ImageURL != nil {
		c.SetImageURL(in.GetImageURL())
	}
	if in.VideoURL != nil {
		c.SetVideoURL(in.GetVideoURL())
	}
	if in.Description != nil {
		c.SetDescription(in.GetDescription())
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.VectorState != nil {
		c.SetVectorState(in.GetVectorState().String())
	} else {
		c.SetVectorState(npool.ConvertState_Default.String())
	}
	if in.VectorID != nil {
		c.SetVectorID(in.GetVectorID())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	if in.IPFSImageURL != nil {
		c.SetIpfsImageURL(in.GetIPFSImageURL())
	}
	if in.ImageSnapshotID != nil {
		c.SetImageSnapshotID(in.GetImageSnapshotID())
	}

	return c
}

func CreateBulk(ctx context.Context, in []*npool.TokenReq) ([]*ent.Token, error) {
	var err error
	rows := []*ent.Token{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.TokenCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Token.Create(), info)
		}
		rows, err = tx.Token.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.TokenReq) (*ent.Token, error) {
	if in == nil {
		return nil, errors.New("input is nil")
	}
	var err error
	var info *ent.Token
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, err
	}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		u := tx.Token.UpdateOneID(id)
		info, err = UpdateSet(u, in).Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:gocyclo
func UpdateSet(u *ent.TokenUpdateOne, in *npool.TokenReq) *ent.TokenUpdateOne {
	if in.ChainType != nil {
		u.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		u.SetChainID(in.GetChainID())
	}
	if in.Contract != nil {
		u.SetContract(in.GetContract())
	}
	if in.TokenType != nil {
		u.SetTokenType(in.GetTokenType().String())
	}
	if in.TokenID != nil {
		u.SetTokenID(in.GetTokenID())
	}
	if in.Owner != nil {
		u.SetOwner(in.GetOwner())
	}
	if in.URI != nil {
		u.SetURI(in.GetURI())
	}
	if in.URIType != nil {
		u.SetURIType(in.GetURIType())
	}
	if in.ImageURL != nil {
		u.SetImageURL(in.GetImageURL())
	}
	if in.VideoURL != nil {
		u.SetVideoURL(in.GetVideoURL())
	}
	if in.Description != nil {
		u.SetDescription(in.GetDescription())
	}
	if in.Name != nil {
		u.SetName(in.GetName())
	}
	if in.VectorState != nil {
		u.SetVectorState(in.GetVectorState().String())
	}
	if in.VectorID != nil {
		u.SetVectorID(in.GetVectorID())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	if in.IPFSImageURL != nil {
		u.SetIpfsImageURL(in.GetIPFSImageURL())
	}
	if in.ImageSnapshotID != nil {
		u.SetImageSnapshotID(in.GetImageSnapshotID())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Token, error) {
	var info *ent.Token
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Token.Query().Where(token.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:funlen,gocyclo
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.TokenQuery, error) {
	stm := cli.Token.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(token.ID(id))
		default:
			return nil, fmt.Errorf("invalid token field")
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
			stm.Where(token.IDIn(ids...))
		}
	}
	if conds.VectorIDs != nil {
		vIDs := []int64{}
		if conds.GetVectorIDs().GetOp() == cruder.IN {
			vIDs = append(vIDs, conds.GetVectorIDs().GetValue()...)
			stm.Where(token.VectorIDIn(vIDs...))
		}
	}
	if conds.ChainType != nil {
		switch conds.GetChainType().GetOp() {
		case cruder.EQ:
			stm.Where(token.ChainType(conds.GetChainType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.ChainID != nil {
		switch conds.GetChainID().GetOp() {
		case cruder.EQ:
			stm.Where(token.ChainID(conds.GetChainID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.Contract != nil {
		switch conds.GetContract().GetOp() {
		case cruder.EQ:
			stm.Where(token.Contract(conds.GetContract().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.TokenType != nil {
		switch conds.GetTokenType().GetOp() {
		case cruder.EQ:
			stm.Where(token.TokenType(conds.GetTokenType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.TokenID != nil {
		switch conds.GetTokenID().GetOp() {
		case cruder.EQ:
			stm.Where(token.TokenID(conds.GetTokenID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.Owner != nil {
		switch conds.GetOwner().GetOp() {
		case cruder.EQ:
			stm.Where(token.Owner(conds.GetOwner().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.URI != nil {
		switch conds.GetURI().GetOp() {
		case cruder.EQ:
			stm.Where(token.URI(conds.GetURI().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}

	if conds.URIType != nil {
		switch conds.GetURIType().GetOp() {
		case cruder.EQ:
			stm.Where(token.URIType(conds.GetURIType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.ImageURL != nil {
		switch conds.GetImageURL().GetOp() {
		case cruder.EQ:
			stm.Where(token.ImageURL(conds.GetImageURL().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.VideoURL != nil {
		switch conds.GetVideoURL().GetOp() {
		case cruder.EQ:
			stm.Where(token.VideoURL(conds.GetVideoURL().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.Description != nil {
		switch conds.GetDescription().GetOp() {
		case cruder.EQ:
			stm.Where(token.Description(conds.GetDescription().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.Name != nil {
		switch conds.GetName().GetOp() {
		case cruder.EQ:
			stm.Where(token.Name(conds.GetName().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.VectorState != nil {
		switch conds.GetVectorState().GetOp() {
		case cruder.EQ:
			stm.Where(token.VectorState(conds.GetVectorState().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.VectorID != nil {
		switch conds.GetVectorID().GetOp() {
		case cruder.EQ:
			stm.Where(token.VectorID(conds.GetVectorID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(token.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.IPFSImageURL != nil {
		switch conds.GetIPFSImageURL().GetOp() {
		case cruder.EQ:
			stm.Where(token.IpfsImageURL(conds.GetIPFSImageURL().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}
	if conds.ImageSnapshotID != nil {
		switch conds.GetImageSnapshotID().GetOp() {
		case cruder.EQ:
			stm.Where(token.ImageSnapshotID(conds.GetImageSnapshotID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid token field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Token, int, error) {
	var err error
	rows := []*ent.Token{}
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
			Order(ent.Desc(token.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.Token, err error) {
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
		exist, err = cli.Token.Query().Where(token.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Token, error) {
	var info *ent.Token
	var err error

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Token.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
