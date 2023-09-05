package block

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

func Ent2Grpc(row *ent.Block) *npool.Block {
	if row == nil {
		return nil
	}

	return &npool.Block{
		ID:          row.ID.String(),
		ChainType:   basetype.ChainType(basetype.ChainType_value[row.ChainType]),
		ChainID:     row.ChainID,
		BlockNumber: row.BlockNumber,
		BlockHash:   row.BlockHash,
		BlockTime:   row.BlockTime,
		ParseState:  basetype.BlockParseState(basetype.BlockParseState_value[row.ParseState]),
		Remark:      row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.Block) []*npool.Block {
	infos := []*npool.Block{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
