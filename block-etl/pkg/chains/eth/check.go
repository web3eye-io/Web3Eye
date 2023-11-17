package eth

import (
	"context"
	"strings"

	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	endpointNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/endpoint"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

// check endpoints and update it`s state on db
func (e *EthIndexer) CheckEndpointAndDeal(ctx context.Context) {
	updateInfos := []*endpoint.EndpointReq{}
	// extract wrong endpoints
	for _, v := range e.Endpoints {
		_, inspectErr := eth.GetEndpointChainID(ctx, v)
		if inspectErr != nil {
			logger.Sugar().Warnf("check the endpoint %v is unavailable,err: %v,has been removed", v, inspectErr)
			conds := &endpoint.Conds{
				ChainType: &web3eye.StringVal{
					Op:    "eq",
					Value: e.ChainType.String(),
				},
				Address: &web3eye.StringVal{
					Op:    "eq",
					Value: v,
				},
			}
			getEResp, err := endpointNMCli.GetEndpoints(ctx, &endpoint.GetEndpointsRequest{Conds: conds})
			if err != nil {
				logger.Sugar().Warnf("get endpoints from nft-meta failed, err: %v", err)
				continue
			}
			for _, info := range getEResp.GetInfos() {
				remark := inspectErr.Error()
				updateInfos = append(updateInfos, &endpoint.EndpointReq{
					ID:        &info.ID,
					ChainType: &info.ChainType,
					ChainID:   &info.ChainID,
					Address:   &info.Address,
					State:     cttype.EndpointState_EndpointError.Enum(),
					Remark:    &remark,
				})
			}
		}
	}

	// clean up the map
	e.BadEndpoints = make(map[string]error)

	// update the infos to db
	updateEResp, err := endpointNMCli.UpdateEndpoints(ctx, &endpoint.UpdateEndpointsRequest{
		Infos: updateInfos,
	})
	if err != nil {
		logger.Sugar().Errorf("get endpoints from nft-meta failed, err: %v", err)
		return
	}
	if len(updateInfos) != 0 {
		for _, v := range updateEResp.Infos {
			logger.Sugar().Warnf("update endpoint %v failed, err: %v", v.ID, v.MSG)
		}
	}
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
