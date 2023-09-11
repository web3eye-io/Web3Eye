package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/utils"
)

type Ss struct {
	B uint
	C string
}

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	s := &Ss{B: 0, C: ""}
	sBytes, err := json.Marshal(s)
	fmt.Println(string(sBytes), err)

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
