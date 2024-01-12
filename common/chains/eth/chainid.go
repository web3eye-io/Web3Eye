package eth

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3eye-io/Web3Eye/common/chains"
)

func CheckEndpointChainID(ctx context.Context, endpoint string) (string, error) {
	useTimes := uint16(2)
	endpoint, err := chains.LockEndpoint(ctx, []string{endpoint}, useTimes)
	if err != nil {
		return "", err
	}

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
