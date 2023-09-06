package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/block"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// go produceNonce("test9")
	// time.Sleep(time.Minute)
	// fmt.Println("sss")
	// go consumeNonce("test9", "s1")
	// go consumeNonce("test9", "s2")
	db.Init()
	chainID := "ssss"
	blockHash := "ssss"
	blockTime := time.Now().Unix()
	blockNum := uint64(111)
	parseState := v1.BlockParseState_BlockTypeFinish
	remark := "ss11"
	fmt.Println(block.Upsert(context.Background(), &npool.BlockReq{
		ChainType:   v1.ChainType_Solana.Enum(),
		ChainID:     &chainID,
		BlockNumber: &blockNum,
		BlockHash:   &blockHash,
		BlockTime:   &blockTime,
		ParseState:  &parseState,
		Remark:      &remark,
	}))
	<-sigchan
	os.Exit(1)
}

func produceNonce(topic string) {
	client, err := ctpulsar.Client()
	if err != nil {
		panic(err)
	}

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:  topic,
		Schema: pulsar.NewInt64Schema(map[string]string{}),
	})
	if err != nil {
		panic(err)
	}
	for i := int64(0); i < 10; i++ {
		payload, err := utils.Uint642Bytes(uint64(i))
		if err != nil {
			panic(err)
		}
		producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: payload,
		})
	}

}

func consumeNonce(topic string, name string) {
	client, err := ctpulsar.Client()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	output := make(chan pulsar.ConsumerMessage)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		EnableDefaultNackBackoffPolicy: true,
		NackRedeliveryDelay:            time.Second,
		Topic:                          topic,
		SubscriptionName:               "mmm1",
		Type:                           pulsar.Shared,
		MessageChannel:                 output,
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Unsubscribe()
	for {
		msg := <-output
		if err != nil {
			panic(err)
		}
		payload, err := utils.Bytes2Uint64(msg.Payload())
		fmt.Println(name, payload, err)
		// if name == "s1" {
		err = consumer.Ack(msg)
		if err != nil {
			fmt.Println(err)
		}
		// }

	}
}

type Ww struct {
	Words string
}
