package synctask

import (
	"context"

	synctaskcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/synctask"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	synctaskent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/synctask"
)

func (h *Handler) ExistSyncTask(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			SyncTask.
			Query().
			Where(
				synctaskent.ID(*h.ID),
				synctaskent.DeletedAt(0),
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

func (h *Handler) ExistSyncTaskConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := synctaskcrud.SetQueryConds(cli.SyncTask.Query(), h.Conds)
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
