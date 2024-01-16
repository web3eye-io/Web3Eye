package synctask

import (
	"context"
	"fmt"

	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	synctaskcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/synctask"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	synctaskent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/synctask"
	synctaskproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
)

type queryHandler struct {
	*Handler
	stm   *ent.SyncTaskSelect
	infos []*synctaskproto.SyncTask
	total uint32
}

func (h *queryHandler) selectSyncTask(stm *ent.SyncTaskQuery) {
	h.stm = stm.Select(
		synctaskent.FieldID,
		synctaskent.FieldEntID,
		synctaskent.FieldChainType,
		synctaskent.FieldChainID,
		synctaskent.FieldStart,
		synctaskent.FieldEnd,
		synctaskent.FieldCurrent,
		synctaskent.FieldTopic,
		synctaskent.FieldSyncState,
		synctaskent.FieldDescription,
		synctaskent.FieldRemark,
		synctaskent.FieldCreatedAt,
		synctaskent.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.ChainType = basetype.ChainType(basetype.ChainType_value[info.ChainTypeStr])
		info.SyncState = basetype.SyncState(basetype.SyncState_value[info.SyncStateStr])
	}
}

func (h *queryHandler) querySyncTask(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.SyncTask.Query().Where(synctaskent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(synctaskent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(synctaskent.EntID(*h.EntID))
	}
	h.selectSyncTask(stm)
	return nil
}

func (h *queryHandler) querySyncTasks(ctx context.Context, cli *ent.Client) error {
	stm, err := synctaskcrud.SetQueryConds(cli.SyncTask.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectSyncTask(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetSyncTask(ctx context.Context) (*synctaskproto.SyncTask, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySyncTask(cli); err != nil {
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
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetSyncTasks(ctx context.Context) ([]*synctaskproto.SyncTask, uint32, error) {
	if h.Conds == nil {
		return nil, 0, fmt.Errorf("the conds is nil")
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySyncTasks(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(synctaskent.FieldID))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
