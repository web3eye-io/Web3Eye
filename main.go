package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db"
	handler "github.com/web3eye-io/Web3Eye/nft-meta/pkg/mw/v1/block"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	blockproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
)

var ret = &blockproto.Block{
	ChainType:   basetype.ChainType_Ethereum,
	ChainID:     "test_block",
	BlockNumber: 10010,
	BlockHash:   "test_block",
	BlockTime:   time.Now().Unix(),
	ParseState:  basetype.BlockParseState_BlockTypeFinish,
	Remark:      "test_block",
}

var req = &blockproto.BlockReq{
	ChainType:   &ret.ChainType,
	ChainID:     &ret.ChainID,
	BlockNumber: &ret.BlockNumber,
	BlockHash:   &ret.BlockHash,
	BlockTime:   &ret.BlockTime,
	ParseState:  &ret.ParseState,
	Remark:      &ret.Remark,
}

func main() {
	err := db.Init()
	if err != nil {
		fmt.Printf("cannot init database: %v \n", err)
		os.Exit(0)
	}

	h, err := handler.NewHandler(
		context.Background(),
		handler.WithChainType(req.ChainType, true),
		handler.WithChainID(req.ChainID, true),
		handler.WithBlockNumber(req.BlockNumber, true),
		handler.WithBlockHash(req.BlockHash, true),
		handler.WithBlockTime(req.BlockTime, true),
		handler.WithParseState(req.ParseState, true),
		handler.WithRemark(req.Remark, true),
	)
	fmt.Println(err)
	info, err := h.CreateBlock(context.Background())
	fmt.Println(err)
	fmt.Println(info)
}
