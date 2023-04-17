package main

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/config"
	"github.com/web3eye-io/Web3Eye/gateway/pkg/streammgr"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
)

// v1 "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
// "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"

func main() {
	logger.Init(logger.DebugLevel, "./")
	config.GetConfig().CloudProxy.IP = "127.0.0.1"

	sc := &streammgr.StreamClient{}
	go sc.Start(context.Background())

	time.Sleep(time.Second)
	contract.UseCloudProxyCC()
	fmt.Println(contract.GetContracts(context.TODO(), nil, 0, 10))
	time.Sleep(time.Second * 10000)
}
