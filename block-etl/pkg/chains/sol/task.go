package sol

import (
	"context"
	"strings"
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
	W3EAddressPrefix       = "w3e_"
)

type SolIndexer struct {
	CurrentBlockNum uint64
	OkEndpoints     []string
	BadEndpoints    map[string]error
	ChainType       basetype.ChainType
	ChainID         string
	// full name: On No Available Endpoints
	ONAEEvents []func()
}

func NewSolIndexer(chainID string) *indexer.Indexer {
	return indexer.NewIndexer(chainID, basetype.ChainType_Solana, &SolIndexer{
		OkEndpoints:     []string{},
		BadEndpoints:    make(map[string]error),
		ChainType:       basetype.ChainType_Solana,
		ChainID:         chainID,
		CurrentBlockNum: 0,
		ONAEEvents:      make([]func(), 0),
	})
}

func (e *SolIndexer) IndexBlock(ctx context.Context, taskBlockNum chan uint64) {
	ctx.Done()
	for {
		e.checkOkEndpoints()
		select {
		case slotNum := <-taskBlockNum:
			block, err := e.CheckBlock(ctx, slotNum)
			if err != nil {
				logger.Sugar().Error(err)
				continue
			}

			if block.ParseState == basetype.BlockParseState_BlockTypeFinish {
				continue
			}

			err = func() error {
				outTransfers1, err := e.IndexTransfer(ctx, slotNum)
				if err != nil {
					return err
				}

				outTransfers2, err := e.IndexToken(ctx, outTransfers1)
				if err != nil {
					return err
				}

				err = e.IndexContract(ctx, outTransfers2, FindContractCreator)
				if err != nil {
					return err
				}
				return nil
			}()

			// not err if slot have no block
			if strings.Contains(err.Error(), "was skipped, or missing in long-term storage") {
				err = nil
			}

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

func (e *SolIndexer) OnNoAvalibleEndpoints(event func()) {
	e.ONAEEvents = append(e.ONAEEvents, event)
}

func (e *SolIndexer) checkOkEndpoints() {
	if len(e.OkEndpoints) == 0 {
		for _, v := range e.ONAEEvents {
			v()
		}
	}
}

func (e *SolIndexer) UpdateEndpoints(endpoints []string) {
	e.OkEndpoints = endpoints
	e.checkOkEndpoints()
}
