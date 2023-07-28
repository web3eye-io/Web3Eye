package sol

import (
	"context"

	"github.com/gagliardetto/solana-go/rpc"
)

func (solCli solClients) GetBlockHeight(ctx context.Context) (uint64, error) {
	height := uint64(0)

	var err error
	err = solCli.WithClient(ctx, func(ctx context.Context, c *rpc.Client) (bool, error) {
		height, err = c.GetBlockHeight(ctx, rpc.CommitmentFinalized)
		if err != nil {
			return false, err
		}
		return false, nil
	})

	return height, err
}

func (solCli solClients) GetBlock(ctx context.Context, slot uint64) (*rpc.GetBlockResult, error) {
	block := &rpc.GetBlockResult{}
	maxSupportedTransactionVersion := uint64(0)
	rewards := false
	var err error
	err = solCli.WithClient(ctx, func(ctx context.Context, c *rpc.Client) (bool, error) {
		block, err = c.GetBlockWithOpts(context.Background(), slot, &rpc.GetBlockOpts{
			MaxSupportedTransactionVersion: &maxSupportedTransactionVersion,
			Rewards:                        &rewards,
			TransactionDetails:             rpc.TransactionDetailsFull,
		})
		if err != nil {
			return true, err
		}
		return false, nil
	})

	return block, err
}
