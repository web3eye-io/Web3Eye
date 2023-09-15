package transfer

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
)

type OrderItem struct {
	ID            uuid.UUID `json:"id,omitempty"`
	CreatedAt     uint32    `json:"created_at,omitempty"`
	UpdatedAt     uint32    `json:"updated_at,omitempty"`
	DeletedAt     uint32    `json:"deleted_at,omitempty"`
	OrderID       string    `json:"order_id,omitempty"`
	OrderItemType string    `json:"order_item_type,omitempty"`
	Contract      string    `json:"contract,omitempty"`
	TokenType     string    `json:"token_type,omitempty"`
	TokenID       string    `json:"token_id,omitempty"`
	Amount        uint64    `json:"amount,omitempty"`
	Remark        string    `json:"remark,omitempty"`
	Name          string    `json:"name,omitempty"`
	Symbol        string    `json:"symbol,omitempty"`
	Decimals      uint32    `json:"decimals,omitempty"`
}

// nolint
func setQueryConds(in *rankernpool.GetTransfersRequest, cli *ent.Client) (*ent.TransferQuery, error) {
	stm := cli.Transfer.Query()
	stm.Where(transfer.ChainType(in.GetChainType().String()))
	stm.Where(transfer.ChainID(in.ChainID))
	stm.Where(transfer.Contract(in.Contract))
	if in.TokenID != nil {
		stm.Where(transfer.TokenID(*in.TokenID))
	}
	return stm, nil
}

func queryOrderItemAndOrder(row *ent.Transfer, cli *ent.Client) *ent.OrderItemSelect {
	return cli.OrderItem.Query().Modify(func(s *sql.Selector) {
		t := sql.Table(order.Table)
		s.
			LeftJoin(t).
			On(
				t.C(order.FieldID),
				s.C(orderitem.FieldOrderID),
			).
			Where(sql.EQ(order.FieldTxHash, row.TxHash)).
			Where(sql.EQ(order.FieldRecipient, row.To)).
			Where(sql.EQ(orderitem.FieldContract, row.Contract)).
			Where(sql.EQ(orderitem.FieldTokenID, row.TokenID))
	})
}

func queryOrderItemsAndContract(ctx context.Context, orderID string, cli *ent.Client) ([]*OrderItem, error) {
	var qOrderItems []*OrderItem
	err := cli.OrderItem.Query().Modify(func(s *sql.Selector) {
		t := sql.Table(contract.Table)
		s.LeftJoin(t).
			On(s.C(orderitem.FieldContract), t.C(contract.FieldAddress)).
			Where(sql.EQ(orderitem.FieldOrderID, orderID)).
			AppendSelect(contract.FieldSymbol, contract.FieldName, contract.FieldDecimals)
	}).Scan(ctx, &qOrderItems)
	if err != nil {
		return nil, err
	}
	return qOrderItems, nil
}

func Rows(ctx context.Context, in *rankernpool.GetTransfersRequest) ([]*rankernpool.Transfer, int, error) {
	var err error
	infos := []*rankernpool.Transfer{}
	rows := []*ent.Transfer{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(in, cli)
		if err != nil {
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		rows, err = stm.
			Offset(int(in.Offset)).
			Order(ent.Desc(transfer.FieldUpdatedAt)).
			Limit(int(in.Limit)).
			All(_ctx)
		if err != nil {
			return err
		}
		rowID_orderID := make(map[string]string, len(rows))
		for _, row := range rows {
			orderItem, err := queryOrderItemAndOrder(row, cli).Only(ctx)
			if err != nil {
				continue
			}
			if orderItem != nil {
				rowID_orderID[row.ID.String()] = orderItem.OrderID
			}
		}

		for _, row := range rows {
			var qOrderItems []*OrderItem
			if id, ok := rowID_orderID[row.ID.String()]; ok {
				qOrderItems, err = queryOrderItemsAndContract(ctx, id, cli)
				if err != nil {
					return err
				}
			}
			infos = append(infos, ent2rpcTransfer(row, qOrderItems))
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func ent2rpcTransfer(row *ent.Transfer, orderItems []*OrderItem) *rankernpool.Transfer {
	amount, _ := utils.DecStr2uint64(row.Amount)
	transfer := &rankernpool.Transfer{
		ID:          row.ID.String(),
		ChainType:   basetype.ChainType(basetype.ChainType_value[row.ChainType]),
		ChainID:     row.ChainID,
		Contract:    row.Contract,
		TokenType:   basetype.TokenType(basetype.TokenType_value[row.TokenType]),
		TokenID:     row.TokenID,
		From:        row.From,
		To:          row.To,
		Amount:      amount,
		BlockNumber: row.BlockNumber,
		TxHash:      row.TxHash,
		BlockHash:   row.BlockHash,
		TxTime:      row.TxTime,
		Remark:      row.Remark,
		TargetItems: []*rankernpool.OrderItem{},
		OfferItems:  []*rankernpool.OrderItem{},
	}
	for _, v := range orderItems {
		orderItem := &rankernpool.OrderItem{
			Contract:      v.Contract,
			TokenType:     basetype.TokenType(basetype.TokenType_value[v.TokenType]),
			TokenID:       v.TokenID,
			Amount:        v.Amount,
			Remark:        v.Remark,
			Name:          v.Name,
			Symbol:        v.Symbol,
			Decimals:      v.Decimals,
			OrderItemType: v.OrderItemType,
		}
		if v.OrderItemType == basetype.OrderItemType_OrderItemOffer.String() {
			transfer.OfferItems = append(transfer.OfferItems, orderItem)
		} else {
			transfer.TargetItems = append(transfer.TargetItems, orderItem)
		}
	}
	return transfer
}
