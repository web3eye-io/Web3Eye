package synctask

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	"github.com/stretchr/testify/assert"

	synctaskproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &synctaskproto.SyncTask{
	ChainType:   basetype.ChainType_Ethereum,
	ChainID:     "test_synctask",
	Start:       10010,
	End:         10086,
	Current:     10011,
	Topic:       "test_synctask",
	SyncState:   basetype.SyncState_Default,
	Description: "test_synctask",
	Remark:      "test_synctask",
}

var req = &synctaskproto.SyncTaskReq{
	ChainType:   &ret.ChainType,
	ChainID:     &ret.ChainID,
	Start:       &ret.Start,
	End:         &ret.End,
	Current:     &ret.Current,
	Topic:       &ret.Topic,
	SyncState:   &ret.SyncState,
	Description: &ret.Description,
	Remark:      &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithStart(req.Start, true),
		WithEnd(req.End, true),
		WithCurrent(req.Current, true),
		WithTopic(req.Topic, true),
		WithSyncState(req.SyncState, true),
		WithDescription(req.Description, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateSyncTask(context.Background())
	if assert.Nil(t, err) {
		req.ID = &info.ID
		req.EntID = &info.EntID
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
		assert.Equal(t, info.SyncStateStr, req.SyncState.String())
	}
}

func update(t *testing.T) {
	req.ChainType = basetype.ChainType_Solana.Enum()
	current := uint64(10020)
	req.Current = &current
	topic := "ssssss"
	req.Topic = &topic

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithChainType(req.ChainType, false),
		WithChainID(req.ChainID, false),
		WithCurrent(req.Current, false),
		WithTopic(req.Topic, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateSyncTask(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
		fmt.Println(info.Current, *req.Current)
		assert.Equal(t, info.Current, *req.Current)
		assert.Equal(t, info.Topic, req.GetTopic())
	}
}

func query(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSyncTask(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&synctaskproto.Conds{
			ChainType: &web3eye.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(*req.ChainType),
			},
			ChainID: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.ChainID,
			},
			Topic: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.Topic,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetSyncTasks(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))
	assert.Equal(t, infos[0], info)

	handler, err = NewHandler(context.Background(),
		WithID(&infos[0].ID, true),
	)
	assert.Nil(t, err)

	_info1, err := handler.DeleteSyncTask(context.Background())
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
