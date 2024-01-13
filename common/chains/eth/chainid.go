package eth

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckEndpointChainID(ctx context.Context, endpoint string) (string, error) {
	cli, err := ethclient.Dial(endpoint)
	if err != nil {
		return "", err
	}
	defer cli.Close()

	syncRet, _err := cli.SyncProgress(ctx)
	if !(_err != nil && strings.Contains(_err.Error(), "Method not found")) {
		if _err != nil {
			return "", _err
		}

		if syncRet != nil {
			return "", fmt.Errorf(
				"node is syncing ,current block %v ,highest block %v ",
				syncRet.CurrentBlock, syncRet.HighestBlock,
			)
		}
	}

	chainID, err := cli.ChainID(ctx)
	if err != nil {
		return "", err
	}

	return chainID.String(), nil
}
