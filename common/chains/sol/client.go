package sol

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/web3eye-io/Web3Eye/common/utils"
)

const (
	MinNodeNum       = 1
	MaxRetries       = 3
	retriesSleepTime = 200 * time.Millisecond
)

type solClients struct {
	endpoints []string
}

func (solCli solClients) GetNode(ctx context.Context) (*rpc.Client, error) {
	endpoint := solCli.endpoints[rand.Intn(len(solCli.endpoints))]

	cli := rpc.New(endpoint)
	return cli, nil
}

func (solCli *solClients) WithClient(ctx context.Context, fn func(ctx context.Context, c *rpc.Client) (bool, error)) error {
	var (
		apiErr, err error
		retry       bool
		client      *rpc.Client
	)

	if err != nil {
		return err
	}

	for i := 0; i < utils.MinInt(MaxRetries, len(solCli.endpoints)); i++ {
		if i > 0 {
			time.Sleep(retriesSleepTime)
		}

		client, err = solCli.GetNode(ctx)

		if err != nil {
			continue
		}

		retry, apiErr = fn(ctx, client)
		client.Close()

		if !retry {
			return apiErr
		}
	}

	if apiErr != nil {
		return apiErr
	}
	return err
}

func Client(endpoints []string) (*solClients, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("endpoints has no item")
	}
	return &solClients{endpoints: endpoints}, nil
}
