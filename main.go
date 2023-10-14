package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	"github.com/web3eye-io/Web3Eye/ranker/api/v1/token"
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// ret, err := sol.GetEndpointChainID(context.TODO(), "https://distinguished-floral-mountain.solana-mainnet.discover.quiknode.pro/c641daff8873a3f24f2f4c90aae89373707c2886/")
	// fmt.Println(ret, err)
	logger.Init(logger.DebugLevel, "./a.log")
	db.Init()
	server := token.Server{}
	ret, err := server.Search(context.Background(), &rankernpool.SearchTokenRequest{
		Vector: []float32{},
		Limit:  100,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(ret.Infos))

	<-sigchan
	os.Exit(1)
}
