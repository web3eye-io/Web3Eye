//nolint
package main

import (
	"context"
	"fmt"

	"github.com/web3eye-io/cyber-tracer/block-etl/pkg/chains/eth"
)

var (
	animationKeywords  = []string{"animation", "video"}
	imageKeywords      = []string{"image"}
	DefaultSearchDepth = 5
)

func main() {
	// test redis
	// redisCli, err := redis.GetClient()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// err = redisCli.Set(context.Background(), "ssss", "sdfasdf", time.Hour).Err()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// redisCli.Close()

	// chains.IndexBlocks()

	num, err := eth.CurrentBlockHeight(context.Background())
	if err != nil {
		panic(err)
	}

	indexer, err := eth.NewIndexer(num - 1000)
	if err != nil {
		fmt.Println(err)
	}
	indexer.IndexTransferToDB(context.Background())

	// 	return false, err
	// })
	// ctx := context.Background()
	// uri := "https://cloneforce.xyz/api/nexus/metadata/0"
	// uri = "ipfs://QmdHkSYMgfMEA6FXx5sJo1QNePELD4jmE1cos7Yc2hzJa1/1654"
	// uri = "ipfs://bafybeidlijy646imwhiqc4uagv3lzsniubsjmwgddp7fayakupdegnf7qa/1152.json"
	// uri = "https://arweave.net/15jkI15e0Gqa5MIczNhxiyy4DwwJ-bp9o_TMk19h3U8"
	// uri = "https://arweave.net/SoRvM8DfwGH29a-5G_S5TDURkraCNZFDH_Ap1g2PwN0"
	// into := &token.TokenMetadata{}
	// fmt.Println("DecodeMetadataFromURI:")
	// fmt.Println(token.DecodeMetadataFromURI(ctx, uri, into))
	// fmt.Println("FindNameAndDescription:")
	// fmt.Println(token.FindNameAndDescription(ctx, *into))
	// fmt.Println("FindImageAndAnimationURLs:")
	// fmt.Println(token.FindImageAndAnimationURLs(ctx, *into, uri, token.AnimationKeywords, token.ImageKeywords, true))
}

// {
// 	"ChainType": "Ethereum",
// 	"ChainID": 0,
// 	"Contract": "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
// 	"TokenType": "ERC-721",
// 	"TokenID": "358226",
// 	"From": "0x0000000000000000000000000000000000000000000000000000000000000000",
// 	"To": "0x00000000000000000000000057c1e0c2adf6eecdb135bcf9ec5f23b319be2c94",
// 	"Amount": 1,
// 	"BlockNumber": 15936888,
// 	"TxHash": "0x58a48b95800fe046264f2fc6648cd6dc1eeff86cb9338c07ac9f04a95598601c",
// 	"BlockHash": "0xe9aa24c2efa047c1f8953022fffa6ac1d625d4d4c407544c51ba4ee12944f564",
// 	"TxTime": 0
//    }
