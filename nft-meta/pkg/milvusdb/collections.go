package milvusdb

import (
	"context"
	"fmt"
	"time"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	"github.com/web3eye-io/Web3Eye/config"
)

const (
	FieldsID         = "ID"
	FieldsVector     = "Vector"
	DefaultPartition = "default"
	VectorDim        = 2048
)

var (
	NFTSchema = &entity.Schema{
		CollectionName: config.GetConfig().NFTMeta.CollectionName,
		Description:    config.GetConfig().NFTMeta.Description,
		AutoID:         true,
		Fields: []*entity.Field{
			{
				Name:       FieldsID,
				DataType:   entity.FieldTypeInt64,
				PrimaryKey: true,
				AutoID:     true,
			},
			{
				Name:     FieldsVector,
				DataType: entity.FieldTypeFloatVector,
				TypeParams: map[string]string{
					entity.TypeParamDim: fmt.Sprintf("%d", VectorDim),
				},
			},
		},
	}
	allSchema = []*entity.Schema{NFTSchema}
)

func initCollections(ctx context.Context, c client.Client) error {
	for _, collection := range allSchema {
		has, err := c.HasCollection(ctx, collection.CollectionName)
		if err != nil {
			return err
		}
		if !has {
			var shardsNum int32 = 8
			err := c.CreateCollection(ctx, collection, shardsNum)
			if err != nil {
				return err
			}

			err = c.CreatePartition(ctx, collection.CollectionName, DefaultPartition)
			if err != nil {
				return err
			}
		}

		indexes, _ := c.DescribeIndex(ctx, collection.CollectionName, FieldsVector)
		haveIndex := false
		for _, index := range indexes {
			_ = c.ReleaseCollection(ctx, collection.CollectionName)
			if index.IndexType() == entity.DISKANN {
				haveIndex = true
				continue
			}
			err := c.DropIndex(ctx, collection.CollectionName, FieldsVector)
			if err != nil {
				return err
			}
		}

		if !haveIndex {
			idx, err := entity.NewIndexDISKANN(entity.L2)
			if err != nil {
				return err
			}
			err = c.CreateIndex(ctx, collection.CollectionName, FieldsVector, idx, false)
			if err != nil {
				return err
			}
		}
		go autoFlush()
		err = c.LoadCollection(ctx, collection.CollectionName, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func autoFlush() {
	lockKey := "milvus_flush_lock"
	refreshTimeout := time.Hour
	for {
		<-time.NewTicker(refreshTimeout).C
		locked, err := ctredis.TryPubLock(lockKey, refreshTimeout)
		if locked && err != nil {
			_ = cli.Flush(context.Background(), c.CollectionName, false)
		}
	}
}
