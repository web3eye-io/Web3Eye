package block

import (
	"context"
	"fmt"

	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	blockcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/block"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	blockent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/block"
	blockproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

type queryHandler struct {
	*Handler
	stm   *ent.BlockSelect
	infos []*blockproto.Block
	total uint32
}

func (h *queryHandler) selectBlock(stm *ent.BlockQuery) {
	h.stm = stm.Select(
		blockent.FieldID,
		blockent.FieldEntID,
		blockent.FieldChainType,
		blockent.FieldChainID,
		blockent.FieldBlockNumber,
		blockent.FieldBlockHash,
		blockent.FieldBlockTime,
		blockent.FieldParseState,
		blockent.FieldRemark,
		blockent.FieldCreatedAt,
		blockent.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.ChainType = basetype.ChainType(basetype.ChainType_value[info.ChainTypeStr])
		info.ParseState = basetype.BlockParseState(basetype.BlockParseState_value[info.ParseStateStr])
	}
}

func (h *queryHandler) queryBlock(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Block.Query().Where(blockent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(blockent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(blockent.EntID(*h.EntID))
	}
	h.selectBlock(stm)
	return nil
}

func (h *queryHandler) queryBlocks(ctx context.Context, cli *ent.Client) error {
	stm, err := blockcrud.SetQueryConds(cli.Block.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectBlock(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetBlock(ctx context.Context) (*blockproto.Block, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryBlock(cli); err != nil {
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

func (h *Handler) GetBlocks(ctx context.Context) ([]*blockproto.Block, uint32, error) {
	if h.Conds == nil {
		return nil, 0, fmt.Errorf("the conds is nil")
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryBlocks(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(blockent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
