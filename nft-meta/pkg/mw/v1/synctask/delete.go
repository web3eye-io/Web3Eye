package synctask

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	synctaskproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
)

func (h *Handler) DeleteSyncTask(ctx context.Context) (*synctaskproto.SyncTask, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}

	info, err := h.GetSyncTask(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return tx.SyncTask.DeleteOne(&ent.SyncTask{ID: *h.ID}).Exec(ctx)
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
