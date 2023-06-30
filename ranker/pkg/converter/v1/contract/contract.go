package contract

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	ranker_npool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/contract"
)

func Ent2Grpc(row *ent.Token) *ranker_npool.ShotToken {
	if row == nil {
		return nil
	}

	return &ranker_npool.ShotToken{
		ID:              row.ID.String(),
		ChainType:       v1.ChainType(v1.ChainType_value[row.ChainType]),
		TokenType:       v1.TokenType(v1.TokenType_value[row.TokenType]),
		TokenID:         row.TokenID,
		Owner:           row.Owner,
		ImageURL:        row.ImageURL,
		Name:            row.Name,
		IPFSImageURL:    row.IpfsImageURL,
		ImageSnapshotID: row.ImageSnapshotID,
	}
}

func Ent2GrpcMany(rows []*ent.Token) []*ranker_npool.ShotToken {
	infos := []*ranker_npool.ShotToken{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
