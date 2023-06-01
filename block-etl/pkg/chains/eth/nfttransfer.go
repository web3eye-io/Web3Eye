package eth

import (
	"context"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cteth "github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/common/chains/eth/contracts"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
)

const (
	// transferEventHash represents the keccak256 hash of Transfer(address,address,uint256)
	transferEventHash eventHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// transferSingleEventHash represents the keccak256 hash of TransferSingle(address,address,address,uint256,uint256)
	transferSingleEventHash eventHash = "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	// transferBatchEventHash represents the keccak256 hash of TransferBatch(address,address,address,uint256[],uint256[])
	transferBatchEventHash eventHash = "0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"

	transferEventArgsLen = 4
	BaseTextNum          = 16
)

// eventHash represents an event keccak256 hash
type eventHash string
type TokenTransfer struct {
	ChainType   basetype.ChainType
	ChainID     string
	Contract    string
	TokenType   basetype.TokenType
	TokenID     string
	From        string
	To          string
	Amount      uint64
	BlockNumber uint64
	TxHash      string
	BlockHash   string
	TxTime      uint32
}

var (
	erc1155ABI, _ = contracts.IERC1155MetaData.GetAbi()
)

//nolint:gocritic
func logsToTransfer(pLogs []types.Log) ([]*TokenTransfer, error) {
	result := make([]*TokenTransfer, 0)
	for _, pLog := range pLogs {
		switch {
		case strings.EqualFold(pLog.Topics[0].Hex(), string(transferEventHash)):
			// filter erc20
			if len(pLog.Topics) < transferEventArgsLen {
				continue
			}

			result = append(result, &TokenTransfer{
				From:        pLog.Topics[1].Hex(),
				To:          pLog.Topics[2].Hex(),
				ChainType:   basetype.ChainType_Ethereum,
				Contract:    pLog.Address.Hex(),
				TokenID:     pLog.Topics[3].Big().String(),
				BlockNumber: pLog.BlockNumber,
				Amount:      1,
				TokenType:   basetype.TokenType_ERC721,
				TxHash:      pLog.TxHash.Hex(),
				BlockHash:   pLog.BlockHash.Hex(),
			})
		case strings.EqualFold(pLog.Topics[0].Hex(), string(transferSingleEventHash)):
			if len(pLog.Topics) < transferEventArgsLen {
				continue
			}
			eventData := map[string]interface{}{}
			err := erc1155ABI.UnpackIntoMap(eventData, "TransferSingle", pLog.Data)
			if err != nil {
				panic(err)
			}

			id, ok := eventData["id"].(*big.Int)
			if !ok {
				panic("Failed to unpack TransferSingle event, id not found")
			}

			value, ok := eventData["value"].(*big.Int)
			if !ok {
				panic("Failed to unpack TransferSingle event, value not found")
			}
			result = append(result, &TokenTransfer{
				From:        pLog.Topics[2].Hex(),
				To:          pLog.Topics[3].Hex(),
				ChainType:   basetype.ChainType_Ethereum,
				Contract:    pLog.Address.Hex(),
				TokenID:     id.String(),
				Amount:      value.Uint64(),
				BlockNumber: pLog.BlockNumber,
				TokenType:   basetype.TokenType_ERC1155,
				TxHash:      pLog.TxHash.Hex(),
				BlockHash:   pLog.BlockHash.Hex(),
			})
		case strings.EqualFold(pLog.Topics[0].Hex(), string(transferBatchEventHash)):
			if len(pLog.Topics) < transferEventArgsLen {
				continue
			}

			eventData := map[string]interface{}{}
			err := erc1155ABI.UnpackIntoMap(eventData, "TransferBatch", pLog.Data)
			if err != nil {
				panic(err)
			}

			ids, ok := eventData["ids"].([]*big.Int)
			if !ok {
				panic("Failed to unpack TransferBatch event, ids not found")
			}

			values, ok := eventData["values"].([]*big.Int)
			if !ok {
				panic("Failed to unpack TransferBatch event, values not found")
			}

			for j := 0; j < len(ids); j++ {
				result = append(result, &TokenTransfer{
					From:        pLog.Topics[2].Hex(),
					To:          pLog.Topics[3].Hex(),
					ChainType:   basetype.ChainType_Ethereum,
					Contract:    pLog.Address.Hex(),
					TokenID:     ids[j].String(),
					Amount:      values[j].Uint64(),
					BlockNumber: pLog.BlockNumber,
					TokenType:   basetype.TokenType_ERC1155,
					TxHash:      pLog.TxHash.Hex(),
					BlockHash:   pLog.BlockHash.Hex(),
				})
			}
		}
	}
	return result, nil
}

func TransferLogs(ctx context.Context, fromBlock, toBlock int64) ([]*TokenTransfer, error) {
	topics := [][]common.Hash{{
		common.HexToHash(string(transferEventHash)),
		common.HexToHash(string(transferSingleEventHash)),
		common.HexToHash(string(transferBatchEventHash)),
	}}

	logs, err := cteth.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Topics:    topics,
	})
	if err != nil {
		return nil, err
	}

	return logsToTransfer(logs)
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
