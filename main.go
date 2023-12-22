package main

// import (
// 	"context"
// 	"time"

// 	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
// 	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/sol"
// )

// func main() {
// 	logger.Init(logger.DebugLevel, "./a.log")
// 	indeser := sol.NewSolIndexer("")
// 	indeser.UpdateEndpoints([]string{"https://ultra-weathered-patina.solana-mainnet.quiknode.pro/6d1f40b3a5315383bc3e9492e0e5e8b0fb4d1073/"})
// 	start := 237150000
// 	numChan := make(chan uint64)
// 	go indeser.IndexBlock(context.Background(), numChan)
// 	for i := 0; i < 10; i++ {
// 		numChan <- uint64(start + i)
// 	}

// 	time.Sleep(10 * time.Second)
// }
