package eth

import (
	"fmt"
	"math/big"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"

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
	OrderFulfilledTopics = []common.Hash{
		common.HexToHash(orderFulfilledEventHash),
	}
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
	OrderAccountItems map[string][]*npool.OrderItem
}

type OrderPricePair struct {
	Recipient   string
	TargetItems []*npool.OrderItem
	OfferItems  []*npool.OrderItem
}

type OrderPriceDetails struct {
	TxHash          string
	OrderPricePairs []OrderPricePair
}

//nolint:gocritic
func LogsToOrders(pLogs []types.Log) []*npool.Order {
	result := make([]*npool.Order, 0)
	for _, pLog := range pLogs {
		orderObj, err := LogToOrderFulfilled(pLog)
		if err != nil {
			logger.Sugar().Warnf("failed to parse OrderFulfilled log,err %v", err)
			continue
		}
		orderAD := TidyOrderAccount(orderObj)
		orderPD := CalOrderPrice(orderAD)
		for _, v := range orderPD.OrderPricePairs {

			result = append(result, &npool.Order{
				TxHash:      pLog.TxHash.String(),
				BlockNumber: pLog.BlockNumber,
				TxIndex:     uint32(pLog.TxIndex),
				LogIndex:    uint32(pLog.Index),
				Recipient:   v.Recipient,
				TargetItems: v.TargetItems,
				OfferItems:  v.OfferItems,
			})
		}
	}
	return result
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
	accOrderItems := make(map[string][]*npool.OrderItem)
	if _, ok := accOrderItems[orderObj.Recipient.String()]; !ok {
		accOrderItems[orderObj.Recipient.String()] = []*npool.OrderItem{}
	}
	if _, ok := accOrderItems[orderObj.Offerer.String()]; !ok {
		accOrderItems[orderObj.Offerer.String()] = []*npool.OrderItem{}
	}
	for _, v := range orderObj.Offer {
		accOrderItems[orderObj.Offerer.String()] = append(accOrderItems[orderObj.Offerer.String()], &npool.OrderItem{
			TokenType: ItemTypeToTokenType[v.ItemType],
			TokenID:   v.Identifier.String(),
			Contract:  v.Token.String(),
			Amount:    -v.Amount.Int64(),
		})
		accOrderItems[orderObj.Recipient.String()] = append(accOrderItems[orderObj.Recipient.String()], &npool.OrderItem{
			TokenType: ItemTypeToTokenType[v.ItemType],
			TokenID:   v.Identifier.String(),
			Contract:  v.Token.String(),
			Amount:    v.Amount.Int64(),
		})
	}
	for _, v := range orderObj.Consideration {
		if _, ok := accOrderItems[v.Recipient.String()]; !ok {
			accOrderItems[v.Recipient.String()] = []*npool.OrderItem{}
		}
		accOrderItems[v.Recipient.String()] = append(accOrderItems[v.Recipient.String()], &npool.OrderItem{
			TokenType: ItemTypeToTokenType[v.ItemType],
			TokenID:   v.Identifier.String(),
			Contract:  v.Token.String(),
			Amount:    v.Amount.Int64(),
		})
		accOrderItems[orderObj.Recipient.String()] = append(accOrderItems[orderObj.Recipient.String()], &npool.OrderItem{
			TokenType: ItemTypeToTokenType[v.ItemType],
			TokenID:   v.Identifier.String(),
			Contract:  v.Token.String(),
			Amount:    v.Amount.Int64(),
		})
	}
	accDetails.OrderAccountItems = accOrderItems
	return accDetails
}

func CalOrderPrice(orderAD *OrderAccountDetails) *OrderPriceDetails {
	collectedOAD := make(map[string][]*npool.OrderItem)
	// collect sample items
	for k, v := range orderAD.OrderAccountItems {
		tokenSet := make(map[string]map[string]*npool.OrderItem)
		for _, item := range v {
			if _, ok := tokenSet[item.Contract]; !ok {
				tokenSet[item.Contract] = make(map[string]*npool.OrderItem)
			}
			if _, ok := tokenSet[item.Contract][item.TokenID]; !ok {
				tokenSet[item.Contract][item.TokenID] = item
			} else {
				lastItem := tokenSet[item.Contract][item.TokenID]
				lastItem.Amount = tokenSet[item.Contract][item.TokenID].Amount + item.Amount
				tokenSet[item.Contract][item.TokenID] = lastItem
			}
		}
		collectedOAD[k] = []*npool.OrderItem{}
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
		targetItems := []*npool.OrderItem{}
		offerItems := []*npool.OrderItem{}
		for _, item := range v {
			if item.TokenType > basetype.TokenType_ERC20 && item.Amount > 0 {
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
