package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/utils"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go produceNonce("test9")
	// time.Sleep(time.Minute)
	go consumeNonce("test9")

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

func consumeNonce(topic string) {
	client, err := ctpulsar.Client()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	output := make(chan pulsar.ConsumerMessage)
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "mmm1",
		Type:             pulsar.Shared,
		MessageChannel:   output,
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Unsubscribe()
	for i := 0; i < 12; i++ {
		msg := <-output
		if err != nil {
			panic(err)
		}
		payload := msg.Payload()
		vv, err := utils.Bytes2Uint64(payload)
		if err != nil {
			panic(err)
		}
		fmt.Println(vv, err)
		err = consumer.Ack(msg)
		fmt.Println(err)
	}
	fmt.Println("finish")
}

type Ww struct {
	Words string
}
