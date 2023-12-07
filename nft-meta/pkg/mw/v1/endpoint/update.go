package endpoint

import (
	"context"
	"fmt"

	endpointcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/endpoint"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	endpointent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/endpoint"
	endpointproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
)

func (h *Handler) UpdateEndpoint(ctx context.Context) (*endpointproto.Endpoint, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Endpoint.
			Query().
			Where(
				endpointent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := endpointcrud.UpdateSet(
			info.Update(),
			&endpointcrud.Req{
				ChainType: h.ChainType,
				ChainID:   h.ChainID,
				Address:   h.Address,
				State:     h.State,
				Remark:    h.Remark,
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

	return h.GetEndpoint(ctx)
}
