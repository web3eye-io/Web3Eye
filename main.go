package main

import (
	"context"
	"fmt"

	v1 "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
)

func main() {
	// contract.SetClientConnInterface(&contract.Po{})
	// fmt.Println(contract.GetContracts(context.TODO(), nil, 0, 10))
	cc := &v1.CloudProxyCC{}
	contract.SetClientConnInterface(cc)
	fmt.Println(contract.GetContracts(context.TODO(), nil, 5, 20))
}
