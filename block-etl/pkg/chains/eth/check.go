package eth

import (
	"context"
	"strings"

	"github.com/web3eye-io/Web3Eye/common/chains/eth"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func (e *EthIndexer) CheckEndpointAndDeal(ctx context.Context) {
	items := []int{}
	for i, v := range e.Endpoints {
		_, err := eth.GetEndpointChainID(ctx, v)
		if err != nil {
			logger.Sugar().Warnf("check the endpoint %v is unavaliable,err: %v,has been removed", v, err)
			e.BadEndpoints[v] = err
			items = append(items, i)
		}
	}

	for i := len(items) - 1; i >= 0; i-- {
		e.Endpoints = append(e.Endpoints[:items[i]], e.Endpoints[items[i]+1:]...)
	}
}

func (e *EthIndexer) checkErr(ctx context.Context, err error) (retry bool) {
	if strings.Contains(err.Error(), "context deadline exceeded") {
		return true
	}
	// TODO:should check the real err
	if strings.Contains(err.Error(), "limit exceeded") {
		e.CheckEndpointAndDeal(ctx)
		if len(e.Endpoints) == 0 {
			e.StopIndex()
			return false
		}
		return true
	}
	return false
}
