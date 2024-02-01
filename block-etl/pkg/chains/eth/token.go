package eth

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/indexer"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/token"
	"github.com/web3eye-io/Web3Eye/common/chains"
	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	blockNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/block"
	contractNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
	tokenNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	transferNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/transfer"
	ctMessage "github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	blockProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
	contractProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
	tokenProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	transferProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

type BlockLogs struct {
	TransferLogs, OrderLogs []*types.Log
}

func (e *EthIndexer) CheckBlock(ctx context.Context, inBlockNum uint64) (*blockProto.Block, error) {
	// ignore err,just adjust if blockOnly is finished
	blockOnly, _ := blockNMCli.GetBlockOnly(ctx, &blockProto.GetBlockOnlyRequest{
		Conds: &blockProto.Conds{
			ChainType:   &ctMessage.Uint32Val{Op: "eq", Value: uint32(e.ChainType)},
			ChainID:     &ctMessage.StringVal{Op: "eq", Value: e.ChainID},
			BlockNumber: &ctMessage.Uint64Val{Op: "eq", Value: inBlockNum},
		},
	})

	if blockOnly != nil &&
		blockOnly.Info != nil &&
		blockOnly.GetInfo().ParseState == basetype.BlockParseState_BlockTypeFinish {
		return blockOnly.GetInfo(), nil
	}

	cli, err := eth.Client(e.OkEndpoints)
	if err != nil {
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}

	block, err := cli.BlockByNumber(ctx, big.NewInt(0).SetUint64(inBlockNum))
	if err != nil {
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}

	number := inBlockNum
	var blockHash = ""
	var blockTime uint64 = 0
	var parseState = basetype.BlockParseState_BlockTypeFailed.Enum()
	remark := ""
	if block != nil {
		blockHash = block.Hash().String()
		blockTime = block.Time()
		parseState = basetype.BlockParseState_BlockTypeStart.Enum()
		remark = "start to parse the block"
	}

	resp, err := blockNMCli.UpsertBlock(ctx, &blockProto.UpsertBlockRequest{
		Info: &blockProto.BlockReq{
			ChainType:   &e.ChainType,
			ChainID:     &e.ChainID,
			BlockNumber: &number,
			BlockHash:   &blockHash,
			BlockTime:   &blockTime,
			ParseState:  parseState,
			Remark:      &remark,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}
	return resp.Info, nil
}

func (e *EthIndexer) IndexBlockLogs(ctx context.Context, inBlockNum uint64) (*BlockLogs, error) {
	cli, err := eth.Client(e.OkEndpoints)
	if err != nil {
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}

	topicsLogs, err := cli.FilterLogsForTopics(context.Background(), int64(inBlockNum), int64(inBlockNum), [][]common.Hash{
		eth.TransfersTopics,
		eth.OrderFulfilledTopics,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot parse logs: %v", err)
	}
	transferLogs := topicsLogs[0]
	orderLogs := topicsLogs[1]

	return &BlockLogs{
		TransferLogs: transferLogs,
		OrderLogs:    orderLogs,
	}, nil
}

func (e *EthIndexer) IndexTransfer(ctx context.Context, logs []*types.Log, blockTime uint64) ([]*chains.TokenTransfer, error) {
	transfers, err := eth.LogsToTransfer(logs)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer logs, err: %v", err)
	}
	if len(transfers) == 0 {
		return nil, nil
	}

	transfersMap := make(map[string]struct{}, len(transfers))
	infos := make([]*transferProto.TransferReq, len(transfers))

	for i := range transfers {
		transIdentifier := indexer.TransferIdentifier(
			transfers[i].Contract,
			transfers[i].TokenID,
			transfers[i].TxHash,
			transfers[i].From,
			transfers[i].LogIndex,
		)
		// just for avoid  repetition,some token will be transfer many times
		if _, ok := transfersMap[transIdentifier]; ok {
			continue
		}
		transfersMap[transIdentifier] = struct{}{}
		infos[i] = &transferProto.TransferReq{
			ChainType:   &e.ChainType,
			ChainID:     &e.ChainID,
			Contract:    &transfers[i].Contract,
			TokenType:   &transfers[i].TokenType,
			TokenID:     &transfers[i].TokenID,
			From:        &transfers[i].From,
			To:          &transfers[i].To,
			Amount:      &transfers[i].Amount,
			BlockNumber: &transfers[i].BlockNumber,
			TxHash:      &transfers[i].TxHash,
			TxTime:      &blockTime,
			LogIndex:    &transfers[i].LogIndex,
			BlockHash:   &transfers[i].BlockHash,
		}
	}

	_, err = transferNMCli.UpsertTransfers(ctx, &transferProto.UpsertTransfersRequest{Infos: infos})
	if err != nil {
		return nil, fmt.Errorf("failed store transfers to db ,err: %v", err)
	}

	return transfers, nil
}

func (e *EthIndexer) IndexToken(ctx context.Context, inTransfers []*chains.TokenTransfer) ([]*ContractMeta, error) {
	outContractMetas := []*ContractMeta{}
	for _, transfer := range inTransfers {
		identifier := indexer.TokenIdentifier(e.ChainType, e.ChainID, transfer.Contract, transfer.TokenID)
		locked, err := ctredis.TryPubLock(identifier, redisExpireDefaultTime)
		if err != nil {
			return nil, fmt.Errorf("lock the token indentifier failed, err: %v", err)
		}

		if !locked {
			continue
		}

		remark := ""
		conds := &tokenProto.Conds{
			ChainType: &ctMessage.Uint32Val{
				Value: uint32(e.ChainType),
				Op:    cruder.EQ,
			},
			ChainID: &ctMessage.StringVal{
				Value: e.ChainID,
				Op:    cruder.EQ,
			},
			Contract: &ctMessage.StringVal{
				Value: transfer.Contract,
				Op:    cruder.EQ,
			},
			TokenID: &ctMessage.StringVal{
				Value: transfer.TokenID,
				Op:    cruder.EQ,
			},
			URIState: &ctMessage.Uint32Val{
				Value: uint32(basetype.BlockParseState_BlockTypeFinish),
				Op:    cruder.EQ,
			},
			VectorState: &ctMessage.Uint32Val{
				Value: uint32(tokenProto.ConvertState_Success),
				Op:    cruder.EQ,
			},
		}

		if resp, err := tokenNMCli.ExistTokenConds(ctx, &tokenProto.ExistTokenCondsRequest{Conds: conds}); err == nil && resp != nil && resp.GetExist() {
			continue
		} else if err != nil {
			return nil, fmt.Errorf("check if the token exist failed, err: %v", err)
		}

		cli, err := eth.Client(e.OkEndpoints)
		if err != nil {
			return nil, fmt.Errorf("cannot get eth client,err: %v", err)
		}

		tokenURI, err := cli.TokenURI(ctx, transfer.TokenType, transfer.Contract, transfer.TokenID, transfer.BlockNumber)
		if err != nil {
			remark = fmt.Sprintf("%v,%v", remark, err)
		}

		tokenURIInfo, err := token.GetTokenURIInfo(ctx, tokenURI)
		if err != nil {
			// if cannot get tokenURIInfo,then set the default value
			tokenURIInfo = &token.TokenURIInfo{}
			remark = fmt.Sprintf("%v,%v", remark, err)
		}

		tokenReq := token.CheckTokenReq(&tokenProto.TokenReq{
			ChainType:   &e.ChainType,
			ChainID:     &e.ChainID,
			Contract:    &transfer.Contract,
			TokenType:   &transfer.TokenType,
			TokenID:     &transfer.TokenID,
			URI:         &tokenURI,
			URIType:     (*string)(&tokenURIInfo.URIType),
			ImageURL:    &tokenURIInfo.ImageURL,
			VideoURL:    &tokenURIInfo.VideoURL,
			Name:        &tokenURIInfo.Name,
			Description: &tokenURIInfo.Description,
			Remark:      &remark,
		})

		_, err = tokenNMCli.UpsertToken(ctx, &tokenProto.UpsertTokenRequest{
			Info: tokenReq,
		})

		if err != nil {
			return nil, fmt.Errorf("create token record failed, %v", err)
		}
		outContractMetas = append(outContractMetas, &ContractMeta{
			TokenType: transfer.TokenType,
			Contract:  transfer.Contract,
		})
	}
	return outContractMetas, nil
}

type ContractMeta struct {
	TokenType basetype.TokenType
	Contract  string
}

func (e *EthIndexer) IndexContract(ctx context.Context, inContracts []*ContractMeta, findContractCreator bool) error {
	for _, contract := range inContracts {
		exist, err := e.checkContract(ctx, contract.Contract)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		if exist {
			continue
		}
		remark := ""
		contractMeta, creator, err := e.getContractInfo(ctx, contract, findContractCreator)
		if err != nil {
			contractMeta = &eth.EthCurrencyMetadata{}
			creator = &chains.ContractCreator{}
			remark = err.Error()
		}
		// store the result
		from := creator.From
		txHash := creator.TxHash
		blockNum := creator.BlockNumber
		txTime := uint32(creator.TxTime)
		_, err = contractNMCli.UpsertContract(ctx, &contractProto.UpsertContractRequest{
			Info: &contractProto.ContractReq{
				ChainType: &e.ChainType,
				ChainID:   &e.ChainID,
				Address:   &contract.Contract,
				Name:      &contractMeta.Name,
				Symbol:    &contractMeta.Symbol,
				Decimals:  &contractMeta.Decimals,
				Creator:   &from,
				BlockNum:  &blockNum,
				TxHash:    &txHash,
				TxTime:    &txTime,
				Remark:    &remark,
			},
		})
		if err != nil {
			return fmt.Errorf("create contract record failed, %v", err)
		}
	}

	return nil
}

func (e *EthIndexer) checkContract(ctx context.Context, contract string) (exist bool, err error) {
	identifier := indexer.ContractIdentifier(e.ChainType, e.ChainID, contract)
	locked, err := ctredis.TryPubLock(identifier, redisExpireDefaultTime)
	if err != nil {
		return false, fmt.Errorf("lock the token indentifier failed, err: %v", err)
	}
	if !locked {
		return true, nil
	}
	// check if the record exist
	conds := &contractProto.Conds{
		ChainType: &ctMessage.Uint32Val{
			Value: uint32(e.ChainType),
			Op:    "eq",
		},
		ChainID: &ctMessage.StringVal{
			Value: e.ChainID,
			Op:    "eq",
		},
		Address: &ctMessage.StringVal{
			Value: contract,
			Op:    "eq",
		},
	}
	if resp, err := contractNMCli.ExistContractConds(ctx, &contractProto.ExistContractCondsRequest{
		Conds: conds,
	}); err == nil && resp != nil && resp.GetExist() {
		return true, nil
	} else if err != nil {
		return false, fmt.Errorf("check if the contract exist failed, err: %v", err)
	}
	return false, nil
}

func (e *EthIndexer) getContractInfo(ctx context.Context, contract *ContractMeta, findContractCreator bool) (*eth.EthCurrencyMetadata, *chains.ContractCreator, error) {
	cli, err := eth.Client(e.OkEndpoints)
	if err != nil {
		return &eth.EthCurrencyMetadata{}, &chains.ContractCreator{}, fmt.Errorf("cannot get eth client,err: %v", err)
	}
	contractMeta := &eth.EthCurrencyMetadata{}
	switch contract.TokenType {
	case basetype.TokenType_Native:
		contractMeta, err = cli.GetCurrencyMetadata(ctx, contract.Contract)
	case basetype.TokenType_ERC20:
		contractMeta, err = cli.GetERC20Metadata(ctx, contract.Contract)
	default:
		contractMeta, err = cli.GetERC721Metadata(ctx, contract.Contract)
	}

	if err != nil {
		return &eth.EthCurrencyMetadata{}, &chains.ContractCreator{}, fmt.Errorf("cannot get eth client,err: %v", err)
	}
	creator := &chains.ContractCreator{}
	// stop get info for creator
	if findContractCreator && contract.TokenType != basetype.TokenType_Native {
		creator, err = cli.GetContractCreator(ctx, contract.Contract)
		if err != nil {
			return &eth.EthCurrencyMetadata{}, &chains.ContractCreator{}, fmt.Errorf("cannot get eth client,err: %v", err)
		}
	}
	return contractMeta, creator, nil
}

func (e *EthIndexer) GetCurrentBlockNum() uint64 {
	return e.CurrentBlockNum
}

func (e *EthIndexer) SyncCurrentBlockNum(ctx context.Context, updateInterval time.Duration) {
	for {
		func() {
			cli, err := eth.Client(e.OkEndpoints)
			if err != nil {
				logger.Sugar().Errorf("eth cannot get eth client,err: %v", err)
				return
			}

			blockNum, err := cli.CurrentBlockNum(ctx)
			if err != nil {
				logger.Sugar().Errorf("eth failed to get current block number: %v", err)
				return
			}

			e.CurrentBlockNum = blockNum
			logger.Sugar().Infof("eth success get current block number: %v", blockNum)
		}()

		select {
		case <-time.NewTicker(updateInterval).C:
			continue
		case <-ctx.Done():
			return
		}
	}
}
