package snapshot

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	snapshotproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

func (h *Handler) DeleteSnapshot(ctx context.Context) (*snapshotproto.Snapshot, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}

	info, err := h.GetSnapshot(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return tx.Snapshot.DeleteOne(&ent.Snapshot{ID: *h.ID}).Exec(ctx)
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
