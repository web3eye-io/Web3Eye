package transfer

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
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"

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
	entTransfer ent.Transfer
	id          string

	transferReq npool.TransferReq
	info        *ent.Transfer
)

func prepareData() {
	entTransfer = ent.Transfer{
		ID:          uuid.New(),
		ChainType:   "eth",
		ChainID:     string(RandInt32()),
		Contract:    "test",
		TokenType:   "sssss",
		TokenID:     strconv.Itoa(RandInt()),
		From:        "f",
		To:          "t",
		Amount:      1,
		BlockNumber: 123,
		TxHash:      uuid.NewString(),
		BlockHash:   uuid.NewString(),
		TxTime:      0,
		Remark:      "",
	}

	id = entTransfer.ID.String()
	chainType := basetype.ChainType(basetype.ChainType_value[entTransfer.ChainType])
	tokenType := basetype.TokenType(basetype.TokenType_value[entTransfer.TokenType])
	transferReq = npool.TransferReq{
		ID: &id,

		ChainType:   &chainType,
		ChainID:     &entTransfer.ChainID,
		Contract:    &entTransfer.Contract,
		TokenType:   &tokenType,
		TokenID:     &entTransfer.TokenID,
		From:        &entTransfer.From,
		To:          &entTransfer.To,
		Amount:      &entTransfer.Amount,
		BlockNumber: &entTransfer.BlockNumber,
		TxHash:      &entTransfer.TxHash,
		BlockHash:   &entTransfer.BlockHash,
		TxTime:      &entTransfer.TxTime,
		Remark:      &entTransfer.Remark,
	}
}

func rowToObject(row *ent.Transfer) *ent.Transfer {
	return &ent.Transfer{
		ID:          row.ID,
		ChainType:   row.ChainType,
		ChainID:     row.ChainID,
		Contract:    row.Contract,
		TokenType:   row.TokenType,
		TokenID:     row.TokenID,
		From:        row.From,
		To:          row.To,
		Amount:      row.Amount,
		BlockNumber: row.BlockNumber,
		TxHash:      row.TxHash,
		BlockHash:   row.BlockHash,
		TxTime:      row.TxTime,
		Remark:      row.Remark,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &transferReq)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entTransfer.ID = info.ID
			id := info.ID.String()
			transferReq.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entTransfer)
	}
}

func createBulk(t *testing.T) {
	entTransfer := []ent.Transfer{
		{
			ID:          uuid.New(),
			ChainType:   "eth",
			ChainID:     string(RandInt32()),
			Contract:    "test",
			TokenType:   "sssss",
			TokenID:     strconv.Itoa(RandInt()),
			From:        "f",
			To:          "t",
			Amount:      1,
			BlockNumber: 123,
			TxHash:      uuid.NewString(),
			BlockHash:   uuid.NewString(),
			TxTime:      0,
			Remark:      "",
		},
		{
			ID:          uuid.New(),
			ChainType:   "eth",
			ChainID:     string(RandInt32()),
			Contract:    "test",
			TokenType:   "sssss",
			TokenID:     strconv.Itoa(RandInt()),
			From:        "f",
			To:          "t",
			Amount:      1,
			BlockNumber: 123,
			TxHash:      uuid.NewString(),
			BlockHash:   uuid.NewString(),
			TxTime:      0,
			Remark:      "",
		},
	}

	transfers := []*npool.TransferReq{}
	for key := range entTransfer {
		id := entTransfer[key].ID.String()
		chainType := basetype.ChainType(basetype.ChainType_value[entTransfer[key].ChainType])
		tokenType := basetype.TokenType(basetype.TokenType_value[entTransfer[key].TokenType])

		transfers = append(transfers, &npool.TransferReq{
			ID:          &id,
			ChainType:   &chainType,
			ChainID:     &entTransfer[key].ChainID,
			Contract:    &entTransfer[key].Contract,
			TokenType:   &tokenType,
			TokenID:     &entTransfer[key].TokenID,
			From:        &entTransfer[key].From,
			To:          &entTransfer[key].To,
			Amount:      &entTransfer[key].Amount,
			BlockNumber: &entTransfer[key].BlockNumber,
			TxHash:      &entTransfer[key].TxHash,
			BlockHash:   &entTransfer[key].BlockHash,
			TxTime:      &entTransfer[key].TxTime,
			Remark:      &entTransfer[key].Remark,
		})
	}
	infos, err := CreateBulk(context.Background(), transfers)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &transferReq)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entTransfer)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entTransfer)
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
		assert.Equal(t, rowToObject(infos[0]), &entTransfer)
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
		assert.Equal(t, rowToObject(info), &entTransfer)
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
		assert.Equal(t, rowToObject(info), &entTransfer)
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
