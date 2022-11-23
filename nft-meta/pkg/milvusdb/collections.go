package milvusdb

// TODO: should use milvus-sdk-go to generate structure
import (
	"context"
	"fmt"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

const (
	FieldsID         = "ID"
	FieldsVector     = "Vector"
	DefaultPartition = "default"
	// TODO: verify rationality
	VectorDim = 2048
	ShardsNum = 3  // for collection:https://milvus.io/docs/v2.1.x/data_processing.md#Data-insertion
	ProbeNum  = 10 // for search
)

var (
	NFTSchema = &entity.Schema{
		CollectionName: "nft_info_images",
		Description:    "this collection for nft-meta",
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
			err = c.CreateCollection(ctx, collection, 2)
			if err != nil {
				return err
			}

			err = c.CreatePartition(ctx, collection.CollectionName, DefaultPartition)
			if err != nil {
				return err
			}

			idx, err := entity.NewIndexFlat(entity.L2, ShardsNum)
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
