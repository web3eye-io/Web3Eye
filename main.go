package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"

	v1 "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
)

// v1 "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
// "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"

func main() {
	// contract.SetClientConnInterface(&contract.Po{})
	// fmt.Println(contract.GetContracts(context.TODO(), nil, 0, 10))
	go func() {
		cli, err := v1.GrpcProxyChannel(context.TODO())
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < 5; i++ {
			req, err := cli.Recv()
			fmt.Println(req, err)

			// cc, err := grpc.Dial(
			// 	"nft-meta:30101",
			// 	grpc.WithTransportCredentials(
			// 		insecure.NewCredentials(),
			// 	),
			// 	grpc.WithDefaultCallOptions(
			// 		grpc.ForceCodec(utils.RawCodec{}),
			// 	),
			// )
			// cc.Invoke(context.Background())

			cli.Send(&npool.ToGrpcProxy{MsgID: req.MsgID, Method: "ssss", RespRaw: []byte(uuid.NewString())})
		}
	}()
	time.Sleep(time.Second)

	fmt.Println(v1.GrpcProxy(context.Background(), &npool.GrpcProxyRequest{MsgID: "123"}))
	fmt.Println(v1.GrpcProxy(context.Background(), &npool.GrpcProxyRequest{MsgID: "345"}))
	time.Sleep(time.Second * 10000)
}
