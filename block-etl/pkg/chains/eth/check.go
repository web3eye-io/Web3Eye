package eth

import (
	"context"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/common/chains/eth"
)

func (e *EthIndexer) CheckEndpointAndDeal(ctx context.Context) {
	// extract wrong endpoints
	for _, v := range e.Endpoints {
		_, inspectErr := eth.GetEndpointChainID(ctx, v)
		if inspectErr != nil {
			logger.Sugar().Warnf("check the endpoint %v is unavailable,err: %v,will be removed", v, inspectErr)
			e.BadEndpoints[v] = inspectErr
		}
	}

	// clean up the map
	e.BadEndpoints = make(map[string]error)
}

// check if endpoints should be stoped by the err,and return weather to retry again
func (e *EthIndexer) checkErr(ctx context.Context, err error) {
	retryErrs := []string{"context deadline exceeded"}
	for _, v := range retryErrs {
		if strings.Contains(err.Error(), v) {
			return
		}
	}

	stopErrs := []string{"401 Unauthorized", "429"}
	for _, v := range stopErrs {
		if strings.Contains(err.Error(), v) {
			e.CheckEndpointAndDeal(ctx)
			if len(e.Endpoints) == 0 {
				e.StopIndex()
			}
			return
		}
	}
}
