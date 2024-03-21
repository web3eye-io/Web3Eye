package transfer

import (
	"context"

	"github.com/google/uuid"
	transfercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/transfer"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	transferent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"
	transferproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func (h *Handler) UpsertTransfer(ctx context.Context) (*transferproto.Transfer, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		id, entID, err := upsertOne(ctx, tx, &transfercrud.Req{
			ID:          h.ID,
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
			LogIndex:    h.LogIndex,
		})
		if err != nil {
			return err
		}
		h.EntID = entID
		h.ID = id
		return nil
	})

	if err != nil {
		return nil, err
	}

	return h.GetTransfer(ctx)
}

func upsertOne(ctx context.Context, tx *ent.Tx, req *transfercrud.Req) (*uint32, *uuid.UUID, error) {
	row, _ := tx.Transfer.Query().Where(
		transferent.Contract(*req.Contract),
		transferent.TokenID(*req.TokenID),
		transferent.TxHash(*req.TxHash),
		transferent.From(*req.From),
	).Only(ctx)
	if row == nil {
		info, err := transfercrud.CreateSet(tx.Transfer.Create(),
			&transfercrud.Req{
				EntID:       req.EntID,
				ChainType:   req.ChainType,
				ChainID:     req.ChainID,
				Contract:    req.Contract,
				TokenType:   req.TokenType,
				TokenID:     req.TokenID,
				From:        req.From,
				To:          req.To,
				Amount:      req.Amount,
				BlockNumber: req.BlockNumber,
				TxHash:      req.TxHash,
				BlockHash:   req.BlockHash,
				TxTime:      req.TxTime,
				Remark:      req.Remark,
				LogIndex:    req.LogIndex,
			},
		).Save(ctx)
		if err != nil {
			return nil, nil, err
		}
		return &info.ID, &info.EntID, nil
	}
	stm, err := transfercrud.UpdateSet(
		row.Update(),
		&transfercrud.Req{
			EntID:       &row.EntID,
			ChainType:   req.ChainType,
			ChainID:     req.ChainID,
			Contract:    req.Contract,
			TokenType:   req.TokenType,
			TokenID:     req.TokenID,
			From:        req.From,
			To:          req.To,
			Amount:      req.Amount,
			BlockNumber: req.BlockNumber,
			TxHash:      req.TxHash,
			BlockHash:   req.BlockHash,
			TxTime:      req.TxTime,
			Remark:      req.Remark,
			LogIndex:    req.LogIndex,
		},
	)
	if err != nil {
		return nil, nil, err
	}

	info, err := stm.Save(ctx)
	if err != nil {
		return nil, nil, err
	}
	return &info.ID, &info.EntID, nil
}

func (h *Handler) UpsertTransfers(ctx context.Context) error {
	entIDs := make([]*uuid.UUID, len(h.Reqs))
	for i, req := range h.Reqs {
		err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			_, entID, err := upsertOne(ctx, tx, req)
			if err != nil {
				return err
			}
			entIDs[i] = entID
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
