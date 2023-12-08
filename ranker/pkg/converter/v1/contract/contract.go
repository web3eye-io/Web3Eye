package contract

import (
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/contract"
)

func Ent2Grpc(row *tokenproto.Token) *rankerproto.ShotToken {
	if row == nil {
		return nil
	}

	return &rankerproto.ShotToken{
		ID:              row.ID,
		ChainType:       row.ChainType,
		TokenType:       row.TokenType,
		TokenID:         row.TokenID,
		Owner:           row.Owner,
		ImageURL:        row.ImageURL,
		Name:            row.Name,
		IPFSImageURL:    row.IPFSImageURL,
		ImageSnapshotID: row.ImageSnapshotID,
	}
}

func Ent2GrpcMany(rows []*tokenproto.Token) []*rankerproto.ShotToken {
	infos := []*rankerproto.ShotToken{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
