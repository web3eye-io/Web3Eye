package synctask

import (
	"context"
	"fmt"

	synctaskcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/synctask"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	synctaskent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/synctask"
	synctaskproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
)

func (h *Handler) UpdateSyncTask(ctx context.Context) (*synctaskproto.SyncTask, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			SyncTask.
			Query().
			Where(
				synctaskent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := synctaskcrud.UpdateSet(
			info.Update(),
			&synctaskcrud.Req{
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				Start:       h.Start,
				End:         h.End,
				Current:     h.Current,
				Topic:       h.Topic,
				SyncState:   h.SyncState,
				Description: h.Description,
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

	return h.GetSyncTask(ctx)
}
