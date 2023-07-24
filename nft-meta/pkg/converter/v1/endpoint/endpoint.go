package endpoint

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
)

func Ent2Grpc(row *ent.Endpoint) *npool.Endpoint {
	if row == nil {
		return nil
	}

	return &npool.Endpoint{
		ID:        row.ID.String(),
		ChainType: basetype.ChainType(basetype.ChainType_value[row.ChainType]),
		ChainID:   row.ChainID,
		Address:   row.Address,
		State:     cttype.EndpointState(cttype.EndpointState_value[row.State]),
		Remark:    row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.Endpoint) []*npool.Endpoint {
	infos := []*npool.Endpoint{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
