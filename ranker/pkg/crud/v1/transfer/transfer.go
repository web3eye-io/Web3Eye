package transfer

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"
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
	ImageURL      string    `json:"image_url,omitempty"`
}

func setQueryConds(in *rankernpool.GetTransfersRequest, cli *ent.Client) *ent.TransferQuery {
	stm := cli.Transfer.Query()
	stm.Where(transfer.ChainType(in.GetChainType().String()))
	stm.Where(transfer.ChainID(in.ChainID))
	stm.Where(transfer.Contract(in.Contract))
	if in.TokenID != nil {
		stm.Where(transfer.TokenID(*in.TokenID))
	}
	return stm
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
			AppendSelect(contract.FieldSymbol, contract.FieldName, contract.FieldDecimals).
			Where(sql.EQ(orderitem.FieldOrderID, orderID))
	}).Scan(ctx, &qOrderItems)
	if err != nil {
		return nil, err
	}

	for _, v := range qOrderItems {
		if v.TokenType == basetype.TokenType_ERC1155.String() ||
			v.TokenType == basetype.TokenType_ERC1155_WITH_CRITERIA.String() ||
			v.TokenType == basetype.TokenType_ERC721.String() ||
			v.TokenType == basetype.TokenType_ERC721_WITH_CRITERIA.String() {
			ret, err := cli.Token.Query().Select(token.FieldImageURL).Where(token.Contract(v.Contract), token.TokenID(v.TokenID)).First(ctx)
			if err != nil {
				return nil, err
			}
			v.ImageURL = ret.ImageURL
		}
	}

	return qOrderItems, nil
}

func Rows(ctx context.Context, in *rankernpool.GetTransfersRequest) ([]*rankernpool.Transfer, int, error) {
	var err error
	infos := []*rankernpool.Transfer{}
	rows := []*ent.Transfer{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := setQueryConds(in, cli)
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
		rowIDOrderID := make(map[uint32]string, len(rows))
		for _, row := range rows {
			orderItem, err := queryOrderItemAndOrder(row, cli).Only(ctx)
			if err != nil {
				continue
			}
			if orderItem != nil {
				rowIDOrderID[row.ID] = orderItem.OrderID.String()
			}
		}

		for _, row := range rows {
			var qOrderItems []*OrderItem
			if id, ok := rowIDOrderID[row.ID]; ok {
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
	rpctransfer := &rankernpool.Transfer{
		ID:          row.ID,
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
		LogIndex:    row.LogIndex,
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
			ImageURL:      v.ImageURL,
		}
		if v.OrderItemType == basetype.OrderItemType_OrderItemOffer.String() {
			rpctransfer.OfferItems = append(rpctransfer.OfferItems, orderItem)
		} else {
			rpctransfer.TargetItems = append(rpctransfer.TargetItems, orderItem)
		}
	}
	return rpctransfer
}
