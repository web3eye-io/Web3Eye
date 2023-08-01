package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/web3eye-io/Web3Eye/block-etl/pkg/token"
	"github.com/web3eye-io/Web3Eye/common/utils"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	uri := "https://ipfs.walken.io/ipfs/QmaBnkkD8KmMLzZuoXG4KrRDeoJBqWSqhYkCHLGev5HmD6"
	tokenURI, err := token.GetTokenURIInfo(context.Background(), uri)
	fmt.Println(err)
	fmt.Println(utils.PrettyStruct(tokenURI))
	fmt.Println(token.TokenURIType(uri))
	<-sigchan
	os.Exit(1)
}

// func TestBlock() {
// 	cli, err := sol.Client([]string{"https://distinguished-floral-mountain.solana-mainnet.discover.quiknode.pro/c641daff8873a3f24f2f4c90aae89373707c2886/"})
// 	fmt.Println(err)
// 	block, err := cli.GetBlock(context.Background(), 208600970)
// 	fmt.Println(err)
// 	fmt.Println(len(sol.GetNFTTransfers(block)))
// }
