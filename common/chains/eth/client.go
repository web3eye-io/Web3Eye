package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/endpoints"
	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/config"
)

const (
	MinNodeNum       = 1
	MaxRetries       = 3
	retriesSleepTime = 200 * time.Millisecond
	dialTimeout      = 3 * time.Second
)

type EClientI interface {
	GetNode(ctx context.Context, endpointmgr *endpoints.Manager) (*ethclient.Client, error)
	WithClient(ctx context.Context, fn func(ctx context.Context, c *ethclient.Client) (bool, error)) error
}

type eClients struct{}

func (eClients eClients) GetNode(ctx context.Context, endpointmgr *endpoints.Manager) (*ethclient.Client, error) {
	endpoint, err := endpointmgr.Peek()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()

	cli, err := ethclient.DialContext(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	syncRet, _err := cli.SyncProgress(ctx)
	if _err != nil {
		cli.Close()
		return nil, _err
	}

	if syncRet != nil {
		cli.Close()
		return nil, fmt.Errorf(
			"node is syncing ,current block %v ,highest block %v ",
			syncRet.CurrentBlock, syncRet.HighestBlock,
		)
	}

	return cli, nil
}

func (eClients *eClients) WithClient(ctx context.Context, fn func(ctx context.Context, c *ethclient.Client) (bool, error)) error {
	var (
		apiErr, err error
		retry       bool
		client      *ethclient.Client
	)

	endpointmgr, err := endpoints.NewManager(config.GetConfig().ETH.Wallets)
	if err != nil {
		return err
	}

	for i := 0; i < utils.MinInt(MaxRetries, endpointmgr.Len()); i++ {
		if i > 0 {
			time.Sleep(retriesSleepTime)
		}

		client, err = eClients.GetNode(ctx, endpointmgr)

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

func Client() EClientI {
	return &eClients{}
}
