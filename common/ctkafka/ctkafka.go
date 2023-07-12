package ctkafka

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/config"
)

const (
	FlushTimeMS         = 15000
	ReadMSGTimeout      = time.Second * 3
	DefaultPartitionNum = 6
)

var (
	topicMap = make(map[string]*CTProducer)
)

func KafkaConfig() kafka.ConfigMap {
	conf := make(map[string]kafka.ConfigValue)
	conf["bootstrap.servers"] = strings.TrimSpace(config.GetConfig().Kafka.BootstrapServers)
	return conf
}

type CTProducer struct {
	Topic    string
	Producer *kafka.Producer
}

// TODO: eventHandle setting may be reasonable
func NewCTProducer(topic string) (*CTProducer, error) {
	if ctP := topicMap[topic]; ctP != nil {
		return ctP, nil
	}

	conf := KafkaConfig()
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		return nil, err
	}

	ctP := &CTProducer{Topic: topic, Producer: p}
	topicMap[topic] = ctP

	return ctP, nil
}

func (ctP *CTProducer) Produce(key, data []byte) error {
	return ctP.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &ctP.Topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          data,
	}, nil)
}

func (ctP *CTProducer) Close() {
	ctP.Producer.Flush(FlushTimeMS)
	ctP.Producer.Close()
}

func CreateTopic(ctx context.Context, topic string) error {
	conf := KafkaConfig()
	admin, err := kafka.NewAdminClient(&conf)
	if err != nil {
		return err
	}
	_, err = admin.CreateTopics(ctx, []kafka.TopicSpecification{
		{
			Topic:         topic,
			NumPartitions: DefaultPartitionNum,
		},
	})
	return err
}

type CTConsumer struct {
	topic    string
	consumer *kafka.Consumer
	start    bool
}

func NewCTConsumerT(topic string, clientID string) (*CTConsumer, error) {
	conf := KafkaConfig()
	conf["group.id"] = "default"
	conf["enable.auto.commit"] = "true"
	conf["auto.offset.reset"] = "earliest"
	conf["client.id"] = clientID

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		return nil, err
	}
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}
	return &CTConsumer{topic: topic, consumer: c, start: false}, nil
}

func NewCTConsumer(topic string) (*CTConsumer, error) {
	conf := KafkaConfig()
	conf["group.id"] = "default"
	conf["enable.auto.commit"] = "true"
	conf["auto.offset.reset"] = "earliest"
	conf["client.id"] = uuid.NewString()

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		return nil, err
	}
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}
	return &CTConsumer{topic: topic, consumer: c, start: false}, nil
}

func (ctC *CTConsumer) Consume(msgHandle func(*kafka.Message), retryHandle func(retryNum int) (exit bool)) error {
	if ctC.start {
		return fmt.Errorf("cannot start to consume,please stop current consumer")
	}
	ctC.start = true
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	retryNum := 0
	// Process messages
	for ctC.start {
		select {
		case sig := <-sigchan:
			logger.Sugar().Errorf("Caught signal %v: terminating\n", sig)
			ctC.start = false
		default:
			ev, err := ctC.consumer.ReadMessage(ReadMSGTimeout)
			if err != nil {
				retryNum++
				exit := retryHandle(retryNum)
				if exit {
					return filterErr(err)
				}
				// Errors are informational and automatically handled by the consumer
				continue
			}
			retryNum = 0
			msgHandle(ev)
		}
	}

	return ctC.consumer.Close()
}

func filterErr(err error) error {
	if strings.Contains(err.Error(), "Timed out") {
		return nil
	}
	return err
}

func (ctC *CTConsumer) Close() {
	ctC.start = false
}
