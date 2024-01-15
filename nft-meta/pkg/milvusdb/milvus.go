package milvusdb

import (
	"context"
	"time"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/web3eye-io/Web3Eye/config"
)

const (
	connectTimeout = time.Second * 5
	contextCancel  = "context canceled"
)

var cli client.Client

func Init(ctx context.Context) error {
	cli, err := Client(ctx)
	if err != nil {
		return err
	}
	return initCollections(ctx, cli)
}

func Client(ctx context.Context) (c client.Client, err error) {
	if cli != nil {
		return cli, nil
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, connectTimeout)
	defer cancel()
	cli, err = client.NewGrpcClient(timeoutCtx, config.GetConfig().Milvus.Address)

	if err != nil {
		return nil, err
	}

	return cli, nil
}
