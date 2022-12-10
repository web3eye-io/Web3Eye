package synctask

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	val "github.com/web3eye-io/cyber-tracer/proto/cybertracer"

	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	cttype "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/cttype"
	npool "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/synctask"

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
	entSyncTask  ent.SyncTask
	id           string
	topicBase    string
	synctaskInfo npool.SyncTaskReq
	info         *ent.SyncTask
)

func prepareData() {
	topicBase = strconv.Itoa(int(time.Now().Unix()))
	entSyncTask = ent.SyncTask{
		ID:          uuid.New(),
		ChainType:   cttype.ChainType_Unkown.String(),
		ChainID:     RandInt32(),
		Start:       665,
		End:         665,
		Current:     665,
		Topic:       "s",
		Description: "asdfasdf",
		SyncState:   "asdfasdf",
		Remark:      "asdfasdf",
	}

	id = entSyncTask.ID.String()
	chainT := cttype.ChainType(cttype.ChainType_value[entSyncTask.ChainType])

	synctaskInfo = npool.SyncTaskReq{
		ID:          &id,
		ChainType:   &chainT,
		ChainID:     &entSyncTask.ChainID,
		Start:       &entSyncTask.Start,
		End:         &entSyncTask.End,
		Current:     &entSyncTask.Current,
		Topic:       &entSyncTask.Topic,
		Description: &entSyncTask.Description,
	}
}

func rowToObject(row *ent.SyncTask) *ent.SyncTask {
	return &ent.SyncTask{
		ID:          row.ID,
		ChainType:   row.ChainType,
		ChainID:     row.ChainID,
		Start:       row.Start,
		End:         row.End,
		Current:     row.Current,
		Topic:       row.Topic,
		Description: row.Description,
		SyncState:   row.SyncState,
		Remark:      row.Remark,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &synctaskInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entSyncTask.ID = info.ID
			id := info.ID.String()
			synctaskInfo.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entSyncTask)
	}
}

func createBulk(t *testing.T) {
	entSyncTask := []ent.SyncTask{
		{
			ID:          uuid.New(),
			ChainType:   "eth",
			ChainID:     RandInt32(),
			Start:       6655,
			End:         6655,
			Current:     6655,
			Topic:       "ss",
			Description: "asdfasdf",
			SyncState:   "asdfasdf",
			Remark:      "asdfasdf",
		},
		{
			ID:          uuid.New(),
			ChainType:   "eth",
			ChainID:     RandInt32(),
			Start:       6656,
			End:         6656,
			Current:     6656,
			Topic:       "sss",
			Description: "asdfasdf",
			SyncState:   "asdfasdf",
			Remark:      "asdfasdf",
		},
	}

	synctasks := []*npool.SyncTaskReq{}
	for key := range entSyncTask {
		id := entSyncTask[key].ID.String()
		chainT := cttype.ChainType(cttype.ChainType_value[entSyncTask[key].ChainType])
		syncS := cttype.SyncState(cttype.SyncState_value[entSyncTask[key].SyncState])

		synctasks = append(synctasks, &npool.SyncTaskReq{
			ID:          &id,
			ChainType:   &chainT,
			ChainID:     &entSyncTask[key].ChainID,
			Start:       &entSyncTask[key].Start,
			End:         &entSyncTask[key].End,
			Current:     &entSyncTask[key].Current,
			Topic:       &entSyncTask[key].Topic,
			Description: &entSyncTask[key].Description,
			SyncState:   &syncS,
			Remark:      &entSyncTask[key].Remark,
		})
	}
	infos, err := CreateBulk(context.Background(), synctasks)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &synctaskInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entSyncTask)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entSyncTask)
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
		assert.Equal(t, rowToObject(infos[0]), &entSyncTask)
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
		assert.Equal(t, rowToObject(info), &entSyncTask)
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
		assert.Equal(t, rowToObject(info), &entSyncTask)
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
