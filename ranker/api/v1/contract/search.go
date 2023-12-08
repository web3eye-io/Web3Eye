package contract

import (
	"context"
	"fmt"

	contracthandler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/contract"
	tokenhandler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/token"

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

	contractHandler, err := contracthandler.NewHandler(ctx, contracthandler.WithConds(contractconds))
	if err != nil {
		return nil, err
	}

	contracts, num, err := contractHandler.GetContracts(ctx)
	if err != nil {
		return nil, err
	}

	if num != 1 {
		return nil, fmt.Errorf("have more then one or have no contract, contract: %v", in.Contract)
	}
	contract := contracts[0]

	tokensconds := &token.Conds{Contract: &web3eye.StringVal{
		Op:    "eq",
		Value: in.Contract,
	}}

	tokenHandler, err := tokenhandler.NewHandler(ctx,
		tokenhandler.WithConds(tokensconds),
		tokenhandler.WithOffset(int32(in.Offset)),
		tokenhandler.WithLimit(int32(in.Limit)),
	)
	if err != nil {
		return nil, err
	}

	tokens, total, err := tokenHandler.GetTokens(ctx)
	if err != nil {
		return nil, err
	}

	transfersconds := &transfer.Conds{
		ChainType: &web3eye.Uint32Val{
			Op:    "eq",
			Value: uint32(*contract.ChainType.Enum()),
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
		contractHandler, err := contracthandler.NewHandler(ctx, contracthandler.WithConds(contractconds))
		if err != nil {
			return nil, err
		}

		_, total, err := contractHandler.GetContracts(ctx)
		if err != nil {
			return nil, err
		}
		v.TransfersNum = total
	}

	return &rankernpool.GetContractAndTokensResp{
		Contract:    contract,
		Tokens:      shotTokens,
		TotalTokens: uint32(total),
	}, nil
}
