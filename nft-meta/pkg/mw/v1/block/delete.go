package block

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	blockproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

func (h *Handler) DeleteBlock(ctx context.Context) (*blockproto.Block, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}

	info, err := h.GetBlock(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return tx.Block.DeleteOne(&ent.Block{ID: *h.ID}).Exec(ctx)
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
