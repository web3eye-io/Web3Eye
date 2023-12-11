package contract

import (
	"context"

	contractcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	contractent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	contractproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
)

//nolint:dupl
func (h *Handler) UpsertContract(ctx context.Context) (*contractproto.Contract, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		row, _ := tx.Contract.Query().Where(
			contractent.ChainType(h.ChainType.String()),
			contractent.ChainID(*h.ChainID),
			contractent.Address(*h.Address),
		).Only(ctx)
		if row == nil {
			info, err := contractcrud.CreateSet(tx.Contract.Create(),
				&contractcrud.Req{
					EntID:       h.EntID,
					ChainType:   h.ChainType,
					ChainID:     h.ChainID,
					Address:     h.Address,
					Name:        h.Name,
					Symbol:      h.Symbol,
					Decimals:    h.Decimals,
					Creator:     h.Creator,
					BlockNum:    h.BlockNum,
					TxHash:      h.TxHash,
					TxTime:      h.TxTime,
					ProfileURL:  h.ProfileURL,
					BaseURL:     h.BaseURL,
					BannerURL:   h.BannerURL,
					Description: h.Description,
					Remark:      h.Remark,
				},
			).Save(ctx)
			if err != nil {
				return err
			}
			h.EntID = &info.EntID
			h.ID = &info.ID
			return nil
		}
		stm, err := contractcrud.UpdateSet(
			row.Update(),
			&contractcrud.Req{
				EntID:       &row.EntID,
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				Address:     h.Address,
				Name:        h.Name,
				Symbol:      h.Symbol,
				Decimals:    h.Decimals,
				Creator:     h.Creator,
				BlockNum:    h.BlockNum,
				TxHash:      h.TxHash,
				TxTime:      h.TxTime,
				ProfileURL:  h.ProfileURL,
				BaseURL:     h.BaseURL,
				BannerURL:   h.BannerURL,
				Description: h.Description,
				Remark:      h.Remark,
			},
		)
		if err != nil {
			return err
		}
		if info, err := stm.Save(_ctx); err == nil {
			h.EntID = &info.EntID
			h.ID = &info.ID
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return h.GetContract(ctx)
}
