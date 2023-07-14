package chains

import (
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/eth"
	common_eth "github.com/web3eye-io/Web3Eye/common/chains/eth"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	"golang.org/x/net/context"
)

type GetEndpointChainID func(ctx context.Context, endpoint string) (ChainID string, err error)

type Index interface {
	StartIndex(ctx context.Context)
	StopIndex()
}

type indexMGR struct {
	UpdateInterval          time.Duration
	EndpointChainIDHandlers map[basetype.ChainType]GetEndpointChainID
	Indexs                  map[basetype.ChainType]map[string]Index
}

type EndpointInfo struct {
	ChainType basetype.ChainType
	Endpoint  string
	chainID   string
}

var (
	pMGR           *indexMGR
	UpdateInterval = time.Second * 10
)

func init() {
	pMGR = &indexMGR{
		UpdateInterval:          UpdateInterval,
		EndpointChainIDHandlers: make(map[basetype.ChainType]GetEndpointChainID),
	}

	pMGR.EndpointChainIDHandlers[basetype.ChainType_Ethereum] = common_eth.GetEndpointChainID
}

func GetIndexMGR() *indexMGR {
	return pMGR
}

func (pmgr *indexMGR) StartRunning(ctx context.Context) {
	eInfos := []EndpointInfo{
		{
			ChainType: basetype.ChainType_Ethereum,
			Endpoint:  "https://mainnet.infura.io/v3/03719c03f3bb46dda13decd1e58537f0",
		},
	}

	for _, info := range eInfos {
		if _, ok := pmgr.EndpointChainIDHandlers[info.ChainType]; !ok {
			logger.Sugar().Errorf("have not support chain type: %v", info.ChainType.String())
		}
		chainID, err := pmgr.EndpointChainIDHandlers[info.ChainType](ctx, info.Endpoint)
		if err != nil {
			logger.Sugar().Errorf("cannot get chainID for chainType: %v,endpoint: %v,err: %v", info.ChainType.String(), info.Endpoint, err)
		}
		info.chainID = chainID
	}
}

func (pmgr *indexMGR) AddEndpoint(info EndpointInfo) {
	if _, ok := pmgr.Indexs[info.ChainType]; !ok {
		pmgr.Indexs[info.ChainType] = make(map[string]Index)
	}

	if _, ok := pmgr.Indexs[info.ChainType][info.chainID]; !ok {
		pmgr.Indexs[info.ChainType][info.chainID] = &eth.EthIndexer{}
	}
}
