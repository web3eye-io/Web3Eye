package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/utils"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	ordercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	orderent "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
	orderproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

type queryHandler struct {
	*Handler
	stm   *ent.OrderSelect
	infos []*orderproto.Order
	total uint32
}

func (h *queryHandler) selectOrder(stm *ent.OrderQuery) {
	h.stm = stm.Select(
		orderent.FieldID,
		orderent.FieldEntID,
		orderent.FieldChainType,
		orderent.FieldChainID,
		orderent.FieldTxHash,
		orderent.FieldBlockNumber,
		orderent.FieldTxIndex,
		orderent.FieldLogIndex,
		orderent.FieldRecipient,
		orderent.FieldRemark,
		orderent.FieldCreatedAt,
		orderent.FieldUpdatedAt,
	)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.ChainType = basetype.ChainType(basetype.ChainType_value[info.ChainTypeStr])
	}
}

func (h *queryHandler) queryOrder(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Order.Query().Where(orderent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(orderent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(orderent.EntID(*h.EntID))
	}
	h.selectOrder(stm)
	return nil
}

func (h *queryHandler) queryOrders(ctx context.Context, cli *ent.Client) error {
	stm, err := ordercrud.SetQueryConds(cli.Order.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectOrder(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func orderItemEnt2Proto(item *ent.OrderItem) *orderproto.OrderItem {
	amount, _ := utils.DecStr2uint64(item.Amount)
	return &orderproto.OrderItem{
		Contract:  item.Contract,
		TokenType: basetype.TokenType(basetype.TokenType_value[item.TokenType]),
		TokenID:   item.TokenID,
		Amount:    amount,
		Remark:    item.Remark,
	}
}
func orderItemsEnt2Proto(items []*ent.OrderItem) []*orderproto.OrderItem {
	_items := make([]*orderproto.OrderItem, len(items))
	for i, v := range items {
		_items[i] = orderItemEnt2Proto(v)
	}
	return _items
}

func (h *Handler) GetOrder(ctx context.Context) (*orderproto.Order, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrder(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		err := handler.scan(_ctx)
		if err != nil {
			return err
		}

		if len(handler.infos) > 1 {
			return fmt.Errorf("too many record")
		} else if len(handler.infos) == 0 {
			return nil
		}

		info := handler.infos[0]
		orderID, err := uuid.Parse(info.EntID)
		if err != nil {
			return err
		}

		targetItems, err := cli.OrderItem.Query().Where(
			orderitem.OrderID(orderID),
			orderitem.OrderItemType(basetype.OrderItemType_OrderItemTarget.String())).
			All(ctx)
		if err != nil {
			return err
		}

		offerItems, err := cli.OrderItem.Query().Where(
			orderitem.OrderID(orderID),
			orderitem.OrderItemType(basetype.OrderItemType_OrderItemOffer.String())).
			All(ctx)
		if err != nil {
			return err
		}

		info.TargetItems = orderItemsEnt2Proto(targetItems)
		info.OfferItems = orderItemsEnt2Proto(offerItems)
		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetOrders(ctx context.Context) ([]*orderproto.Order, uint32, error) {
	if h.Conds == nil {
		return nil, 0, fmt.Errorf("the conds is nil")
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrders(ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(orderent.FieldUpdatedAt))
		err := handler.scan(_ctx)
		if err != nil {
			return err
		}

		for _, info := range handler.infos {
			orderID, err := uuid.Parse(info.EntID)
			if err != nil {
				return err
			}

			targetItems, err := cli.OrderItem.Query().Where(
				orderitem.OrderID(orderID),
				orderitem.OrderItemType(basetype.OrderItemType_OrderItemTarget.String())).
				All(ctx)
			if err != nil {
				return err
			}

			offerItems, err := cli.OrderItem.Query().Where(
				orderitem.OrderID(orderID),
				orderitem.OrderItemType(basetype.OrderItemType_OrderItemOffer.String())).
				All(ctx)
			if err != nil {
				return err
			}

			info.TargetItems = orderItemsEnt2Proto(targetItems)
			info.OfferItems = orderItemsEnt2Proto(offerItems)
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
