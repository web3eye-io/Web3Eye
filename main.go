package main

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
)

func main() {
	contract.SetClientConnInterface(&contract.Po{})

	fmt.Println(contract.GetContracts(context.TODO(), nil, 0, 10))
}
