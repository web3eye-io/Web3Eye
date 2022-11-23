package transfer

import (
	npool "github.com/NpoolPlatform/message/npool/nftmeta/v1/transfer"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"
)

func Ent2Grpc(row *ent.Transfer) *npool.Transfer {
	if row == nil {
		return nil
	}

	return &npool.Transfer{
		ID:          row.ID.String(),
		ChainType:   row.ChainType,
		ChainID:     row.ChainID,
		Contract:    row.Contract,
		TokenType:   row.TokenType,
		TokenID:     row.TokenID,
		From:        row.From,
		To:          row.To,
		Amount:      row.Amount,
		BlockNumber: row.BlockNumber,
		TxHash:      row.TxHash,
		BlockHash:   row.BlockHash,
		TxTime:      row.TxTime,
		Remark:      row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.Transfer) []*npool.Transfer {
	infos := []*npool.Transfer{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
