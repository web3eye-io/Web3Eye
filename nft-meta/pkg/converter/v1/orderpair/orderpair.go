package orderpair

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/orderpair"
)

func Ent2Grpc(row *ent.OrderPair) *npool.OrderPair {
	if row == nil {
		return nil
	}

	return &npool.OrderPair{
		ID:        row.ID.String(),
		TxHash:    row.TxHash,
		Recipient: row.Recipient,
		TargetID:  row.TargetID,
		OfferID:   row.OfferID,
		Remark:    row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.OrderPair) []*npool.OrderPair {
	infos := []*npool.OrderPair{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
