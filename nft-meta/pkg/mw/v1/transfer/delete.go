package transfer

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	transferproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func (h *Handler) DeleteTransfer(ctx context.Context) (*transferproto.Transfer, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}

	info, err := h.GetTransfer(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return tx.Transfer.DeleteOne(&ent.Transfer{ID: *h.ID}).Exec(ctx)
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
