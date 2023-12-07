package transfer

import (
	"context"
	"fmt"

	transfercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/transfer"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	transferent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"
	transferproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func (h *Handler) UpdateTransfer(ctx context.Context) (*transferproto.Transfer, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Transfer.
			Query().
			Where(
				transferent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := transfercrud.UpdateSet(
			info.Update(),
			&transfercrud.Req{
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
		if _, err := stm.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTransfer(ctx)
}
