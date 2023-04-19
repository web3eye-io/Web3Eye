package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/web3eye-io/Web3Eye/config"
// 	"github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/token"
// )

// func main() {
// 	config.GetConfig().Ranker.IP = "172.20.235.201"
// 	token.UseCloudProxyCC()
// 	start := time.Now()
// 	fmt.Println(token.Search(context.Background(), []float32{1}, 10))
// 	fmt.Println(token.CountTokens(context.Background(), nil))
// 	fmt.Println(time.Now().UnixMilli() - start.UnixMilli())
// }
