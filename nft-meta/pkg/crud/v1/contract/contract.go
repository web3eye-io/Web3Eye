package contract

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
)

func Create(ctx context.Context, in *npool.ContractReq) (*ent.Contract, error) {
	var info *ent.Contract
	var err error

	if in == nil {
		return nil, errors.New("input is nil")
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		c := tx.Contract.Create()
		info, err = CreateSet(c, in).Save(ctx)
		return err
	})

	if err != nil {
		return nil, err
	}

	return info, nil
}

func Upsert(ctx context.Context, in *npool.ContractReq) (*ent.Contract, error) {
	if in == nil {
		return nil, errors.New("input is nil")
	}
	var info *ent.Contract
	var err error
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		row, _ := tx.Contract.Query().Where(
			contract.ChainType(in.GetChainType().String()),
			contract.ChainID(in.GetChainID()),
			contract.Address(in.GetAddress()),
		).Only(ctx)
		if row == nil {
			info, err = CreateSet(tx.Contract.Create(), in).Save(ctx)
			return err
		}
		info, err = UpdateSet(tx.Contract.UpdateOneID(row.ID), in).Save(ctx)
		return err
	})
	return info, err
}

// nolint
func CreateSet(c *ent.ContractCreate, in *npool.ContractReq) *ent.ContractCreate {
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
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.Symbol != nil {
		c.SetSymbol(in.GetSymbol())
	}
	if in.Decimals != nil {
		c.SetDecimals(in.GetDecimals())
	}
	if in.Creator != nil {
		c.SetCreator(in.GetCreator())
	}
	if in.BlockNum != nil {
		c.SetBlockNum(in.GetBlockNum())
	}
	if in.TxHash != nil {
		c.SetTxHash(in.GetTxHash())
	}
	if in.TxTime != nil {
		c.SetTxTime(in.GetTxTime())
	}
	if in.ProfileURL != nil {
		c.SetProfileURL(in.GetProfileURL())
	}
	if in.BaseURL != nil {
		c.SetBaseURL(in.GetBaseURL())
	}
	if in.BannerURL != nil {
		c.SetBannerURL(in.GetBannerURL())
	}
	if in.Description != nil {
		c.SetDescription(in.GetDescription())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*npool.ContractReq) ([]*ent.Contract, error) {
	var err error
	rows := []*ent.Contract{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.ContractCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Contract.Create(), info)
		}
		return tx.Contract.CreateBulk(bulk...).OnConflict().UpdateNewValues().Exec(ctx)
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// nolint
func Update(ctx context.Context, in *npool.ContractReq) (*ent.Contract, error) {
	if in == nil {
		return nil, errors.New("input is nil")
	}

	var info *ent.Contract
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, err
	}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		contractUpdate := tx.Contract.UpdateOneID(id)
		info, err = UpdateSet(contractUpdate, in).Save(ctx)
		return err
	})

	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.ContractUpdateOne, in *npool.ContractReq) *ent.ContractUpdateOne {
	if in.ChainType != nil {
		u.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		u.SetChainID(in.GetChainID())
	}
	if in.Address != nil {
		u.SetAddress(in.GetAddress())
	}
	if in.Name != nil {
		u.SetName(in.GetName())
	}
	if in.Symbol != nil {
		u.SetSymbol(in.GetSymbol())
	}
	if in.Decimals != nil {
		u.SetDecimals(in.GetDecimals())
	}
	if in.Creator != nil {
		u.SetCreator(in.GetCreator())
	}
	if in.BlockNum != nil {
		u.SetBlockNum(in.GetBlockNum())
	}
	if in.TxHash != nil {
		u.SetTxHash(in.GetTxHash())
	}
	if in.TxTime != nil {
		u.SetTxTime(in.GetTxTime())
	}
	if in.ProfileURL != nil {
		u.SetProfileURL(in.GetProfileURL())
	}
	if in.BaseURL != nil {
		u.SetBaseURL(in.GetBaseURL())
	}
	if in.BannerURL != nil {
		u.SetBannerURL(in.GetBannerURL())
	}
	if in.Description != nil {
		u.SetDescription(in.GetDescription())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Contract, error) {
	var info *ent.Contract
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Contract.Query().Where(contract.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.ContractQuery, error) {
	stm := cli.Contract.Query()
	if conds == nil {
		return stm, nil
	}

	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(contract.ID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
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
			stm.Where(contract.IDIn(ids...))
		}
	}

	if conds.ChainType != nil {
		switch conds.GetChainType().GetOp() {
		case cruder.EQ:
			stm.Where(contract.ChainType(conds.GetChainType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}

	if conds.ChainID != nil {
		switch conds.GetChainID().GetOp() {
		case cruder.EQ:
			stm.Where(contract.ChainID(conds.GetChainID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}

	if conds.Address != nil {
		switch conds.GetAddress().GetOp() {
		case cruder.EQ:
			stm.Where(contract.Address(conds.GetAddress().GetValue()))
		default:
			return nil, fmt.Errorf("invalid Address field")
		}
	}
	if conds.Name != nil {
		switch conds.GetName().GetOp() {
		case cruder.EQ:
			stm.Where(contract.Name(conds.GetName().GetValue()))
		default:
			return nil, fmt.Errorf("invalid Name field")
		}
	}
	if conds.Symbol != nil {
		switch conds.GetSymbol().GetOp() {
		case cruder.EQ:
			stm.Where(contract.Symbol(conds.GetSymbol().GetValue()))
		default:
			return nil, fmt.Errorf("invalid Symbol field")
		}
	}

	if conds.Decimals != nil {
		switch conds.GetDecimals().GetOp() {
		case cruder.EQ:
			stm.Where(contract.Decimals(conds.GetDecimals().GetValue()))
		default:
			return nil, fmt.Errorf("invalid Decimals field")
		}
	}

	if conds.Creator != nil {
		switch conds.GetCreator().GetOp() {
		case cruder.EQ:
			stm.Where(contract.Creator(conds.GetCreator().GetValue()))
		default:
			return nil, fmt.Errorf("invalid Creator field")
		}
	}
	if conds.BlockNum != nil {
		switch conds.GetBlockNum().GetOp() {
		case cruder.EQ:
			stm.Where(contract.BlockNum(conds.GetBlockNum().GetValue()))
		default:
			return nil, fmt.Errorf("invalid BlockNum field")
		}
	}
	if conds.TxHash != nil {
		switch conds.GetTxHash().GetOp() {
		case cruder.EQ:
			stm.Where(contract.TxHash(conds.GetTxHash().GetValue()))
		default:
			return nil, fmt.Errorf("invalid TxHash field")
		}
	}
	if conds.TxTime != nil {
		switch conds.GetTxTime().GetOp() {
		case cruder.EQ:
			stm.Where(contract.TxTime(conds.GetTxTime().GetValue()))
		default:
			return nil, fmt.Errorf("invalid TxTime field")
		}
	}
	if conds.ProfileURL != nil {
		switch conds.GetProfileURL().GetOp() {
		case cruder.EQ:
			stm.Where(contract.ProfileURL(conds.GetProfileURL().GetValue()))
		default:
			return nil, fmt.Errorf("invalid ProfileURL field")
		}
	}
	if conds.BaseURL != nil {
		switch conds.GetBaseURL().GetOp() {
		case cruder.EQ:
			stm.Where(contract.BaseURL(conds.GetBaseURL().GetValue()))
		default:
			return nil, fmt.Errorf("invalid BaseURL field")
		}
	}
	if conds.BannerURL != nil {
		switch conds.GetBannerURL().GetOp() {
		case cruder.EQ:
			stm.Where(contract.BannerURL(conds.GetBannerURL().GetValue()))
		default:
			return nil, fmt.Errorf("invalid BannerURL field")
		}
	}
	if conds.Description != nil {
		switch conds.GetDescription().GetOp() {
		case cruder.EQ:
			stm.Where(contract.Description(conds.GetDescription().GetValue()))
		default:
			return nil, fmt.Errorf("invalid Description field")
		}
	}

	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(contract.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Contract, int, error) {
	var err error
	rows := []*ent.Contract{}
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
			Order(ent.Desc(contract.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.Contract, err error) {
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
		exist, err = cli.Contract.Query().Where(contract.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Contract, error) {
	var info *ent.Contract
	var err error

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Contract.UpdateOneID(id).SetDeletedAt(uint32(time.Now().Unix())).Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
