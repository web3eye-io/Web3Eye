package eth

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	orderNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/order"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

// index blocks [start,start+limit]
func (e *EthIndexer) IndexOrder(ctx context.Context, logs []*types.Log) ([]*ContractMeta, error) {
	orders := eth.LogsToOrders(logs)
	if len(orders) == 0 {
		return nil, nil
	}

	chainType := e.ChainType
	chainID := e.ChainID
	ordersReq := make([]*order.OrderReq, len(orders))
	for i, v := range orders {
		ordersReq[i] = &order.OrderReq{
			ChainType:   &chainType,
			ChainID:     &chainID,
			TxHash:      &v.TxHash,
			BlockNumber: &v.BlockNumber,
			TxIndex:     &v.TxIndex,
			LogIndex:    &v.LogIndex,
			Recipient:   &v.Recipient,
			TargetItems: v.TargetItems,
			OfferItems:  v.OfferItems,
			Remark:      &v.Remark,
		}
	}

	_, err := orderNMCli.CreateOrders(ctx, &order.CreateOrdersRequest{Infos: ordersReq})
	if err != nil {
		return nil, fmt.Errorf("failed store orders to db,err: %v", err)
	}

	contractSet := make(map[string]struct{})
	contractList := []*ContractMeta{}
	for _, item := range orders {
		items := []*order.OrderItem{}
		copy(items, item.TargetItems)
		items = append(items, item.OfferItems...)
		for _, item := range items {
			if _, ok := contractSet[item.Contract]; ok {
				continue
			}
			contractSet[item.Contract] = struct{}{}
			contractList = append(contractList, &ContractMeta{TokenType: item.TokenType, Contract: item.Contract})
		}
	}

	return contractList, nil
}
