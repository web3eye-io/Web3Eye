package eth

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/token"
	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/common/ctkafka"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	"github.com/web3eye-io/Web3Eye/common/utils"
	contractNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
	synctaskNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/synctask"
	tokenNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	transferNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/transfer"
	ctMessage "github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
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

type EthIndexer struct {
	taskChan      chan uint64
	taskMap       map[string]struct{}
	currentHeight uint64
	updateTime    int64
}

func NewIndexer() (*EthIndexer, error) {
	return &EthIndexer{
		taskChan:      make(chan uint64),
		taskMap:       make(map[string]struct{}),
		currentHeight: 0,
		updateTime:    0,
	}, nil
}

func (e *EthIndexer) IndexTasks(ctx context.Context) {
	logger.Sugar().Info("start to index task for ethereum")
	e.indexTransferToDB(ctx)

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

	// check wheather have task in kafka
	resp, err := synctaskNMCli.GetSyncTasks(ctx, &synctask.GetSyncTasksRequest{Conds: conds, Offset: 0, Limit: 0})
	if err != nil {
		logger.Sugar().Error(err)
	}
	for _, v := range resp.GetInfos() {
		err = e.addTask(v.Topic)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
	}

	conds.SyncState.Value = cttype.SyncState_Start.String()

	for {
		// TODO: limit sync topic number
		resp, err := synctaskNMCli.GetSyncTasks(ctx, &synctask.GetSyncTasksRequest{Conds: conds, Offset: 0, Limit: 0})
		if err != nil {
			logger.Sugar().Error(err)
		}
		for _, v := range resp.GetInfos() {
			err = e.addTask(v.Topic)
			if err != nil {
				logger.Sugar().Error(err)
				continue
			}
		}
		time.Sleep(CheckTopicInterval)
	}
}

func (e *EthIndexer) addTask(topic string) error {
	pLock.Lock()
	defer pLock.Unlock()
	if _, ok := e.taskMap[topic]; ok {
		return nil
	}
	e.taskMap[topic] = struct{}{}
	logger.Sugar().Infof("start to consumer task topic: %v", topic)
	consumer, err := ctkafka.NewCTConsumer(topic)
	if err != nil {
		logger.Sugar().Error(err)
		delete(e.taskMap, topic)
		return err
	}
	go func() {
		err = consumer.Consume(func(m *kafka.Message) {
			num, err := utils.Bytes2Uint64(m.Value)
			if err != nil {
				logger.Sugar().Error(err)
				return
			}

			e.taskChan <- num
			if num%uint64(LogIntervalHeight) == 0 {
				logger.Sugar().Infof("start sync %v height", num)
			}
		}, func(retryNum int) {
			if retryNum < MaxRetriesThenTriggerTask {
				return
			} else {
				_, err := synctaskNMCli.TriggerSyncTask(context.Background(), &synctask.TriggerSyncTaskRequest{Topic: topic})
				if err != nil {
					logger.Sugar().Error(err)
				}
			}

			if retryNum >= MaxRetriesThenStopConsume {
				logger.Sugar().Warnf("cannot consume topic %v,retry %v times", topic, MaxRetriesThenStopConsume)
				delete(e.taskMap, topic)
				consumer.Close()
			}
		})
		if err != nil {
			logger.Sugar().Error(err)
			delete(e.taskMap, topic)
			consumer.Close()
		}
	}()

	return err
}

func (e *EthIndexer) indexTransfer(ctx context.Context, handler func([]*TokenTransfer) error) {
	for i := 0; i < MaxDealBlockNum; i++ {
		wg.Add(1)
		go func() {
			for num := range e.taskChan {
				transfers, err := TransferLogs(ctx, int64(num), int64(num))
				if err != nil && containErr(err.Error()) {
					logger.Sugar().Errorf("will retry anlysis height %v for parsing transfer logs failed, %v", num, err)
					// TODO: this will be a bug,when all consumer wait for retry,deadlock occurs
					e.taskChan <- num
					continue
				}

				if err != nil {
					logger.Sugar().Errorf("parse transfer logs failed, %v", err)
				}

				err = handler(transfers)
				if err != nil {
					logger.Sugar().Errorf("handle transfers failed, %v", err)
				}
			}
			wg.Done()
		}()
	}
}

// TODO: task flow should use channel
func (e *EthIndexer) indexTransferToDB(ctx context.Context) {
	e.indexTransfer(ctx, func(transfers []*TokenTransfer) error {
		transferErr := e.transferToDB(ctx, transfers)
		if transferErr != nil {
			return transferErr
		}
		e.tokenInfoToDB(ctx, transfers)
		return nil
	})
}

func (e *EthIndexer) transferToDB(ctx context.Context, transfers []*TokenTransfer) error {
	tt := make([]*transferProto.TransferReq, len(transfers))
	for i := range transfers {
		chainType := transfers[i].ChainType
		tokenType := string(transfers[i].TokenType)
		tt[i] = &transferProto.TransferReq{
			ChainType:   &chainType,
			ChainID:     eth.CurrentChainID,
			Contract:    &transfers[i].Contract,
			TokenType:   &tokenType,
			TokenID:     &transfers[i].TokenID,
			From:        &transfers[i].From,
			To:          &transfers[i].To,
			Amount:      &transfers[i].Amount,
			BlockNumber: &transfers[i].BlockNumber,
			TxHash:      &transfers[i].TxHash,
			BlockHash:   &transfers[i].BlockHash,
			TxTime:      &transfers[i].TxTime,
		}
	}

	if len(tt) == 0 {
		return nil
	}

	for i := 0; i < Retries; i++ {
		_, err := transferNMCli.UpsertTransfers(ctx, &transferProto.UpsertTransfersRequest{Infos: tt})
		if err != nil && containErr(err.Error()) {
			logger.Sugar().Errorf("will retry for creating transfer record failed, %v", err)
			continue
		}

		if err != nil {
			return fmt.Errorf("create transfer record failed, %v", err)
		}
		break
	}
	return nil
}

func (e *EthIndexer) tokenInfoToDB(ctx context.Context, transfers []*TokenTransfer) {
	for _, transfer := range transfers {
		identifier := tokenIdentifier(transfer.ChainType, transfer.ChainID, transfer.Contract, transfer.TokenID)
		if err := ctredis.TryPubLock(identifier, redisExpireDefaultTime); err != nil {
			continue
		}

		remark := ""
		conds := &tokenProto.Conds{
			ChainType: &ctMessage.StringVal{
				Value: transfer.ChainType.String(),
				Op:    "eq",
			},
			ChainID: &ctMessage.StringVal{
				Value: transfer.ChainID,
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

		tokenURI, err := eth.TokenURI(ctx, transfer.TokenType, transfer.Contract, transfer.TokenID, transfer.BlockNumber)
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
					ChainType:   &transfer.ChainType,
					ChainID:     eth.CurrentChainID,
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

func (e *EthIndexer) contractToDB(ctx context.Context, transfer *TokenTransfer) error {
	identifier := contractIdentifier(transfer.ChainType, transfer.ChainID, transfer.Contract)
	if err := ctredis.TryPubLock(identifier, redisExpireDefaultTime); err != nil {
		return nil
	}

	conds := &contractProto.Conds{
		ChainType: &ctMessage.StringVal{
			Value: transfer.ChainType.String(),
			Op:    "eq",
		},
		ChainID: &ctMessage.StringVal{
			Value: transfer.ChainID,
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

	contractMeta, err := eth.GetERC721Metadata(ctx, transfer.Contract)
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
				ChainType: &transfer.ChainType,
				ChainID:   eth.CurrentChainID,
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
