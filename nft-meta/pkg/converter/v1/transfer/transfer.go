package transfer

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func Ent2Grpc(row *ent.Transfer) *npool.Transfer {
	if row == nil {
		return nil
	}

	return &npool.Transfer{
		ID:          row.ID.String(),
		ChainType:   basetype.ChainType(basetype.ChainType_value[row.ChainType]),
		ChainID:     row.ChainID,
		Contract:    row.Contract,
		TokenType:   basetype.TokenType(basetype.TokenType_value[row.TokenType]),
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
