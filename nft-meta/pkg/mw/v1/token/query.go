package token

import (
	"context"
	"fmt"

	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	tokencrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	tokenent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

type queryHandler struct {
	*Handler
	stm   *ent.TokenSelect
	infos []*tokenproto.Token
	total uint32
}

func (h *queryHandler) selectToken(stm *ent.TokenQuery) {
	h.stm = stm.Select(
		tokenent.FieldID,
		tokenent.FieldEntID,
		tokenent.FieldChainType,
		tokenent.FieldChainID,
		tokenent.FieldContract,
		tokenent.FieldTokenType,
		tokenent.FieldTokenID,
		tokenent.FieldOwner,
		tokenent.FieldURI,
		tokenent.FieldURIType,
		tokenent.FieldImageURL,
		tokenent.FieldVideoURL,
		tokenent.FieldName,
		tokenent.FieldDescription,
		tokenent.FieldVectorState,
		tokenent.FieldVectorID,
		tokenent.FieldIpfsImageURL,
		tokenent.FieldImageSnapshotID,
		tokenent.FieldRemark,
		tokenent.FieldCreatedAt,
		tokenent.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.ChainType = basetype.ChainType(basetype.ChainType_value[info.ChainTypeStr])
		info.TokenType = basetype.TokenType(basetype.TokenType_value[info.TokenTypeStr])
		info.VectorState = tokenproto.ConvertState(tokenproto.ConvertState_value[info.VectorStateStr])
	}
}

func (h *queryHandler) queryToken(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Token.Query().Where(tokenent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(tokenent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(tokenent.EntID(*h.EntID))
	}
	h.selectToken(stm)
	return nil
}

func (h *queryHandler) queryTokens(ctx context.Context, cli *ent.Client) error {
	stm, err := tokencrud.SetQueryConds(cli.Token.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectToken(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetToken(ctx context.Context) (*tokenproto.Token, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryToken(cli); err != nil {
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

func (h *Handler) GetTokens(ctx context.Context) ([]*tokenproto.Token, uint32, error) {
	if h.Conds == nil {
		return nil, 0, fmt.Errorf("the conds is nil")
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTokens(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(tokenent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
