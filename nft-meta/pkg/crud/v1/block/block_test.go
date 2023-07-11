package block

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
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"

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
	entBlock ent.Block
	id       string

	blockInfo npool.BlockReq
	info      *ent.Block
)

func prepareData() {
	entBlock = ent.Block{
		ID:          uuid.New(),
		ChainType:   "eth",
		ChainID:     string(RandInt32()),
		BlockNumber: uint64(5),
		BlockHash:   "1155",
		BlockTime:   1222,
	}

	id = entBlock.ID.String()
	chainType := basetype.ChainType(basetype.ChainType_value[entBlock.ChainType])
	blockInfo = npool.BlockReq{
		ID:          &id,
		ChainType:   &chainType,
		ChainID:     &entBlock.ChainID,
		BlockNumber: &entBlock.BlockNumber,
		BlockHash:   &entBlock.BlockHash,
		BlockTime:   &entBlock.BlockTime,
	}
}

func rowToObject(row *ent.Block) *ent.Block {
	return &ent.Block{
		ID:          row.ID,
		ChainType:   row.ChainType,
		ChainID:     row.ChainID,
		BlockNumber: row.BlockNumber,
		BlockHash:   row.BlockHash,
		BlockTime:   row.BlockTime,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &blockInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entBlock.ID = info.ID
			id := info.ID.String()
			blockInfo.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entBlock)
	}
}

func createBulk(t *testing.T) {
	entBlock := []ent.Block{
		{
			ID:          uuid.New(),
			ChainType:   "eth",
			ChainID:     string(RandInt32()),
			BlockNumber: uint64(5),
			BlockHash:   "1155",
			BlockTime:   1122,
		},
		{
			ID:          uuid.New(),
			ChainType:   "eth",
			ChainID:     string(RandInt32()),
			BlockNumber: uint64(5),
			BlockHash:   "1155",
			BlockTime:   1122,
		},
	}

	blocks := []*npool.BlockReq{}
	for key := range entBlock {
		id := entBlock[key].ID.String()
		chainType := basetype.ChainType(basetype.ChainType_value[entBlock[key].ChainType])
		blocks = append(blocks, &npool.BlockReq{
			ID:          &id,
			ChainType:   &chainType,
			ChainID:     &entBlock[key].ChainID,
			BlockNumber: &entBlock[key].BlockNumber,
			BlockHash:   &entBlock[key].BlockHash,
			BlockTime:   &entBlock[key].BlockTime,
		})
	}
	infos, err := CreateBulk(context.Background(), blocks)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &blockInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entBlock)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entBlock)
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
		assert.Equal(t, rowToObject(infos[0]), &entBlock)
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
		assert.Equal(t, rowToObject(info), &entBlock)
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
		assert.Equal(t, rowToObject(info), &entBlock)
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
