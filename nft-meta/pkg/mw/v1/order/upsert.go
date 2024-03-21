package order

import (
	"context"

	ordercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	orderent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	orderproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) UpsertOrder(ctx context.Context) (*orderproto.Order, error) {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		id, entID, err := upsertOne(_ctx, tx, &ordercrud.Req{
			EntID:       h.EntID,
			ChainType:   h.ChainType,
			ChainID:     h.ChainID,
			TxHash:      h.TxHash,
			BlockNumber: h.BlockNumber,
			TxIndex:     h.TxIndex,
			LogIndex:    h.LogIndex,
			Recipient:   h.Recipient,
			TargetItems: h.TargetItems,
			OfferItems:  h.OfferItems,
			Remark:      h.Remark,
		})
		if err != nil {
			return err
		}
		h.ID = id
		h.EntID = entID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOrder(ctx)
}

func upsertOne(_ctx context.Context, tx *ent.Tx, req *ordercrud.Req) (*uint32, *uuid.UUID, error) {
	row, _ := tx.Order.Query().Where(
		orderent.TxHash(*req.TxHash),
		orderent.Recipient(*req.Recipient),
		orderent.LogIndex(*req.LogIndex),
	).Only(_ctx)
	var id *uint32
	var entID *uuid.UUID
	if row == nil {
		info, err := ordercrud.CreateSet(
			tx.Order.Create(),
			req,
		).Save(_ctx)
		if err != nil {
			return id, entID, err
		}

		id = &info.ID
		entID = &info.EntID
	} else {
		req.EntID = &row.EntID
		updateOne, err := ordercrud.UpdateSet(
			row.Update(),
			req,
		)
		if err != nil {
			return id, entID, err
		}
		err = updateOne.Exec(_ctx)
		if err != nil {
			return id, entID, err
		}

		id = &row.ID
		entID = &row.EntID
	}

	targetBulk := createOrderItemsSet(entID, tx, req.TargetItems, basetype.OrderItemType_OrderItemTarget)
	err := tx.OrderItem.CreateBulk(targetBulk...).OnConflict().UpdateNewValues().Exec(_ctx)
	if err != nil {
		return id, entID, err
	}

	offerBulk := createOrderItemsSet(entID, tx, req.OfferItems, basetype.OrderItemType_OrderItemOffer)
	err = tx.OrderItem.CreateBulk(offerBulk...).OnConflict().UpdateNewValues().Exec(_ctx)
	return id, entID, err
}

func (h *Handler) UpsertOrders(ctx context.Context) ([]*orderproto.Order, error) {
	entIDs := []uuid.UUID{}
	for _, req := range h.Reqs {
		err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			_, entID, err := upsertOne(_ctx, tx, req)
			if err != nil {
				return err
			}

			entIDs = append(entIDs, *entID)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	h.Conds = &ordercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: entIDs},
	}
	h.Offset = 0
	h.Limit = int32(len(entIDs))

	infos, _, err := h.GetOrders(ctx)
	return infos, err
}
