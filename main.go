package main

import (
	"context"
	"fmt"
	"time"

	"github.com/web3eye-io/Web3Eye/common/utils"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
)

// v1 "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
// "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/contract"

func main() {

	go func() {
		cli, err := v1.GrpcProxyChannel(context.TODO())
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < 5; i++ {
			req, err := cli.Recv()
			fmt.Println(req, err)

			cc, err := grpc.Dial(
				req.Info.TargetServer,
				grpc.WithTransportCredentials(
					insecure.NewCredentials(),
				),
				grpc.WithDefaultCallOptions(
					grpc.ForceCodec(utils.RawCodec{}),
				),
			)
			reply := &[]byte{}
			cc.Invoke(context.Background(), req.Info.Method, req.Info.RawData, reply)

			cli.Send(&npool.ToGrpcProxy{MsgID: req.MsgID, Info: &npool.GrpcInfo{
				TargetServer: req.Info.TargetServer,
				Method:       req.Info.Method,
				RawData:      *reply,
			}})
		}
	}()
	time.Sleep(time.Second)
	contract.UseCloudProxyCC()
	fmt.Println(contract.GetContracts(context.TODO(), nil, 0, 10))
	time.Sleep(time.Second * 10000)
}
