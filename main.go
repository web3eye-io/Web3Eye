package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/web3eye-io/Web3Eye/common/ctredis"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println(ctredis.TryLock("sssss", time.Second))
	// go func() {
	// 	mgr := milvusdb.NewNFTConllectionMGR()
	// 	out, err := mgr.Search(context.Background(), [][2048]float32{imageconvert.ToArrayVector([]float32{1, 1})}, 100)
	// 	fmt.Println(utils.PrettyStruct(out))
	// 	fmt.Println(err)
	// }()

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
