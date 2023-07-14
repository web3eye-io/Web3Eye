package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEndpointChainID(ctx context.Context, endpoint string) (string, error) {
	cli, err := ethclient.Dial(endpoint)
	if err != nil {
		return "", err
	}

	chainID, err := cli.ChainID(ctx)
	if err != nil {
		return "", err
	}

	return chainID.String(), nil
}
