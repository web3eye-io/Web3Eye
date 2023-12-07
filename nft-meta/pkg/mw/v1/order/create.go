package order

import (
	"context"

	ordercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/order"
	orderitemcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/orderitem"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	orderproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func createOrderItemsSet(orderID *uuid.UUID, tx *ent.Tx, items []*orderproto.OrderItem, itemType basetype.OrderItemType) []*ent.OrderItemCreate {
	bulk := make([]*ent.OrderItemCreate, len(items))

	for i, item := range items {
		bulk[i] = orderitemcrud.CreateSet(
			tx.OrderItem.Create(),
			&orderitemcrud.Req{
				OrderID:   orderID,
				Contract:  &item.Contract,
				TokenType: &item.TokenType,
				TokenID:   &item.TokenID,
				Amount:    &item.Amount,
				Remark:    &item.Remark,
			},
		).SetOrderItemType(itemType.String())
	}
	return bulk
}

func (h *Handler) CreateOrder(ctx context.Context) (*orderproto.Order, error) {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := ordercrud.CreateSet(
			tx.Order.Create(),
			&ordercrud.Req{
				EntID:       h.EntID,
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				TxHash:      h.TxHash,
				BlockNumber: h.BlockNumber,
				TxIndex:     h.TxIndex,
				LogIndex:    h.LogIndex,
				Recipient:   h.Recipient,
				Remark:      h.Remark,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		h.EntID = &info.EntID

		targetBulk := createOrderItemsSet(&info.EntID, tx, h.TargetItems, basetype.OrderItemType_OrderItemTarget)
		_, err = tx.OrderItem.CreateBulk(targetBulk...).Save(ctx)
		if err != nil {
			return err
		}

		offerBulk := createOrderItemsSet(&info.EntID, tx, h.OfferItems, basetype.OrderItemType_OrderItemOffer)
		_, err = tx.OrderItem.CreateBulk(offerBulk...).Save(ctx)

		return err
	})
	if err != nil {
		return nil, err
	}

	return h.GetOrder(ctx)
}

func (h *Handler) CreateOrders(ctx context.Context) ([]*orderproto.Order, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := ordercrud.CreateSet(
				tx.Order.Create(),
				&ordercrud.Req{
					EntID:       req.EntID,
					ChainType:   req.ChainType,
					ChainID:     req.ChainID,
					TxHash:      req.TxHash,
					BlockNumber: req.BlockNumber,
					TxIndex:     req.TxIndex,
					LogIndex:    req.LogIndex,
					Recipient:   req.Recipient,
					Remark:      req.Remark,
				},
			).Save(ctx)
			if err != nil {
				return err
			}

			targetBulk := createOrderItemsSet(&info.EntID, tx, h.TargetItems, basetype.OrderItemType_OrderItemTarget)
			_, err = tx.OrderItem.CreateBulk(targetBulk...).Save(ctx)
			if err != nil {
				return err
			}

			offerBulk := createOrderItemsSet(&info.EntID, tx, h.OfferItems, basetype.OrderItemType_OrderItemOffer)
			_, err = tx.OrderItem.CreateBulk(offerBulk...).Save(ctx)
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

	h.Conds = &ordercrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetOrders(ctx)
	return infos, err
}
