package eth

import (
	"context"
	"strings"

	"github.com/web3eye-io/Web3Eye/common/chains/eth"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

// check endpoints and update it`s state on db
func (e *EthIndexer) CheckEndpointAndDeal(ctx context.Context) {
	okEndpints := []string{}
	// extract wrong endpoints
	for _, v := range e.Endpoints {
		_, inspectErr := eth.GetEndpointChainID(ctx, v)
		if inspectErr != nil {
			logger.Sugar().Warnf("check the endpoint %v is unavailable,err: %v,has been removed", v, inspectErr)
			e.BadEndpoints[v] = inspectErr
		} else {
			okEndpints = append(okEndpints, v)
		}
	}

	e.UpdateEndpoints(okEndpints)
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
	needCheck := false
	for _, v := range stopErrs {
		if strings.Contains(err.Error(), v) {
			needCheck = true
			break
		}
	}

	if len(e.Endpoints) == 0 {
		e.StopIndex()
	}
}
