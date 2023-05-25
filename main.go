package main

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
)

func main() {
	cli, err := milvusdb.Client(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cli.HasCollection(context.Background(), milvusdb.NFTSchema.CollectionName))
	fmt.Println(cli.DropCollection(context.Background(), milvusdb.NFTSchema.CollectionName))
	fmt.Println(cli.HasCollection(context.Background(), milvusdb.NFTSchema.CollectionName))
}
