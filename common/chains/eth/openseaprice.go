package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/web3eye-io/Web3Eye/common/chains/eth/contracts"
)

const (
	// orderFulfilled EventHash represents the keccak256 hash of
	// OrderFulfilled (bytes32 orderHash, index_topic_1 address offerer, index_topic_2 address zone, address recipient, tuple[] offer, tuple[] consideration)
	orderFulfilledEventHash   = "0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31"
	orderFulfilledTopicLength = 3
)

var (
	openseaABI, _       = contracts.OpenseaMetaData.GetAbi()
	ItemTypeToTokenType = map[uint8]basetype.TokenType{
		0: basetype.TokenType_Native,
		1: basetype.TokenType_ERC20,
		2: basetype.TokenType_ERC721,
		3: basetype.TokenType_ERC1155,
		4: basetype.TokenType_ERC721_WITH_CRITERIA,
		5: basetype.TokenType_ERC1155_WITH_CRITERIA,
	}
)

type OrderItem struct {
	PayType       basetype.TokenType
	TokenContract string
	TokenID       string
	Amount        *big.Int
}

type OrderAccountDetails struct {
	TxHash            string
	OrderAccountItems map[string][]OrderItem
}

type OrderPricePair struct {
	Recipient   string
	TargetItems []OrderItem
	OfferItems  []OrderItem
}

type OrderPriceDetails struct {
	TxHash          string
	OrderPricePairs []OrderPricePair
}

//nolint:gocritic
func LogsToPrice(pLogs []types.Log) []*OrderPriceDetails {
	result := make([]*OrderPriceDetails, 0)
	for _, pLog := range pLogs {
		orderObj, err := LogToOrderFulfilled(pLog)
		if err != nil {
			logger.Sugar().Warnf("failed to parse OrderFulfilled log,err %v", err)
			continue
		}
		orderAD := TidyOrderAccount(orderObj)
		orderPD := CalOrderPrice(orderAD)
		result = append(result, orderPD)
	}
	return result
}

func (ethCli *ethClients) OrderFulfilledLogs(ctx context.Context, fromBlock, toBlock int64) ([]*OrderPriceDetails, error) {
	topics := [][]common.Hash{{
		common.HexToHash(string(orderFulfilledEventHash)),
	}}

	logs, err := ethCli.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Topics:    topics,
	})
	if err != nil {
		return nil, err
	}

	return LogsToPrice(logs), nil
}

func LogToOrderFulfilled(orderLog types.Log) (*contracts.OpenseaOrderFulfilled, error) {
	orderObj := contracts.OpenseaOrderFulfilled{}
	err := openseaABI.UnpackIntoInterface(&orderObj, "OrderFulfilled", orderLog.Data)
	if err != nil {
		return nil, err
	}

	if len(orderLog.Topics) < orderFulfilledTopicLength {
		return nil, fmt.Errorf("expect topics length is >= %v,but the topics length is %v", orderFulfilledTopicLength, len(orderLog.Topics))
	}

	orderObj.Offerer = common.HexToAddress(orderLog.Topics[1].Hex())
	orderObj.Zone = common.HexToAddress(string(orderLog.Topics[2].Hex()))
	orderObj.Raw.TxHash = orderLog.TxHash
	orderObj.Raw.BlockHash = orderLog.BlockHash
	orderObj.Raw.BlockNumber = orderLog.BlockNumber
	return &orderObj, nil
}

func TidyOrderAccount(orderObj *contracts.OpenseaOrderFulfilled) *OrderAccountDetails {
	accDetails := &OrderAccountDetails{
		TxHash: orderObj.Raw.TxHash.String(),
	}
	accOrderItems := make(map[string][]OrderItem)
	if _, ok := accOrderItems[orderObj.Recipient.String()]; !ok {
		accOrderItems[orderObj.Recipient.String()] = []OrderItem{}
	}
	if _, ok := accOrderItems[orderObj.Offerer.String()]; !ok {
		accOrderItems[orderObj.Offerer.String()] = []OrderItem{}
	}
	for _, v := range orderObj.Offer {
		accOrderItems[orderObj.Offerer.String()] = append(accOrderItems[orderObj.Offerer.String()], OrderItem{
			PayType:       ItemTypeToTokenType[v.ItemType],
			TokenID:       v.Identifier.String(),
			TokenContract: v.Token.String(),
			Amount:        big.NewInt(0).Neg(v.Amount),
		})
		accOrderItems[orderObj.Recipient.String()] = append(accOrderItems[orderObj.Recipient.String()], OrderItem{
			PayType:       ItemTypeToTokenType[v.ItemType],
			TokenID:       v.Identifier.String(),
			TokenContract: v.Token.String(),
			Amount:        v.Amount,
		})
	}
	for _, v := range orderObj.Consideration {
		if _, ok := accOrderItems[v.Recipient.String()]; !ok {
			accOrderItems[v.Recipient.String()] = []OrderItem{}
		}
		accOrderItems[v.Recipient.String()] = append(accOrderItems[v.Recipient.String()], OrderItem{
			PayType:       ItemTypeToTokenType[v.ItemType],
			TokenID:       v.Identifier.String(),
			TokenContract: v.Token.String(),
			Amount:        v.Amount,
		})
		accOrderItems[orderObj.Recipient.String()] = append(accOrderItems[orderObj.Recipient.String()], OrderItem{
			PayType:       ItemTypeToTokenType[v.ItemType],
			TokenID:       v.Identifier.String(),
			TokenContract: v.Token.String(),
			Amount:        big.NewInt(0).Neg(v.Amount),
		})
	}
	accDetails.OrderAccountItems = accOrderItems
	return accDetails
}

func CalOrderPrice(orderAD *OrderAccountDetails) *OrderPriceDetails {
	collectedOAD := make(map[string][]OrderItem)
	// collect sample items
	for k, v := range orderAD.OrderAccountItems {
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

	oderPD := &OrderPriceDetails{
		TxHash:          orderAD.TxHash,
		OrderPricePairs: []OrderPricePair{},
	}

	for k, v := range collectedOAD {
		targetItems := []OrderItem{}
		offerItems := []OrderItem{}
		for _, item := range v {
			if item.PayType > basetype.TokenType_ERC20 && item.Amount.Sign() > 0 {
				targetItems = append(targetItems, item)
			} else {
				offerItems = append(offerItems, item)
			}
		}
		if len(targetItems) > 0 {
			oderPD.OrderPricePairs = append(oderPD.OrderPricePairs, OrderPricePair{
				Recipient:   k,
				TargetItems: targetItems,
				OfferItems:  offerItems,
			})
		}
	}

	return oderPD
}
