package transfer

import (
	"context"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func Create(ctx context.Context, in *npool.TransferReq) (*ent.Transfer, error) {
	var info *ent.Transfer
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Transfer.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.TransferCreate, in *npool.TransferReq) *ent.TransferCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
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
		c.SetTokenType(in.GetTokenType())
	}
	if in.TokenID != nil {
		c.SetTokenID(in.GetTokenID())
	}
	if in.From != nil {
		c.SetFrom(in.GetFrom())
	}
	if in.To != nil {
		c.SetTo(in.GetTo())
	}
	if in.BlockNumber != nil {
		c.SetBlockNumber(in.GetBlockNumber())
	}
	if in.Amount != nil {
		c.SetAmount(in.GetAmount())
	}
	if in.TxHash != nil {
		c.SetTxHash(in.GetTxHash())
	}
	if in.BlockHash != nil {
		c.SetBlockHash(in.GetBlockHash())
	}
	if in.TxTime != nil {
		c.SetTxTime(in.GetTxTime())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*npool.TransferReq) ([]*ent.Transfer, error) {
	var err error
	rows := []*ent.Transfer{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.TransferCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Transfer.Create(), info)
		}
		rows, err = tx.Transfer.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

//nolint:gocyclo
func Update(ctx context.Context, in *npool.TransferReq) (*ent.Transfer, error) {
	var err error
	var info *ent.Transfer

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.Transfer.UpdateOneID(uuid.MustParse(in.GetID()))
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
			u.SetTokenType(in.GetTokenType())
		}
		if in.TokenID != nil {
			u.SetTokenID(in.GetTokenID())
		}
		if in.From != nil {
			u.SetFrom(in.GetFrom())
		}
		if in.To != nil {
			u.SetTo(in.GetTo())
		}
		if in.BlockNumber != nil {
			u.SetBlockNumber(in.GetBlockNumber())
		}
		if in.Amount != nil {
			u.SetAmount(in.GetAmount())
		}
		if in.TxHash != nil {
			u.SetTxHash(in.GetTxHash())
		}
		if in.BlockHash != nil {
			u.SetBlockHash(in.GetBlockHash())
		}
		if in.TxTime != nil {
			u.SetTxTime(in.GetTxTime())
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

// func UpdateSet(u *ent.TransferUpdateOne, in *npool.TransferReq) *ent.TransferUpdateOne {
// 	if in.VectorID != nil {
// 		u.SetVectorID(in.GetVectorID())
// 	}
// 	if in.Remark != nil {
// 		u.SetRemark(in.GetRemark())
// 	}
// 	return u
// }

func Row(ctx context.Context, id uuid.UUID) (*ent.Transfer, error) {
	var info *ent.Transfer
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Transfer.Query().Where(transfer.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

// nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.TransferQuery, error) {
	stm := cli.Transfer.Query()
	if conds == nil {
		return stm, nil
	}

	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.ID(id))
		default:
			return nil, fmt.Errorf("invalid transfer field")
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
			stm.Where(transfer.IDIn(ids...))
		}
	}

	if conds.ChainType != nil {
		switch conds.GetChainType().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.ChainType(conds.GetChainType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.ChainID != nil {
		switch conds.GetChainID().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.ChainID(conds.GetChainID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.Contract != nil {
		switch conds.GetContract().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.Contract(conds.GetContract().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.TokenType != nil {
		switch conds.GetTokenType().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.TokenType(conds.GetTokenType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.TokenID != nil {
		switch conds.GetTokenID().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.TokenID(conds.GetTokenID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.TokenID != nil {
		switch conds.GetTokenID().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.TokenID(conds.GetTokenID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.From != nil {
		switch conds.GetFrom().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.From(conds.GetFrom().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.To != nil {
		switch conds.GetTo().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.To(conds.GetTo().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.Amount != nil {
		switch conds.GetAmount().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.Amount(conds.GetAmount().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.BlockNumber != nil {
		switch conds.GetBlockNumber().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.BlockNumber(conds.GetBlockNumber().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.TxHash != nil {
		switch conds.GetTxHash().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.TxHash(conds.GetTxHash().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.BlockHash != nil {
		switch conds.GetBlockHash().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.BlockHash(conds.GetBlockHash().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.TxTime != nil {
		switch conds.GetTxTime().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.TxTime(conds.GetTxTime().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}
	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(transfer.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid transfer field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Transfer, int, error) {
	var err error
	rows := []*ent.Transfer{}
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
			Order(ent.Desc(transfer.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (info *ent.Transfer, err error) {
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
		exist, err = cli.Transfer.Query().Where(transfer.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Transfer, error) {
	var info *ent.Transfer
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Transfer.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
