package eth

import (
	"context"
	"errors"
	"math/rand"
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

var (
	// TODO:should check the chainID of erver endpoints,if them is not equal,panic
	CurrentChainID *string
)

type eClients struct {
	endpoints []string
}

func (eClients eClients) GetNode(ctx context.Context) (*ethclient.Client, error) {
	endpoint := eClients.endpoints[rand.Intn(len(eClients.endpoints))]

	ctx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()

	cli, err := ethclient.DialContext(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	if CurrentChainID == nil {
		_chainID, err := cli.ChainID(ctx)
		if err != nil {
			return nil, err
		}
		chainID := _chainID.String()
		CurrentChainID = &chainID
	}

	// sync check is to many,so will be canceled
	// syncRet, _err := cli.SyncProgress(ctx)
	// if _err != nil {
	// 	cli.Close()
	// 	return nil, _err
	// }

	// if syncRet != nil {
	// 	cli.Close()
	// 	return nil, fmt.Errorf(
	// 		"node is syncing ,current block %v ,highest block %v ",
	// 		syncRet.CurrentBlock, syncRet.HighestBlock,
	// 	)
	// }

	return cli, nil
}

func (eClients *eClients) WithClient(ctx context.Context, fn func(ctx context.Context, c *ethclient.Client) (bool, error)) error {
	var (
		apiErr, err error
		retry       bool
		client      *ethclient.Client
	)

	if err != nil {
		return err
	}

	for i := 0; i < utils.MinInt(MaxRetries, len(eClients.endpoints)); i++ {
		if i > 0 {
			time.Sleep(retriesSleepTime)
		}

		client, err = eClients.GetNode(ctx)

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

func Client(endpoints []string) (*eClients, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("endpoints has no item")
	}
	return &eClients{endpoints: endpoints}, nil
}
