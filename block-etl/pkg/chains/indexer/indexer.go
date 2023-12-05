package indexer

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/utils"
	synctaskNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/synctask"
	ctMessage "github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
)

const (
	CheckTopicInterval     = time.Second * 10
	FindContractCreator    = false
	maxTopicNum            = 5
	maxParseGoroutineNum   = 5
	updateBlockNumInterval = time.Minute
)

type IIndexer interface {
	SyncCurrentBlockNum(ctx context.Context, updateBlockNumInterval time.Duration)
	IndexBlock(ctx context.Context, taskBlockNum chan uint64)
	GetCurrentBlockNum() uint64
	UpdateEndpoints(endpoints []string)
	OnNoAvalibleEndpoints(func())
}

type Indexer struct {
	ChainType basetype.ChainType
	ChainID   string
	onIndex   bool
	cancel    *context.CancelFunc
	IIndexer
}

func NewIndexer(chainID string, chainType basetype.ChainType, iIndexer IIndexer) *Indexer {
	return &Indexer{
		ChainType: chainType,
		ChainID:   chainID,
		onIndex:   false,
		IIndexer:  iIndexer,
	}
}

func (e *Indexer) StartIndex(ctx context.Context) {
	logger.Sugar().Infof("start the indexer chainType: %v, chainID: %v", e.ChainType, e.ChainID)
	ctx, cancel := context.WithCancel(ctx)
	e.cancel = &cancel
	go e.SyncCurrentBlockNum(ctx, updateBlockNumInterval)

	e.OnNoAvalibleEndpoints(func() {
		e.StopIndex()
	})
	e.onIndex = true

	taskBlockNum, err := e.pullTaskTopics(ctx)
	if err != nil {
		logger.Sugar().Error(err)
		panic(err)
	}
	for i := 0; i < maxParseGoroutineNum; i++ {
		go e.IndexBlock(ctx, taskBlockNum)
	}
}

func (e *Indexer) IsOnIndex() bool {
	return e.onIndex
}

func (e *Indexer) StopIndex() {
	if e.cancel != nil {
		logger.Sugar().Infof("stop the indexer chainType: %v, chainID: %v", e.ChainType, e.ChainID)
		(*e.cancel)()
		e.cancel = nil
		e.onIndex = false
	}
}

func (e *Indexer) pullTaskTopics(ctx context.Context) (outBlockNum chan uint64, err error) {
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
			Value: basetype.SyncState_Start.String(),
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

				for _, info := range resp.GetInfos() {
					e.indexTopicTasks(ctx, pulsarCli, info, outBlockNum)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return outBlockNum, nil
}

func (e *Indexer) indexTopicTasks(ctx context.Context, pulsarCli pulsar.Client, task *synctask.SyncTask, outBlockNum chan uint64) {
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

			err = consumer.Ack(msg)
			if err != nil {
				logger.Sugar().Errorf("ack SyncTask msg failed ,err: %v", err)
				continue
			}
			retries = 0
		case <-time.NewTicker(CheckTopicInterval).C:
			if retries > maxRetries {
				return
			}
			resp, err := synctaskNMCli.TriggerSyncTask(ctx, &synctask.TriggerSyncTaskRequest{Topic: task.Topic, CurrentBlockNum: e.GetCurrentBlockNum()})
			if err != nil {
				logger.Sugar().Errorf("triggerSyncTask failed ,err: %v", err)
				continue
			}
			if resp.Info.SyncState != basetype.SyncState_Start {
				return
			}
			retries++
		}
	}
}

func TransferIdentifier(contract, tokenID, txHash, from string) string {
	return fmt.Sprintf("%v:%v:%v:%v", contract, tokenID, txHash, from)
}

func TokenIdentifier(chain basetype.ChainType, chainID, contract, tokenID string) string {
	return fmt.Sprintf("%v:%v:%v:%v", chain, chainID, contract, tokenID)
}

func ContractIdentifier(chain basetype.ChainType, chainID, contract string) string {
	return fmt.Sprintf("%v:%v:%v", chain, chainID, contract)
}