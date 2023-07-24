package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/metaplex/tokenmeta"
	"github.com/web3eye-io/Web3Eye/common/utils"
)

func init() {
	logger.Init(logger.DebugLevel, "./a.log")
}

func main() {
	// doSomeThing()
	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// <-sigchan
	// os.Exit(1)
	TestSDK()
}

func doSomeThing() {
	cli := rpc.New("https://distinguished-floral-mountain.solana-mainnet.discover.quiknode.pro/c641daff8873a3f24f2f4c90aae89373707c2886/")

	// // out, err := cli.GetBlockHeight(context.Background(), rpc.CommitmentFinalized)
	// // if err != nil {
	// // 	logger.Sugar().Error(err)
	// // 	os.Exit(1)
	// // }

	// // fmt.Println(out)

	maxSupportedTransactionVersion := uint64(0)
	rewards := false
	block, err := cli.GetBlockWithOpts(context.Background(), 207205793, &rpc.GetBlockOpts{
		MaxSupportedTransactionVersion: &maxSupportedTransactionVersion,
		Rewards:                        &rewards,
		TransactionDetails:             rpc.TransactionDetailsFull,
	})
	if err != nil {
		logger.Sugar().Error(err)
		os.Exit(1)
	}
	for i, v := range block.Transactions {
		if v.Meta.Err != nil {
			fmt.Println(i + 1)
			fmt.Println(len(v.Meta.PreBalances))
			fmt.Println(len(v.Meta.PostBalances))
			fmt.Println(i + 1)
		}
	}
}

type NFTTokenTransfer struct {
	From    string
	To      string
	TokenID string
}

func GetNftTransfers(info *rpc.TransactionMeta, filterErrTX bool) {

	if info == nil || (filterErrTX && info.Err != nil) {
		return
	}

	preNftTransfers := []rpc.TokenBalance{}
	postNftTransfers := []rpc.TokenBalance{}
	nftTransfers := []NFTTokenTransfer{}

	for _, v := range info.PreTokenBalances {
		if v.UiTokenAmount.Amount == "1" && v.UiTokenAmount.Decimals == 0 {
			preNftTransfers = append(preNftTransfers, v)
		}
	}

	for _, v := range info.PostTokenBalances {
		if v.UiTokenAmount.Amount == "1" && v.UiTokenAmount.Decimals == 0 {
			postNftTransfers = append(postNftTransfers, v)
		}
	}

	for _, v1 := range postNftTransfers {
		from := "0"
		to := v1.Mint.String()
		tokenID := v1.Owner.String()
		for _, v2 := range preNftTransfers {
			if v1.Mint == v2.Mint {
				from = v2.Owner.String()
				continue
			}
		}

		nftTransfers = append(nftTransfers, NFTTokenTransfer{
			From:    from,
			To:      to,
			TokenID: tokenID,
		})
	}

	fmt.Println("utils.PrettyStruct(info.PreTokenBalances)")
	fmt.Println(utils.PrettyStruct(preNftTransfers))
	fmt.Println("utils.PrettyStruct(info.PostTokenBalances)")
	fmt.Println(utils.PrettyStruct(postNftTransfers))

	fmt.Println("utils.PrettyStruct(nftTransfers)")
	fmt.Println(utils.PrettyStruct(nftTransfers))
}

func GetTokenMeta(mint string) {

}

func TestSDK() {
	// NFT in solana is a normal mint but only mint 1.
	// If you want to get its metadata, you need to know where it stored.
	// and you can use `tokenmeta.GetTokenMetaPubkey` to get the metadata account key
	// here I take a random Degenerate Ape Academy as an example
	mint := common.PublicKeyFromString("DSwfRF1jhhu6HpSuzaig1G19kzP73PfLZBPLofkw6fLD")
	metadataAccount, err := tokenmeta.GetTokenMetaPubkey(mint)
	if err != nil {
		log.Fatalf("faield to get metadata account, err: %v", err)
	}

	mAcc, err := solana.PublicKeyFromBase58(metadataAccount.ToBase58())
	if err != nil {
		log.Fatalf("faield to get metadata account, err: %v", err)
	}

	// new a client
	endpoint := "https://distinguished-floral-mountain.solana-mainnet.discover.quiknode.pro/c641daff8873a3f24f2f4c90aae89373707c2886/"
	c := rpc.New(endpoint)
	// c := client.NewClient(rpc.MainnetRPCEndpoint)

	// get data which stored in metadataAccount
	accountInfo, err := c.GetAccountInfo(context.Background(), mAcc)
	if err != nil {
		log.Fatalf("failed to get accountInfo, err: %v", err)
	}

	// parse it
	metadata, err := tokenmeta.MetadataDeserialize(accountInfo.Bytes())
	if err != nil {
		log.Fatalf("failed to parse metaAccount, err: %v", err)
	}

	fmt.Println(utils.PrettyStruct(metadata))
}
