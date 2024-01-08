package sol

import (
	"context"

	solana "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/token_metadata"
)

// solana have block and slot,but we use slot repalce to block
func (solCli solClients) GetSlotHeight(ctx context.Context) (uint64, error) {
	height := uint64(0)

	var err error
	var useTimes uint16 = 1
	err = solCli.WithClient(ctx, useTimes, func(ctx context.Context, c *rpc.Client) (bool, error) {
		height, err = c.GetSlot(ctx, rpc.CommitmentFinalized)
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
	var useTimes uint16 = 1
	err = solCli.WithClient(ctx, useTimes, func(ctx context.Context, c *rpc.Client) (bool, error) {
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

func (solCli solClients) GetTX(ctx context.Context, txSig solana.Signature) (*rpc.GetTransactionResult, error) {
	tx := &rpc.GetTransactionResult{}
	maxSupportedTransactionVersion := uint64(0)
	var err error
	var useTimes uint16 = 1
	err = solCli.WithClient(ctx, useTimes, func(ctx context.Context, c *rpc.Client) (bool, error) {
		tx, err = c.GetTransaction(ctx, txSig, &rpc.GetTransactionOpts{MaxSupportedTransactionVersion: &maxSupportedTransactionVersion})
		if err != nil {
			return true, err
		}
		return false, nil
	})

	return tx, err
}

func (solCli solClients) GetMetadata(ctx context.Context, mint string) (*token_metadata.Metadata, error) {
	mintAcc := common.PublicKeyFromString(mint)
	metadataAccount, err := token_metadata.GetTokenMetaPubkey(mintAcc)
	if err != nil {
		return nil, err
	}

	metadataAcc, err := solana.PublicKeyFromBase58(metadataAccount.ToBase58())
	if err != nil {
		return nil, err
	}

	accountInfo := &rpc.GetAccountInfoResult{}
	var useTimes uint16 = 1
	err = solCli.WithClient(ctx, useTimes, func(ctx context.Context, c *rpc.Client) (bool, error) {
		accountInfo, err = c.GetAccountInfo(ctx, metadataAcc)
		if err != nil {
			return true, err
		}
		return false, nil
	})

	if err != nil {
		return nil, err
	}

	metadata, err := token_metadata.MetadataDeserialize(accountInfo.Bytes())
	return &metadata, err
}
