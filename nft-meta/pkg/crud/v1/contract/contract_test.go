package contract

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
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"

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
	entContract ent.Contract
	id          string

	contractInfo npool.ContractReq
	info         *ent.Contract
)

func prepareData() {
	entContract = ent.Contract{
		ID:        uuid.New(),
		ChainType: "eth",
		ChainID:   string(RandInt32()),
		Address:   "test",
		Name:      "ssss",
		Symbol:    "slsl",
		Creator:   "sssdf",
		Remark:    "",
	}

	id = entContract.ID.String()
	chainType := basetype.ChainType(basetype.ChainType_value[entContract.ChainType])
	contractInfo = npool.ContractReq{
		ID:        &id,
		ChainType: &chainType,
		ChainID:   &entContract.ChainID,
		Address:   &entContract.Address,
		Name:      &entContract.Name,
		Symbol:    &entContract.Symbol,
		Creator:   &entContract.Creator,
		Remark:    &entContract.Remark,
	}
}

func rowToObject(row *ent.Contract) *ent.Contract {
	return &ent.Contract{
		ID:        row.ID,
		ChainType: row.ChainType,
		ChainID:   row.ChainID,
		Address:   row.Address,
		Name:      row.Name,
		Symbol:    row.Symbol,
		Creator:   row.Creator,
		Remark:    row.Remark,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &contractInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entContract.ID = info.ID
			id := info.ID.String()
			contractInfo.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entContract)
	}
}

func createBulk(t *testing.T) {
	entContract := []ent.Contract{
		{
			ID:        uuid.New(),
			ChainType: "eth",
			ChainID:   string(RandInt32()),
			Address:   "",
			Name:      "",
			Symbol:    "",
			Creator:   "",
			Remark:    "",
		},
		{
			ID:        uuid.New(),
			ChainType: "eth",
			ChainID:   string(RandInt32()),
			Address:   "tensor",
			Name:      "tensor",
			Symbol:    "tensor",
			Creator:   "tensor",
			Remark:    "tensor",
		},
	}

	contracts := []*npool.ContractReq{}
	for key := range entContract {
		id := entContract[key].ID.String()
		chainType := basetype.ChainType(basetype.ChainType_value[entContract[key].ChainType])

		contracts = append(contracts, &npool.ContractReq{
			ID:        &id,
			ChainType: &chainType,
			ChainID:   &entContract[key].ChainID,
			Address:   &entContract[key].Address,
			Name:      &entContract[key].Name,
			Symbol:    &entContract[key].Symbol,
			Creator:   &entContract[key].Creator,
			Remark:    &entContract[key].Remark,
		})
	}
	infos, err := CreateBulk(context.Background(), contracts)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &contractInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entContract)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entContract)
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
		assert.Equal(t, rowToObject(infos[0]), &entContract)
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
		assert.Equal(t, rowToObject(info), &entContract)
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
		assert.Equal(t, rowToObject(info), &entContract)
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
