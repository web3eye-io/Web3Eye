package eth

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/indexer"
	blockNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/block"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	blockProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

const (
	CheckTopicInterval     = time.Second * 10
	FindContractCreator    = false
	redisExpireDefaultTime = time.Second * 10
)

type EthIndexer struct {
	CurrentBlockNum uint64
	OkEndpoints     []string
	BadEndpoints    map[string]error
	ChainType       basetype.ChainType
	ChainID         string
	// full name: On No Available Endpoints
	ONAEEvents []func()
}

func NewEthIndexer(chainID string) *indexer.Indexer {
	return indexer.NewIndexer(chainID, basetype.ChainType_Ethereum, &EthIndexer{
		OkEndpoints:     []string{},
		BadEndpoints:    make(map[string]error),
		ChainType:       basetype.ChainType_Ethereum,
		ChainID:         chainID,
		CurrentBlockNum: 0,
		ONAEEvents:      make([]func(), 0),
	})
}

func (e *EthIndexer) IndexBlock(ctx context.Context, taskBlockNum chan uint64) {
	ctx.Done()
	for {
		select {
		case num := <-taskBlockNum:
			block, err := e.CheckBlock(ctx, num)
			if err != nil {
				logger.Sugar().Error(err)
				continue
			}

			if block.ParseState == basetype.BlockParseState_BlockTypeFinish {
				continue
			}

			err = func() error {
				blockLogs, err := e.IndexBlockLogs(ctx, num)
				if err != nil {
					return err
				}

				filteredT1, err := e.IndexTransfer(ctx, blockLogs.TransferLogs)
				if err != nil {
					return err
				}

				contractT1, err := e.IndexToken(ctx, filteredT1)
				if err != nil {
					return err
				}

				err = e.IndexContract(ctx, contractT1, FindContractCreator)
				if err != nil {
					return err
				}

				contractT2, err := e.IndexOrder(ctx, blockLogs.OrderLogs)
				if err != nil {
					return err
				}

				err = e.IndexContract(ctx, contractT2, FindContractCreator)
				if err != nil {
					return err
				}
				return nil
			}()

			if err != nil {
				logger.Sugar().Error(err)
			}

			remark := ""
			parseState := basetype.BlockParseState_BlockTypeFinish
			if err != nil {
				remark = err.Error()
				parseState = basetype.BlockParseState_BlockTypeFailed
			}

			_, err = blockNMCli.UpdateBlock(ctx, &blockProto.UpdateBlockRequest{
				Info: &blockProto.BlockReq{
					ID:         &block.ID,
					ParseState: &parseState,
					Remark:     &remark,
				},
			})
			if err != nil {
				logger.Sugar().Error(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (e *EthIndexer) OnNoAvalibleEndpoints(event func()) {
	e.ONAEEvents = append(e.ONAEEvents, event)
}

func (e *EthIndexer) UpdateEndpoints(endpoints []string) {
	e.OkEndpoints = endpoints
	if len(e.OkEndpoints) == 0 {
		for _, v := range e.ONAEEvents {
			v()
		}
	}
}
