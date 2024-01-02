package snapshot

import (
	"context"

	snapshotcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/snapshot"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	snapshotproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateSnapshot(ctx context.Context) (*snapshotproto.Snapshot, error) {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := snapshotcrud.CreateSet(
			cli.Snapshot.Create(),
			&snapshotcrud.Req{
				EntID:         h.EntID,
				Index:         h.Index,
				SnapshotCommP: h.SnapshotCommP,
				SnapshotRoot:  h.SnapshotRoot,
				SnapshotURI:   h.SnapshotURI,
				BackupState:   h.BackupState,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSnapshot(ctx)
}

func (h *Handler) CreateSnapshots(ctx context.Context) ([]*snapshotproto.Snapshot, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := snapshotcrud.CreateSet(tx.Snapshot.Create(), req).Save(_ctx)
			if err != nil {
				return err
			}
			ids = append(ids, info.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &snapshotcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetSnapshots(ctx)
	return infos, err
}
