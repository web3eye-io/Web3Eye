package contract

import (
	"context"
	"fmt"

	contractcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	contractent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	contractproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
)

func (h *Handler) UpdateContract(ctx context.Context) (*contractproto.Contract, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Contract.
			Query().
			Where(
				contractent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := contractcrud.UpdateSet(
			info.Update(),
			&contractcrud.Req{
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				Address:     h.Address,
				Name:        h.Name,
				Symbol:      h.Symbol,
				Decimals:    h.Decimals,
				Creator:     h.Creator,
				BlockNum:    h.BlockNum,
				TxHash:      h.TxHash,
				TxTime:      h.TxTime,
				ProfileURL:  h.ProfileURL,
				BaseURL:     h.BaseURL,
				BannerURL:   h.BannerURL,
				Description: h.Description,
				Remark:      h.Remark,
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

	return h.GetContract(ctx)
}
