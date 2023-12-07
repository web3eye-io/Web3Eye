package block

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/stretchr/testify/assert"

	blockproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &blockproto.Block{
	ChainType:   basetype.ChainType_Ethereum,
	ChainID:     "test_block",
	BlockNumber: 10010,
	BlockHash:   "test_block",
	BlockTime:   time.Now().Unix(),
	ParseState:  basetype.BlockParseState_BlockTypeFinish,
	Remark:      "test_block",
}

var req = &blockproto.BlockReq{
	ChainType:   &ret.ChainType,
	ChainID:     &ret.ChainID,
	BlockNumber: &ret.BlockNumber,
	BlockHash:   &ret.BlockHash,
	BlockTime:   &ret.BlockTime,
	ParseState:  &ret.ParseState,
	Remark:      &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithBlockNumber(req.BlockNumber, true),
		WithBlockHash(req.BlockHash, true),
		WithBlockTime(req.BlockTime, true),
		WithParseState(req.ParseState, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateBlock(context.Background())
	if assert.Nil(t, err) {
		req.ID = &info.ID
		req.EntID = &info.EntID
		assert.Equal(t, info.ParseStateStr, req.ParseState.String())
	}
}

func update(t *testing.T) {
	req.ChainType = basetype.ChainType_Solana.Enum()
	req.ParseState = basetype.BlockParseState_BlockTypeStart.Enum()

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithChainType(req.ChainType, false),
		WithChainID(req.ChainID, false),
		WithParseState(req.ParseState, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateBlock(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
	}

	req.ParseState = basetype.BlockParseState_BlockTypeFailed.Enum()

	handler, err = NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithChainID(req.ChainID, false),
		WithParseState(req.ParseState, false),
	)
	assert.Nil(t, err)

	_, err = handler.UpdateBlock(context.Background())
	assert.Nil(t, err)
}

func query(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetBlock(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&blockproto.Conds{
			ChainType: &web3eye.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(*req.ChainType),
			},
			ChainID: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.ChainID,
			},
			BlockNumber: &web3eye.Uint64Val{
				Op:    cruder.EQ,
				Value: *req.BlockNumber,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetBlocks(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))
	assert.Equal(t, infos[0], info)

	handler, err = NewHandler(context.Background(),
		WithID(&infos[0].ID, true),
	)
	assert.Nil(t, err)

	_info1, err := handler.DeleteBlock(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, infos[0].ID, _info1.ID)
}

func upsert(t *testing.T) {
	// update
	blockHash := "sssssss"
	handler, err := NewHandler(context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithBlockNumber(req.BlockNumber, true),
		WithBlockHash(&blockHash, true),
		WithBlockTime(req.BlockTime, true),
		WithParseState(req.ParseState, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info1, err := handler.UpsertBlock(context.Background())
	if assert.Nil(t, err) {
		assert.NotNil(t, info1)
		assert.Equal(t, blockHash, info1.BlockHash)
		req.BlockHash = &blockHash
	}

	// create
	chainID := "sssssss"
	handler, err = NewHandler(context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(&chainID, true),
		WithBlockNumber(req.BlockNumber, true),
		WithBlockHash(&blockHash, true),
		WithBlockTime(req.BlockTime, true),
		WithParseState(req.ParseState, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)
	info2, err := handler.UpsertBlock(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, blockHash, info2.BlockHash)
	assert.NotEqual(t, req.ChainID, info2.ChainID)
	assert.NotEqual(t, req.ID, info2.ID)
	assert.Equal(t, chainID, info2.ChainID)

	handler, err = NewHandler(context.Background(),
		WithID(&info2.ID, true),
	)
	assert.Nil(t, err)
	_info2, err := handler.DeleteBlock(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, _info2.ID, info2.ID)

}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	err := db.Init()
	if err != nil {
		fmt.Printf("cannot init database: %v \n", err)
		os.Exit(0)
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("upsert", upsert)
	t.Run("query", query)
}
