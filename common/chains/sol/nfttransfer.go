package sol

import (
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/web3eye-io/Web3Eye/common/chains"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
)

const (
	transferEventArgsLen = 4
	BaseTextNum          = 16
)

// eventHash represents an event keccak256 hash

type NFTTokenTransfer struct {
	From    string
	To      string
	TokenID string
}

func GetNftTransfersFromTX(info rpc.TransactionWithMeta, filterErrTX bool) []*NFTTokenTransfer {
	if filterErrTX && info.Meta.Err != nil {
		return nil
	}

	preNftTransfers := []rpc.TokenBalance{}
	postNftTransfers := []rpc.TokenBalance{}
	nftTransfers := []*NFTTokenTransfer{}

	for _, v := range info.Meta.PreTokenBalances {
		if v.UiTokenAmount.Amount == "1" && v.UiTokenAmount.Decimals == 0 {
			preNftTransfers = append(preNftTransfers, v)
		}
	}

	for _, v := range info.Meta.PostTokenBalances {
		if v.UiTokenAmount.Amount == "1" && v.UiTokenAmount.Decimals == 0 {
			postNftTransfers = append(postNftTransfers, v)
		}
	}

	for _, v1 := range postNftTransfers {
		from := "0"
		to := v1.Mint.String()
		tokenID := v1.Owner.String()
		for _, v2 := range preNftTransfers {
			if v1.Mint == v2.Mint {
				from = v2.Owner.String()
				continue
			}
		}

		nftTransfers = append(nftTransfers, &NFTTokenTransfer{
			From:    from,
			To:      to,
			TokenID: tokenID,
		})
	}
	return nftTransfers
}

func GetNFTTransfers(block *rpc.GetBlockResult) []*chains.TokenTransfer {
	nftTransfers := []*chains.TokenTransfer{}
	for i, tx := range block.Transactions {
		_nftTransfers := GetNftTransfersFromTX(tx, true)
		for _, transfer := range _nftTransfers {
			nftTransfers = append(nftTransfers,
				&chains.TokenTransfer{
					TokenType:   basetype.TokenType_Metaplex,
					TokenID:     transfer.TokenID,
					From:        transfer.From,
					To:          transfer.To,
					Amount:      1,
					BlockNumber: tx.Slot,
					TxHash:      block.Signatures[i].String(),
					BlockHash:   block.Blockhash.String(),
				})
		}
	}
	return nftTransfers
}

type Token struct {
	ChainType basetype.ChainType
	ChainID   int32
	Contract  string
	TokenType basetype.TokenType
	TokenID   string
	TokenURI  string
	Owner     string
	MediaURL  string
	MediaType string
}
