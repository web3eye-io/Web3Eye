package main

import (
	"context"
	"fmt"
	"os"

	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
)

func main() {

	cli, err := milvusdb.Client(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = cli.DropCollection(context.Background(), config.GetConfig().NFTMeta.CollectionName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = milvusdb.Init(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
