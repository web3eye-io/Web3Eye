package block

import (
	"context"
	"fmt"

	blockcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/block"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	blockent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/block"
	blockproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

func (h *Handler) UpdateBlock(ctx context.Context) (*blockproto.Block, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Block.
			Query().
			Where(
				blockent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := blockcrud.UpdateSet(
			info.Update(),
			&blockcrud.Req{
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
		if _, err := stm.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetBlock(ctx)
}
