package token

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

	tokenproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &tokenproto.Token{
	ChainType:       basetype.ChainType_Ethereum,
	ChainID:         "test_token",
	Contract:        "test_token",
	TokenType:       basetype.TokenType_ERC20,
	TokenID:         "test_token",
	Owner:           "test_token",
	URI:             "test_token",
	URIType:         "test_token",
	ImageURL:        "test_token",
	VideoURL:        "test_token",
	Name:            "test_token",
	Description:     "test_token",
	VectorState:     tokenproto.ConvertState_Default,
	VectorID:        10010,
	IPFSImageURL:    "test_token",
	ImageSnapshotID: 11111,
	Remark:          "test_token",
}

var req = &tokenproto.TokenReq{
	ChainType:       &ret.ChainType,
	ChainID:         &ret.ChainID,
	Contract:        &ret.Contract,
	TokenType:       &ret.TokenType,
	TokenID:         &ret.TokenID,
	Owner:           &ret.Owner,
	URI:             &ret.URI,
	URIType:         &ret.URIType,
	ImageURL:        &ret.ImageURL,
	VideoURL:        &ret.VideoURL,
	Name:            &ret.Name,
	Description:     &ret.Description,
	VectorState:     &ret.VectorState,
	VectorID:        &ret.VectorID,
	IPFSImageURL:    &ret.IPFSImageURL,
	ImageSnapshotID: &ret.ImageSnapshotID,
	Remark:          &ret.Remark,
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithContract(req.Contract, true),
		WithTokenType(req.TokenType, true),
		WithTokenID(req.TokenID, true),
		WithOwner(req.Owner, true),
		WithURI(req.URI, true),
		WithURIType(req.URIType, true),
		WithImageURL(req.ImageURL, true),
		WithVideoURL(req.VideoURL, true),
		WithName(req.Name, true),
		WithDescription(req.Description, true),
		WithVectorState(req.VectorState, true),
		WithVectorID(req.VectorID, true),
		WithIPFSImageURL(req.IPFSImageURL, true),
		WithImageSnapshotID(req.ImageSnapshotID, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateToken(context.Background())
	if assert.Nil(t, err) {
		req.ID = &info.ID
		req.EntID = &info.EntID
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
	}
}

func update(t *testing.T) {
	req.ChainType = basetype.ChainType_Solana.Enum()
	url := "google"
	req.IPFSImageURL = &url

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithChainType(req.ChainType, false),
		WithChainID(req.ChainID, false),
		WithIPFSImageURL(req.IPFSImageURL, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateToken(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info.IPFSImageURL, req.GetIPFSImageURL())
		assert.Equal(t, info.ChainTypeStr, req.ChainType.String())
		assert.Equal(t, info.TokenTypeStr, req.TokenType.String())
		assert.Equal(t, info.VectorStateStr, req.VectorState.String())
	}
}

func upsert(t *testing.T) {
	// just update
	url := "sssssss"
	handler, err := NewHandler(context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithContract(req.Contract, true),
		WithTokenID(req.TokenID, true),
		WithIPFSImageURL(&url, true),
		WithDescription(req.Description, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)

	info1, err := handler.UpsertToken(context.Background())
	if assert.Nil(t, err) {
		assert.NotNil(t, info1)
		assert.Equal(t, info1.IPFSImageURL, url)
		req.IPFSImageURL = &url
	}

	// create new record
	tokenid := "sssssss"
	handler, err = NewHandler(context.Background(),
		WithChainType(req.ChainType, true),
		WithChainID(req.ChainID, true),
		WithContract(req.Contract, true),
		WithTokenType(req.TokenType, true),
		WithTokenID(&tokenid, true),
		WithOwner(req.Owner, true),
		WithURI(req.URI, true),
		WithURIType(req.URIType, true),
		WithImageURL(req.ImageURL, true),
		WithVideoURL(req.VideoURL, true),
		WithName(req.Name, true),
		WithDescription(req.Description, true),
		WithVectorState(req.VectorState, true),
		WithVectorID(req.VectorID, true),
		WithIPFSImageURL(req.IPFSImageURL, true),
		WithImageSnapshotID(req.ImageSnapshotID, true),
		WithRemark(req.Remark, true),
	)
	assert.Nil(t, err)
	info2, err := handler.UpsertToken(context.Background())
	assert.Nil(t, err)
	assert.NotEqual(t, req.TokenID, info2.TokenID)
	assert.NotEqual(t, req.ID, info2.ID)
	assert.Equal(t, req.GetIPFSImageURL(), info2.IPFSImageURL)
	assert.Equal(t, info1.ChainID, info2.ChainID)

	handler, err = NewHandler(context.Background(),
		WithID(&info2.ID, true),
	)
	assert.Nil(t, err)
	_info2, err := handler.DeleteToken(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, _info2.ID, info2.ID)
}

func query(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetToken(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&tokenproto.Conds{
			Contract: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: req.GetContract(),
			},
			TokenID: &web3eye.StringVal{
				Op:    cruder.EQ,
				Value: *req.TokenID,
			},
		}),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetTokens(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, total, uint32(1))
	assert.Equal(t, infos[0], info)

	handler, err = NewHandler(context.Background(),
		WithID(&infos[0].ID, true),
	)
	assert.Nil(t, err)

	_info1, err := handler.DeleteToken(context.Background())
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
	t.Run("upsert", upsert)
	t.Run("query", query)
}
