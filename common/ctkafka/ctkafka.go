package ctkafka

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"github.com/web3eye-io/cyber-tracer/config"
)

const (
	FlushTimeMS         = 15000
	ReadMSGTimeout      = time.Second
	DefaultPartitionNum = 6
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

// TODO: auto create topic
func NewCTProducer(topic string, eventHandle func(kafka.Event)) (*CTProducer, error) {
	conf := KafkaConfig()
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		return nil, err
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	// TODO: it should be close by some way
	go func() {
		for e := range p.Events() {
			eventHandle(e)
		}
	}()

	return &CTProducer{Topic: topic, Producer: p}, nil
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
	topic     string
	consumer  *kafka.Consumer
	msgHandle func(*kafka.Message)
	start     bool
}

func NewCTConsumer(topic string, msgHandle func(*kafka.Message)) (*CTConsumer, error) {
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
	return &CTConsumer{topic: topic, consumer: c, start: false, msgHandle: msgHandle}, nil
}

func (ctC *CTConsumer) Consume() error {
	if ctC.start {
		return fmt.Errorf("cannot start to consume,please stop current consumer")
	}
	ctC.start = true
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		// Process messages

		for ctC.start {
			select {
			case sig := <-sigchan:
				fmt.Printf("Caught signal %v: terminating\n", sig)
				ctC.start = false
			default:
				ev, err := ctC.consumer.ReadMessage(ReadMSGTimeout)
				if err != nil {
					// Errors are informational and automatically handled by the consumer
					continue
				}
				ctC.msgHandle(ev)
			}
		}
		ctC.consumer.Close()
	}()
	return nil
}

func (ctC *CTConsumer) Close() {
	ctC.start = false
}
