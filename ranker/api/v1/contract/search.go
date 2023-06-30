package contract

import (
	"context"

	contractcrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/contract"
	tokencrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	transfercrud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/transfer"

	nftmetaconverter "github.com/web3eye-io/Web3Eye/nft-meta/pkg/converter/v1/contract"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/contract"
	rankerconverter "github.com/web3eye-io/Web3Eye/ranker/pkg/converter/v1/contract"
)

func (s *Server) GetContractAndTokens(ctx context.Context, in *rankernpool.GetContractAndTokensReq) (*rankernpool.GetContractAndTokensResp, error) {
	contractconds := &contract.Conds{Address: &web3eye.StringVal{
		Op:    "eq",
		Value: in.Contract,
	}}

	contract, err := contractcrud.RowOnly(ctx, contractconds)
	if err != nil {
		return nil, err
	}

	tokensconds := &token.Conds{Contract: &web3eye.StringVal{
		Op:    "eq",
		Value: in.Contract,
	}}
	tokens, total, err := tokencrud.Rows(ctx, tokensconds, int(in.Offset), int(in.Limit))
	if err != nil {
		return nil, err
	}

	transfersconds := &transfer.Conds{
		ChainType: &web3eye.StringVal{
			Op:    "eq",
			Value: contract.ChainType,
		},
		Contract: &web3eye.StringVal{
			Op:    "eq",
			Value: contract.Address,
		},
		ChainID: &web3eye.StringVal{
			Op:    "eq",
			Value: contract.ChainID,
		},
	}
	shotTokens := rankerconverter.Ent2GrpcMany(tokens)
	for _, v := range shotTokens {
		transfersconds.TokenID = &web3eye.StringVal{
			Op:    "eq",
			Value: v.TokenID,
		}
		total, err := transfercrud.Count(ctx, transfersconds)
		if err != nil {
			return nil, err
		}
		v.TransfersNum = total
	}

	return &rankernpool.GetContractAndTokensResp{
		Contract:    nftmetaconverter.Ent2Grpc(contract),
		Toknes:      shotTokens,
		TotalTokens: uint32(total),
	}, nil
}
