package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/chains"
	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/utils"
	orderNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/order"
	synctaskNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/synctask"
	ctMessage "github.com/web3eye-io/Web3Eye/proto/web3eye"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"
	orderhead "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
)

func (e *EthIndexer) PullParsePriceTasks(ctx context.Context, outBlockNum chan uint64) {
	logger.Sugar().Info("start to index parse price task for ethereum")
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
		logger.Sugar().Error(err)
		panic(err)
	}
	defer pulsarCli.Close()

	for {
		select {
		case <-time.NewTicker(CheckTopicInterval).C:
			resp, err := synctaskNMCli.GetSyncTasks(ctx, &synctask.GetSyncTasksRequest{Conds: conds, Offset: 0, Limit: maxTopicNum})
			if err != nil {
				logger.Sugar().Error(err)
			}

			for _, v := range resp.GetInfos() {
				e.indexTask(ctx, pulsarCli, v, outBlockNum)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (e *EthIndexer) indexPPTask(ctx context.Context, pulsarCli pulsar.Client, task *synctask.SyncTask, outBlockNum chan uint64) {
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

// index blocks [start,start+limit]
func (e *EthIndexer) IndexPrice(ctx context.Context, start uint64, limit uint64) ([]*chains.TokenTransfer, error) {
	cli, err := eth.Client(e.Endpoints)
	if err != nil {
		return nil, fmt.Errorf("cannot get eth client,err: %v", err)
	}

	transfers, err := cli.OrderFulfilledLogs(ctx, int64(start), int64(start+limit))
	if err != nil {
		e.checkErr(ctx, err)
		return nil, fmt.Errorf("failed to get transfer logs, err: %v, block: %v to %v", err, start, start+limit)
	}
	if len(transfers) == 0 {
		return nil, nil
	}

	for _, v := range transfers {
		v.ChainType = e.ChainType
		v.ChainID = e.ChainID
	}

	_, err = orderNMCli.CreateOrders(ctx, &orderhead.CreateOrdersRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed store transfers to db for block number: %v,err: %v, block: %v to %v", err, start, start+limit)
	}

	return nil, nil
}
