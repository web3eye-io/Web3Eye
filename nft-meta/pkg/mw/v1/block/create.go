package block

import (
	"context"

	blockcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/block"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	blockproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateBlock(ctx context.Context) (*blockproto.Block, error) {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := blockcrud.CreateSet(
			cli.Block.Create(),
			&blockcrud.Req{
				EntID:       h.EntID,
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				BlockNumber: h.BlockNumber,
				BlockHash:   h.BlockHash,
				BlockTime:   h.BlockTime,
				ParseState:  h.ParseState,
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

	return h.GetBlock(ctx)
}

func (h *Handler) CreateBlocks(ctx context.Context) ([]*blockproto.Block, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := blockcrud.CreateSet(tx.Block.Create(), req).Save(_ctx)
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

	h.Conds = &blockcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetBlocks(ctx)
	return infos, err
}
