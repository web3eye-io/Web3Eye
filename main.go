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
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	db.Init()
	// id := "042c9730-3570-4336-9704-3f80eb9f3244"
	// chainType := basetype.ChainType_Ethereum.String()
	// chainID := "1"
	// txHash := "sssssss"
	// blockNumber := uint64(123)
	// txIndex := uint32(123)
	// logIndex := uint32(125)
	// logIndex1 := uint32(126)
	// recipient := "123"
	// targetItems := []*npool.OrderItem{
	// 	{
	// 		Contract:  "contract",
	// 		TokenType: "sajdiofjasiodfjois",
	// 		TokenID:   "sss",
	// 		Amount:    uint64(11),
	// 		Remark:    "ssss",
	// 	},
	// }
	// offerItems := []*npool.OrderItem{
	// 	{
	// 		Contract:  "contract",
	// 		TokenType: "sajdiofjasiodfjois",
	// 		TokenID:   "ssddfsds",
	// 		Amount:    uint64(11),
	// 		Remark:    "ssss",
	// 	},
	// }
	// remark := "123"

	order, _, err := order.Rows(context.Background(), nil, 0, 100)
	fmt.Println(utils.PrettyStruct(order), err)

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
