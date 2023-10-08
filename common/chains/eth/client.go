package eth

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3eye-io/Web3Eye/common/utils"
)

const (
	MinNodeNum       = 1
	MaxRetries       = 3
	retriesSleepTime = 200 * time.Millisecond
	dialTimeout      = 3 * time.Second
)

type ethClients struct {
	endpoints []string
}

func (ethCli ethClients) GetNode(ctx context.Context) (*ethclient.Client, error) {
	randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(ethCli.endpoints))))
	if err != nil {
		return nil, err
	}
	endpoint := ethCli.endpoints[randIndex.Int64()]

	ctx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()

	cli, err := ethclient.DialContext(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (ethCli *ethClients) WithClient(ctx context.Context, fn func(ctx context.Context, c *ethclient.Client) (bool, error)) error {
	var (
		apiErr, err error
		retry       bool
		client      *ethclient.Client
	)

	if err != nil {
		return err
	}

	for i := 0; i < utils.MinInt(MaxRetries, len(ethCli.endpoints)); i++ {
		if i > 0 {
			time.Sleep(retriesSleepTime)
		}

		client, err = ethCli.GetNode(ctx)

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

func Client(endpoints []string) (*ethClients, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("endpoints has no item")
	}
	return &ethClients{endpoints: endpoints}, nil
}
