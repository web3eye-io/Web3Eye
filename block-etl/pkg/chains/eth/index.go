package eth

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/cyber-tracer/block-etl/pkg/token"
	"github.com/web3eye-io/cyber-tracer/common/ctredis"
	contractNMCli "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/client/v1/contract"
	tokenNMCli "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/client/v1/token"
	transferNMCli "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/client/v1/transfer"
	ctMessage "github.com/web3eye-io/cyber-tracer/proto/cybertracer"
	contractProto "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/contract"
	tokenProto "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/token"
	transferProto "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/transfer"
)

const (
	Retries                = 3
	confirmedBlockNum      = 5
	redisExpireDefaultTime = time.Minute * 5
	maxTokenURILen         = 256
)

var (
	connectTimeout = "context deadline exceeded"
	retryErrs      = []string{connectTimeout}
)

var (
	// manage process order
	pLock = &sync.Mutex{}
	// manager refresh current block height
	gLock = &sync.Mutex{}
	wg    = &sync.WaitGroup{}
)

type EthIndexer struct {
	// If MaxDealBlockNum is 5 and MaxLogTaskSizeInOneBlock is 10
	// ,than num of total goroutines is 5*10
	MaxDealBlockNum          int
	MaxLogTaskSizeInOneBlock int
	LogIntervalHeight        uint16
	fromBlock                uint64
	toBlock                  uint64
	finishBlocks             sortedBlockAim
	blockAim                 map[uint64]*processAim
	currentHeight            uint64
	updateTime               int64
}

type sortedBlockAim []uint64

type processAim struct {
	transfer bool
	token    bool
	contract bool
}

func NewIndexer(fromBlock uint64) (*EthIndexer, error) {
	return &EthIndexer{
		MaxDealBlockNum:          5,
		MaxLogTaskSizeInOneBlock: 10,
		LogIntervalHeight:        10,
		fromBlock:                fromBlock,
		toBlock:                  fromBlock,
		finishBlocks:             []uint64{},
		blockAim:                 make(map[uint64]*processAim),
		currentHeight:            0,
		updateTime:               0,
	}, nil
}

func (e *EthIndexer) IndexTransfer(ctx context.Context, handler func([]*TokenTransfer) error) {
	blockNum := make(chan uint64, e.MaxDealBlockNum)
	defer close(blockNum)

	for i := 0; i < e.MaxDealBlockNum; i++ {
		wg.Add(1)
		go func() {
			for num := range blockNum {
				transfers, err := TransferLogs(ctx, int64(num), int64(num))
				if err != nil && containErr(err.Error()) {
					logger.Sugar().Errorf("will retry anlysis height %v for parsing transfer logs failed, %v", num, err)
					blockNum <- num
				}

				if err != nil {
					logger.Sugar().Errorf("parse transfer logs failed, %v", err)
				}

				if len(transfers) == 0 {
					e.finishAll(num)
					continue
				}

				err = handler(transfers)
				if err != nil {
					logger.Sugar().Errorf("handle transfers failed, %v", err)
					continue
				}
				e.finishAll(num)
			}
			wg.Done()
		}()
	}

	logger.Sugar().Infof("start sync history ethereum log,from %v height", e.fromBlock)
	// fast catch up
	for {
		num := e.getNextBlock()
		blockNum <- num
		if int(num) >= int(e.getCurrentConfirmedHeight(ctx)) {
			break
		}
	}

	logger.Sugar().Infof("start sync recent ethereum log,from %v height", e.fromBlock)
	// follow the height
	for {
		time.Sleep(time.Minute)

		for {
			num := e.getNextBlock()
			blockNum <- num
			if num >= e.getCurrentConfirmedHeight(ctx) {
				break
			}
		}
	}
	// TODO：should take signal of system,than close goroutine
}

func (e *EthIndexer) IndexTransferToDB(ctx context.Context) {
	e.IndexTransfer(ctx, func(transfers []*TokenTransfer) error {
		transferErr := e.transferToDB(ctx, transfers)
		if transferErr != nil {
			return transferErr
		}
		tokenErr := e.tokenInfoToDB(ctx, transfers)
		if tokenErr != nil {
			return tokenErr
		}
		return nil
	})
}

func (e *EthIndexer) transferToDB(ctx context.Context, transfers []*TokenTransfer) error {
	tt := make([]*transferProto.TransferReq, len(transfers))
	for i := range transfers {
		chainType := string(transfers[i].ChainType)
		tokenType := string(transfers[i].TokenType)
		tt[i] = &transferProto.TransferReq{
			ChainType:   &chainType,
			ChainID:     &transfers[i].ChainID,
			Contract:    &transfers[i].Contract,
			TokenType:   &tokenType,
			TokenID:     &transfers[i].TokenID,
			From:        &transfers[i].From,
			To:          &transfers[i].To,
			Amount:      &transfers[i].Amount,
			BlockNumber: &transfers[i].BlockNumber,
			TxHash:      &transfers[i].TxHash,
			BlockHash:   &transfers[i].BlockHash,
			TxTime:      &transfers[i].TxTime,
		}
	}

	for i := 0; i < Retries; i++ {
		_, err := transferNMCli.CreateTransfers(ctx, tt)
		if err != nil && containErr(err.Error()) {
			logger.Sugar().Errorf("will retry for creating transfer record failed, %v", err)
			continue
		}

		if err != nil {
			return fmt.Errorf("create transfer record failed, %v", err)
		}
		break
	}
	return nil
}

//nolint:gocyclo
func (e *EthIndexer) tokenInfoToDB(ctx context.Context, transfers []*TokenTransfer) error {
	for _, transfer := range transfers {
		identifier := tokenIdentifier(transfer.ChainType, transfer.ChainID, transfer.Contract, transfer.TokenID)
		if _, err := ctredis.Get(identifier); err == nil {
			return nil
		}

		tokenType := string(transfer.TokenType)
		remark := ""
		conds := &tokenProto.Conds{
			ChainType: &ctMessage.StringVal{
				Value: string(transfer.ChainType),
				Op:    "eq",
			},
			ChainID: &ctMessage.Int32Val{
				Value: transfer.ChainID,
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

		if exist, err := tokenNMCli.ExistTokenConds(ctx, conds); exist && err == nil {
			err = ctredis.Set(identifier, true, redisExpireDefaultTime)
			if err != nil {
				logger.Sugar().Error(err)
			}
			return nil
		}

		err := ctredis.Set(identifier, false, redisExpireDefaultTime)
		if err != nil {
			logger.Sugar().Error(err)
			err = ctredis.Del(identifier)
			if err != nil {
				logger.Sugar().Error(err)
			}
			return nil
		}

		err = e.contractToDB(ctx, transfer)
		if err != nil {
			logger.Sugar().Error(err)
		}

		tokenURI, err := TokenURI(ctx, transfer.TokenType, transfer.Contract, transfer.TokenID, transfer.BlockNumber)
		if err != nil {
			remark = err.Error()
		}

		tokenURIInfo, err := token.GetTokenURIInfo(ctx, tokenURI)
		if err != nil {
			tokenURIInfo = &token.TokenURIInfo{}
		}

		if len(tokenURI) > maxTokenURILen {
			tokenURI = ""
		}

		for i := 0; i < Retries; i++ {
			_, err = tokenNMCli.CreateToken(ctx, &tokenProto.TokenReq{
				ChainType:   (*string)(&transfer.ChainType),
				ChainID:     &transfer.ChainID,
				Contract:    &transfer.Contract,
				TokenType:   &tokenType,
				TokenID:     &transfer.TokenID,
				URI:         &tokenURI,
				URIType:     (*string)(&tokenURIInfo.URIType),
				ImageURL:    &tokenURIInfo.ImageURL,
				VideoURL:    &tokenURIInfo.VideoURL,
				Name:        &tokenURIInfo.Name,
				Description: &tokenURIInfo.Description,
				VectorState: tokenProto.ConvertState_Waiting.Enum(),
				Remark:      &remark,
			})
			if err != nil && containErr(err.Error()) {
				logger.Sugar().Errorf("will retry for creating token record failed, %v", err)
				continue
			}

			if err != nil {
				return fmt.Errorf("create token record failed, %v", err)
			}
			break
		}
		return nil
	}
	return nil
}

func (e *EthIndexer) contractToDB(ctx context.Context, transfer *TokenTransfer) error {
	identifier := contractIdentifier(transfer.ChainType, transfer.ChainID, transfer.Contract)
	if _, err := ctredis.Get(identifier); err == nil {
		return nil
	}

	conds := &contractProto.Conds{
		ChainType: &ctMessage.StringVal{
			Value: string(transfer.ChainType),
			Op:    "eq",
		},
		ChainID: &ctMessage.Int32Val{
			Value: transfer.ChainID,
			Op:    "eq",
		},
		Address: &ctMessage.StringVal{
			Value: transfer.Contract,
			Op:    "eq",
		},
	}

	if exist, err := contractNMCli.ExistContractConds(ctx, conds); exist && err == nil {
		err = ctredis.Set(identifier, true, redisExpireDefaultTime)
		if err != nil {
			logger.Sugar().Error(err)
		}
		return nil
	}
	err := ctredis.Set(identifier, false, redisExpireDefaultTime)
	if err != nil {
		logger.Sugar().Error(err)
		err = ctredis.Del(identifier)
		if err != nil {
			logger.Sugar().Error(err)
		}
		return nil
	}

	remark := ""
	creator, err := GetContractCreator(ctx, transfer.Contract)
	if err != nil {
		creator = &ContractCreator{}
		remark = err.Error()
	}

	contractMeta, err := GetERC721Metadata(ctx, transfer.Contract)
	if err != nil {
		contractMeta = &ERC721Metadata{}
		remark = fmt.Sprintf("%v,%v", remark, err)
	}

	from := creator.From.String()
	txHash := creator.TxHash.Hex()
	blockNum := creator.BlockNumber
	txTime := uint32(creator.TxTime)
	for i := 0; i < Retries; i++ {
		_, err = contractNMCli.CreateContract(ctx, &contractProto.ContractReq{
			ChainType: (*string)(&transfer.ChainType),
			ChainID:   &transfer.ChainID,
			Address:   &transfer.Contract,
			Name:      &contractMeta.Name,
			Symbol:    &contractMeta.Symbol,
			Creator:   &from,
			BlockNum:  &blockNum,
			TxHash:    &txHash,
			TxTime:    &txTime,
			Remark:    &remark,
		})
		if err != nil && containErr(err.Error()) {
			logger.Sugar().Errorf("will retry for creating token record failed, %v", err)
			continue
		}

		if err != nil {
			return fmt.Errorf("create token record failed, %v", err)
		}
		break
	}
	return nil
}

func (e *EthIndexer) getNextBlock() uint64 {
	pLock.Lock()
	defer pLock.Unlock()

	sort.Sort(e.finishBlocks)
	i := 0
	for ; i < e.finishBlocks.Len(); i++ {
		if e.fromBlock+uint64(i) == e.finishBlocks[i] {
			continue
		}
		break
	}

	e.finishBlocks = e.finishBlocks[i:]
	e.fromBlock += uint64(i)

	num := e.toBlock
	e.toBlock += 1
	e.blockAim[num] = &processAim{}
	return num
}

// func (e *EthIndexer) finishTransfer(blockNum uint64) {
// 	e.blockAim[blockNum].transfer = true
// 	e._finishAim(blockNum)
// }

// func (e *EthIndexer) finishToken(blockNum uint64) {
// 	e.blockAim[blockNum].token = true
// 	e._finishAim(blockNum)
// }

// func (e *EthIndexer) finishContract(blockNum uint64) {
// 	e.blockAim[blockNum].contract = true
// 	e._finishAim(blockNum)
// }

func (e *EthIndexer) finishAll(blockNum uint64) {
	e.blockAim[blockNum].contract = true
	e.blockAim[blockNum].token = true
	e.blockAim[blockNum].transfer = true
	e._finishAim(blockNum)
}

// not allow to be called by func with upper case letter
func (e *EthIndexer) _finishAim(blockNum uint64) {
	if e.blockAim[blockNum].token && e.blockAim[blockNum].transfer && e.blockAim[blockNum].contract {
		e._finishBlock(blockNum)
	}
}

// not allow to be called by func with upper case letter
func (e *EthIndexer) _finishBlock(blockNum uint64) {
	pLock.Lock()
	defer pLock.Unlock()

	if blockNum < e.fromBlock || blockNum > e.toBlock {
		return
	}
	for _, v := range e.finishBlocks {
		if blockNum == v {
			return
		}
	}

	e.finishBlocks = append(e.finishBlocks, blockNum)
	if blockNum%uint64(e.LogIntervalHeight) == 0 {
		logger.Sugar().Infof("sync block height pass by %v", blockNum)
	}
}

func (e *EthIndexer) getCurrentConfirmedHeight(ctx context.Context) uint64 {
	gLock.Lock()
	defer gLock.Unlock()

	if e.updateTime > time.Now().Unix() {
		return e.currentHeight
	}

	num, err := CurrentBlockHeight(ctx)
	if err != nil {
		logger.Sugar().Errorf("get block height failed, %v", err)
		return e.currentHeight
	}

	if num > confirmedBlockNum {
		e.currentHeight = num - confirmedBlockNum
	}

	e.updateTime = time.Now().Add(time.Minute).Unix()

	return num
}

func (a sortedBlockAim) Len() int {
	return len(a)
}

func (a sortedBlockAim) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a sortedBlockAim) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func containErr(errStr string) bool {
	for _, v := range retryErrs {
		if strings.Contains(errStr, v) {
			return true
		}
	}
	return false
}

func tokenIdentifier(chain ChainType, chainID int32, contract, tokenID string) string {
	return fmt.Sprintf("%v:%v:%v:%v", chain, chainID, contract, tokenID)
}

func contractIdentifier(chain ChainType, chainID int32, contract string) string {
	return fmt.Sprintf("%v+%v+%v", chain, chainID, contract)
}
