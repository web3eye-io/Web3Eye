package block

import (
	"context"

	blockcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/block"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	blockent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/block"
	blockproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

func (h *Handler) UpsertBlock(ctx context.Context) (*blockproto.Block, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		row, _ := tx.Block.Query().Where(
			blockent.ChainType(h.ChainType.String()),
			blockent.ChainID(*h.ChainID),
			blockent.BlockNumber(*h.BlockNumber),
		).Only(ctx)
		if row == nil {
			info, err := blockcrud.CreateSet(tx.Block.Create(),
				&blockcrud.Req{
					EntID:       h.EntID,
					ChainType:   h.ChainType,
					ChainID:     h.ChainID,
					BlockNumber: h.BlockNumber,
					BlockHash:   h.BlockHash,
					BlockTime:   h.BlockTime,
					ParseState:  h.ParseState,
					Remark:      h.Remark,
				},
			).Save(ctx)
			if err == nil {
				h.EntID = &info.EntID
				h.ID = &info.ID
			}
			return err
		}

		stm, err := blockcrud.UpdateSet(
			row.Update(),
			&blockcrud.Req{
				EntID:       &row.EntID,
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				BlockNumber: h.BlockNumber,
				BlockHash:   h.BlockHash,
				BlockTime:   h.BlockTime,
				ParseState:  h.ParseState,
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

	return h.GetBlock(ctx)
}
