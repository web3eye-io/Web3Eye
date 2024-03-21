package sol

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/web3eye-io/Web3Eye/common/chains"
)

const (
	MinNodeNum       = 1
	MaxRetries       = 3
	retriesSleepTime = 500 * time.Millisecond
)

type solClients struct {
	endpoints []string
}

func (solCli solClients) GetNode(ctx context.Context, useTimes uint16) (*rpc.Client, string, error) {
	if len(solCli.endpoints) == 0 {
		return nil, "", fmt.Errorf("have no available endpoints")
	}

	endpoint, err := chains.LockEndpoint(ctx, solCli.endpoints, useTimes)
	if err != nil {
		return nil, "", err
	}

	cli := rpc.New(endpoint)
	return cli, endpoint, nil
}

func (solCli *solClients) WithClient(ctx context.Context, useTimes uint16, fn func(ctx context.Context, c *rpc.Client) (bool, error)) error {
	var (
		apiErr, nodeErr error
		retry           bool
	)

	for i := 0; i < MaxRetries; i++ {
		if i > 0 {
			time.Sleep(retriesSleepTime)
		}

		client, endpoint, err := solCli.GetNode(ctx, useTimes)

		if err != nil {
			nodeErr = err
			continue
		}

		retry, apiErr = fn(ctx, client)
		client.Close()

		if !retry {
			return apiErr
		}
		if apiErr != nil {
			go checkEndpoint(context.Background(), endpoint, apiErr)
		}
	}

	if apiErr != nil {
		return apiErr
	}
	return nodeErr
}

func checkEndpoint(ctx context.Context, endpoint string, err error) {
	if err == nil {
		return
	}
	cli, err := Client([]string{endpoint})
	if err != nil {
		return
	}

	_, err = cli.GetChainID(ctx)
	if err == nil {
		return
	}

	err = chains.GetEndpintIntervalMGR().BackoffEndpoint(endpoint)
	if err != nil {
		logger.Sugar().Warnw("checkEndpoint", "Msg", "failed to backoffEndpoint", "Endpoint", endpoint, "Error", err)
	}
}

func Client(endpoints []string) (*solClients, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("endpoints has no item")
	}
	return &solClients{endpoints: endpoints}, nil
}
