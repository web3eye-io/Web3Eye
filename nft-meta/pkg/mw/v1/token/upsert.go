package token

import (
	"context"

	tokencrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	tokenent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

func (h *Handler) UpsertToken(ctx context.Context) (*tokenproto.Token, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		row, _ := tx.Token.Query().Where(
			tokenent.ChainType(h.ChainType.String()),
			tokenent.Contract(*h.Contract),
			tokenent.TokenID(*h.TokenID),
		).Only(ctx)
		if row == nil {
			info, err := tokencrud.CreateSet(tx.Token.Create(),
				&tokencrud.Req{
					EntID:           h.EntID,
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
			).Save(ctx)
			if err != nil {
				return err
			}
			h.EntID = &info.EntID
			h.ID = &info.ID
			return nil
		}
		stm, err := tokencrud.UpdateSet(
			row.Update(),
			&tokencrud.Req{
				EntID:           &row.EntID,
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
		if info, err := stm.Save(_ctx); err == nil {
			h.EntID = &info.EntID
			h.ID = &info.ID
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return h.GetToken(ctx)
}
