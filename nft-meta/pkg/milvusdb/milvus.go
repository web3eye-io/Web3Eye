package milvusdb

import (
	"context"
	"time"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

const (
	milvusAddr     = `localhost:19530`
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
	go func() {
		defer cancel()
		cli, err = client.NewGrpcClient(ctx, milvusAddr)
	}()

	<-timeoutCtx.Done()
	ctxErr := timeoutCtx.Err()

	if ctxErr.Error() != contextCancel {
		return nil, ctxErr
	}

	if err != nil {
		return nil, err
	}

	return cli, nil
}
