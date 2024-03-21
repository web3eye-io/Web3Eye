package token

import (
	"context"

	tokencrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateToken(ctx context.Context) (*tokenproto.Token, error) {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := tokencrud.CreateSet(
			cli.Token.Create(),
			&tokencrud.Req{
				EntID:           h.EntID,
				ChainType:       h.ChainType,
				ChainID:         h.ChainID,
				Contract:        h.Contract,
				TokenType:       h.TokenType,
				TokenID:         h.TokenID,
				Owner:           h.Owner,
				URI:             h.URI,
				URIState:        h.URIState,
				URIType:         h.URIType,
				ImageURL:        h.ImageURL,
				VideoURL:        h.VideoURL,
				Name:            h.Name,
				Description:     h.Description,
				VectorState:     h.VectorState,
				VectorID:        h.VectorID,
				IPFSImageURL:    h.IPFSImageURL,
				ImageSnapshotID: h.ImageSnapshotID,
				Remark:          h.Remark,
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

	return h.GetToken(ctx)
}

func (h *Handler) CreateTokens(ctx context.Context) ([]*tokenproto.Token, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := tokencrud.CreateSet(tx.Token.Create(), req).Save(_ctx)
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

	h.Conds = &tokencrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetTokens(ctx)
	return infos, err
}
