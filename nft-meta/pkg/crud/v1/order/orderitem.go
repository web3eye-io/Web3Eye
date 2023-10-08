package order

import (
	"context"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

func ItemCreateSet(c *ent.OrderItemCreate, orderID uuid.UUID, in *npool.OrderItem, itemType basetype.OrderItemType) *ent.OrderItemCreate {
	c.SetID(uuid.New())
	c.SetOrderID(orderID.String())
	c.SetOrderItemType(itemType.String())
	c.SetContract(in.GetContract())
	c.SetTokenType(in.GetTokenType().String())
	c.SetTokenID(in.GetTokenID())
	c.SetAmount(utils.Uint64ToDecStr(in.GetAmount()))
	c.SetRemark(in.GetRemark())
	return c
}

func ItemCreateBulkSet(tx *ent.Tx, orderID uuid.UUID, in []*npool.OrderItem, itemType basetype.OrderItemType) []*ent.OrderItemCreate {
	cBulk := make([]*ent.OrderItemCreate, len(in))
	for i, v := range in {
		cBulk[i] = ItemCreateSet(tx.OrderItem.Create(), orderID, v, itemType)
	}
	return cBulk
}

func ItemsUpdate(ctx context.Context, tx *ent.Tx, orderID uuid.UUID, in []*npool.OrderItem, orderItemType basetype.OrderItemType) ([]*ent.OrderItem, error) {
	items, err := tx.OrderItem.Query().Where(
		orderitem.OrderID(orderID.String()),
		orderitem.OrderItemType(orderItemType.String())).
		All(ctx)
	if err != nil {
		return nil, err
	}

	if len(in) == 0 {
		return items, nil
	}

	ids := make([]uuid.UUID, len(items))
	for i, v := range items {
		ids[i] = v.ID
	}

	_, err = tx.OrderItem.Delete().Where(orderitem.IDIn(ids...)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	bulk := ItemCreateBulkSet(tx, orderID, in, orderItemType)
	items, err = tx.OrderItem.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return items, err
}
