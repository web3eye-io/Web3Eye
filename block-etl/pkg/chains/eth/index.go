package eth

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/web3eye-io/cyber-tracer/block-etl/pkg/token"
	"github.com/web3eye-io/cyber-tracer/common/ctkafka"
	"github.com/web3eye-io/cyber-tracer/common/ctredis"
	"github.com/web3eye-io/cyber-tracer/common/utils"
	contractNMCli "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/client/v1/contract"
	synctaskNMCli "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/client/v1/synctask"
	tokenNMCli "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/client/v1/token"
	transferNMCli "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/client/v1/transfer"
	ctMessage "github.com/web3eye-io/cyber-tracer/proto/cybertracer"
	contractProto "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/contract"
	"github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/cttype"
	"github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/synctask"
	tokenProto "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/token"
	transferProto "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/transfer"
)

const (
	Retries                   = 3
	MaxDealBlockNum           = 5
	LogIntervalHeight         = 10
	MaxRetriesThenTriggerTask = 5
	MaxRetriesThenStopConsume = 5
	// confirmedBlockNum      = 5
	redisExpireDefaultTime = time.Minute * 5
	maxTokenURILen         = 256
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
		taskChan:      make(chan uint64, MaxDealBlockNum),
		taskMap:       make(map[string]struct{}),
		currentHeight: 0,
		updateTime:    0,
	}, nil
}

func (e *EthIndexer) indexTransfer(ctx context.Context, handler func([]*TokenTransfer) error) {
	for i := 0; i < MaxDealBlockNum; i++ {
		wg.Add(1)
		go func() {
			for num := range e.taskChan {
				transfers, err := TransferLogs(ctx, int64(num), int64(num))
				if err != nil && containErr(err.Error()) {
					logger.Sugar().Errorf("will retry anlysis height %v for parsing transfer logs failed, %v", num, err)
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
		tokenErr := e.tokenInfoToDB(ctx, transfers)
		if tokenErr != nil {
			return tokenErr
		}
		return nil
	})
}

func (e *EthIndexer) transferToDB(ctx context.Context, transfers []*TokenTransfer) error {
	tt := make([]*transferProto.TransferReq, len(transfers))
	for i := range transfers {
		chainType := string(transfers[i].ChainType)
		tokenType := string(transfers[i].TokenType)
		tt[i] = &transferProto.TransferReq{
			ChainType:   &chainType,
			ChainID:     &transfers[i].ChainID,
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

	for i := 0; i < Retries; i++ {
		_, err := transferNMCli.CreateTransfers(ctx, tt)
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

//nolint:gocyclo
func (e *EthIndexer) tokenInfoToDB(ctx context.Context, transfers []*TokenTransfer) error {
	for _, transfer := range transfers {
		identifier := tokenIdentifier(transfer.ChainType, transfer.ChainID, transfer.Contract, transfer.TokenID)
		if _, err := ctredis.Get(identifier); err == nil {
			return nil
		}

		tokenType := string(transfer.TokenType)
		remark := ""
		conds := &tokenProto.Conds{
			ChainType: &ctMessage.StringVal{
				Value: string(transfer.ChainType),
				Op:    "eq",
			},
			ChainID: &ctMessage.Int32Val{
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

		if exist, err := tokenNMCli.ExistTokenConds(ctx, conds); exist && err == nil {
			err = ctredis.Set(identifier, true, redisExpireDefaultTime)
			if err != nil {
				logger.Sugar().Error(err)
			}
			return nil
		}

		err := ctredis.Set(identifier, false, redisExpireDefaultTime)
		if err != nil {
			logger.Sugar().Error(err)
			err = ctredis.Del(identifier)
			if err != nil {
				logger.Sugar().Error(err)
			}
			return nil
		}

		err = e.contractToDB(ctx, transfer)
		if err != nil {
			logger.Sugar().Error(err)
		}

		tokenURI, err := TokenURI(ctx, transfer.TokenType, transfer.Contract, transfer.TokenID, transfer.BlockNumber)
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
			_, err = tokenNMCli.CreateToken(ctx, &tokenProto.TokenReq{
				ChainType:   (*string)(&transfer.ChainType),
				ChainID:     &transfer.ChainID,
				Contract:    &transfer.Contract,
				TokenType:   &tokenType,
				TokenID:     &transfer.TokenID,
				URI:         &tokenURI,
				URIType:     (*string)(&tokenURIInfo.URIType),
				ImageURL:    &tokenURIInfo.ImageURL,
				VideoURL:    &tokenURIInfo.VideoURL,
				Name:        &tokenURIInfo.Name,
				Description: &tokenURIInfo.Description,
				VectorState: tokenProto.ConvertState_Waiting.Enum(),
				Remark:      &remark,
			})
			if err != nil && containErr(err.Error()) {
				logger.Sugar().Errorf("will retry for creating token record failed, %v", err)
				continue
			}

			if err != nil {
				return fmt.Errorf("create token record failed, %v", err)
			}
			break
		}
		return nil
	}
	return nil
}

func (e *EthIndexer) contractToDB(ctx context.Context, transfer *TokenTransfer) error {
	identifier := contractIdentifier(transfer.ChainType, transfer.ChainID, transfer.Contract)
	if _, err := ctredis.Get(identifier); err == nil {
		return nil
	}

	conds := &contractProto.Conds{
		ChainType: &ctMessage.StringVal{
			Value: string(transfer.ChainType),
			Op:    "eq",
		},
		ChainID: &ctMessage.Int32Val{
			Value: transfer.ChainID,
			Op:    "eq",
		},
		Address: &ctMessage.StringVal{
			Value: transfer.Contract,
			Op:    "eq",
		},
	}

	if exist, err := contractNMCli.ExistContractConds(ctx, conds); exist && err == nil {
		err = ctredis.Set(identifier, true, redisExpireDefaultTime)
		if err != nil {
			logger.Sugar().Error(err)
		}
		return nil
	}
	err := ctredis.Set(identifier, false, redisExpireDefaultTime)
	if err != nil {
		logger.Sugar().Error(err)
		err = ctredis.Del(identifier)
		if err != nil {
			logger.Sugar().Error(err)
		}
		return nil
	}

	remark := ""
	creator, err := GetContractCreator(ctx, transfer.Contract)
	if err != nil {
		creator = &ContractCreator{}
		remark = err.Error()
	}

	contractMeta, err := GetERC721Metadata(ctx, transfer.Contract)
	if err != nil {
		contractMeta = &ERC721Metadata{}
		remark = fmt.Sprintf("%v,%v", remark, err)
	}

	from := creator.From.String()
	txHash := creator.TxHash.Hex()
	blockNum := creator.BlockNumber
	txTime := uint32(creator.TxTime)
	for i := 0; i < Retries; i++ {
		_, err = contractNMCli.CreateContract(ctx, &contractProto.ContractReq{
			ChainType: (*string)(&transfer.ChainType),
			ChainID:   &transfer.ChainID,
			Address:   &transfer.Contract,
			Name:      &contractMeta.Name,
			Symbol:    &contractMeta.Symbol,
			Creator:   &from,
			BlockNum:  &blockNum,
			TxHash:    &txHash,
			TxTime:    &txTime,
			Remark:    &remark,
		})
		if err != nil && containErr(err.Error()) {
			logger.Sugar().Errorf("will retry for creating token record failed, %v", err)
			continue
		}

		if err != nil {
			return fmt.Errorf("create token record failed, %v", err)
		}
		break
	}
	return nil
}

func (e *EthIndexer) IndexTasks(ctx context.Context) {
	logger.Sugar().Info("start to index task for ethereum")
	e.indexTransferToDB(ctx)

	conds := &synctask.Conds{
		ChainType: &ctMessage.StringVal{
			Value: cttype.ChainType_Ethereum.String(),
			Op:    "eq",
		},
		ChainID: &ctMessage.Int32Val{
			Value: 1,
			Op:    "eq",
		},
		SyncState: &ctMessage.StringVal{
			Value: cttype.SyncState_Start.String(),
			Op:    "eq",
		},
	}

	for {
		// TODO: limit sync topic number
		tasks, _, err := synctaskNMCli.GetSyncTasks(ctx, conds, 0, 0)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		for _, v := range tasks {
			err = e.addTask(v.Topic)
			if err != nil {
				logger.Sugar().Error(err)
				continue
			}
		}
		time.Sleep(time.Second * 3)
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
			if retryNum >= MaxRetriesThenTriggerTask {
				_task, err := synctaskNMCli.TriggerSyncTask(context.Background(), topic)
				if err != nil {
					logger.Sugar().Error(err)
					return
				}

				if _task.SyncState != cttype.SyncState_Start {
					delete(e.taskMap, topic)
					consumer.Close()
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

// func (e *EthIndexer) getCurrentConfirmedHeight(ctx context.Context) uint64 {
// 	gLock.Lock()
// 	defer gLock.Unlock()

// 	if e.updateTime > time.Now().Unix() {
// 		return e.currentHeight
// 	}

// 	num, err := CurrentBlockHeight(ctx)
// 	if err != nil {
// 		logger.Sugar().Errorf("get block height failed, %v", err)
// 		return e.currentHeight
// 	}

// 	if num > confirmedBlockNum {
// 		e.currentHeight = num - confirmedBlockNum
// 	}

// 	e.updateTime = time.Now().Add(time.Minute).Unix()

// 	return num
// }

func containErr(errStr string) bool {
	for _, v := range retryErrs {
		if strings.Contains(errStr, v) {
			return true
		}
	}
	return false
}

func tokenIdentifier(chain ChainType, chainID int32, contract, tokenID string) string {
	return fmt.Sprintf("%v:%v:%v:%v", chain, chainID, contract, tokenID)
}

func contractIdentifier(chain ChainType, chainID int32, contract string) string {
	return fmt.Sprintf("%v+%v+%v", chain, chainID, contract)
}
