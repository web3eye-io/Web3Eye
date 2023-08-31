package orderitem

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

	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/orderitem"

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
	entOrderItem ent.OrderItem
	id           string

	orderitemReq npool.OrderItemReq
	info         *ent.OrderItem
)

func prepareData() {
	entOrderItem = ent.OrderItem{
		ID:         uuid.New(),
		Contract:   "test",
		TokenType:  "sssss",
		TokenID:    strconv.Itoa(RandInt()),
		Amount:     1,
		PortionNum: 5,
		Remark:     "",
	}

	id = entOrderItem.ID.String()
	orderitemReq = npool.OrderItemReq{
		ID:         &id,
		Contract:   &entOrderItem.Contract,
		TokenType:  &entOrderItem.TokenType,
		TokenID:    &entOrderItem.TokenID,
		Amount:     &entOrderItem.Amount,
		PortionNum: &entOrderItem.PortionNum,
		Remark:     &entOrderItem.Remark,
	}
}

func rowToObject(row *ent.OrderItem) *ent.OrderItem {
	return &ent.OrderItem{
		ID:         row.ID,
		Contract:   row.Contract,
		TokenType:  row.TokenType,
		TokenID:    row.TokenID,
		Amount:     row.Amount,
		PortionNum: row.PortionNum,
		Remark:     row.Remark,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &orderitemReq)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entOrderItem.ID = info.ID
			id := info.ID.String()
			orderitemReq.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entOrderItem)
	}
}

func createBulk(t *testing.T) {
	entOrderItem := []ent.OrderItem{
		{
			ID:         uuid.New(),
			Contract:   "test",
			TokenType:  "sssss",
			TokenID:    strconv.Itoa(RandInt()),
			Amount:     1,
			PortionNum: 123,
			Remark:     "",
		},
		{
			ID:         uuid.New(),
			Contract:   "test",
			TokenType:  "sssss",
			TokenID:    strconv.Itoa(RandInt()),
			Amount:     1,
			PortionNum: 111,
			Remark:     "",
		},
	}

	orderitems := []*npool.OrderItemReq{}
	for key := range entOrderItem {
		id := entOrderItem[key].ID.String()
		orderitems = append(orderitems, &npool.OrderItemReq{
			ID:         &id,
			Contract:   &entOrderItem[key].Contract,
			TokenType:  &entOrderItem[key].TokenType,
			TokenID:    &entOrderItem[key].TokenID,
			Amount:     &entOrderItem[key].Amount,
			PortionNum: &entOrderItem[key].PortionNum,
			Remark:     &entOrderItem[key].Remark,
		})
	}
	infos, err := CreateBulk(context.Background(), orderitems)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &orderitemReq)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entOrderItem)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entOrderItem)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entOrderItem)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entOrderItem)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entOrderItem)
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
