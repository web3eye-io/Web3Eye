package contract

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
)

func Ent2Grpc(row *ent.Contract) *npool.Contract {
	if row == nil {
		return nil
	}

	return &npool.Contract{
		ID:          row.ID.String(),
		ChainType:   row.ChainType,
		ChainID:     row.ChainID,
		Address:     row.Address,
		Name:        row.Name,
		Symbol:      row.Symbol,
		Creator:     row.Creator,
		BlockNum:    row.BlockNum,
		TxHash:      row.TxHash,
		TxTime:      row.TxTime,
		ProfileURL:  row.ProfileURL,
		BaseURL:     row.BaseURL,
		BannerURL:   row.BannerURL,
		Description: row.Description,
		Remark:      row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.Contract) []*npool.Contract {
	infos := []*npool.Contract{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
