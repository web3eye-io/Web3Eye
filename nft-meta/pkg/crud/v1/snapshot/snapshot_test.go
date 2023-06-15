package snapshot

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"

	val "github.com/web3eye-io/Web3Eye/proto/web3eye"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	//nolint
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	if err == nil && runByGithubAction {
		return
	}
}

var (
	entSnapshot ent.Snapshot
	id          string

	snapshotInfo       npool.SnapshotRequest
	updateSnapshotInfo npool.UpdateSnapshotRequest
	info               *ent.Snapshot
)

func prepareData() {
	entSnapshot = ent.Snapshot{
		ID:            uuid.New(),
		Index:         RandUInt64(),
		SnapshotCommP: "eeeeeeeeeeeeee",
		SnapshotRoot:  fmt.Sprint(RandUInt64()),
		SnapshotURI:   "test",
		BackupState:   npool.BackupState_BackupStateCreated.String(),
	}

	id = entSnapshot.ID.String()
	snapshotInfo = npool.SnapshotRequest{
		Index:         entSnapshot.Index,
		SnapshotCommP: entSnapshot.SnapshotCommP,
		SnapshotRoot:  entSnapshot.SnapshotRoot,
		SnapshotURI:   entSnapshot.SnapshotURI,
		BackupState:   npool.BackupState(npool.BackupState_value[entSnapshot.BackupState]),
	}
	updateSnapshotInfo = npool.UpdateSnapshotRequest{
		Index:         &snapshotInfo.Index,
		SnapshotCommP: &snapshotInfo.SnapshotCommP,
		SnapshotRoot:  &snapshotInfo.SnapshotRoot,
		SnapshotURI:   &snapshotInfo.SnapshotURI,
		BackupState:   &snapshotInfo.BackupState,
	}
}

func rowToObject(row *ent.Snapshot) *ent.Snapshot {
	return &ent.Snapshot{
		ID:            row.ID,
		Index:         row.Index,
		SnapshotCommP: row.SnapshotCommP,
		SnapshotRoot:  row.SnapshotRoot,
		SnapshotURI:   row.SnapshotURI,
		BackupState:   row.BackupState,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &snapshotInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entSnapshot)
		updateSnapshotInfo.ID = info.ID.String()
	}
}

func createBulk(t *testing.T) {
	entSnapshots := []ent.Snapshot{
		{
			ID:            uuid.New(),
			Index:         RandUInt64(),
			SnapshotCommP: "eeeeeeeeeeeeee1",
			SnapshotRoot:  fmt.Sprint(RandUInt64()),
			SnapshotURI:   "test",
			BackupState:   npool.BackupState_BackupStateCreated.String(),
		},
		{
			ID:            uuid.New(),
			Index:         RandUInt64(),
			SnapshotCommP: "eeeeeeeeeeeeee2",
			SnapshotRoot:  fmt.Sprint(RandUInt64()),
			SnapshotURI:   "test",
			BackupState:   npool.BackupState_BackupStateCreated.String(),
		},
	}

	snapshots := []*npool.SnapshotRequest{}
	for key := range entSnapshots {
		snapshots = append(snapshots, &npool.SnapshotRequest{
			Index:         entSnapshots[key].Index,
			SnapshotCommP: entSnapshots[key].SnapshotCommP,
			SnapshotRoot:  entSnapshots[key].SnapshotRoot,
			SnapshotURI:   entSnapshots[key].SnapshotURI,
			BackupState:   npool.BackupState(npool.BackupState_value[entSnapshots[key].BackupState]),
		})
	}
	infos, err := CreateBulk(context.Background(), snapshots)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &updateSnapshotInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entSnapshot)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entSnapshot)
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
		assert.Equal(t, rowToObject(infos[0]), &entSnapshot)
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
		assert.Equal(t, rowToObject(info), &entSnapshot)
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
		assert.Equal(t, rowToObject(info), &entSnapshot)
	}
}

func RandUInt64() uint64 {
	MaxUint64 := ^uint64(0)
	MaxInt64 := int64(MaxUint64 >> 1)
	randInt, err := rand.Int(rand.Reader, big.NewInt(MaxInt64))
	if err != nil {
		return 0
	}
	return randInt.Uint64()
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
