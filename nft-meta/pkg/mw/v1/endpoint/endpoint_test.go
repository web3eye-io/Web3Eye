package endpoint

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

	endpointproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &endpointproto.Endpoint{
	ChainType: basetype.ChainType_Ethereum,
	ChainID:   "test_endpoint",
	Address:   "test_endpoint",
	State:     basetype.EndpointState_EndpointDefault,
	Remark:    "test_endpoint",
}

var req = &endpointproto.EndpointReq{
	ChainType: &ret.ChainType,
	ChainID:   &ret.ChainID,
	Address:   &ret.Address,
	State:     &ret.State,
	Remark:    &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithAddress(req.Address, true),
		WithState(req.State, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateEndpoint(context.Background())
	if assert.Nil(t, err) {
		req.ID = &info.ID
		req.EntID = &info.EntID
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
	}
}

func update(t *testing.T) {
	req.ChainType = basetype.ChainType_Solana.Enum()
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithChainType(req.ChainType, false),
		WithChainID(req.ChainID, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateEndpoint(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
	}
}

func query(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetEndpoint(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&endpointproto.Conds{
			ChainType: &web3eye.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(*req.ChainType),
			},
			ChainID: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.ChainID,
			},
			Address: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.Address,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetEndpoints(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))
	assert.Equal(t, infos[0], info)

	handler, err = NewHandler(context.Background(),
		WithID(&infos[0].ID, true),
	)
	assert.Nil(t, err)

	_info1, err := handler.DeleteEndpoint(context.Background())
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
