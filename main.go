package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/gagliardetto/solana-go"
// 	sol_cli "github.com/web3eye-io/Web3Eye/common/chains/sol"
// 	"github.com/web3eye-io/Web3Eye/common/utils"

// 	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
// 	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/sol"
// )

// func main() {
// logger.Init(logger.DebugLevel, "./a.log")
// indeser := sol.NewSolIndexer("5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d")
// indeser.UpdateEndpoints([]string{"https://ultra-weathered-patina.solana-mainnet.quiknode.pro/6d1f40b3a5315383bc3e9492e0e5e8b0fb4d1073/"})

// cli, err := sol_cli.Client([]string{"https://ultra-weathered-patina.solana-mainnet.quiknode.pro/6d1f40b3a5315383bc3e9492e0e5e8b0fb4d1073/"})
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(0)
// }
// ret, err := cli.GetTX(context.Background(), solana.MustSignatureFromBase58("4XpUb2q8xdSv1nFk3U6WV3specXvzh8poV5upPxxUW9G39UpZ3YGQ1ipKESASBErPgyKUYDpeZb3ha15br7iYGCE"))
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(0)
// }
// fmt.Println(utils.PrettyStruct(ret))
// taskBlockNum := make(chan uint64)
// go indeser.IndexBlock(context.Background(), taskBlockNum)

// for i := 0; i < 10; i++ {
// 	taskBlockNum <- 236004010 + uint64(i)
// }

// time.Sleep(time.Second * 10)

// }
