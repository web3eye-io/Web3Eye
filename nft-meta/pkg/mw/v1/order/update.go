package order

import (
	"context"
	"fmt"

	ordercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	orderent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	orderproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

type updateHandler struct {
	*Handler
}

func (uh *updateHandler) ItemsUpdate(ctx context.Context, tx *ent.Tx, in []*orderproto.OrderItem, orderItemType basetype.OrderItemType) ([]*ent.OrderItem, error) {
	items, err := tx.OrderItem.Query().Where(
		orderitem.OrderID(*uh.EntID),
		orderitem.OrderItemType(orderItemType.String())).
		All(ctx)
	if err != nil {
		return nil, err
	}

	if len(in) == 0 {
		return items, nil
	}

	ids := make([]uint32, len(items))
	for i, v := range items {
		ids[i] = v.ID
	}

	_, err = tx.OrderItem.Delete().Where(orderitem.IDIn(ids...)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	bulk := createOrderItemsSet(uh.EntID, tx, in, orderItemType)
	items, err = tx.OrderItem.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return items, err
}

func (h *Handler) UpdateOrder(ctx context.Context) (*orderproto.Order, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	uh := &updateHandler{Handler: h}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Order.
			Query().
			Where(
				orderent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := ordercrud.UpdateSet(
			info.Update(),
			&ordercrud.Req{
				ChainType:   h.ChainType,
				ChainID:     h.ChainID,
				TxHash:      h.TxHash,
				BlockNumber: h.BlockNumber,
				TxIndex:     h.TxIndex,
				LogIndex:    h.LogIndex,
				Recipient:   h.Recipient,
				Remark:      h.Remark,
			},
		)
		if err != nil {
			return err
		}
		if _, err := stm.Save(_ctx); err != nil {
			return err
		}
		uh.EntID = &info.EntID

		_, err = uh.ItemsUpdate(ctx, tx, h.TargetItems, basetype.OrderItemType_OrderItemTarget)
		if err != nil {
			return err
		}

		_, err = uh.ItemsUpdate(ctx, tx, h.OfferItems, basetype.OrderItemType_OrderItemOffer)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOrder(ctx)
}
