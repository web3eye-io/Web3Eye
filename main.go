package main

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
// 	"github.com/web3eye-io/Web3Eye/common/chains/sol"
// 	"github.com/web3eye-io/Web3Eye/common/utils"
// )

// func main() {
// 	logger.Init(logger.DebugLevel, "./a.log")
// 	// chainID := "5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d "
// 	endpoint := "https://ultra-weathered-patina.solana-mainnet.quiknode.pro/6d1f40b3a5315383bc3e9492e0e5e8b0fb4d1073/"
// 	cli, err := sol.Client([]string{endpoint})
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}

// 	metaData, err := cli.GetMetadata(context.Background(), "7oLsDzr125vmw5gPTMEGvGFS2VrRbtiFcgnQXs3RdF2")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}
// 	fmt.Printf("%v", utils.PrettyStruct(metaData))
// }
