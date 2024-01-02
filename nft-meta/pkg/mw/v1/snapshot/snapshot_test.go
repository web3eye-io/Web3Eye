package snapshot

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"

	"github.com/stretchr/testify/assert"

	snapshotproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &snapshotproto.Snapshot{
	Index:         10010,
	SnapshotCommP: "test_snapshot",
	SnapshotRoot:  "test_snapshot",
	SnapshotURI:   "test_snapshot",
	BackupState:   "test_snapshot",
}

var req = &snapshotproto.SnapshotReq{
	Index:         &ret.Index,
	SnapshotCommP: &ret.SnapshotCommP,
	SnapshotRoot:  &ret.SnapshotRoot,
	SnapshotURI:   &ret.SnapshotURI,
	BackupState:   &ret.BackupState,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithIndex(req.Index, true),
		WithSnapshotCommP(req.SnapshotCommP, true),
		WithSnapshotRoot(req.SnapshotRoot, true),
		WithSnapshotURI(req.SnapshotURI, true),
		WithBackupState(req.BackupState, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateSnapshot(context.Background())
	if assert.Nil(t, err) {
		req.ID = &info.ID
		req.EntID = &info.EntID
		assert.Equal(t, info.SnapshotURI, req.GetSnapshotURI())
	}
}

func update(t *testing.T) {
	index := uint64(10086)
	req.Index = &index

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithIndex(req.Index, false),
		WithSnapshotURI(req.SnapshotURI, false),
		WithBackupState(req.BackupState, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateSnapshot(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info.Index, req.GetIndex())
	}
}

func query(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSnapshot(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&snapshotproto.Conds{
			Index: &web3eye.Uint64Val{
				Op:    cruder.EQ,
				Value: *req.Index,
			},
			SnapshotCommP: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.SnapshotCommP,
			},
			BackupState: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.BackupState,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetSnapshots(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))
	assert.Equal(t, infos[0], info)

	handler, err = NewHandler(context.Background(),
		WithID(&infos[0].ID, true),
	)
	assert.Nil(t, err)

	_info1, err := handler.DeleteSnapshot(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, infos[0].ID, _info1.ID)
}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	err := db.Init()
	if err != nil {
		fmt.Printf("cannot init database: %v \n", err)
		os.Exit(0)
	}

	t.Run("create", create)
	t.Run("update", update)
	t.Run("query", query)
}
