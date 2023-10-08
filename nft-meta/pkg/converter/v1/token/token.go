package token

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

func Ent2Grpc(row *ent.Token) *npool.Token {
	if row == nil {
		return nil
	}

	return &npool.Token{
		ID:              row.ID.String(),
		ChainType:       basetype.ChainType(basetype.ChainType_value[row.ChainType]),
		ChainID:         row.ChainID,
		Contract:        row.Contract,
		TokenID:         row.TokenID,
		Owner:           row.Owner,
		URI:             row.URI,
		URIType:         row.URIType,
		ImageURL:        row.ImageURL,
		VideoURL:        row.VideoURL,
		Description:     row.Description,
		Name:            row.Name,
		TokenType:       basetype.TokenType(basetype.TokenType_value[row.TokenType]),
		VectorState:     npool.ConvertState(npool.ConvertState_value[row.VectorState]),
		VectorID:        row.VectorID,
		IPFSImageURL:    row.IpfsImageURL,
		ImageSnapshotID: row.ImageSnapshotID,
		Remark:          row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.Token) []*npool.Token {
	infos := []*npool.Token{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
