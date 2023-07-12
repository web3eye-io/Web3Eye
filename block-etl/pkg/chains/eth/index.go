package eth

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/token"
	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	blockNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/block"
	contractNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
	synctaskNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/synctask"
	tokenNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	transferNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/transfer"
	ctMessage "github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	blockProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
	contractProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
	tokenProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	transferProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

const (
	Retries                   = 3
	MaxDealBlockNum           = 5
	LogIntervalHeight         = 10
	MaxRetriesThenTriggerTask = 5
	MaxRetriesThenStopConsume = 10
	CheckTopicInterval        = time.Second * 10
	// confirmedBlockNum      = 5
	redisExpireDefaultTime = time.Second * 10
	// redisExpireDefaultTime = time.Minute * 5
	maxTokenURILen = 256
	maxTopicNum    = 5
)

var (
	connectTimeout = "context deadline exceeded"
	retryErrs      = []string{connectTimeout}
)

var (
	// manage process order
	pLock = &sync.Mutex{}
	wg    = &sync.WaitGroup{}
)

var GetEndpointChainID = func(ctx context.Context, endpoint string) (string, error) {
	cli, err := ethclient.DialContext(ctx, endpoint)
	if err != nil {
		return "", err
	}

	chainID, err := cli.ChainID(ctx)
	if err != nil {
		return "", err
	}

	return chainID.String(), nil
}

type EthIndexer struct {
	Endpoints     []string
	ChainID       string
	ChainType     basetype.ChainType
	taskChan      chan uint64
	taskMap       map[string]struct{}
	currentHeight uint64
	updateTime    int64
}

func NewIndexer(endpoints []string, chainID string) *EthIndexer {
	return &EthIndexer{
		ChainType:     basetype.ChainType_Ethereum,
		Endpoints:     endpoints,
		ChainID:       chainID,
		taskChan:      make(chan uint64),
		taskMap:       make(map[string]struct{}),
		currentHeight: 0,
		updateTime:    0,
	}
}

func (e *EthIndexer) StartIndex(ctx context.Context) {
	taskBlockNum := make(chan uint64)
	indexBlockNum := make(chan uint64)
	go e.IndexTasks(ctx, taskBlockNum)
	go e.IndexBlock(ctx, taskBlockNum, indexBlockNum)
	for v := range indexBlockNum {
		fmt.Println(v)
	}
}

func (e *EthIndexer) IndexTasks(ctx context.Context, outBlockNum chan uint64) {
	logger.Sugar().Info("start to index task for ethereum")

	conds := &synctask.Conds{
		ChainType: &ctMessage.StringVal{
			Value: basetype.ChainType_Ethereum.String(),
			Op:    "eq",
		},
		SyncState: &ctMessage.StringVal{
			Value: cttype.SyncState_Finish.String(),
			Op:    "eq",
		},
	}

	logger.Sugar().Info("check if there are syncTasks with state is finsh in kafka")
	// check wheather have task in kafka
	resp, err := synctaskNMCli.GetSyncTasks(ctx, &synctask.GetSyncTasksRequest{Conds: conds, Offset: 0, Limit: 0})
	if err != nil {
		logger.Sugar().Error(err)
	}
	for _, v := range resp.GetInfos() {
		err = e.addTask(v.Topic, outBlockNum)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
	}

	conds.SyncState.Value = cttype.SyncState_Start.String()
	logger.Sugar().Info("start get syncTasks")
	for {
		resp, err := synctaskNMCli.GetSyncTasks(ctx, &synctask.GetSyncTasksRequest{Conds: conds, Offset: 0, Limit: maxTopicNum})
		if err != nil {
			logger.Sugar().Error(err)
		}
		for _, v := range resp.GetInfos() {
			err = e.addTask(v.Topic, outBlockNum)
			if err != nil {
				logger.Sugar().Error(err)
				continue
			}
		}
		time.Sleep(CheckTopicInterval)
	}
}

func (e *EthIndexer) addTask(topic string, outBlockNum chan uint64) error {
	pLock.Lock()
	defer pLock.Unlock()

	if _, ok := e.taskMap[topic]; ok {
		return nil
	}
	e.taskMap[topic] = struct{}{}
	defer delete(e.taskMap, topic)

	resp, err := synctaskNMCli.TriggerSyncTask(context.Background(), &synctask.TriggerSyncTaskRequest{Topic: topic})
	if err != nil {
		return err
	}

	for _, v := range resp.BlockNums {
		outBlockNum <- v
	}

	return nil
}

func (e *EthIndexer) IndexBlock(ctx context.Context, inBlockNum, outBlockNum chan uint64) {
	for blockNum := range inBlockNum {
		cli, err := eth.Client(e.Endpoints)
		if err != nil {
			logger.Sugar().Errorf("cannot get eth client,err: %v", err)
			continue
		}
		block, err := cli.BlockByNumber(ctx, blockNum)
		if err != nil {
			logger.Sugar().Errorf("cannot get eth client,err: %v", err)
			continue
		}
		number := block.Number().Uint64()
		blockHash := block.Hash().String()
		blockTime := int64(block.Time())
		_, err = blockNMCli.UpsertBlock(ctx, &blockProto.UpsertBlockRequest{
			Info: &blockProto.BlockReq{
				ChainType:   basetype.ChainType_Ethereum.Enum(),
				ChainID:     &e.ChainID,
				BlockNumber: &number,
				BlockHash:   &blockHash,
				BlockTime:   &blockTime,
			},
		})
		if err != nil {
			logger.Sugar().Errorf("cannot get eth client,err: %v", err)
			continue
		}
		outBlockNum <- block.Number().Uint64()
	}
}

func (e *EthIndexer) IndexTransfer(ctx context.Context, inBlockNum chan uint64, outTransfers chan []*eth.TokenTransfer) {
	for num := range inBlockNum {
		cli, err := eth.Client(e.Endpoints)
		if err != nil {
			logger.Sugar().Errorf("cannot get eth client,err: %v", err)
			continue
		}
		transfers, err := cli.TransferLogs(ctx, int64(num), int64(num))
		if err != nil {
			logger.Sugar().Errorf("failed to get transfer logs, err: %v, block: %v", err, num)
			continue
		}
		fmt.Println(len(transfers))
		if len(transfers) == 0 {
			continue
		}

		infos := make([]*transferProto.TransferReq, len(transfers))
		for i := range transfers {
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
			logger.Sugar().Errorf("failed store transfers to db for block number: %v,err: %v", num, err)
			continue
		}

		outTransfers <- transfers
	}
}

func (e *EthIndexer) tokenInfoToDB(ctx context.Context, transfers []*eth.TokenTransfer) {
	for _, transfer := range transfers {
		identifier := tokenIdentifier(e.ChainType, e.ChainID, transfer.Contract, transfer.TokenID)
		if err := ctredis.TryPubLock(identifier, redisExpireDefaultTime); err != nil {
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
			logger.Sugar().Error(err)
		}

		// TODO: use channel
		err := e.contractToDB(ctx, transfer)
		if err != nil {
			logger.Sugar().Error(err)
		}

		cli, err := eth.Client(e.Endpoints)
		if err != nil {
			logger.Sugar().Error(err)
		}

		tokenURI, err := cli.TokenURI(ctx, transfer.TokenType, transfer.Contract, transfer.TokenID, transfer.BlockNumber)
		if err != nil {
			remark = err.Error()
		}

		tokenURIInfo, err := token.GetTokenURIInfo(ctx, tokenURI)
		if err != nil {
			tokenURIInfo = &token.TokenURIInfo{}
		}

		if len(tokenURI) > maxTokenURILen {
			tokenURI = ""
		}

		for i := 0; i < Retries; i++ {
			_, err = tokenNMCli.CreateToken(ctx, &tokenProto.CreateTokenRequest{
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
			if err != nil && containErr(err.Error()) {
				logger.Sugar().Errorf("will retry for creating token record failed, %v", err)
				continue
			}

			if err != nil {
				logger.Sugar().Errorf("create token record failed, %v", err)
			}
			break
		}
	}
}

func (e *EthIndexer) contractToDB(ctx context.Context, transfer *eth.TokenTransfer) error {
	identifier := contractIdentifier(e.ChainType, e.ChainID, transfer.Contract)
	if err := ctredis.TryPubLock(identifier, redisExpireDefaultTime); err != nil {
		return nil
	}

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
		return nil
	} else if err != nil {
		logger.Sugar().Error(err)
	}

	remark := ""
	cli, err := eth.Client(e.Endpoints)
	if err != nil {
		logger.Sugar().Error(err)
	}

	contractMeta, err := cli.GetERC721Metadata(ctx, transfer.Contract)
	if err != nil {
		contractMeta = &eth.ERC721Metadata{}
		remark = fmt.Sprintf("%v,%v", remark, err)
	}

	// stop get info for creator
	// creator, err := cteth.GetContractCreator(ctx, transfer.Contract)
	// if err != nil {
	// 	creator = &cteth.ContractCreator{}
	// 	remark = err.Error()
	// }

	// from := creator.From.String()
	// txHash := creator.TxHash.Hex()
	// blockNum := creator.BlockNumber
	// txTime := uint32(creator.TxTime)
	for i := 0; i < Retries; i++ {
		_, err = contractNMCli.CreateContract(ctx, &contractProto.CreateContractRequest{
			Info: &contractProto.ContractReq{
				ChainType: &e.ChainType,
				ChainID:   &e.ChainID,
				Address:   &transfer.Contract,
				Name:      &contractMeta.Name,
				Symbol:    &contractMeta.Symbol,
				// Creator:   &from,
				// BlockNum:  &blockNum,
				// TxHash:    &txHash,
				// TxTime:    &txTime,
				Remark: &remark,
			},
		})
		if err != nil && containErr(err.Error()) {
			logger.Sugar().Errorf("will retry for creating contract record failed, %v", err)
			continue
		}

		if err != nil {
			return fmt.Errorf("create contract record failed, %v", err)
		}
		break
	}
	return nil
}

func containErr(errStr string) bool {
	for _, v := range retryErrs {
		if strings.Contains(errStr, v) {
			return true
		}
	}
	return false
}

func tokenIdentifier(chain basetype.ChainType, chainID, contract, tokenID string) string {
	return fmt.Sprintf("%v:%v:%v:%v", chain, chainID, contract, tokenID)
}

func contractIdentifier(chain basetype.ChainType, chainID, contract string) string {
	return fmt.Sprintf("%v+%v+%v", chain, chainID, contract)
}
