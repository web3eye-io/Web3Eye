package contract

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

	contractproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &contractproto.Contract{
	ChainType:   basetype.ChainType_Ethereum,
	ChainID:     "test_contract",
	Address:     "test_contract",
	Name:        "test_contract",
	Symbol:      "test_contract",
	Decimals:    6,
	Creator:     "test_contract",
	BlockNum:    10010,
	TxHash:      "test_contract",
	TxTime:      1988,
	ProfileURL:  "test_contract",
	BaseURL:     "test_contract",
	BannerURL:   "test_contract",
	Description: "test_contract",
	Remark:      "test_contract",
}

var req = &contractproto.ContractReq{
	ChainType:   &ret.ChainType,
	ChainID:     &ret.ChainID,
	Address:     &ret.Address,
	Name:        &ret.Name,
	Symbol:      &ret.Symbol,
	Decimals:    &ret.Decimals,
	Creator:     &ret.Creator,
	BlockNum:    &ret.BlockNum,
	TxHash:      &ret.TxHash,
	TxTime:      &ret.TxTime,
	ProfileURL:  &ret.ProfileURL,
	BaseURL:     &ret.BaseURL,
	BannerURL:   &ret.BannerURL,
	Description: &ret.Description,
	Remark:      &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithAddress(req.Address, true),
		WithName(req.Name, true),
		WithSymbol(req.Symbol, true),
		WithDecimals(req.Decimals, true),
		WithCreator(req.Creator, true),
		WithBlockNum(req.BlockNum, true),
		WithTxHash(req.TxHash, true),
		WithTxTime(req.TxTime, true),
		WithProfileURL(req.ProfileURL, true),
		WithBaseURL(req.BaseURL, true),
		WithBannerURL(req.BannerURL, true),
		WithDescription(req.Description, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateContract(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info.ChainTypeStr, ret.ChainType.String())
	}
}

func update(t *testing.T) {
	ret.ChainType = basetype.ChainType_Solana
	url := "google"
	req.ProfileURL = &url

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithChainType(req.ChainType, false),
		WithChainID(req.ChainID, false),
		WithProfileURL(req.ProfileURL, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateContract(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info.ChainTypeStr, ret.ChainType.String())
	}
}

func query(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetContract(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&contractproto.Conds{
			ChainType: &web3eye.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(*req.ChainType),
			},
			ChainID: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.ChainID,
			},
			Address: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.Address,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetContracts(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))
	assert.Equal(t, infos[0], info)

	handler, err = NewHandler(context.Background(),
		WithID(&infos[0].ID, true),
	)
	assert.Nil(t, err)

	_info1, err := handler.DeleteContract(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, infos[0].ID, _info1.ID)
}

func upsert(t *testing.T) {
	// cannot success for upsert
	url := "sssssss"
	handler, err := NewHandler(context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithAddress(req.Address, true),
		WithBannerURL(&url, true),
		WithDecimals(req.Decimals, true),
		WithDescription(req.Description, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info1, err := handler.UpsertContract(context.Background())
	if assert.Nil(t, err) {
		assert.NotNil(t, info1)
		assert.Equal(t, info1.BannerURL, url)
		req.BannerURL = &url
	}

	// can success for upsert
	chainID := "sssssss"
	handler, err = NewHandler(context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(&chainID, true),
		WithAddress(req.Address, true),
		WithName(req.Name, true),
		WithSymbol(req.Symbol, true),
		WithDecimals(req.Decimals, true),
		WithCreator(req.Creator, true),
		WithBlockNum(req.BlockNum, true),
		WithTxHash(req.TxHash, true),
		WithTxTime(req.TxTime, true),
		WithProfileURL(req.ProfileURL, true),
		WithBaseURL(req.BaseURL, true),
		WithBannerURL(req.BannerURL, true),
		WithDescription(req.Description, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)
	info2, err := handler.UpsertContract(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, *req.BannerURL, info2.BannerURL)
	assert.NotEqual(t, req.ChainID, info2.ChainID)
	assert.NotEqual(t, req.ID, info2.ID)
	assert.Equal(t, chainID, info2.ChainID)

	handler, err = NewHandler(context.Background(),
		WithID(&info2.ID, true),
	)
	assert.Nil(t, err)
	_info2, err := handler.DeleteContract(context.Background())
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
