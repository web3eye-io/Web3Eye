package transfer

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/stretchr/testify/assert"

	transferproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &transferproto.Transfer{
	ChainType:   basetype.ChainType_Ethereum,
	ChainID:     "test_transfer",
	Contract:    "test_transfer",
	TokenType:   basetype.TokenType_ERC20,
	TokenID:     "test_transfer",
	From:        "test_transfer",
	To:          "test_transfer",
	Amount:      10010,
	BlockNumber: 10086,
	TxHash:      "test_transfer",
	BlockHash:   "test_transfer",
	TxTime:      10011,
	Remark:      "test_transfer",
}

var req = &transferproto.TransferReq{
	ChainType:   &ret.ChainType,
	ChainID:     &ret.ChainID,
	Contract:    &ret.Contract,
	TokenType:   &ret.TokenType,
	TokenID:     &ret.TokenID,
	From:        &ret.From,
	To:          &ret.To,
	Amount:      &ret.Amount,
	BlockNumber: &ret.BlockNumber,
	TxHash:      &ret.TxHash,
	BlockHash:   &ret.BlockHash,
	TxTime:      &ret.TxTime,
	Remark:      &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithContract(req.Contract, true),
		WithTokenType(req.TokenType, true),
		WithTokenID(req.TokenID, true),
		WithFrom(req.From, true),
		WithTo(req.To, true),
		WithAmount(req.Amount, true),
		WithBlockNumber(req.BlockNumber, true),
		WithTxHash(req.TxHash, true),
		WithBlockHash(req.BlockHash, true),
		WithTxTime(req.TxTime, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateTransfer(context.Background())
	if assert.Nil(t, err) {
		req.ID = &info.ID
		req.EntID = &info.EntID
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
	}
}

func update(t *testing.T) {
	req.ChainType = basetype.ChainType_Solana.Enum()
	txHash := "google"
	req.TxHash = &txHash

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithChainType(req.ChainType, false),
		WithChainID(req.ChainID, false),
		WithTxHash(req.TxHash, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateTransfer(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
		assert.Equal(t, info.TokenTypeStr, req.TokenType.String())
		assert.Equal(t, info.TxHash, req.GetTxHash())
	}
}

func upsert(t *testing.T) {
	// just update
	remark := "sssssss"
	handler, err := NewHandler(context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithContract(req.Contract, true),
		WithTokenID(req.TokenID, true),
		WithTxHash(req.TxHash, true),
		WithFrom(req.From, true),
		WithRemark(&remark, true),
	)
	assert.Nil(t, err)

	info1, err := handler.UpsertTransfer(context.Background())
	if assert.Nil(t, err) {
		assert.NotNil(t, info1)
		assert.Equal(t, info1.Remark, remark)
		req.Remark = &remark
	}

	// can success for upsert
	tokenID := "sssssss"
	handler, err = NewHandler(context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithContract(req.Contract, true),
		WithTokenType(req.TokenType, true),
		WithTokenID(&tokenID, true),
		WithFrom(req.From, true),
		WithTo(req.To, true),
		WithAmount(req.Amount, true),
		WithBlockNumber(req.BlockNumber, true),
		WithTxHash(req.TxHash, true),
		WithBlockHash(req.BlockHash, true),
		WithTxTime(req.TxTime, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)
	info2, err := handler.UpsertTransfer(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, req.GetFrom(), info2.From)
	assert.NotEqual(t, req.GetTokenID(), info2.TokenID)
	assert.NotEqual(t, req.GetID(), info2.ID)

	handler, err = NewHandler(context.Background(),
		WithID(&info2.ID, true),
	)
	assert.Nil(t, err)
	_info2, err := handler.DeleteTransfer(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, _info2.ID, info2.ID)

}

func query(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetTransfer(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&transferproto.Conds{
			Contract: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.Contract,
			},
			TokenID: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.TokenID,
			},
			TxHash: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.TxHash,
			},
			From: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.From,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetTransfers(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))
	assert.Equal(t, infos[0], info)

	handler, err = NewHandler(context.Background(),
		WithID(&infos[0].ID, true),
	)
	assert.Nil(t, err)

	_info1, err := handler.DeleteTransfer(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, infos[0].ID, _info1.ID)
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
