package milvusdb

import (
	"context"
	"fmt"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
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
			err = c.CreateCollection(ctx, collection, shardsNum)
			if err != nil {
				return err
			}

			err = c.CreatePartition(ctx, collection.CollectionName, DefaultPartition)
			if err != nil {
				return err
			}

			idx, err := entity.NewIndexDISKANN(entity.L2)
			if err != nil {
				return err
			}
			err = c.CreateIndex(ctx, collection.CollectionName, FieldsVector, idx, false)
			if err != nil {
				return err
			}
		}

		err = c.LoadCollection(ctx, collection.CollectionName, false)
		if err != nil {
			return err
		}
	}
	return nil
}
