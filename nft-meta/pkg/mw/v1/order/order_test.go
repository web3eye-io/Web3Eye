package order

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

	orderproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &orderproto.Order{
	ChainType:   basetype.ChainType_Ethereum,
	ChainID:     "test_order",
	TxHash:      "test_order",
	BlockNumber: 10010,
	TxIndex:     10086,
	LogIndex:    10011,
	Recipient:   "test_order",
	TargetItems: []*orderproto.OrderItem{
		{
			Contract:  "test_order_item",
			TokenType: basetype.TokenType_ERC20,
			TokenID:   "test_order_item",
			Amount:    1111,
			Remark:    "test_order_item",
		},
		{
			Contract:  "test_order_item",
			TokenType: basetype.TokenType_ERC20,
			TokenID:   "test_ordem",
			Amount:    1111,
			Remark:    "test_order_item",
		},
	},
	OfferItems: []*orderproto.OrderItem{
		{
			Contract:  "test_order_item",
			TokenType: basetype.TokenType_ERC20,
			TokenID:   "test_order_item",
			Amount:    1111,
			Remark:    "test_order_item",
		},
		{
			Contract:  "test_order_item",
			TokenType: basetype.TokenType_ERC20,
			TokenID:   "test_ordem",
			Amount:    1111,
			Remark:    "test_order_item",
		},
	},
	Remark: "test_order",
}

var req = &orderproto.OrderReq{
	ChainType:   &ret.ChainType,
	ChainID:     &ret.ChainID,
	TxHash:      &ret.TxHash,
	BlockNumber: &ret.BlockNumber,
	TxIndex:     &ret.TxIndex,
	LogIndex:    &ret.LogIndex,
	Recipient:   &ret.Recipient,
	TargetItems: ret.TargetItems,
	OfferItems:  ret.OfferItems,
	Remark:      &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithTxHash(req.TxHash, true),
		WithBlockNumber(req.BlockNumber, true),
		WithTxIndex(req.TxIndex, true),
		WithLogIndex(req.LogIndex, true),
		WithRecipient(req.Recipient, true),
		WithTargetItems(req.TargetItems, true),
		WithOfferItems(req.OfferItems, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateOrder(context.Background())
	if assert.Nil(t, err) {
		req.ID = &info.ID
		req.EntID = &info.EntID
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
		assert.Equal(t, len(info.TargetItems), len(req.TargetItems))
		assert.Equal(t, len(info.OfferItems), len(req.OfferItems))
	}
}

func update(t *testing.T) {
	req.ChainType = basetype.ChainType_Solana.Enum()
	recipient := "google"
	req.Recipient = &recipient
	req.TargetItems = []*orderproto.OrderItem{
		{
			Contract:  "sssss",
			TokenType: basetype.TokenType_ERC20,
			TokenID:   "ssssss",
			Amount:    6666,
			Remark:    "ssss",
		},
	}
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithChainType(req.ChainType, false),
		WithChainID(req.ChainID, false),
		WithTargetItems(req.TargetItems, false),
		WithOfferItems(req.OfferItems, false),
		WithRecipient(req.Recipient, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateOrder(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
		assert.Equal(t, info.Recipient, req.GetRecipient())
		assert.Equal(t, info.TargetItems[0], req.TargetItems[0])
		assert.Equal(t, info.OfferItems[0], req.OfferItems[0])
	}
}

func query(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetOrder(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&orderproto.Conds{
			ChainType: &web3eye.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(*req.ChainType),
			},
			ChainID: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.ChainID,
			},
			TxHash: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.TxHash,
			},
			LogIndex: &web3eye.Uint32Val{
				Op:    cruder.EQ,
				Value: *req.LogIndex,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetOrders(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))
	assert.Equal(t, infos[0], info)

	handler, err = NewHandler(context.Background(),
		WithID(&infos[0].ID, true),
	)
	assert.Nil(t, err)

	_info1, err := handler.DeleteOrder(context.Background())
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
	t.Run("query", query)
}
