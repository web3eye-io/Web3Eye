package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/token"
	"github.com/web3eye-io/Web3Eye/common/chains"
	"github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	blockNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/block"
	contractNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
	synctaskNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/synctask"
	tokenNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	transferNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/transfer"
	ctMessage "github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	blockProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
	contractProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
	tokenProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	transferProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

const (
	CheckTopicInterval     = time.Second * 10
	FindContractCreator    = false
	redisExpireDefaultTime = time.Second * 10
	maxTopicNum            = 5
	updateBlockNumInterval = time.Minute
)

type EthIndexer struct {
	Endpoints       []string
	BadEndpoints    map[string]error
	ChainType       basetype.ChainType
	ChainID         string
	CurrentBlockNum uint64
	onIndex         bool
	cancel          *context.CancelFunc
}

func NewIndexer(chainID string) *EthIndexer {
	return &EthIndexer{
		BadEndpoints:    make(map[string]error),
		ChainType:       basetype.ChainType_Ethereum,
		ChainID:         chainID,
		CurrentBlockNum: 0,
	}
}

func (e *EthIndexer) StartIndex(ctx context.Context) {
	logger.Sugar().Infof("start the indexer chainType: %v, chainID: %v", e.ChainType, e.ChainID)
	ctx, cancel := context.WithCancel(ctx)
	e.cancel = &cancel
	go e.GetCurrentBlockNum(ctx, updateBlockNumInterval)

	taskBlockNum := make(chan uint64)
	indexBlockNum := make(chan uint64)
	outTransfers := make(chan []*chains.TokenTransfer)
	outTransfer := make(chan *chains.TokenTransfer)
	e.onIndex = true

	// TODO: can be muilti goroutine
	go e.IndexTasks(ctx, taskBlockNum)
	go e.IndexBlock(ctx, taskBlockNum, indexBlockNum)
	go e.IndexTransfer(ctx, indexBlockNum, outTransfers)
	go e.IndexToken(ctx, outTransfers, outTransfer)
	go e.IndexContract(ctx, outTransfer, FindContractCreator)
}

func (e *EthIndexer) UpdateEndpoints(endpoints []string) {
	e.Endpoints = endpoints
}

func (e *EthIndexer) IsOnIndex() bool {
	return e.onIndex
}

func (e *EthIndexer) StopIndex() {
	if e.cancel != nil {
		logger.Sugar().Infof("stop the indexer chainType: %v, chainID: %v", e.ChainType, e.ChainID)
		(*e.cancel)()
		e.cancel = nil
		e.BadEndpoints = nil
		e.Endpoints = nil
		e.onIndex = false
	}
}

func (e *EthIndexer) IndexTasks(ctx context.Context, outBlockNum chan uint64) {
	logger.Sugar().Info("start to index task for ethereum")
	conds := &synctask.Conds{
		ChainType: &ctMessage.StringVal{
			Value: e.ChainType.String(),
			Op:    "eq",
		},
		ChainID: &ctMessage.StringVal{
			Value: e.ChainID,
			Op:    "eq",
		},
		SyncState: &ctMessage.StringVal{
			Value: cttype.SyncState_Start.String(),
			Op:    "eq",
		},
	}
	for {
		select {
		case <-time.NewTicker(CheckTopicInterval).C:
			resp, err := synctaskNMCli.GetSyncTasks(ctx, &synctask.GetSyncTasksRequest{Conds: conds, Offset: 0, Limit: maxTopicNum})
			if err != nil {
				logger.Sugar().Error(err)
			}
			for _, v := range resp.GetInfos() {
				resp, err := synctaskNMCli.TriggerSyncTask(ctx, &synctask.TriggerSyncTaskRequest{Topic: v.Topic, CurrentBlockNum: e.CurrentBlockNum})
				if err != nil {
					logger.Sugar().Errorf("triggerSyncTask failed ,err: %v", err)
				}

				for _, v := range resp.BlockNums {
					outBlockNum <- v
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (e *EthIndexer) IndexBlock(ctx context.Context, inBlockNum, outBlockNum chan uint64) {
	for {
		select {
		case blockNum := <-inBlockNum:
			cli, err := eth.Client(e.Endpoints)
			if err != nil {
				logger.Sugar().Errorf("cannot get eth client,err: %v", err)
				continue
			}
			block, err := cli.BlockByNumber(ctx, blockNum)
			if err != nil {
				e.checkErr(ctx, err)
				logger.Sugar().Errorf("cannot get eth client,err: %v", err)
				continue
			}
			number := block.Number().Uint64()
			blockHash := block.Hash().String()
			blockTime := int64(block.Time())
			_, err = blockNMCli.UpsertBlock(ctx, &blockProto.UpsertBlockRequest{
				Info: &blockProto.BlockReq{
					ChainType:   &e.ChainType,
					ChainID:     &e.ChainID,
					BlockNumber: &number,
					BlockHash:   &blockHash,
					BlockTime:   &blockTime,
				},
			})
			if err != nil {
				logger.Sugar().Errorf("cannot get eth client,err: %v", err)
				continue
			}

			outBlockNum <- block.Number().Uint64()
		case <-ctx.Done():
			return
		}
	}
}

func (e *EthIndexer) IndexTransfer(ctx context.Context, inBlockNum chan uint64, outTransfers chan []*chains.TokenTransfer) {
	for {
		select {
		case num := <-inBlockNum:
			cli, err := eth.Client(e.Endpoints)
			if err != nil {
				logger.Sugar().Errorf("cannot get eth client,err: %v", err)
				continue
			}
			transfers, err := cli.TransferLogs(ctx, int64(num), int64(num))
			if err != nil {
				e.checkErr(ctx, err)
				logger.Sugar().Errorf("failed to get transfer logs, err: %v, block: %v", err, num)
				continue
			}
			if len(transfers) == 0 {
				continue
			}

			infos := make([]*transferProto.TransferReq, len(transfers))
			for i := range transfers {
				tokenType := string(transfers[i].TokenType)
				infos[i] = &transferProto.TransferReq{
					ChainType:   &e.ChainType,
					ChainID:     &e.ChainID,
					Contract:    &transfers[i].Contract,
					TokenType:   &tokenType,
					TokenID:     &transfers[i].TokenID,
					From:        &transfers[i].From,
					To:          &transfers[i].To,
					Amount:      &transfers[i].Amount,
					BlockNumber: &transfers[i].BlockNumber,
					TxHash:      &transfers[i].TxHash,
					BlockHash:   &transfers[i].BlockHash,
				}
			}

			_, err = transferNMCli.UpsertTransfers(ctx, &transferProto.UpsertTransfersRequest{Infos: infos})
			if err != nil {
				logger.Sugar().Errorf("failed store transfers to db for block number: %v,err: %v", num, err)
				continue
			}

			outTransfers <- transfers
		case <-ctx.Done():
			return
		}
	}
}

func (e *EthIndexer) IndexToken(ctx context.Context, inTransfers chan []*chains.TokenTransfer, outTransfer chan *chains.TokenTransfer) {
	for {
		select {
		case transfers := <-inTransfers:
			for _, transfer := range transfers {
				identifier := tokenIdentifier(e.ChainType, e.ChainID, transfer.Contract, transfer.TokenID)
				locked, err := ctredis.TryPubLock(identifier, redisExpireDefaultTime)
				if err != nil {
					logger.Sugar().Errorf("lock the token indentifier failed, err: %v", err)
					continue
				}

				if !locked {
					continue
				}

				remark := ""
				conds := &tokenProto.Conds{
					ChainType: &ctMessage.StringVal{
						Value: e.ChainType.String(),
						Op:    "eq",
					},
					ChainID: &ctMessage.StringVal{
						Value: e.ChainID,
						Op:    "eq",
					},
					Contract: &ctMessage.StringVal{
						Value: transfer.Contract,
						Op:    "eq",
					},
					TokenID: &ctMessage.StringVal{
						Value: transfer.TokenID,
						Op:    "eq",
					},
				}

				if resp, err := tokenNMCli.ExistTokenConds(ctx, &tokenProto.ExistTokenCondsRequest{Conds: conds}); err == nil && resp != nil && resp.GetExist() {
					continue
				} else if err != nil {
					logger.Sugar().Errorf("check if the token exist failed, err: %v", err)
					continue
				}

				cli, err := eth.Client(e.Endpoints)
				if err != nil {
					logger.Sugar().Errorf("cannot get eth client,err: %v", err)
					continue
				}

				tokenURI, err := cli.TokenURI(ctx, transfer.TokenType, transfer.Contract, transfer.TokenID, transfer.BlockNumber)
				if err != nil {
					e.checkErr(ctx, err)
					logger.Sugar().Warnf("cannot get tokenURI,err: %v", err)
					remark = err.Error()
				}

				tokenURIInfo, err := token.GetTokenURIInfo(ctx, tokenURI)
				if err != nil {
					tokenURIInfo = &token.TokenURIInfo{}
				}

				_, err = tokenNMCli.CreateToken(ctx, &tokenProto.CreateTokenRequest{
					Info: &tokenProto.TokenReq{
						ChainType:   &e.ChainType,
						ChainID:     &e.ChainID,
						Contract:    &transfer.Contract,
						TokenType:   &transfer.TokenType,
						TokenID:     &transfer.TokenID,
						URI:         &tokenURI,
						URIType:     (*string)(&tokenURIInfo.URIType),
						ImageURL:    &tokenURIInfo.ImageURL,
						VideoURL:    &tokenURIInfo.VideoURL,
						Name:        &tokenURIInfo.Name,
						Description: &tokenURIInfo.Description,
						VectorState: tokenProto.ConvertState_Waiting.Enum(),
						Remark:      &remark,
					},
				})

				if err != nil {
					logger.Sugar().Errorf("create token record failed, %v", err)
					continue
				}
				outTransfer <- transfer
			}
		case <-ctx.Done():
			return
		}
	}
}

func (e *EthIndexer) IndexContract(ctx context.Context, inTransfer chan *chains.TokenTransfer, findContractCreator bool) {
	for {
		select {
		case transfer := <-inTransfer:
			identifier := contractIdentifier(e.ChainType, e.ChainID, transfer.Contract)
			locked, err := ctredis.TryPubLock(identifier, redisExpireDefaultTime)
			if err != nil {
				logger.Sugar().Errorf("lock the token indentifier failed, err: %v", err)
				continue
			}

			if !locked {
				continue
			}

			conds := &contractProto.Conds{
				ChainType: &ctMessage.StringVal{
					Value: e.ChainType.String(),
					Op:    "eq",
				},
				ChainID: &ctMessage.StringVal{
					Value: e.ChainID,
					Op:    "eq",
				},
				Address: &ctMessage.StringVal{
					Value: transfer.Contract,
					Op:    "eq",
				},
			}

			if resp, err := contractNMCli.ExistContractConds(ctx, &contractProto.ExistContractCondsRequest{
				Conds: conds,
			}); err == nil && resp != nil && resp.GetExist() {
				continue
			} else if err != nil {
				logger.Sugar().Errorf("check if the contract exist failed, err: %v", err)
				continue
			}

			remark := ""
			cli, err := eth.Client(e.Endpoints)
			if err != nil {
				logger.Sugar().Errorf("cannot get eth client,err: %v", err)
				continue
			}

			contractMeta, err := cli.GetERC721Metadata(ctx, transfer.Contract)
			if err != nil {
				e.checkErr(ctx, err)
				logger.Sugar().Warnf("transfer cannot get ,err: %v", err)
				contractMeta = &eth.ERC721Metadata{}
				remark = fmt.Sprintf("%v,%v", remark, err)
			}

			creator := &eth.ContractCreator{}
			// stop get info for creator
			if findContractCreator {
				creator, err = cli.GetContractCreator(ctx, transfer.Contract)
				if err != nil {
					e.checkErr(ctx, err)
					remark = err.Error()
				}
			}

			from := creator.From.String()
			txHash := creator.TxHash.Hex()
			blockNum := creator.BlockNumber
			txTime := uint32(creator.TxTime)
			_, err = contractNMCli.CreateContract(ctx, &contractProto.CreateContractRequest{
				Info: &contractProto.ContractReq{
					ChainType: &e.ChainType,
					ChainID:   &e.ChainID,
					Address:   &transfer.Contract,
					Name:      &contractMeta.Name,
					Symbol:    &contractMeta.Symbol,
					Creator:   &from,
					BlockNum:  &blockNum,
					TxHash:    &txHash,
					TxTime:    &txTime,
					Remark:    &remark,
				},
			})
			if err != nil {
				logger.Sugar().Errorf("create contract record failed, %v", err)
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}

func (e *EthIndexer) GetCurrentBlockNum(ctx context.Context, updateInterval time.Duration) {
	for {
		func() {
			cli, err := eth.Client(e.Endpoints)
			if err != nil {
				logger.Sugar().Errorf("cannot get eth client,err: %v", err)
				return
			}

			blockNum, err := cli.CurrentBlockNum(ctx)
			if err != nil {
				e.checkErr(ctx, err)
				logger.Sugar().Errorf("failed to get current block number: %v", err)
				return
			}

			e.CurrentBlockNum = blockNum
			logger.Sugar().Infof("success get current block number: %v", blockNum)
		}()

		select {
		case <-time.NewTicker(updateInterval).C:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func tokenIdentifier(chain basetype.ChainType, chainID, contract, tokenID string) string {
	return fmt.Sprintf("%v:%v:%v:%v", chain, chainID, contract, tokenID)
}

func contractIdentifier(chain basetype.ChainType, chainID, contract string) string {
	return fmt.Sprintf("%v+%v+%v", chain, chainID, contract)
}
