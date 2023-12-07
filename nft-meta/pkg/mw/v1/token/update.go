package token

import (
	"context"
	"fmt"

	tokencrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	tokenent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

func (h *Handler) UpdateToken(ctx context.Context) (*tokenproto.Token, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Token.
			Query().
			Where(
				tokenent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := tokencrud.UpdateSet(
			info.Update(),
			&tokencrud.Req{
				ChainType:       h.ChainType,
				ChainID:         h.ChainID,
				Contract:        h.Contract,
				TokenType:       h.TokenType,
				TokenID:         h.TokenID,
				Owner:           h.Owner,
				URI:             h.URI,
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
		)
		if err != nil {
			return err
		}
		if _, err := stm.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetToken(ctx)
}
