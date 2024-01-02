package snapshot

import (
	"context"
	"fmt"

	snapshotcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/snapshot"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	snapshotent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/snapshot"
	snapshotproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

func (h *Handler) UpdateSnapshot(ctx context.Context) (*snapshotproto.Snapshot, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Snapshot.
			Query().
			Where(
				snapshotent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := snapshotcrud.UpdateSet(
			info.Update(),
			&snapshotcrud.Req{
				EntID:         h.EntID,
				Index:         h.Index,
				SnapshotCommP: h.SnapshotCommP,
				SnapshotRoot:  h.SnapshotRoot,
				SnapshotURI:   h.SnapshotURI,
				BackupState:   h.BackupState,
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

	return h.GetSnapshot(ctx)
}
