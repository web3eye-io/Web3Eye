package blocknumber

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	val "github.com/web3eye-io/cyber-tracer/message/cybertracer"

	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/web3eye-io/cyber-tracer/message/cybertracer/nftmeta/v1/blocknumber"

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
	entBlockNumber  ent.BlockNumber
	id              string
	topicBase       string
	blocknumberInfo npool.BlockNumberReq
	info            *ent.BlockNumber
)

func prepareData() {
	topicBase = strconv.Itoa(int(time.Now().Unix()))
	entBlockNumber = ent.BlockNumber{
		ID:          uuid.New(),
		ChainType:   "eth",
		ChainID:     RandInt32(),
		Identifier:  "",
		CurrentNum:  665,
		Topic:       "s",
		Description: "asdfasdf",
	}

	id = entBlockNumber.ID.String()

	blocknumberInfo = npool.BlockNumberReq{
		ID:          &id,
		ChainType:   &entBlockNumber.ChainType,
		ChainID:     &entBlockNumber.ChainID,
		Identifier:  &entBlockNumber.Identifier,
		CurrentNum:  &entBlockNumber.CurrentNum,
		Topic:       &entBlockNumber.Topic,
		Description: &entBlockNumber.Description,
	}
}

func rowToObject(row *ent.BlockNumber) *ent.BlockNumber {
	return &ent.BlockNumber{
		ID:          row.ID,
		ChainType:   row.ChainType,
		ChainID:     row.ChainID,
		Identifier:  row.Identifier,
		CurrentNum:  row.CurrentNum,
		Topic:       row.Topic,
		Description: row.Description,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &blocknumberInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entBlockNumber.ID = info.ID
			id := info.ID.String()
			blocknumberInfo.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entBlockNumber)
	}
}

func createBulk(t *testing.T) {
	entBlockNumber := []ent.BlockNumber{
		{
			ID:          uuid.New(),
			ChainType:   "eth",
			ChainID:     RandInt32(),
			Identifier:  "",
			CurrentNum:  6655,
			Topic:       "ss",
			Description: "asdfasdf",
		},
		{
			ID:          uuid.New(),
			ChainType:   "eth",
			ChainID:     RandInt32(),
			Identifier:  "",
			CurrentNum:  6656,
			Topic:       "sss",
			Description: "asdfasdf",
		},
	}

	blocknumbers := []*npool.BlockNumberReq{}
	for key := range entBlockNumber {
		id := entBlockNumber[key].ID.String()

		blocknumbers = append(blocknumbers, &npool.BlockNumberReq{
			ID:          &id,
			ChainType:   &entBlockNumber[key].ChainType,
			ChainID:     &entBlockNumber[key].ChainID,
			Identifier:  &entBlockNumber[key].Identifier,
			CurrentNum:  &entBlockNumber[key].CurrentNum,
			Topic:       &entBlockNumber[key].Topic,
			Description: &entBlockNumber[key].Description,
		})
	}
	infos, err := CreateBulk(context.Background(), blocknumbers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &blocknumberInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entBlockNumber)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entBlockNumber)
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
		assert.Equal(t, rowToObject(infos[0]), &entBlockNumber)
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
		assert.Equal(t, rowToObject(info), &entBlockNumber)
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
		assert.Equal(t, rowToObject(info), &entBlockNumber)
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
