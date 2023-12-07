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

type queryHandler struct {
	*Handler
	stm   *ent.SnapshotSelect
	infos []*snapshotproto.Snapshot
	total uint32
}

func (h *queryHandler) selectSnapshot(stm *ent.SnapshotQuery) {
	h.stm = stm.Select(
		snapshotent.FieldID,
		snapshotent.FieldEntID,
		snapshotent.FieldIndex,
		snapshotent.FieldSnapshotCommP,
		snapshotent.FieldSnapshotRoot,
		snapshotent.FieldSnapshotURI,
		snapshotent.FieldBackupState,
		snapshotent.FieldCreatedAt,
		snapshotent.FieldUpdatedAt,
	)
}

func (h *queryHandler) querySnapshot(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Snapshot.Query().Where(snapshotent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(snapshotent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(snapshotent.EntID(*h.EntID))
	}
	h.selectSnapshot(stm)
	return nil
}

func (h *queryHandler) querySnapshots(ctx context.Context, cli *ent.Client) error {
	stm, err := snapshotcrud.SetQueryConds(cli.Snapshot.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectSnapshot(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetSnapshot(ctx context.Context) (*snapshotproto.Snapshot, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySnapshot(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return handler.infos[0], nil
}

func (h *Handler) GetSnapshots(ctx context.Context) ([]*snapshotproto.Snapshot, uint32, error) {
	if h.Conds == nil {
		return nil, 0, fmt.Errorf("the conds is nil")
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySnapshots(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(snapshotent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	return handler.infos, handler.total, nil
}
