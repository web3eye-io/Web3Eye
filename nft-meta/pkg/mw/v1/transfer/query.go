package transfer

import (
	"context"
	"fmt"

	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	transfercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/transfer"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	transferent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"
	transferproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

type queryHandler struct {
	*Handler
	stm   *ent.TransferSelect
	infos []*transferproto.Transfer
	total uint32
}

func (h *queryHandler) selectTransfer(stm *ent.TransferQuery) {
	h.stm = stm.Select(
		transferent.FieldID,
		transferent.FieldEntID,
		transferent.FieldChainType,
		transferent.FieldChainID,
		transferent.FieldContract,
		transferent.FieldTokenType,
		transferent.FieldTokenID,
		transferent.FieldFrom,
		transferent.FieldTo,
		transferent.FieldAmount,
		transferent.FieldBlockNumber,
		transferent.FieldTxHash,
		transferent.FieldBlockHash,
		transferent.FieldTxTime,
		transferent.FieldRemark,
		transferent.FieldCreatedAt,
		transferent.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.ChainType = basetype.ChainType(basetype.ChainType_value[info.ChainTypeStr])
		info.TokenType = basetype.TokenType(basetype.TokenType_value[info.TokenTypeStr])
	}
}

func (h *queryHandler) queryTransfer(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Transfer.Query().Where(transferent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(transferent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(transferent.EntID(*h.EntID))
	}
	h.selectTransfer(stm)
	return nil
}

func (h *queryHandler) queryTransfers(ctx context.Context, cli *ent.Client) error {
	stm, err := transfercrud.SetQueryConds(cli.Transfer.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectTransfer(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetTransfer(ctx context.Context) (*transferproto.Transfer, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTransfer(cli); err != nil {
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

func (h *Handler) GetTransfers(ctx context.Context) ([]*transferproto.Transfer, uint32, error) {
	if h.Conds == nil {
		return nil, 0, fmt.Errorf("the conds is nil")
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTransfers(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(transferent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
