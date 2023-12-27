package eth

import (
	"context"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/web3eye-io/Web3Eye/common/chains"
	"github.com/web3eye-io/Web3Eye/common/chains/eth/contracts"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
)

const (
	// transferEventHash represents the keccak256 hash of Transfer(address,address,uint256)
	transferEventHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	// transferSingleEventHash represents the keccak256 hash of TransferSingle(address,address,address,uint256,uint256)
	transferSingleEventHash = "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	// transferBatchEventHash represents the keccak256 hash of TransferBatch(address,address,address,uint256[],uint256[])
	transferBatchEventHash = "0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"

	transferEventArgsLen = 4
	BaseTextNum          = 16
)

var (
	erc1155ABI, _   = contracts.IERC1155MetaData.GetAbi()
	TransfersTopics = []common.Hash{
		common.HexToHash(transferEventHash),
		// common.HexToHash(transferSingleEventHash),
		// common.HexToHash(transferBatchEventHash),
	}
)

func LogsToTransfer(pLogs []*types.Log) ([]*chains.TokenTransfer, error) {
	result := make([]*chains.TokenTransfer, 0)
	for _, pLog := range pLogs {
		switch {
		case strings.EqualFold(pLog.Topics[0].Hex(), string(transferEventHash)):
			// filter erc20
			if len(pLog.Topics) < transferEventArgsLen {
				continue
			}

			result = append(result, &chains.TokenTransfer{
				From:        common.HexToAddress(pLog.Topics[1].Hex()).String(),
				To:          common.HexToAddress(pLog.Topics[2].Hex()).String(),
				Contract:    pLog.Address.Hex(),
				TokenID:     pLog.Topics[3].Big().String(),
				BlockNumber: pLog.BlockNumber,
				Amount:      1,
				TokenType:   basetype.TokenType_ERC721,
				TxHash:      pLog.TxHash.Hex(),
				BlockHash:   pLog.BlockHash.Hex(),
				LogIndex:    uint32(pLog.Index),
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
			result = append(result, &chains.TokenTransfer{
				From:        common.HexToAddress(pLog.Topics[2].Hex()).String(),
				To:          common.HexToAddress(pLog.Topics[3].Hex()).String(),
				Contract:    pLog.Address.Hex(),
				TokenID:     id.String(),
				Amount:      value.Uint64(),
				BlockNumber: pLog.BlockNumber,
				TokenType:   basetype.TokenType_ERC1155,
				TxHash:      pLog.TxHash.Hex(),
				BlockHash:   pLog.BlockHash.Hex(),
				LogIndex:    uint32(pLog.Index),
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
				result = append(result, &chains.TokenTransfer{
					From:        common.HexToAddress(pLog.Topics[2].Hex()).String(),
					To:          common.HexToAddress(pLog.Topics[3].Hex()).String(),
					Contract:    pLog.Address.Hex(),
					TokenID:     ids[j].String(),
					Amount:      values[j].Uint64(),
					BlockNumber: pLog.BlockNumber,
					TokenType:   basetype.TokenType_ERC1155,
					TxHash:      pLog.TxHash.Hex(),
					BlockHash:   pLog.BlockHash.Hex(),
					LogIndex:    uint32(pLog.Index),
				})
			}
		}
	}
	return result, nil
}

// get the logs from chain
// overwrite the ethclient.FilterLogs,for filter diffent topics from logs
// if topics={{}} retrun {{all logs}}
// if topics={{A}} retrun {{A logs}}
// if topics={{A,B,C}} return {{A or B or C logs}}
// if topics={{A,B,C},{D}} return {{A or B or C logs},{D logs}}
func (ethCli *ethClients) FilterLogsForTopics(ctx context.Context, fromBlock, toBlock int64, topics [][]common.Hash) ([][]*types.Log, error) {
	// the native func filterLogs cannot filter {{A...},{B...}}
	logs, err := ethCli.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Topics:    [][]common.Hash{},
	})

	if err != nil {
		return nil, err
	}

	if len(topics) == 0 {
		return [][]*types.Log{logs}, err
	}

	// init topicSets
	topicSets := make([]map[common.Hash]struct{}, len(topics))
	topicLogs := make([][]*types.Log, len(topics))
	allSets := make(map[common.Hash]struct{})
	for i, items := range topics {
		topicSets[i] = make(map[common.Hash]struct{}, len(items))
		topicLogs[i] = make([]*types.Log, 0)
		for _, item := range items {
			topicSets[i][item] = struct{}{}
			allSets[item] = struct{}{}
		}
	}

	for _, v := range logs {
		if len(v.Topics) == 0 {
			continue
		}
		if _, ok := allSets[v.Topics[0]]; !ok {
			continue
		}
		for i, topic := range topicSets {
			if _, ok := topic[v.Topics[0]]; ok {
				topicLogs[i] = append(topicLogs[i], v)
			}
		}
	}

	return topicLogs, nil
}
