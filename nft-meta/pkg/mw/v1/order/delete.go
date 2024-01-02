package order

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	orderproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

func (h *Handler) DeleteOrder(ctx context.Context) (*orderproto.Order, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}

	info, err := h.GetOrder(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return tx.Order.DeleteOne(&ent.Order{ID: *h.ID}).Exec(ctx)
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
