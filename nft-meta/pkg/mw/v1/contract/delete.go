package contract

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	contractproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
)

func (h *Handler) DeleteContract(ctx context.Context) (*contractproto.Contract, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}

	info, err := h.GetContract(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return tx.Contract.DeleteOne(&ent.Contract{ID: *h.ID}).Exec(ctx)
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
