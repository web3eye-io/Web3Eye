package token

import (
	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
)

func Ent2Grpc(row *nftmetaproto.Token) *rankerproto.SearchToken {
	if row == nil {
		return nil
	}

	return &rankerproto.SearchToken{
		ID:              row.ID,
		EntID:           row.EntID,
		ChainType:       row.ChainType,
		ChainID:         row.ChainID,
		Contract:        row.Contract,
		TokenID:         row.TokenID,
		TokenType:       row.TokenType,
		Owner:           row.Owner,
		URI:             row.URI,
		URIType:         row.URIType,
		ImageURL:        row.ImageURL,
		VideoURL:        row.VideoURL,
		Description:     row.Description,
		Name:            row.Name,
		VectorState:     row.VectorState,
		VectorID:        row.VectorID,
		Remark:          row.Remark,
		IPFSImageURL:    row.IPFSImageURL,
		ImageSnapshotID: row.ImageSnapshotID,
	}
}

func Ent2GrpcMany(rows []*nftmetaproto.Token) []*rankerproto.SearchToken {
	infos := []*rankerproto.SearchToken{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
