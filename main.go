package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/common/utils"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	price()

	// balance()
	<-sigchan
	os.Exit(1)
}

func balance() {
	cli, err := ethclient.Dial("https://mainnet.infura.io/v3/8cc70eaecd7c40d9817b6f4747f0e2f7")
	if err != nil {
		panic(err)
	}
	b1, err := cli.BalanceAt(context.Background(), common.HexToAddress("0x3c310fA26D8Da144474842f902699DDc5800840e"), big.NewInt(18024891))
	b2, err := cli.BalanceAt(context.Background(), common.HexToAddress("0x3c310fA26D8Da144474842f902699DDc5800840e"), big.NewInt(18024892))
	b3, err := cli.BalanceAt(context.Background(), common.HexToAddress("0x3c310fA26D8Da144474842f902699DDc5800840e"), big.NewInt(18024893))
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
}

func price() {
	OrderFulfilledTopic := "0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31"
	cli, err := ethclient.Dial("https://mainnet.infura.io/v3/8cc70eaecd7c40d9817b6f4747f0e2f7")
	if err != nil {
		panic(err)
	}

	topics := [][]common.Hash{{
		common.HexToHash(string(OrderFulfilledTopic)),
	}}

	logs, err := cli.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(18024821),
		ToBlock:   big.NewInt(18024821),
		Topics:    topics,
	})
	if err != nil {
		panic(err)
	}
	orderPD := eth.LogsToPrice(logs)
	fmt.Println(utils.PrettyStruct(orderPD))

	// fmt.Println(string(logs[0].Data))
	// fmt.Println(utils.PrettyStruct(logs[0]))
	fmt.Println("utils.PrettyStruct(logs[1])")
	// fmt.Println(utils.PrettyStruct(logs[1]))
}
