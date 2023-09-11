package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	TransferLogs, OrderLogs []types.Log
}

func (e *EthIndexer) CheckBlock(ctx context.Context, inBlockNum uint64) (*blockProto.Block, error) {
	blockOnly, err := blockNMCli.GetBlockOnly(ctx, &blockProto.GetBlockOnlyRequest{
		Conds: &blockProto.Conds{
			ChainType:   &ctMessage.StringVal{Op: "eq", Value: e.ChainType.String()},
			ChainID:     &ctMessage.StringVal{Op: "eq", Value: e.ChainID},
			BlockNumber: &ctMessage.Uint64Val{Op: "eq", Value: inBlockNum},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get block only,err: %v", err)
	}

	if blockOnly.Info != nil && blockOnly.GetInfo().ParseState == basetype.BlockParseState_BlockTypeFinish {
		return blockOnly.GetInfo(), nil
	}

	cli, err := eth.Client(e.Endpoints)
	if err != nil {
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}

	block, err := cli.BlockByNumber(ctx, inBlockNum)
	if err != nil {
		e.checkErr(ctx, err)
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}

	number := block.Number().Uint64()
	blockHash := block.Hash().String()
	blockTime := int64(block.Time())
	remark := ""
	resp, err := blockNMCli.UpsertBlock(ctx, &blockProto.UpsertBlockRequest{
		Info: &blockProto.BlockReq{
			ChainType:   &e.ChainType,
			ChainID:     &e.ChainID,
			BlockNumber: &number,
			BlockHash:   &blockHash,
			BlockTime:   &blockTime,
			ParseState:  basetype.BlockParseState_BlockTypeStart.Enum(),
			Remark:      &remark,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}
	return resp.Info, nil
}

func (e *EthIndexer) IndexBlockLogs(ctx context.Context, inBlockNum uint64) (*BlockLogs, error) {
	cli, err := eth.Client(e.Endpoints)
	if err != nil {
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}

	topicsLogs, err := cli.FilterLogsForTopics(context.Background(), int64(inBlockNum), int64(inBlockNum), [][]common.Hash{
		eth.TransfersTopics,
		eth.OrderFulfilledTopics,
	})
	if err != nil {
		e.checkErr(ctx, err)
		return nil, fmt.Errorf("cannot parse logs: %v", err)
	}
	transferLogs := topicsLogs[0]
	orderLogs := topicsLogs[1]

	return &BlockLogs{
		TransferLogs: transferLogs,
		OrderLogs:    orderLogs,
	}, nil
}

func (e *EthIndexer) IndexTransfer(ctx context.Context, logs []types.Log) ([]*chains.TokenTransfer, error) {
	transfers, err := eth.LogsToTransfer(logs)
	if err != nil {
		e.checkErr(ctx, err)
		return nil, fmt.Errorf("failed to get transfer logs, err: %v", err)
	}
	if len(transfers) == 0 {
		return nil, nil
	}

	transfersMap := make(map[string]struct{}, len(transfers))
	infos := make([]*transferProto.TransferReq, len(transfers))

	for i := range transfers {

		transIdentifier := transferIdentifier(
			transfers[i].Contract,
			transfers[i].TokenID,
			transfers[i].TxHash,
			transfers[i].From)
		// just for avoid  repetition,some token will be transfer many times
		if _, ok := transfersMap[transIdentifier]; ok {
			continue
		}
		transfersMap[transIdentifier] = struct{}{}

		tokenType := string(transfers[i].TokenType)
		infos[i] = &transferProto.TransferReq{
			ChainType:   &e.ChainType,
			ChainID:     &e.ChainID,
			Contract:    &transfers[i].Contract,
			TokenType:   &tokenType,
			TokenID:     &transfers[i].TokenID,
			From:        &transfers[i].From,
			To:          &transfers[i].To,
			Amount:      &transfers[i].Amount,
			BlockNumber: &transfers[i].BlockNumber,
			TxHash:      &transfers[i].TxHash,
			BlockHash:   &transfers[i].BlockHash,
		}
	}

	_, err = transferNMCli.UpsertTransfers(ctx, &transferProto.UpsertTransfersRequest{Infos: infos})
	if err != nil {
		return nil, fmt.Errorf("failed store transfers to db ,err: %v", err)
	}

	return transfers, nil
}

func (e *EthIndexer) IndexToken(ctx context.Context, inTransfers []*chains.TokenTransfer) ([]*chains.TokenTransfer, error) {
	outTransfer := []*chains.TokenTransfer{}
	for _, transfer := range inTransfers {
		identifier := tokenIdentifier(e.ChainType, e.ChainID, transfer.Contract, transfer.TokenID)
		locked, err := ctredis.TryPubLock(identifier, redisExpireDefaultTime)
		if err != nil {
			return nil, fmt.Errorf("lock the token indentifier failed, err: %v", err)
		}

		if !locked {
			continue
		}

		remark := ""
		conds := &tokenProto.Conds{
			ChainType: &ctMessage.StringVal{
				Value: e.ChainType.String(),
				Op:    "eq",
			},
			ChainID: &ctMessage.StringVal{
				Value: e.ChainID,
				Op:    "eq",
			},
			Contract: &ctMessage.StringVal{
				Value: transfer.Contract,
				Op:    "eq",
			},
			TokenID: &ctMessage.StringVal{
				Value: transfer.TokenID,
				Op:    "eq",
			},
		}

		if resp, err := tokenNMCli.ExistTokenConds(ctx, &tokenProto.ExistTokenCondsRequest{Conds: conds}); err == nil && resp != nil && resp.GetExist() {
			continue
		} else if err != nil {
			return nil, fmt.Errorf("check if the token exist failed, err: %v", err)
		}

		cli, err := eth.Client(e.Endpoints)
		if err != nil {
			return nil, fmt.Errorf("cannot get eth client,err: %v", err)
		}

		tokenURI, err := cli.TokenURI(ctx, transfer.TokenType, transfer.Contract, transfer.TokenID, transfer.BlockNumber)
		if err != nil {
			e.checkErr(ctx, err)
			logger.Sugar().Warnf("cannot get tokenURI,err: %v", err)
			remark = fmt.Sprintf("%v,%v", remark, err)
		}

		tokenURIInfo, err := token.GetTokenURIInfo(ctx, tokenURI)
		if err != nil {
			tokenURIInfo = &token.TokenURIInfo{}
			remark = fmt.Sprintf("%v,%v", remark, err)
		}

		_, err = tokenNMCli.UpsertToken(ctx, &tokenProto.UpsertTokenRequest{
			Info: &tokenProto.TokenReq{
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
				VectorState: tokenProto.ConvertState_Waiting.Enum(),
				Remark:      &remark,
			},
		})

		if err != nil {
			fmt.Println(tokenURI)
			return nil, fmt.Errorf("create token record failed, %v", err)
		}
		outTransfer = append(outTransfer, transfer)
	}
	return outTransfer, nil
}

func (e *EthIndexer) IndexContract(ctx context.Context, inTransfers []*chains.TokenTransfer, findContractCreator bool) error {
	for _, transfer := range inTransfers {
		identifier := contractIdentifier(e.ChainType, e.ChainID, transfer.Contract)
		locked, err := ctredis.TryPubLock(identifier, redisExpireDefaultTime)
		if err != nil {
			return fmt.Errorf("lock the token indentifier failed, err: %v", err)
		}
		if !locked {
			continue
		}
		// check if the record exist
		conds := &contractProto.Conds{
			ChainType: &ctMessage.StringVal{
				Value: e.ChainType.String(),
				Op:    "eq",
			},
			ChainID: &ctMessage.StringVal{
				Value: e.ChainID,
				Op:    "eq",
			},
			Address: &ctMessage.StringVal{
				Value: transfer.Contract,
				Op:    "eq",
			},
		}
		if resp, err := contractNMCli.ExistContractConds(ctx, &contractProto.ExistContractCondsRequest{
			Conds: conds,
		}); err == nil && resp != nil && resp.GetExist() {
			continue
		} else if err != nil {
			return fmt.Errorf("check if the contract exist failed, err: %v", err)
		}

		// parse the token metadata
		remark := ""
		cli, err := eth.Client(e.Endpoints)
		if err != nil {
			return fmt.Errorf("cannot get eth client,err: %v", err)
		}

		contractMeta, err := cli.GetERC721Metadata(ctx, transfer.Contract)
		if err != nil {
			e.checkErr(ctx, err)
			logger.Sugar().Warnf("cannot get contrcact metadata,err: %v", err)
			contractMeta = &eth.ERC721Metadata{}
			remark = fmt.Sprintf("%v,%v", remark, err)
		}

		creator := &eth.ContractCreator{}
		// stop get info for creator
		if findContractCreator {
			creator, err = cli.GetContractCreator(ctx, transfer.Contract)
			if err != nil {
				e.checkErr(ctx, err)
				remark = err.Error()
			}
		}

		// store the result
		from := creator.From.String()
		txHash := creator.TxHash.Hex()
		blockNum := creator.BlockNumber
		txTime := uint32(creator.TxTime)
		_, err = contractNMCli.UpsertContract(ctx, &contractProto.UpsertContractRequest{
			Info: &contractProto.ContractReq{
				ChainType: &e.ChainType,
				ChainID:   &e.ChainID,
				Address:   &transfer.Contract,
				Name:      &contractMeta.Name,
				Symbol:    &contractMeta.Symbol,
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

func (e *EthIndexer) GetCurrentBlockNum(ctx context.Context, updateInterval time.Duration) {
	for {
		func() {
			cli, err := eth.Client(e.Endpoints)
			if err != nil {
				logger.Sugar().Errorf("cannot get eth client,err: %v", err)
				return
			}

			blockNum, err := cli.CurrentBlockNum(ctx)
			if err != nil {
				e.checkErr(ctx, err)
				logger.Sugar().Errorf("failed to get current block number: %v", err)
				return
			}

			e.CurrentBlockNum = blockNum
			logger.Sugar().Infof("success get current block number: %v", blockNum)
		}()

		select {
		case <-time.NewTicker(updateInterval).C:
			continue
		case <-ctx.Done():
			return
		}
	}
}
