package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3eye-io/Web3Eye/common/chains/eth/contracts"
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

type OrderPayType uint8

const (
	OrderPayType_NATIVE                OrderPayType = iota
	OrderPayType_ERC20                 OrderPayType = iota
	OrderPayType_ERC721                OrderPayType = iota
	OrderPayType_ERC1155               OrderPayType = iota
	OrderPayType_ERC721_WITH_CRITERIA  OrderPayType = iota
	OrderPayType_ERC1155_WITH_CRITERIA OrderPayType = iota
)

type OrderItem struct {
	PayType       OrderPayType
	TokenContract string
	TokenID       string
	Amount        *big.Int
}

type OrderAccountDetails map[string][]OrderItem

func CountAccount(orderObj *contracts.ContractsOrderFulfilled) OrderAccountDetails {
	accDetails := make(OrderAccountDetails)
	if _, ok := accDetails[orderObj.Recipient.String()]; !ok {
		accDetails[orderObj.Recipient.String()] = []OrderItem{}
	}
	if _, ok := accDetails[orderObj.Offerer.String()]; !ok {
		accDetails[orderObj.Offerer.String()] = []OrderItem{}
	}
	for _, v := range orderObj.Offer {
		accDetails[orderObj.Offerer.String()] = append(accDetails[orderObj.Offerer.String()], OrderItem{
			PayType:       OrderPayType(v.ItemType),
			TokenID:       v.Identifier.String(),
			TokenContract: v.Token.String(),
			Amount:        big.NewInt(0).Neg(v.Amount),
		})
		accDetails[orderObj.Recipient.String()] = append(accDetails[orderObj.Recipient.String()], OrderItem{
			PayType:       OrderPayType(v.ItemType),
			TokenID:       v.Identifier.String(),
			TokenContract: v.Token.String(),
			Amount:        v.Amount,
		})
	}
	for _, v := range orderObj.Consideration {
		if _, ok := accDetails[v.Recipient.String()]; !ok {
			accDetails[v.Recipient.String()] = []OrderItem{}
		}
		accDetails[v.Recipient.String()] = append(accDetails[v.Recipient.String()], OrderItem{
			PayType:       OrderPayType(v.ItemType),
			TokenID:       v.Identifier.String(),
			TokenContract: v.Token.String(),
			Amount:        v.Amount,
		})
		accDetails[orderObj.Recipient.String()] = append(accDetails[orderObj.Recipient.String()], OrderItem{
			PayType:       OrderPayType(v.ItemType),
			TokenID:       v.Identifier.String(),
			TokenContract: v.Token.String(),
			Amount:        big.NewInt(0).Neg(v.Amount),
		})
	}
	return accDetails
}

func GetOrderFulfilled(orderLog types.Log) (*contracts.ContractsOrderFulfilled, error) {
	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.ContractsMetaData.ABI)))
	if err != nil {
		panic(err)
	}
	orderObj := contracts.ContractsOrderFulfilled{}
	err = contractAbi.UnpackIntoInterface(&orderObj, "OrderFulfilled", orderLog.Data)
	if err != nil {
		return nil, err
	}

	if len(orderLog.Topics) < 3 {
		return nil, fmt.Errorf("expect topics length is >= 3,but the topics length is %v", len(orderLog.Topics))
	}

	orderObj.Offerer = common.HexToAddress(orderLog.Topics[1].Hex())
	orderObj.Zone = common.HexToAddress(string(orderLog.Topics[2].Hex()))
	return &orderObj, nil
}

func GetTokenPrice(oad OrderAccountDetails) {
	collectedOAD := make(OrderAccountDetails)
	// collect sample items
	for k, v := range oad {
		tokenSet := make(map[string]map[string]OrderItem)
		for _, item := range v {
			if _, ok := tokenSet[item.TokenContract]; !ok {
				tokenSet[item.TokenContract] = make(map[string]OrderItem)
			}
			if _, ok := tokenSet[item.TokenContract][item.TokenID]; !ok {
				tokenSet[item.TokenContract][item.TokenID] = item
			} else {
				lastItem := tokenSet[item.TokenContract][item.TokenID]
				lastItem.Amount = big.NewInt(0).Add(tokenSet[item.TokenContract][item.TokenID].Amount, item.Amount)
				tokenSet[item.TokenContract][item.TokenID] = lastItem
			}
		}
		collectedOAD[k] = []OrderItem{}
		for _, items := range tokenSet {
			for _, item := range items {
				collectedOAD[k] = append(collectedOAD[k], item)
			}
		}
	}

	// fmt.Println(utils.PrettyStruct(collectedOAD))
	pricePair := [][2][]OrderItem{}

	for _, v := range collectedOAD {
		tokens := []OrderItem{}
		payments := []OrderItem{}
		for _, item := range v {
			if item.PayType > OrderPayType_ERC20 && item.Amount.Sign() > 0 {
				tokens = append(tokens, item)
			} else {
				payments = append(payments, item)
			}
		}
		if len(tokens) > 0 {
			pricePair = append(pricePair, [2][]OrderItem{tokens, payments})
		}
	}
	fmt.Println(utils.PrettyStruct(pricePair))
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
		FromBlock: big.NewInt(18031213),
		ToBlock:   big.NewInt(18031213),
		Topics:    topics,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range logs {
		if v.TxHash.String() == "0xadf7e1261b906b41f1c43a73d8eeeeb8de11b5fe5b59fa4adaa3510306774048" {
			orderObj, err := GetOrderFulfilled(v)
			if err != nil {
				panic(err)
			}
			fmt.Println(utils.PrettyStruct(orderObj))

			if len(orderObj.Offer) > 1 {
				fmt.Println(v.TxHash)
			}
			accDetails := CountAccount(orderObj)
			GetTokenPrice(accDetails)
		}
	}

	// fmt.Println(string(logs[0].Data))
	// fmt.Println(utils.PrettyStruct(logs[0]))
	fmt.Println("utils.PrettyStruct(logs[1])")
	// fmt.Println(utils.PrettyStruct(logs[1]))
}
