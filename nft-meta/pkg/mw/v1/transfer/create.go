package transfer

import (
	"context"

	transfercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/transfer"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	transferproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateTransfer(ctx context.Context) (*transferproto.Transfer, error) {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := transfercrud.CreateSet(
			cli.Transfer.Create(),
			&transfercrud.Req{
				EntID:       h.EntID,
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				Contract:    h.Contract,
				TokenType:   h.TokenType,
				TokenID:     h.TokenID,
				From:        h.From,
				To:          h.To,
				Amount:      h.Amount,
				BlockNumber: h.BlockNumber,
				TxHash:      h.TxHash,
				BlockHash:   h.BlockHash,
				TxTime:      h.TxTime,
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

	return h.GetTransfer(ctx)
}

func (h *Handler) CreateTransfers(ctx context.Context) ([]*transferproto.Transfer, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := transfercrud.CreateSet(tx.Transfer.Create(), req).Save(_ctx)
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

	h.Conds = &transfercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetTransfers(ctx)
	return infos, err
}
