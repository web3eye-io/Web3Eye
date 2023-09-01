package orderitem

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/orderitem"
)

func Ent2Grpc(row *ent.OrderItem) *npool.OrderItem {
	if row == nil {
		return nil
	}

	return &npool.OrderItem{
		ID:         row.ID.String(),
		Contract:   row.Contract,
		TokenType:  row.TokenType,
		TokenID:    row.TokenID,
		Amount:     row.Amount,
		PortionNum: row.PortionNum,
		Remark:     row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.OrderItem) []*npool.OrderItem {
	infos := []*npool.OrderItem{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
