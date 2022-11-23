package token

import (
	npool "github.com/NpoolPlatform/message/npool/nftmeta/v1/token"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"
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
