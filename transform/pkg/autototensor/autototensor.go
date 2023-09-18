package autototensor

import (
	"context"
	"os"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/config"
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	"github.com/web3eye-io/Web3Eye/transform/pkg/filegetter"
	"github.com/web3eye-io/Web3Eye/transform/pkg/model"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
)

const (
	TaskCheckInterval = time.Second
	// this mean MaxTaskCheckInterval=TaskCheckInterval*2^MaxTCIIndex
	// If TaskCheckInterval start as 1 second and MaxTCIIndex is 10,the max TaskCheckInterval is 1024=1*2^10
	MaxTCIIndex = 10
)

func Run() {
	logger.Sugar().Info("start to check and deal the vector_state")
	taskCheckInterval := TaskCheckInterval
	index := 0
	for {
		<-time.NewTicker(taskCheckInterval).C
		if err := autoToTensor(context.Background()); err != nil && index < MaxTCIIndex {
			index++
		} else if err == nil && index > 0 {
			index = 0
		}
		// TODO:check value of taskCheckInterval, maybe it will overflow
		taskCheckInterval = TaskCheckInterval << index
	}
}

func autoToTensor(ctx context.Context) error {
	// check the token service can be reached,then continue
	if _, err := token.LookupHost(); err != nil {
		return err
	}

	pulsarCli, err := ctpulsar.Client()
	if err != nil {
		return err
	}
	defer pulsarCli.Close()

	output := make(chan pulsar.ConsumerMessage)
	_, err = pulsarCli.Subscribe(pulsar.ConsumerOptions{
		Topic:            config.GetConfig().Pulsar.TopicTransformImage,
		SubscriptionName: "TransformImageConsummer",
		Type:             pulsar.Shared,
		MessageChannel:   output,
	})
	if err != nil {
		return err
	}

	for msg := range output {
		id := msg.Key()
		imgURL := string(msg.Message.Payload())
		errRecord := ""
		var vector []float32

		func() {
			filename := uuid.NewString()
			path, err := filegetter.GetFileFromURL(imgURL, config.GetConfig().Transform.DataDir, filename)
			if err != nil {
				logger.Sugar().Errorf("failed to download file form url, err %v", err)
				errRecord = err.Error()
				return
			}
			defer os.Remove(*path)

			vector, err = model.ToImageVector(ctx, *path)
			if err != nil {
				logger.Sugar().Errorf("failed to transform url to vector, err %v", err)
				errRecord = err.Error()
			}
		}()

		_, err := token.UpdateImageVector(ctx, &tokenproto.UpdateImageVectorRequest{
			ID:     id,
			Vector: vector,
			Remark: errRecord,
		})
		if err != nil {
			logger.Sugar().Errorf("failed to update token to nftmeta, err %v", err)
			return err
		}
	}

	return nil
}
