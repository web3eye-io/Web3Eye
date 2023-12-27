package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
// 	"github.com/apache/pulsar-client-go/pulsar"
// 	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
// )

// func main() {
// 	logger.Init(logger.DebugLevel, "./a.log")

// 	go produce()
// 	consumer()
// }
// func consumer() {
// 	cli, err := ctpulsar.Client()
// 	if err != nil {
// 		logger.Sugar().Error(err)
// 		os.Exit(0)
// 	}
// 	defer cli.Close()

// 	consumChan := make(chan pulsar.ConsumerMessage)

// 	_, err = cli.Subscribe(pulsar.ConsumerOptions{
// 		Topic:            "token-image-bucket",
// 		SubscriptionName: "ssss",
// 		MessageChannel:   consumChan,
// 	})
// 	if err != nil {
// 		logger.Sugar().Error(err)
// 		os.Exit(0)
// 	}

// 	for {
// 		num := <-consumChan
// 		fmt.Println(num.Key())
// 		// if num.Key() == "5" {
// 		// 	continue
// 		// }
// 		// consum.AckID(num.ID())
// 	}
// }

// func produce() {
// 	cli, err := ctpulsar.Client()
// 	if err != nil {
// 		logger.Sugar().Error(err)
// 		os.Exit(0)
// 	}
// 	defer cli.Close()
// 	producer, err := cli.CreateProducer(pulsar.ProducerOptions{
// 		Topic: "token-image-bucket",
// 	})
// 	if err != nil {
// 		logger.Sugar().Error(err)
// 		os.Exit(0)
// 	}
// 	for i := 0; i < 10; i++ {
// 		_, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
// 			Key: fmt.Sprintf("%v", i),
// 		})
// 		if err != nil {
// 			logger.Sugar().Error(err)
// 			os.Exit(0)
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	fmt.Println("I finish produce")
// }

// func tableView() {
// 	cli, err := ctpulsar.Client()
// 	if err != nil {
// 		logger.Sugar().Error(err)
// 		os.Exit(0)
// 	}
// 	defer cli.Close()
// 	tableView, err := cli.CreateTableView(pulsar.TableViewOptions{
// 		Topic: "token-image-bucket",
// 	})
// 	if err != nil {
// 		logger.Sugar().Error(err)
// 		os.Exit(0)
// 	}
// 	fmt.Println(tableView.Keys())
// 	fmt.Println(tableView.Size())
// }
