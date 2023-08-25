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

	// go produceNonce()
	// time.Sleep(time.Minute)
	// go consumeNonce()

	<-sigchan
	os.Exit(1)
}

func produceNonce() {
	client, err := ctpulsar.Client()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "test01",
	})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {

		msgID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte("ssss"),
			Key:     fmt.Sprintf("ss%v", i),
		})
		fmt.Println(msgID, err)
	}
}

func consumeNonce() {
	client, err := ctpulsar.Client()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "test01",
		SubscriptionName: "mmm1",
		Type:             pulsar.Shared,
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Unsubscribe()
	for i := 0; i < 100; i++ {
		msg, err := consumer.Receive(context.Background())
		fmt.Println(utils.PrettyStruct(msg))
		fmt.Println(msg, err)
		err = consumer.Ack(msg)
		fmt.Println(err)
	}
}

type Ww struct {
	Words string
}
