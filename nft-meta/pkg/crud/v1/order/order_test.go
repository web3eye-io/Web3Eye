package order

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"strconv"
	"testing"

	val "github.com/web3eye-io/Web3Eye/proto/web3eye"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var (
	entOrder ent.Order
	id       string

	orderReq npool.OrderReq
	info     *OrderDetail
)

func prepareData() {
	entOrder = ent.Order{
		ID:          uuid.New(),
		ChainType:   "test",
		ChainID:     "sssss",
		TxHash:      strconv.Itoa(RandInt()),
		BlockNumber: 111,
		TxIndex:     111,
		LogIndex:    111,
		Recipient:   "sssS",
		Remark:      "",
	}

	id = entOrder.ID.String()
	chainType := basetype.ChainType(basetype.ChainType_value[entOrder.ChainType])
	orderReq = npool.OrderReq{
		ID:          &id,
		ChainType:   &chainType,
		ChainID:     &entOrder.ChainID,
		TxHash:      &entOrder.TxHash,
		BlockNumber: &entOrder.BlockNumber,
		TxIndex:     &entOrder.TxIndex,
		LogIndex:    &entOrder.LogIndex,
		Recipient:   &entOrder.Recipient,
		Remark:      &entOrder.Remark,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &orderReq)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.Order.ID, uuid.UUID{}.String()) {
			entOrder.ID = info.Order.ID
			id := info.Order.ID.String()
			orderReq.ID = &id
		}
		assert.Equal(t, info, &entOrder)
	}
}

func createBulk(t *testing.T) {
	entOrder := []ent.Order{
		{
			ID:          uuid.New(),
			ChainType:   "test",
			ChainID:     "sssss",
			TxHash:      strconv.Itoa(RandInt()),
			BlockNumber: 111,
			TxIndex:     111,
			LogIndex:    111,
			Recipient:   "sssS",
			Remark:      "",
		},
		{
			ID:          uuid.New(),
			ChainType:   "test",
			ChainID:     "sssss",
			TxHash:      strconv.Itoa(RandInt()),
			BlockNumber: 111,
			TxIndex:     111,
			LogIndex:    111,
			Recipient:   "sssS",
			Remark:      "",
		},
	}

	orders := []*npool.OrderReq{}
	for key := range entOrder {
		id := entOrder[key].ID.String()
		chainType := basetype.ChainType(basetype.ChainType_value[entOrder[key].ChainType])

		orders = append(orders, &npool.OrderReq{
			ID:          &id,
			ChainType:   &chainType,
			ChainID:     &entOrder[key].ChainID,
			TxHash:      &entOrder[key].TxHash,
			BlockNumber: &entOrder[key].BlockNumber,
			TxIndex:     &entOrder[key].TxIndex,
			LogIndex:    &entOrder[key].LogIndex,
			Recipient:   &entOrder[key].Recipient,
			Remark:      &entOrder[key].Remark,
		})
	}
	infos, err := CreateBulk(context.Background(), orders)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].Order.ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].Order.ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &orderReq)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &entOrder)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.Order.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &entOrder)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.Order.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0], &entOrder)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.Order.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &entOrder)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.Order.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), info.Order.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.Order.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.Order.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &entOrder)
	}
}

func RandInt64() int64 {
	MaxUint64 := ^uint64(0)
	MaxInt64 := int64(MaxUint64 >> 1)
	randInt, err := rand.Int(rand.Reader, big.NewInt(MaxInt64))
	if err != nil {
		return 0
	}
	return randInt.Int64()
}

func RandInt() int {
	return int(RandInt64())
}

func RandInt32() int32 {
	return int32(RandInt64())
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	prepareData()
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("update", update)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("delete", deleteT)
	t.Run("count", count)
}
