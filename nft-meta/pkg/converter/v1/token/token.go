package token

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

func Ent2Grpc(row *ent.Token) *npool.Token {
	if row == nil {
		return nil
	}

	return &npool.Token{
		ID:          row.ID.String(),
		ChainType:   row.ChainType,
		ChainID:     row.ChainID,
		Contract:    row.Contract,
		TokenID:     row.TokenID,
		Owner:       row.Owner,
		URI:         row.URI,
		URIType:     row.URIType,
		ImageURL:    row.ImageURL,
		VideoURL:    row.VideoURL,
		Description: row.Description,
		Name:        row.Name,
		VectorState: npool.ConvertState(npool.ConvertState_value[row.VectorState]),
		VectorID:    row.VectorID,
		Remark:      row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.Token) []*npool.Token {
	infos := []*npool.Token{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
