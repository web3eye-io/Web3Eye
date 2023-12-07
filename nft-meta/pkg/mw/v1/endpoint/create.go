package endpoint

import (
	"context"

	endpointcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/endpoint"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	endpointproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateEndpoint(ctx context.Context) (*endpointproto.Endpoint, error) {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := endpointcrud.CreateSet(
			cli.Endpoint.Create(),
			&endpointcrud.Req{
				EntID:     h.EntID,
				ChainType: h.ChainType,
				ChainID:   h.ChainID,
				Address:   h.Address,
				State:     h.State,
				Remark:    h.Remark,
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

	return h.GetEndpoint(ctx)
}

func (h *Handler) CreateEndpoints(ctx context.Context) ([]*endpointproto.Endpoint, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := endpointcrud.CreateSet(tx.Endpoint.Create(), req).Save(_ctx)
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

	h.Conds = &endpointcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetEndpoints(ctx)
	return infos, err
}
