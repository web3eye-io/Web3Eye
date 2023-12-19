package main

// import (
// 	"context"
// 	"time"

// 	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
// 	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/sol"
// )

// func main() {
// 	logger.Init(logger.DebugLevel, "./a.log")
// 	indeser := sol.NewSolIndexer("5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d ")
// 	indeser.UpdateEndpoints([]string{"https://ultra-weathered-patina.solana-mainnet.quiknode.pro/6d1f40b3a5315383bc3e9492e0e5e8b0fb4d1073/"})
// 	taskBlockNum := make(chan uint64)
// 	go indeser.IndexBlock(context.Background(), taskBlockNum)

// 	for i := 0; i < 10; i++ {
// 		taskBlockNum <- 236004010 + uint64(i)
// 	}

// 	time.Sleep(time.Second * 10)
// }
