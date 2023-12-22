package sol

import (
	"context"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/common/chains/sol"
)

// check if endpoints should be stoped by the err,and return weather to retry again
func (e *SolIndexer) checkErr(ctx context.Context, err error) {
	// Code -32007 slot not have block
	retryErrs := []string{"context deadline exceeded", "context canceled", "Code: (int) -32007"}
	for _, v := range retryErrs {
		if strings.Contains(err.Error(), v) {
			return
		}
	}

	okEndpints := []string{}
	// extract wrong endpoints
	for _, v := range e.OkEndpoints {
		_, inspectErr := sol.GetEndpointChainID(ctx, v)
		if inspectErr != nil {
			logger.Sugar().Warnf("check the endpoint %v is unavailable,err: %v,has been removed", v, inspectErr)
			e.BadEndpoints[v] = inspectErr
		} else {
			okEndpints = append(okEndpints, v)
		}
	}

	e.UpdateEndpoints(okEndpints)
}
