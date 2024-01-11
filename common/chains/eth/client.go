package eth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3eye-io/Web3Eye/common/chains"
)

const (
	MinNodeNum       = 1
	MaxRetries       = 3
	retriesSleepTime = 500 * time.Millisecond
	dialTimeout      = 3 * time.Second
)

type ethClients struct {
	endpoints []string
}

func (ethCli *ethClients) GetNode(ctx context.Context, useTimes uint16) (*ethclient.Client, string, error) {
	endpoints := ethCli.endpoints
	if len(endpoints) == 0 {
		return nil, "", fmt.Errorf("have no available endpoints")
	}

	endpoint, err := chains.LockEndpoint(ctx, endpoints, useTimes)
	if err != nil {
		return nil, "", err
	}

	ctx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()

	cli, err := ethclient.DialContext(ctx, endpoint)

	if err != nil {
		go checkEndpoint(context.Background(), endpoint, err)
		return nil, "", err
	}

	return cli, endpoint, nil
}

func (ethCli *ethClients) WithClient(ctx context.Context, useTimes uint16, fn func(ctx context.Context, c *ethclient.Client) (bool, error)) error {
	var (
		apiErr, nodeErr error
		retry           bool
	)

	for i := 0; i < MaxRetries; i++ {
		if i > 0 {
			time.Sleep(retriesSleepTime)
		}

		client, endpoint, err := ethCli.GetNode(ctx, useTimes)
		if err != nil {
			nodeErr = err
			continue
		}

		retry, apiErr = fn(ctx, client)
		client.Close()

		if apiErr != nil {
			go checkEndpoint(context.Background(), endpoint, apiErr)
		}

		if !retry {
			return apiErr
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

	_, err = CheckEndpointChainID(ctx, endpoint)
	if err == nil {
		return
	}

	err = chains.GetEndpintIntervalMGR().BackoffEndpoint(endpoint)
	if err != nil {
		logger.Sugar().Warnw("checkEndpoint", "Msg", "failed to backoffEndpoint", "Endpoint", endpoint, "Error", err)
	}
}

func Client(endpoints []string) (*ethClients, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("endpoints has no item")
	}
	return &ethClients{endpoints: endpoints}, nil
}
