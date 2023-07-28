package sol

import (
	"context"

	"github.com/gagliardetto/solana-go/rpc"
)

func GetEndpointChainID(ctx context.Context, endpoint string) (string, error) {
	cli := rpc.New(endpoint)

	out, err := cli.GetGenesisHash(ctx)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
