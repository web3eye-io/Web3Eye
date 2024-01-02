package sol

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/portto/solana-go-sdk/program/metaplex/token_metadata"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/indexer"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/token"
	"github.com/web3eye-io/Web3Eye/common/chains"
	"github.com/web3eye-io/Web3Eye/common/chains/sol"
	"github.com/web3eye-io/Web3Eye/common/ctredis"
	blockNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/block"
	contractNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
	tokenNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/token"
	transferNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/transfer"
	ctMessage "github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	blockProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
	contractProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
	tokenProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	transferProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func (e *SolIndexer) CheckBlock(ctx context.Context, inBlockNum uint64) (*blockProto.Block, error) {
	// ignore err,just adjust if blockOnly is finished
	blockOnly, _ := blockNMCli.GetBlockOnly(ctx, &blockProto.GetBlockOnlyRequest{
		Conds: &blockProto.Conds{
			ChainType:   &ctMessage.Uint32Val{Op: "eq", Value: uint32(e.ChainType)},
			ChainID:     &ctMessage.StringVal{Op: "eq", Value: e.ChainID},
			BlockNumber: &ctMessage.Uint64Val{Op: "eq", Value: inBlockNum},
		},
	})

	if blockOnly != nil &&
		blockOnly.Info != nil &&
		blockOnly.GetInfo().ParseState == basetype.BlockParseState_BlockTypeFinish {
		return blockOnly.GetInfo(), nil
	}

	cli, err := sol.Client(e.OkEndpoints)
	if err != nil {
		return nil, fmt.Errorf("cannot get sol client,err: %v", err)
	}

	block, err := cli.GetBlock(ctx, inBlockNum)
	if err != nil {
		e.checkErr(ctx, err)
		return nil, fmt.Errorf("cannot get sol block,err: %v", err)
	}

	blockHash := block.Blockhash.String()
	blockTime := uint64(block.BlockTime.Time().Unix())
	remark := ""
	resp, err := blockNMCli.UpsertBlock(ctx, &blockProto.UpsertBlockRequest{
		Info: &blockProto.BlockReq{
			ChainType: &e.ChainType,
			ChainID:   &e.ChainID,
			// replace block height with slot height
			BlockNumber: &inBlockNum,
			BlockHash:   &blockHash,
			BlockTime:   &blockTime,
			ParseState:  basetype.BlockParseState_BlockTypeStart.Enum(),
			Remark:      &remark,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("cannot get sol client,err: %v", err)
	}
	return resp.Info, nil
}

func (e *SolIndexer) IndexTransfer(ctx context.Context, inBlockNum uint64) ([]*chains.TokenTransfer, error) {
	cli, err := sol.Client(e.OkEndpoints)
	if err != nil {
		return nil, fmt.Errorf("cannot get sol client,err: %v", err)
	}
	block, err := cli.GetBlock(ctx, inBlockNum)
	if err != nil {
		e.checkErr(ctx, err)
		return nil, fmt.Errorf("cannot get sol block,err: %v", err)
	}
	txTime := uint64(block.BlockTime.Time().Unix())

	transfers := sol.GetNFTTransfers(block)
	if len(transfers) == 0 {
		return nil, nil
	}

	infos := make([]*transferProto.TransferReq, len(transfers))
	for i := range transfers {
		metadata, err := cli.GetMetadata(ctx, transfers[i].TokenID)
		if err != nil {
			e.checkErr(ctx, err)
			return nil, fmt.Errorf("cannot get sol token metadata,err: %v", err)
		}
		transfers[i].Contract = GenCollectionAddr(metadata)

		infos[i] = &transferProto.TransferReq{
			ChainType:   &e.ChainType,
			ChainID:     &e.ChainID,
			Contract:    &transfers[i].Contract,
			TokenType:   &transfers[i].TokenType,
			TokenID:     &transfers[i].TokenID,
			From:        &transfers[i].From,
			To:          &transfers[i].To,
			Amount:      &transfers[i].Amount,
			BlockNumber: &transfers[i].BlockNumber,
			TxHash:      &transfers[i].TxHash,
			TxTime:      &txTime,
			BlockHash:   &transfers[i].BlockHash,
		}
	}

	_, err = transferNMCli.UpsertTransfers(ctx, &transferProto.UpsertTransfersRequest{Infos: infos})
	if err != nil {
		return nil, fmt.Errorf("failed store transfers to db for block number: %v,err: %v", inBlockNum, err)
	}

	return transfers, nil
}

func (e *SolIndexer) IndexToken(ctx context.Context, inTransfers []*chains.TokenTransfer) ([]*chains.TokenTransfer, error) {
	outTransfers := []*chains.TokenTransfer{}
	for _, transfer := range inTransfers {
		identifier := indexer.TokenIdentifier(e.ChainType, e.ChainID, transfer.Contract, transfer.TokenID)
		locked, err := ctredis.TryPubLock(identifier, redisExpireDefaultTime)
		if err != nil {
			return nil, fmt.Errorf("lock the token indentifier failed, err: %v", err)
		}

		if !locked {
			continue
		}

		conds := &tokenProto.Conds{
			ChainType: &ctMessage.Uint32Val{
				Value: uint32(e.ChainType),
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
			return outTransfers, fmt.Errorf("check if the token exist failed, err: %v", err)
		}

		cli, err := sol.Client(e.OkEndpoints)
		if err != nil {
			return outTransfers, fmt.Errorf("cannot get sol client,err: %v", err)
		}
		remark := ""

		uriState := basetype.TokenURIState_TokenURIFinish
		metadata, err := cli.GetMetadata(ctx, transfer.TokenID)
		if err != nil {
			uriState = basetype.TokenURIState_TokenURIError
			e.checkErr(ctx, err)
			logger.Sugar().Warnf("cannot get metadata,err: %v, tokenID: %v", err, transfer.TokenID)
			remark = fmt.Sprintf("%v,%v", remark, err)
		}

		tokenURIInfo := &token.TokenURIInfo{}
		if metadata != nil {
			var complete bool
			tokenURIInfo, complete, err = token.GetTokenURIInfo(ctx, metadata.Data.Uri)
			if err != nil {
				uriState = basetype.TokenURIState_TokenURIError
				tokenURIInfo = &token.TokenURIInfo{}
				remark = fmt.Sprintf("%v,%v", remark, err)
			} else if !complete {
				uriState = basetype.TokenURIState_TokenURIIncomplete
			}
		} else {
			uriState = basetype.TokenURIState_TokenURIError
			// if cannot get metadata,then set the default value
			metadata = &token_metadata.Metadata{}
		}

		if len(metadata.Data.Uri) > indexer.MaxTokenURILength {
			remark = fmt.Sprintf("%v,tokenURI too long(length: %v),skip to store it", remark, len(metadata.Data.Uri))
			metadata.Data.Uri = metadata.Data.Uri[:indexer.OverLimitStoreLength]
		}

		_, err = tokenNMCli.UpsertToken(ctx, &tokenProto.UpsertTokenRequest{
			Info: &tokenProto.TokenReq{
				ChainType:   &e.ChainType,
				ChainID:     &e.ChainID,
				Contract:    &transfer.Contract,
				TokenType:   &transfer.TokenType,
				TokenID:     &transfer.TokenID,
				URI:         &metadata.Data.Uri,
				URIState:    &uriState,
				URIType:     (*string)(&tokenURIInfo.URIType),
				ImageURL:    &tokenURIInfo.ImageURL,
				VideoURL:    &tokenURIInfo.VideoURL,
				Name:        &metadata.Data.Name,
				Description: &metadata.Data.Symbol,
				VectorState: tokenProto.ConvertState_Waiting.Enum(),
				Remark:      &remark,
			},
		})

		if err != nil {
			return outTransfers, fmt.Errorf("create token record failed, %v", err)
		}

		outTransfers = append(outTransfers, transfer)
	}
	return outTransfers, nil
}

func (e *SolIndexer) IndexContract(ctx context.Context, inTransfers []*chains.TokenTransfer, findContractCreator bool) error {
	for _, transfer := range inTransfers {
		exist, err := e.checkContract(ctx, transfer)
		if err != nil {
			return err
		}
		if exist {
			continue
		}

		contractMeta, creator, remark := e.getContractInfo(ctx, transfer)

		from := creator.From
		txHash := creator.TxHash
		blockNum := creator.BlockNumber
		decimal := uint32(0)
		txTime := uint32(creator.TxTime)
		_, err = contractNMCli.UpsertContract(ctx, &contractProto.UpsertContractRequest{
			Info: &contractProto.ContractReq{
				ChainType: &e.ChainType,
				ChainID:   &e.ChainID,
				Address:   &transfer.Contract,
				Name:      &contractMeta.Data.Name,
				Symbol:    &contractMeta.Data.Symbol,
				Decimals:  &decimal,
				Creator:   &from,
				BlockNum:  &blockNum,
				TxHash:    &txHash,
				TxTime:    &txTime,
				Remark:    &remark,
			},
		})
		if err != nil {
			return fmt.Errorf("create contract record failed, %v", err)
		}
	}
	return nil
}

func (e *SolIndexer) checkContract(ctx context.Context, transfer *chains.TokenTransfer) (exist bool, err error) {
	identifier := indexer.ContractIdentifier(e.ChainType, e.ChainID, transfer.Contract)
	locked, err := ctredis.TryPubLock(identifier, redisExpireDefaultTime)
	if err != nil {
		return false, fmt.Errorf("lock the token indentifier failed, err: %v", err)
	}

	if !locked {
		return true, nil
	}

	conds := &contractProto.Conds{
		ChainType: &ctMessage.Uint32Val{
			Op:    "eq",
			Value: uint32(e.ChainType),
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

	if resp, err := contractNMCli.ExistContractConds(ctx, &contractProto.ExistContractCondsRequest{Conds: conds}); err == nil && resp != nil && resp.GetExist() {
		return true, nil
	} else if err != nil {
		return false, fmt.Errorf("check if the contract exist failed, err: %v", err)
	}
	return false, nil
}

func (e *SolIndexer) getContractInfo(ctx context.Context, transfer *chains.TokenTransfer) (*token_metadata.Metadata, *chains.ContractCreator, string) {
	remark := ""
	contractMeta := &token_metadata.Metadata{}

	// TODO: support find creator
	creator := &chains.ContractCreator{}

	cli, err := sol.Client(e.OkEndpoints)
	if err != nil {
		return contractMeta, creator, fmt.Sprintf("cannot get sol client,err: %v", err)
	}

	if IsW3EAddress(transfer.Contract) {
		contractMeta, err = cli.GetMetadata(ctx, transfer.TokenID)
	} else {
		contractMeta, err = cli.GetMetadata(ctx, transfer.Contract)
	}

	if err != nil {
		e.checkErr(ctx, err)
		remark = err.Error()
	}

	if contractMeta != nil && contractMeta.Data.Creators != nil && len(*contractMeta.Data.Creators) > 0 {
		creator.From = (*contractMeta.Data.Creators)[0].Address.String()
	}

	return contractMeta, creator, remark
}

func (e *SolIndexer) GetCurrentBlockNum() uint64 {
	return e.CurrentBlockNum
}

func (e *SolIndexer) SyncCurrentBlockNum(ctx context.Context, updateInterval time.Duration) {
	for {
		func() {
			cli, err := sol.Client(e.OkEndpoints)
			if err != nil {
				logger.Sugar().Errorf("sol cannot get sol client,err: %v", err)
				return
			}

			blockNum, err := cli.GetSlotHeight(ctx)
			if err != nil {
				e.checkErr(ctx, err)
				logger.Sugar().Errorf("sol failed to get current block number: %v", err)
				return
			}

			e.CurrentBlockNum = blockNum
			logger.Sugar().Infof("sol success get current block number: %v", blockNum)
		}()

		select {
		case <-time.NewTicker(updateInterval).C:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func IsW3EAddress(addr string) bool {
	return strings.HasPrefix(addr, W3EAddressPrefix)
}

// if have no collection info ,we will build it for token
func GenCollectionAddr(info *token_metadata.Metadata) string {
	if info.Collection != nil {
		return info.Collection.Key.String()
	}
	h224 := sha256.New224()
	h224.Write([]byte(info.Data.Name))
	h224.Write([]byte(info.Data.Symbol))
	if info.Data.Creators != nil && len(*info.Data.Creators) > 0 {
		h224.Write([]byte((*info.Data.Creators)[0].Address.String()))
	}
	return fmt.Sprintf("%v%x", W3EAddressPrefix, h224.Sum(nil))
}
