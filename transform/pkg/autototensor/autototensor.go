package autototensor

import (
	"context"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/ctpulsar"
	"github.com/web3eye-io/Web3Eye/common/oss"
	"github.com/web3eye-io/Web3Eye/config"
	gen_car "github.com/web3eye-io/Web3Eye/gen-car/pkg/client/v1"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/gencar/v1"
	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	"github.com/web3eye-io/Web3Eye/transform/pkg/filegetter"
	"github.com/web3eye-io/Web3Eye/transform/pkg/model"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
)

const (
	TaskPauseInterval = 3 * time.Second
)

func Run(ctx context.Context) {
	go func() {
		err := autoToTensor(ctx)
		if err != nil {
			logger.Sugar().Error(err)
			panic(err)
		}
	}()
	<-ctx.Done()
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
	consumer, err := pulsarCli.Subscribe(pulsar.ConsumerOptions{
		Topic:            config.GetConfig().Pulsar.TopicTransformImage,
		SubscriptionName: "TransformImageConsummer",
		Type:             pulsar.Shared,
		MessageChannel:   output,
	})
	if err != nil {
		return err
	}

	for msg := range output {
		_id := msg.Key()
		//nolint:gomnd
		idUint, err := strconv.ParseUint(_id, 10, 32)
		if err != nil {
			logger.Sugar().Errorf("failed to update token to nftmeta, err %v", err)
			return err
		}
		id := uint32(idUint)

		imgURL := string(msg.Message.Payload())
		errRecord := ""
		retry := false
		var vector []float32
		func() {
			filename := uuid.NewString()
			var filePath *string
			filePath, retry, err = filegetter.GetFileFromURL(ctx, imgURL, config.GetConfig().Transform.DataDir, filename)
			if err != nil {
				logger.Sugar().Errorf("failed to download file form url, err %v", err)
				errRecord = err.Error()
				return
			}
			err = updateForGenCar(ctx, id, *filePath)
			if err != nil {
				logger.Sugar().Errorf("failed to update for gen-car file form url, err %v", err)
				errRecord = err.Error()
			} else {
				defer os.Remove(*filePath)
			}

			vector, err = model.ToImageVector(ctx, *filePath)
			if err != nil {
				logger.Sugar().Errorf("failed to transform url to vector, err %v", err)
				errRecord = err.Error()
			}
		}()
		if retry {
			consumer.NackID(msg.ID())
			continue
		}

		_, err = token.UpdateImageVector(ctx, &tokenproto.UpdateImageVectorRequest{
			ID:     id,
			Vector: vector,
			Remark: errRecord,
		})

		if err != nil {
			logger.Sugar().Errorf("failed to update token to nftmeta, err %v", err)
			time.Sleep(TaskPauseInterval)
			continue
		}

		err = consumer.AckID(msg.ID())
		if err != nil {
			logger.Sugar().Errorf("failed to ask id to pulsar, err %v", err)
			time.Sleep(TaskPauseInterval)
			continue
		}
	}

	return nil
}

func updateForGenCar(ctx context.Context, id uint32, filePath string) error {
	s3key := path.Base(filePath)
	err := oss.UploadFile(ctx, filePath, config.GetConfig().Minio.TokenImageBucket, s3key)
	if err != nil {
		return err
	}
	_, err = gen_car.ReportFile(ctx, &v1.ReportFileRequest{
		ID:    id,
		S3Key: s3key,
	})
	return err
}
