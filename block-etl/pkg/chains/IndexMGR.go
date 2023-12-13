package chains

import (
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/eth"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/sol"
	endpointNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/endpoint"

	common_eth "github.com/web3eye-io/Web3Eye/common/chains/eth"
	common_sol "github.com/web3eye-io/Web3Eye/common/chains/sol"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
	"golang.org/x/net/context"
)

type GetEndpointChainID func(ctx context.Context, endpoint string) (ChainID string, err error)

type Index interface {
	StartIndex(ctx context.Context)
	IsOnIndex() bool
	UpdateEndpoints(endpoints []string)
	StopIndex()
}

type indexMGR struct {
	UpdateInterval          time.Duration
	EndpointChainIDHandlers map[basetype.ChainType]GetEndpointChainID
	Indexs                  map[basetype.ChainType]map[string]Index
	endpointGroups          map[basetype.ChainType]map[string][]string
}

var (
	pMGR              *indexMGR
	UpdateInterval    = time.Second * 10
	maxUsingEndpoints = 10
)

func init() {
	pMGR = &indexMGR{
		UpdateInterval:          UpdateInterval,
		Indexs:                  make(map[basetype.ChainType]map[string]Index),
		EndpointChainIDHandlers: make(map[basetype.ChainType]GetEndpointChainID),
		endpointGroups:          make(map[basetype.ChainType]map[string][]string),
	}

	// TODO:should be registered
	pMGR.EndpointChainIDHandlers[basetype.ChainType_Ethereum] = common_eth.GetEndpointChainID
	pMGR.EndpointChainIDHandlers[basetype.ChainType_Solana] = common_sol.GetEndpointChainID
}

func GetIndexMGR() *indexMGR {
	return pMGR
}

func (pmgr *indexMGR) Run(ctx context.Context) {
	for {
		pmgr.checkNewEndpoints(ctx)
		pmgr.checkAvaliableEndpoints(ctx)
		<-time.NewTicker(UpdateInterval).C
	}
}

// check for the newly created endpoints
func (pmgr *indexMGR) checkNewEndpoints(ctx context.Context) {
	conds := &endpoint.Conds{
		State: &web3eye.Uint32Val{
			Op:    "eq",
			Value: uint32(basetype.EndpointState_EndpointDefault),
		},
	}

	getEResp, err := endpointNMCli.GetEndpoints(ctx, &endpoint.GetEndpointsRequest{
		Conds:  conds,
		Limit:  int32(maxUsingEndpoints),
		Offset: 0,
	})
	if err != nil {
		logger.Sugar().Errorf("get endpoints from nft-meta failed, err: %v", err)
		return
	}

	infos := getEResp.GetInfos()
	updateInfos := []*endpoint.EndpointReq{}
	for _, info := range infos {
		func() {
			handler, ok := pmgr.EndpointChainIDHandlers[info.ChainType]
			if !ok {
				logger.Sugar().Warnf("have not handler for chain type: %v", info.ChainType)
				info.State = basetype.EndpointState_EndpointError
				return
			}

			chainID, err := handler(ctx, info.Address)
			if err != nil {
				info.State = basetype.EndpointState_EndpointError
				return
			}

			info.ChainID = chainID
			info.State = basetype.EndpointState_EndpointAvaliable
		}()

		updateInfos = append(updateInfos, &endpoint.EndpointReq{
			ID:        &info.ID,
			ChainType: &info.ChainType,
			ChainID:   &info.ChainID,
			Address:   &info.Address,
			State:     &info.State,
		})
	}

	if len(updateInfos) == 0 {
		return
	}

	updateEResp, err := endpointNMCli.UpdateEndpoints(ctx, &endpoint.UpdateEndpointsRequest{
		Infos: updateInfos,
	})
	if err != nil {
		logger.Sugar().Errorf("update endpoints to nft-meta failed, err: %v", err)
		return
	}
	if len(updateInfos) != 0 {
		for _, v := range updateEResp.Infos {
			logger.Sugar().Errorf("update endpoint %v failed, err: %v", v.ID, v.MSG)
		}
	}
}

// check erver chantype-chainid available endpoints and update it to indexer
func (pmgr *indexMGR) checkAvaliableEndpoints(ctx context.Context) {
	conds := &endpoint.Conds{
		State: &web3eye.Uint32Val{
			Op:    "eq",
			Value: uint32(basetype.EndpointState_EndpointAvaliable),
		},
	}

	getEResp, err := endpointNMCli.GetEndpoints(ctx, &endpoint.GetEndpointsRequest{Conds: conds})
	if err != nil {
		logger.Sugar().Errorf("get endpoints from nft-meta failed, err: %v", err)
		return
	}

	endpointGroups := make(map[basetype.ChainType]map[string][]string)

	infos := getEResp.GetInfos()
	for _, info := range infos {
		if _, ok := endpointGroups[info.ChainType]; !ok {
			endpointGroups[info.ChainType] = make(map[string][]string)
		}
		if _, ok := endpointGroups[info.ChainType][info.ChainID]; !ok {
			endpointGroups[info.ChainType][info.ChainID] = []string{}
		}
		endpointGroups[info.ChainType][info.ChainID] = append(endpointGroups[info.ChainType][info.ChainID], info.Address)
	}

	// check if have no endpoints,will be stop
	for chainType, v := range pmgr.endpointGroups {
		for chainID := range v {
			if _, ok := endpointGroups[chainType][chainID]; !ok {
				pmgr.Indexs[chainType][chainID].StopIndex()
			}
		}
	}

	// update endpoints for ever indexer
	pmgr.endpointGroups = endpointGroups
	for chainType, v := range pmgr.endpointGroups {
		for chainID, endpoints := range v {
			pmgr.updateEndpoints(ctx, chainType, chainID, endpoints)
		}
	}
}

// update to indexer
func (pmgr *indexMGR) updateEndpoints(ctx context.Context, chanType basetype.ChainType, chainID string, endpoints []string) {
	if _, ok := pmgr.Indexs[chanType]; !ok {
		pmgr.Indexs[chanType] = make(map[string]Index)
	}

	if _, ok := pmgr.Indexs[chanType][chainID]; !ok {
		switch chanType {
		case basetype.ChainType_Ethereum:
			pmgr.Indexs[chanType][chainID] = eth.NewEthIndexer(chainID)
		case basetype.ChainType_Solana:
			pmgr.Indexs[chanType][chainID] = sol.NewIndexer(chainID)
		default:
			logger.Sugar().Warnf("have no chainType: %v chainID: %v indexer type,will skip", chanType, chainID)
			return
		}
	}

	pmgr.Indexs[chanType][chainID].UpdateEndpoints(endpoints)
	if !pmgr.Indexs[chanType][chainID].IsOnIndex() {
		pmgr.Indexs[chanType][chainID].StartIndex(ctx)
	}
}
