package order

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

type OrderDetail struct {
	Order       *ent.Order
	TargetItems []*ent.OrderItem
	OfferItems  []*ent.OrderItem
}

func Create(ctx context.Context, in *npool.OrderReq) (*OrderDetail, error) {
	var order *ent.Order
	var targetItems []*ent.OrderItem
	var offerItems []*ent.OrderItem
	var err error

	if in == nil {
		return nil, errors.New("input is nil")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		id := uuid.NewString()
		in.ID = &id
		// Create
		order, err = CreateSet(tx.Order.Create(), in).Save(ctx)
		if err != nil {
			return err
		}
		targetBulk := ItemCreateBulkSet(tx, order.ID, in.TargetItems, v1.OrderItemType_OrderItemTarget)
		targetItems, err = tx.OrderItem.CreateBulk(targetBulk...).Save(ctx)
		if err != nil {
			return err
		}
		offerBulk := ItemCreateBulkSet(tx, order.ID, in.OfferItems, v1.OrderItemType_OrderItemOffer)
		offerItems, err = tx.OrderItem.CreateBulk(offerBulk...).Save(ctx)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return &OrderDetail{
		Order:       order,
		TargetItems: targetItems,
		OfferItems:  offerItems,
	}, nil
}

func CreateSet(c *ent.OrderCreate, in *npool.OrderReq) *ent.OrderCreate {
	if in.ID != nil {
		id, err := uuid.Parse(*in.ID)
		if err != nil {
			id = uuid.New()
		}
		c.SetID(id)
	}
	if in.ChainType != nil {
		c.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		c.SetChainID(in.GetChainID())
	}
	if in.TxHash != nil {
		c.SetTxHash(in.GetTxHash())
	}
	if in.BlockNumber != nil {
		c.SetBlockNumber(in.GetBlockNumber())
	}
	if in.TxIndex != nil {
		c.SetTxIndex(in.GetTxIndex())
	}
	if in.LogIndex != nil {
		c.SetLogIndex(in.GetLogIndex())
	}
	if in.Recipient != nil {
		c.SetRecipient(in.GetRecipient())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func CreateBulk(ctx context.Context, infos []*npool.OrderReq) ([]*OrderDetail, error) {
	var err error
	rows := make([]*OrderDetail, len(infos))

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for i, in := range infos {
			id := uuid.NewString()
			in.ID = &id
			// Create
			order, err := CreateSet(tx.Order.Create(), in).Save(ctx)
			if err != nil {
				return err
			}
			targetBulk := ItemCreateBulkSet(tx, order.ID, in.TargetItems, v1.OrderItemType_OrderItemTarget)
			targetItems, err := tx.OrderItem.CreateBulk(targetBulk...).Save(ctx)
			if err != nil {
				return err
			}
			offerBulk := ItemCreateBulkSet(tx, order.ID, in.OfferItems, v1.OrderItemType_OrderItemOffer)
			offerItems, err := tx.OrderItem.CreateBulk(offerBulk...).Save(ctx)
			if err != nil {
				return err
			}
			rows[i] = &OrderDetail{
				Order:       order,
				TargetItems: targetItems,
				OfferItems:  offerItems,
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

//nolint:gocyclo
func Update(ctx context.Context, in *npool.OrderReq) (*OrderDetail, error) {
	var err error
	var order *ent.Order
	var targetItems []*ent.OrderItem
	var offerItems []*ent.OrderItem
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		u := tx.Order.UpdateOneID(uuid.MustParse(in.GetID()))
		order, err = UpdateSet(u, in).Save(_ctx)
		if err != nil {
			return err
		}

		targetItems, err = ItemsUpdate(ctx, tx, order.ID, in.TargetItems, v1.OrderItemType_OrderItemTarget)
		if err != nil {
			return err
		}

		offerItems, err = ItemsUpdate(ctx, tx, order.ID, in.OfferItems, v1.OrderItemType_OrderItemOffer)
		if err != nil {
			return err
		}

		return err
	})
	if err != nil {
		return nil, err
	}

	return &OrderDetail{
		Order:       order,
		TargetItems: targetItems,
		OfferItems:  offerItems,
	}, nil
}

func UpdateSet(u *ent.OrderUpdateOne, in *npool.OrderReq) *ent.OrderUpdateOne {
	if in.ChainType != nil {
		u.SetChainType(in.GetChainType().String())
	}
	if in.ChainID != nil {
		u.SetChainID(in.GetChainID())
	}
	if in.TxHash != nil {
		u.SetTxHash(in.GetTxHash())
	}
	if in.BlockNumber != nil {
		u.SetBlockNumber(in.GetBlockNumber())
	}
	if in.TxIndex != nil {
		u.SetTxIndex(in.GetTxIndex())
	}
	if in.LogIndex != nil {
		u.SetLogIndex(in.GetLogIndex())
	}
	if in.Recipient != nil {
		u.SetRecipient(in.GetRecipient())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*OrderDetail, error) {
	var info *ent.Order
	var targetItems []*ent.OrderItem
	var offerItems []*ent.OrderItem
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Order.Query().Where(order.ID(id)).Only(_ctx)
		if err != nil {
			return err
		}
		targetItems, err = cli.OrderItem.Query().Where(
			orderitem.OrderID(info.ID.String()),
			orderitem.OrderItemType(v1.OrderItemType_OrderItemTarget.String())).
			All(ctx)
		if err != nil {
			return err
		}
		offerItems, err = cli.OrderItem.Query().Where(
			orderitem.OrderID(info.ID.String()),
			orderitem.OrderItemType(v1.OrderItemType_OrderItemOffer.String())).
			All(ctx)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return &OrderDetail{
		Order:       info,
		TargetItems: targetItems,
		OfferItems:  offerItems,
	}, nil
}

// nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.OrderQuery, error) {
	stm := cli.Order.Query()
	if conds == nil {
		return stm, nil
	}

	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(order.ID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}

	if conds.IDs != nil {
		if conds.GetIDs().GetOp() == cruder.IN {
			var ids []uuid.UUID
			for _, val := range conds.GetIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(order.IDIn(ids...))
		}
	}

	if conds.ChainType != nil {
		switch conds.GetChainType().GetOp() {
		case cruder.EQ:
			stm.Where(order.ChainType(conds.GetChainType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.ChainID != nil {
		switch conds.GetChainID().GetOp() {
		case cruder.EQ:
			stm.Where(order.ChainID(conds.GetChainID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.TxHash != nil {
		switch conds.GetTxHash().GetOp() {
		case cruder.EQ:
			stm.Where(order.TxHash(conds.GetTxHash().GetValue()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.BlockNumber != nil {
		switch conds.GetBlockNumber().GetOp() {
		case cruder.EQ:
			stm.Where(order.BlockNumber(conds.GetBlockNumber().GetValue()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.TxIndex != nil {
		switch conds.GetTxIndex().GetOp() {
		case cruder.EQ:
			stm.Where(order.TxIndex(conds.GetTxIndex().GetValue()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.LogIndex != nil {
		switch conds.GetLogIndex().GetOp() {
		case cruder.EQ:
			stm.Where(order.LogIndex(conds.GetLogIndex().GetValue()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.Recipient != nil {
		switch conds.GetRecipient().GetOp() {
		case cruder.EQ:
			stm.Where(order.Recipient(conds.GetRecipient().GetValue()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(order.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*OrderDetail, int, error) {
	var err error
	infos := []*OrderDetail{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		rows, err := stm.
			Offset(offset).
			Order(ent.Desc(order.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}
		for _, info := range rows {
			targetItems, err := cli.OrderItem.Query().Where(
				orderitem.OrderID(info.ID.String()),
				orderitem.OrderItemType(v1.OrderItemType_OrderItemTarget.String())).
				All(ctx)
			if err != nil {
				return err
			}
			offerItems, err := cli.OrderItem.Query().Where(
				orderitem.OrderID(info.ID.String()),
				orderitem.OrderItemType(v1.OrderItemType_OrderItemOffer.String())).
				All(ctx)
			if err != nil {
				return err
			}
			infos = append(infos, &OrderDetail{
				Order:       info,
				TargetItems: targetItems,
				OfferItems:  offerItems,
			})
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*OrderDetail, error) {
	var info *ent.Order
	var targetItems []*ent.OrderItem
	var offerItems []*ent.OrderItem
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		targetItems, err = cli.OrderItem.Query().Where(
			orderitem.OrderID(info.ID.String()),
			orderitem.OrderItemType(v1.OrderItemType_OrderItemTarget.String())).
			All(ctx)
		if err != nil {
			return err
		}
		offerItems, err = cli.OrderItem.Query().Where(
			orderitem.OrderID(info.ID.String()),
			orderitem.OrderItemType(v1.OrderItemType_OrderItemOffer.String())).
			All(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &OrderDetail{
		Order:       info,
		TargetItems: targetItems,
		OfferItems:  offerItems,
	}, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error

	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Order.Query().Where(order.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error

	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*OrderDetail, error) {
	var order *ent.Order
	var err error

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		_, err = tx.OrderItem.Update().
			Where(
				orderitem.OrderID(id.String())).
			SetDeletedAt(uint32(time.Now().Unix())).Save(ctx)
		if err != nil {
			return err
		}
		order, err = tx.Order.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &OrderDetail{
		Order: order,
	}, nil
}
