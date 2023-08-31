package orderpair

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

	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/orderpair"

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
	entOrderPair ent.OrderPair
	id           string

	orderpairReq npool.OrderPairReq
	info         *ent.OrderPair
)

func prepareData() {
	entOrderPair = ent.OrderPair{
		ID:        uuid.New(),
		TxHash:    "test",
		Recipient: "sssss",
		TargetID:  strconv.Itoa(RandInt()),
		BarterID:  "1",
		Remark:    "",
	}

	id = entOrderPair.ID.String()
	orderpairReq = npool.OrderPairReq{
		ID:        &id,
		TxHash:    &entOrderPair.TxHash,
		Recipient: &entOrderPair.Recipient,
		TargetID:  &entOrderPair.TargetID,
		BarterID:  &entOrderPair.BarterID,
		Remark:    &entOrderPair.Remark,
	}
}

func rowToObject(row *ent.OrderPair) *ent.OrderPair {
	return &ent.OrderPair{
		ID:        row.ID,
		TxHash:    row.TxHash,
		Recipient: row.Recipient,
		TargetID:  row.TargetID,
		BarterID:  row.BarterID,
		Remark:    row.Remark,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &orderpairReq)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entOrderPair.ID = info.ID
			id := info.ID.String()
			orderpairReq.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entOrderPair)
	}
}

func createBulk(t *testing.T) {
	entOrderPair := []ent.OrderPair{
		{
			ID:        uuid.New(),
			TxHash:    "test",
			Recipient: "sssss",
			TargetID:  strconv.Itoa(RandInt()),
			BarterID:  "1",
			Remark:    "",
		},
		{
			ID:        uuid.New(),
			TxHash:    "test",
			Recipient: "sssss",
			TargetID:  strconv.Itoa(RandInt()),
			BarterID:  "1",
			Remark:    "",
		},
	}

	orderpairs := []*npool.OrderPairReq{}
	for key := range entOrderPair {
		id := entOrderPair[key].ID.String()
		orderpairs = append(orderpairs, &npool.OrderPairReq{
			ID:        &id,
			TxHash:    &entOrderPair[key].TxHash,
			Recipient: &entOrderPair[key].Recipient,
			TargetID:  &entOrderPair[key].TargetID,
			BarterID:  &entOrderPair[key].BarterID,
			Remark:    &entOrderPair[key].Remark,
		})
	}
	infos, err := CreateBulk(context.Background(), orderpairs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &orderpairReq)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entOrderPair)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entOrderPair)
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
		assert.Equal(t, rowToObject(infos[0]), &entOrderPair)
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
		assert.Equal(t, rowToObject(info), &entOrderPair)
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
		assert.Equal(t, rowToObject(info), &entOrderPair)
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
