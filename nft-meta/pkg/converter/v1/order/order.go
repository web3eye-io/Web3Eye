package order

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/order"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

func Ent2Grpc(row *order.OrderDetail) *npool.Order {
	if row == nil {
		return nil
	}

	targetItems := make([]*npool.OrderItem, len(row.TargetItems))
	for i, v := range row.TargetItems {
		targetItems[i] = &npool.OrderItem{
			Contract:  v.Contract,
			TokenType: basetype.TokenType(basetype.TokenType_value[v.TokenType]),
			TokenID:   v.TokenID,
			Amount:    v.Amount,
			Remark:    v.Remark,
		}
	}

	offerItems := make([]*npool.OrderItem, len(row.OfferItems))
	for i, v := range row.OfferItems {
		offerItems[i] = &npool.OrderItem{
			Contract:  v.Contract,
			TokenType: basetype.TokenType(basetype.TokenType_value[v.TokenType]),
			TokenID:   v.TokenID,
			Amount:    v.Amount,
			Remark:    v.Remark,
		}
	}

	return &npool.Order{
		ID:          row.Order.ID.String(),
		ChainType:   basetype.ChainType(basetype.ChainType_value[row.Order.ChainType]),
		ChainID:     row.Order.ChainID,
		TxHash:      row.Order.TxHash,
		BlockNumber: row.Order.BlockNumber,
		TxIndex:     row.Order.TxIndex,
		LogIndex:    row.Order.LogIndex,
		Recipient:   row.Order.Recipient,
		TargetItems: targetItems,
		OfferItems:  offerItems,
		Remark:      row.Order.Remark,
	}
}

func Ent2GrpcMany(rows []*order.OrderDetail) []*npool.Order {
	infos := []*npool.Order{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
