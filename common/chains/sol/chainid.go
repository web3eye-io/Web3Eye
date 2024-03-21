package sol

import (
	"context"

	"github.com/gagliardetto/solana-go/rpc"
)

func CheckEndpointChainID(ctx context.Context, endpoint string) (string, error) {
	cli := rpc.New(endpoint)

	_, err := cli.GetHealth(ctx)
	if err != nil {
		return "", err
	}

	out, err := cli.GetGenesisHash(ctx)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
