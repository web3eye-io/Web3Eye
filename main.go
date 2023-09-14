package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
	"github.com/web3eye-io/Web3Eye/ranker/pkg/crud/v1/transfer"
)

type Ss struct {
	B uint
	C string
}

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	db.Init()

	ret, n, err := transfer.Rows(context.Background(), &rankernpool.GetTransfersRequest{
		ChainType: basetype.ChainType_Ethereum,
		ChainID:   "1",
		Contract:  "0x880644ddF208E471C6f2230d31f9027578FA6FcC",
		TokenID:   "4205",
	})

	fmt.Println(utils.PrettyStruct(ret))
	fmt.Println(n)
	fmt.Println(err)
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
