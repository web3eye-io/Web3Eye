package blocknumber

import (
	npool "github.com/NpoolPlatform/message/npool/nftmeta/v1/blocknumber"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"
)

func Ent2Grpc(row *ent.BlockNumber) *npool.BlockNumber {
	if row == nil {
		return nil
	}

	return &npool.BlockNumber{
		ID:          row.ID.String(),
		ChainType:   row.ChainType,
		ChainID:     row.ChainID,
		Identifier:  row.Identifier,
		CurrentNum:  row.CurrentNum,
		Topic:       row.Topic,
		Description: row.Description,
	}
}

func Ent2GrpcMany(rows []*ent.BlockNumber) []*npool.BlockNumber {
	infos := []*npool.BlockNumber{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
