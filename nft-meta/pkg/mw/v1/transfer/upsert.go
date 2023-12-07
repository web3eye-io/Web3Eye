package transfer

import (
	"context"

	transfercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/transfer"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	transferent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"
	transferproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func (h *Handler) UpsertTransfer(ctx context.Context) (*transferproto.Transfer, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		row, _ := tx.Transfer.Query().Where(
			transferent.Contract(*h.Contract),
			transferent.TokenID(*h.TokenID),
			transferent.TxHash(*h.TxHash),
			transferent.From(*h.From),
		).Only(ctx)
		if row == nil {
			info, err := transfercrud.CreateSet(tx.Transfer.Create(),
				&transfercrud.Req{
					EntID:       h.EntID,
					ChainType:   h.ChainType,
					ChainID:     h.ChainID,
					Contract:    h.Contract,
					TokenType:   h.TokenType,
					TokenID:     h.TokenID,
					From:        h.From,
					To:          h.To,
					Amount:      h.Amount,
					BlockNumber: h.BlockNumber,
					TxHash:      h.TxHash,
					BlockHash:   h.BlockHash,
					TxTime:      h.TxTime,
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
		stm, err := transfercrud.UpdateSet(
			row.Update(),
			&transfercrud.Req{
				EntID:       &row.EntID,
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				Contract:    h.Contract,
				TokenType:   h.TokenType,
				TokenID:     h.TokenID,
				From:        h.From,
				To:          h.To,
				Amount:      h.Amount,
				BlockNumber: h.BlockNumber,
				TxHash:      h.TxHash,
				BlockHash:   h.BlockHash,
				TxTime:      h.TxTime,
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

	return h.GetTransfer(ctx)
}
