package contract

import (
	"context"

	contractcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	contractproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

//nolint:dupl
func (h *Handler) CreateContract(ctx context.Context) (*contractproto.Contract, error) {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := contractcrud.CreateSet(
			cli.Contract.Create(),
			&contractcrud.Req{
				EntID:       h.EntID,
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

	return h.GetContract(ctx)
}

func (h *Handler) CreateContracts(ctx context.Context) ([]*contractproto.Contract, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := contractcrud.CreateSet(tx.Contract.Create(), req).Save(_ctx)
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

	h.Conds = &contractcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetContracts(ctx)
	return infos, err
}
