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
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// go produceNonce("test9")
	// time.Sleep(time.Minute)
	// fmt.Println("sss")
	// go consumeNonce("test9", "s1")
	// go consumeNonce("test9", "s2")
	now := time.Now()
	err := milvusdb.Init(context.Background())
	if err != nil {
		panic(fmt.Errorf("milvus init err: %v", err))
	}
	milvusmgr := milvusdb.NewNFTConllectionMGR()

	_vector := imageconvert.ToArrayVector([]float32{})
	vectors := [][milvusdb.VectorDim]float32{}
	fmt.Println(time.Now().Sub(now))
	for i := 0; i < 100; i++ {
		vectors = append(vectors, _vector)
	}
	ret, err := milvusmgr.Create(context.Background(), vectors)
	fmt.Println(ret, err)

	fmt.Println(time.Now().Sub(now))
	_scores, err := milvusmgr.Search(context.Background(), [][milvusdb.VectorDim]float32{_vector}, 5)
	fmt.Println(_scores, err)
	fmt.Println(time.Now().Sub(now))

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
