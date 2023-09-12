package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/utils"
	blockNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/block"
	synctaskNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/synctask"
	ctMessage "github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	blockProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
)

const (
	CheckTopicInterval     = time.Second * 10
	FindContractCreator    = false
	redisExpireDefaultTime = time.Second * 10
	maxTopicNum            = 5
	maxParseGoroutineNum   = 5
	updateBlockNumInterval = time.Minute
)

type EthIndexer struct {
	Endpoints       []string
	BadEndpoints    map[string]error
	ChainType       basetype.ChainType
	ChainID         string
	CurrentBlockNum uint64
	onIndex         bool
	cancel          *context.CancelFunc
}

func NewIndexer(chainID string) *EthIndexer {
	return &EthIndexer{
		BadEndpoints:    make(map[string]error),
		ChainType:       basetype.ChainType_Ethereum,
		ChainID:         chainID,
		CurrentBlockNum: 0,
	}
}

func (e *EthIndexer) StartIndex(ctx context.Context) {
	logger.Sugar().Infof("start the indexer chainType: %v, chainID: %v", e.ChainType, e.ChainID)
	ctx, cancel := context.WithCancel(ctx)
	e.cancel = &cancel
	go e.GetCurrentBlockNum(ctx, updateBlockNumInterval)

	e.onIndex = true

	taskBlockNum, err := e.PullTaskTopics(ctx)
	if err != nil {
		logger.Sugar().Error(err)
		panic(err)
	}
	for i := 0; i < maxParseGoroutineNum; i++ {
		go e.IndexBlock(ctx, taskBlockNum)
	}
	time.Sleep(time.Minute * 3)
}

func (e *EthIndexer) UpdateEndpoints(endpoints []string) {
	e.Endpoints = endpoints
}

func (e *EthIndexer) IsOnIndex() bool {
	return e.onIndex
}

func (e *EthIndexer) StopIndex() {
	if e.cancel != nil {
		logger.Sugar().Infof("stop the indexer chainType: %v, chainID: %v", e.ChainType, e.ChainID)
		(*e.cancel)()
		e.cancel = nil
		e.BadEndpoints = nil
		e.Endpoints = nil
		e.onIndex = false
	}
}

func (e *EthIndexer) PullTaskTopics(ctx context.Context) (outBlockNum chan uint64, err error) {
	logger.Sugar().Info("start to index task for ethereum")
	conds := &synctask.Conds{
		ChainType: &ctMessage.StringVal{
			Value: e.ChainType.String(),
			Op:    "eq",
		},
		ChainID: &ctMessage.StringVal{
			Value: e.ChainID,
			Op:    "eq",
		},
		SyncState: &ctMessage.StringVal{
			Value: cttype.SyncState_Start.String(),
			Op:    "eq",
		},
	}

	pulsarCli, err := ctpulsar.Client()
	if err != nil {
		return nil, err
	}
	defer pulsarCli.Close()

	outBlockNum = make(chan uint64)
	go func() {
		defer close(outBlockNum)
		for {
			select {
			case <-time.NewTicker(CheckTopicInterval).C:
				resp, err := synctaskNMCli.GetSyncTasks(ctx, &synctask.GetSyncTasksRequest{Conds: conds, Offset: 0, Limit: maxTopicNum})
				if err != nil {
					logger.Sugar().Error(err)
				}

				for _, v := range resp.GetInfos() {
					e.indexTopicTasks(ctx, pulsarCli, v, outBlockNum)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return outBlockNum, nil
}

func (e *EthIndexer) indexTopicTasks(ctx context.Context, pulsarCli pulsar.Client, task *synctask.SyncTask, outBlockNum chan uint64) {
	output := make(chan pulsar.ConsumerMessage)
	consumer, err := pulsarCli.Subscribe(pulsar.ConsumerOptions{
		Topic:            task.Topic,
		SubscriptionName: "TaskConsummer",
		Type:             pulsar.Shared,
		MessageChannel:   output,
	})
	if err != nil {
		logger.Sugar().Errorf("consummer SyncTask failed ,err: %v", err)
		return
	}
	defer consumer.Close()

	retries := 0
	maxRetries := 3
	for {
		select {
		case msg := <-output:
			blockNum, err := utils.Bytes2Uint64(msg.Payload())
			if err != nil {
				logger.Sugar().Errorf("consummer SyncTask failed ,err: %v", err)
				continue
			}
			outBlockNum <- blockNum

			consumer.Ack(msg)
			retries = 0
		case <-time.NewTicker(CheckTopicInterval).C:
			if retries > maxRetries {
				return
			}
			resp, err := synctaskNMCli.TriggerSyncTask(ctx, &synctask.TriggerSyncTaskRequest{Topic: task.Topic, CurrentBlockNum: e.CurrentBlockNum})
			if err != nil {
				logger.Sugar().Errorf("triggerSyncTask failed ,err: %v", err)
				continue
			}
			if resp.Info.SyncState != cttype.SyncState_Start {
				return
			}
			retries++
		}
	}
}

func (e *EthIndexer) IndexBlock(ctx context.Context, taskBlockNum chan uint64) {
	for {
		select {
		case num := <-taskBlockNum:
			block, err := e.CheckBlock(ctx, num)
			if err != nil {
				logger.Sugar().Error(err)
				continue
			}

			if block.ParseState == basetype.BlockParseState_BlockTypeFinish {
				continue
			}

			err = func() error {
				blockLogs, err := e.IndexBlockLogs(ctx, num)
				if err != nil {
					return err
				}

				filteredT1, err := e.IndexTransfer(ctx, blockLogs.TransferLogs)
				if err != nil {
					return err
				}

				contractT1, err := e.IndexToken(ctx, filteredT1)
				if err != nil {
					return err
				}

				err = e.IndexContract(ctx, contractT1, FindContractCreator)
				if err != nil {
					return err
				}

				contractT2, err := e.IndexOrder(ctx, blockLogs.OrderLogs)
				if err != nil {
					return err
				}

				err = e.IndexContract(ctx, contractT2, FindContractCreator)
				if err != nil {
					return err
				}
				return nil
			}()

			if err != nil {
				logger.Sugar().Error(err)
			}

			remark := ""
			parseState := basetype.BlockParseState_BlockTypeFinish
			if err != nil {
				remark = err.Error()
				parseState = basetype.BlockParseState_BlockTypeFailed
			}

			blockNMCli.UpdateBlock(ctx, &blockProto.UpdateBlockRequest{
				Info: &blockProto.BlockReq{
					ID:         &block.ID,
					ParseState: &parseState,
					Remark:     &remark,
				},
			})
		case <-ctx.Done():
			return
		}
	}
}

func transferIdentifier(contract, tokenID, txHash, from string) string {
	return fmt.Sprintf("%v:%v:%v:%v", contract, tokenID, txHash, from)
}

func tokenIdentifier(chain basetype.ChainType, chainID, contract, tokenID string) string {
	return fmt.Sprintf("%v:%v:%v:%v", chain, chainID, contract, tokenID)
}

func contractIdentifier(chain basetype.ChainType, chainID, contract string) string {
	return fmt.Sprintf("%v:%v:%v", chain, chainID, contract)
}
