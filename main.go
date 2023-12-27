package main

// import (
// 	"context"
// 	"fmt"
// 	"math/big"

// 	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/web3eye-io/Web3Eye/common/chains/eth"
// 	"github.com/web3eye-io/Web3Eye/common/utils"
// )

// func main() {
// 	logger.Init(logger.DebugLevel, "./a.log")
// 	cli, err := ethclient.DialContext(context.Background(), "https://mainnet.infura.io/v3/8cc70eaecd7c40d9817b6f4747f0e2f7")
// 	fmt.Println(err)
// 	ret, err := cli.FilterLogs(context.Background(), ethereum.FilterQuery{
// 		FromBlock: big.NewInt(18874059),
// 		ToBlock:   big.NewInt(18874059),
// 		Topics: [][]common.Hash{
// 			eth.TransfersTopics,
// 			eth.OrderFulfilledTopics,
// 		},
// 	})
// 	fmt.Println(err)
// 	fmt.Println(utils.PrettyStruct(ret))

// 	ecli, err := eth.Client([]string{"https://mainnet.infura.io/v3/8cc70eaecd7c40d9817b6f4747f0e2f7"})
// 	fmt.Println(err)
// 	eret, err := ecli.FilterLogsForTopics(context.Background(), 18874059, 18874059, [][]common.Hash{
// 		eth.TransfersTopics,
// 		eth.OrderFulfilledTopics})
// 	fmt.Println(utils.PrettyStruct(eret))
// }
