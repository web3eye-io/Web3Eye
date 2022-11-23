package milvusdb

import (
	"context"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

const (
	milvusAddr = `localhost:19530`
)

var cli client.Client

func Init(ctx context.Context) error {
	cli, err := Client(ctx)
	if err != nil {
		return err
	}
	return initCollections(ctx, cli)
}

func Client(ctx context.Context) (client.Client, error) {
	if cli != nil {
		return cli, nil
	}
	c, err := client.NewGrpcClient(ctx, milvusAddr)
	if err != nil {
		return nil, err
	}
	cli = c
	return cli, nil
}
