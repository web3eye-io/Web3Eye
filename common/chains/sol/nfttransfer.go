package sol

import (
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/web3eye-io/Web3Eye/common/chains"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
)

const (
	BaseTextNum = 16
)

// eventHash represents an event keccak256 hash

type NFTTokenTransfer struct {
	From    string
	To      string
	TokenID string
}

func GetNftTransfersFromTX(info *rpc.TransactionMeta, filterErrTX bool) []*NFTTokenTransfer {
	if filterErrTX && info.Err != nil {
		return nil
	}

	preNftTransfers := []rpc.TokenBalance{}
	postNftTransfers := []rpc.TokenBalance{}
	nftTransfers := []*NFTTokenTransfer{}

	for _, v := range info.PreTokenBalances {
		if v.UiTokenAmount.Amount == "1" && v.UiTokenAmount.Decimals == 0 {
			preNftTransfers = append(preNftTransfers, v)
		}
	}

	for _, v := range info.PostTokenBalances {
		if v.UiTokenAmount.Amount == "1" && v.UiTokenAmount.Decimals == 0 {
			postNftTransfers = append(postNftTransfers, v)
		}
	}

	for _, v1 := range postNftTransfers {
		from := "0"
		to := v1.Owner.String()
		tokenID := v1.Mint.String()
		for _, v2 := range preNftTransfers {
			if v1.Mint == v2.Mint {
				from = v2.Owner.String()
				continue
			}
		}
		if from == to {
			continue
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
	for _, txWithMeta := range block.Transactions {
		signature := ""
		tx, err := txWithMeta.GetTransaction()
		if err == nil && len(tx.Signatures) > 0 {
			// The signature that can be found is the first in the signature list of the tx
			signature = tx.Signatures[0].String()
		}

		_nftTransfers := GetNftTransfersFromTX(txWithMeta.Meta, true)

		for _, transfer := range _nftTransfers {
			nftTransfers = append(nftTransfers,
				&chains.TokenTransfer{
					TokenType: basetype.TokenType_Metaplex,
					TokenID:   transfer.TokenID,
					From:      transfer.From,
					To:        transfer.To,
					Amount:    1,
					// replace block height with slot height
					BlockNumber: block.ParentSlot + 1,
					TxHash:      signature,
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
