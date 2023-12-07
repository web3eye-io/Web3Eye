package endpoint

import (
	"context"
	"fmt"

	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	endpointcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/endpoint"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	endpointent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/endpoint"
	endpointproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
)

type queryHandler struct {
	*Handler
	stm   *ent.EndpointSelect
	infos []*endpointproto.Endpoint
	total uint32
}

func (h *queryHandler) selectEndpoint(stm *ent.EndpointQuery) {
	h.stm = stm.Select(
		endpointent.FieldID,
		endpointent.FieldEntID,
		endpointent.FieldChainType,
		endpointent.FieldChainID,
		endpointent.FieldAddress,
		endpointent.FieldState,
		endpointent.FieldRemark,
		endpointent.FieldCreatedAt,
		endpointent.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.ChainType = basetype.ChainType(basetype.ChainType_value[info.ChainTypeStr])
	}
}

func (h *queryHandler) queryEndpoint(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Endpoint.Query().Where(endpointent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(endpointent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(endpointent.EntID(*h.EntID))
	}
	h.selectEndpoint(stm)
	return nil
}

func (h *queryHandler) queryEndpoints(ctx context.Context, cli *ent.Client) error {
	stm, err := endpointcrud.SetQueryConds(cli.Endpoint.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectEndpoint(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetEndpoint(ctx context.Context) (*endpointproto.Endpoint, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEndpoint(cli); err != nil {
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

func (h *Handler) GetEndpoints(ctx context.Context) ([]*endpointproto.Endpoint, uint32, error) {
	if h.Conds == nil {
		return nil, 0, fmt.Errorf("the conds is nil")
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEndpoints(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(endpointent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
