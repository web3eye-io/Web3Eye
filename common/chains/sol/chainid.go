package sol

import (
	"context"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/web3eye-io/Web3Eye/common/chains"
)

func CheckEndpointChainID(ctx context.Context, endpoint string) (string, error) {
	useTimes := uint16(2)
	endpoint, err := chains.LockEndpoint(ctx, []string{endpoint}, useTimes)
	if err != nil {
		return "", err
	}

	cli := rpc.New(endpoint)

	_, err = cli.GetHealth(ctx)
	if err != nil {
		return "", err
	}

	out, err := cli.GetGenesisHash(ctx)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
