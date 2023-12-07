package contract

import (
	"context"
	"fmt"

	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	contractcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	contractent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	contractproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
)

type queryHandler struct {
	*Handler
	stm   *ent.ContractSelect
	infos []*contractproto.Contract
	total uint32
}

func (h *queryHandler) selectContract(stm *ent.ContractQuery) {
	h.stm = stm.Select(
		contractent.FieldID,
		contractent.FieldEntID,
		contractent.FieldChainType,
		contractent.FieldChainID,
		contractent.FieldAddress,
		contractent.FieldName,
		contractent.FieldSymbol,
		contractent.FieldDecimals,
		contractent.FieldCreator,
		contractent.FieldBlockNum,
		contractent.FieldTxHash,
		contractent.FieldTxTime,
		contractent.FieldProfileURL,
		contractent.FieldBaseURL,
		contractent.FieldBannerURL,
		contractent.FieldDescription,
		contractent.FieldRemark,
		contractent.FieldCreatedAt,
		contractent.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.ChainType = basetype.ChainType(basetype.ChainType_value[info.ChainTypeStr])
	}
}

func (h *queryHandler) queryContract(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Contract.Query().Where(contractent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(contractent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(contractent.EntID(*h.EntID))
	}
	h.selectContract(stm)
	return nil
}

func (h *queryHandler) queryContracts(ctx context.Context, cli *ent.Client) error {
	stm, err := contractcrud.SetQueryConds(cli.Contract.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectContract(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetContract(ctx context.Context) (*contractproto.Contract, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryContract(cli); err != nil {
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

func (h *Handler) GetContracts(ctx context.Context) ([]*contractproto.Contract, uint32, error) {
	if h.Conds == nil {
		return nil, 0, fmt.Errorf("the conds is nil")
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryContracts(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(contractent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
