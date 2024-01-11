package eth

import (
	"context"
	"fmt"
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
				logger.Sugar().Errorw("IndexBlock", "BlockNum", num, "Error", err)
				continue
			}

			if block.ParseState != basetype.BlockParseState_BlockTypeStart {
				continue
			}

			err = func() error {
				err := e.checkOkEndpoints()
				if err != nil {
					return err
				}

				blockLogs, err := e.IndexBlockLogs(ctx, block.BlockNumber)
				if err != nil {
					return err
				}

				filteredT1, err := e.IndexTransfer(ctx, blockLogs.TransferLogs, block.BlockTime)
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

			remark := ""
			parseState := basetype.BlockParseState_BlockTypeFinish
			if err != nil {
				logger.Sugar().Errorw("IndexBlock", "BlockNum", num, "Error", err)
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
				logger.Sugar().Errorw("IndexBlock", "BlockNum", num, "Error", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (e *EthIndexer) OnNoAvalibleEndpoints(event func()) {
	e.ONAEEvents = append(e.ONAEEvents, event)
}

func (e *EthIndexer) checkOkEndpoints() error {
	if len(e.OkEndpoints) == 0 {
		for _, v := range e.ONAEEvents {
			v()
		}
		return fmt.Errorf("have no available endpoints")
	}
	return nil
}

func (e *EthIndexer) UpdateEndpoints(endpoints []string) {
	e.OkEndpoints = endpoints
	_ = e.checkOkEndpoints()
}
