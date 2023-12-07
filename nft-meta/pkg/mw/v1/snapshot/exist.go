package snapshot

import (
	"context"

	snapshotcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/snapshot"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	snapshotent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/snapshot"
)

func (h *Handler) ExistSnapshot(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Snapshot.
			Query().
			Where(
				snapshotent.EntID(*h.EntID),
				snapshotent.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistSnapshotConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := snapshotcrud.SetQueryConds(cli.Snapshot.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
