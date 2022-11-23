package token

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"strconv"
	"testing"

	val "github.com/NpoolPlatform/message/npool"

	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/nftmeta/v1/token"

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
	entToken ent.Token
	id       string

	tokenInfo npool.TokenReq
	info      *ent.Token
)

func prepareData() {
	entToken = ent.Token{
		ID:        uuid.New(),
		VectorID:  RandInt64(),
		ChainType: "eth",
		ChainID:   RandInt32(),
		Contract:  "test",
		TokenType: "1155",
		TokenID:   strconv.Itoa(RandInt()),
		URI:       "uri",
		URIType:   "ss",
		ImageURL:  "ss",
		VideoURL:  "ss",
		Remark:    "",
	}

	id = entToken.ID.String()

	tokenInfo = npool.TokenReq{
		ID:        &id,
		VectorID:  &entToken.VectorID,
		ChainType: &entToken.ChainType,
		ChainID:   &entToken.ChainID,
		Contract:  &entToken.Contract,
		TokenType: &entToken.TokenType,
		TokenID:   &entToken.TokenID,
		URI:       &entToken.URI,
		URIType:   &entToken.URIType,
		ImageURL:  &entToken.ImageURL,
		VideoURL:  &entToken.VideoURL,
		Remark:    &entToken.Remark,
	}
}

func rowToObject(row *ent.Token) *ent.Token {
	return &ent.Token{
		ID:        row.ID,
		VectorID:  row.VectorID,
		ChainType: row.ChainType,
		ChainID:   row.ChainID,
		Contract:  row.Contract,
		TokenType: row.TokenType,
		TokenID:   row.TokenID,
		URI:       row.URI,
		URIType:   row.URIType,
		ImageURL:  row.ImageURL,
		VideoURL:  row.VideoURL,
		Remark:    row.Remark,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &tokenInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entToken.ID = info.ID
			id := info.ID.String()
			tokenInfo.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entToken)
	}
}

func createBulk(t *testing.T) {
	entToken := []ent.Token{
		{
			ID:        uuid.New(),
			VectorID:  RandInt64(),
			ChainType: "eth",
			ChainID:   RandInt32(),
			Contract:  "test",
			TokenType: "1155",
			TokenID:   strconv.Itoa(RandInt()),
			URI:       "uri",
			URIType:   "ss",
			ImageURL:  "ss",
			VideoURL:  "ss",
			Remark:    "",
		},
		{
			ID:        uuid.New(),
			VectorID:  RandInt64(),
			ChainType: "eth",
			ChainID:   RandInt32(),
			Contract:  "test",
			TokenType: "1155",
			TokenID:   strconv.Itoa(RandInt()),
			URI:       "uri",
			URIType:   "ss",
			ImageURL:  "ss",
			VideoURL:  "ss",
			Remark:    "tensor",
		},
	}

	tokens := []*npool.TokenReq{}
	for key := range entToken {
		id := entToken[key].ID.String()

		tokens = append(tokens, &npool.TokenReq{
			ID:        &id,
			VectorID:  &entToken[key].VectorID,
			ChainType: &entToken[key].ChainType,
			ChainID:   &entToken[key].ChainID,
			Contract:  &entToken[key].Contract,
			TokenType: &entToken[key].TokenType,
			TokenID:   &entToken[key].TokenID,
			URI:       &entToken[key].URI,
			URIType:   &entToken[key].URIType,
			ImageURL:  &entToken[key].ImageURL,
			VideoURL:  &entToken[key].VideoURL,
			Remark:    &entToken[key].Remark,
		})
	}
	infos, err := CreateBulk(context.Background(), tokens)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &tokenInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entToken)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entToken)
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
		assert.Equal(t, rowToObject(infos[0]), &entToken)
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
		assert.Equal(t, rowToObject(info), &entToken)
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
		assert.Equal(t, rowToObject(info), &entToken)
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
