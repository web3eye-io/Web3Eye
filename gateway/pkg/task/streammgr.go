package task

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	v1 "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
	"github.com/web3eye-io/Web3Eye/common/utils"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func startChannel() {
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
}

const (
	reqTimeout = time.Second * 3
)

type streamClient struct {
	client    npool.Manager_GrpcProxyChannelClient
	toProxy   chan *npool.ToGrpcProxy
	closeSend chan struct{}
	closeRecv chan struct{}
}

func (sc *streamClient) Start() {

}

func (sc *streamClient) recv(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	go func() {
		for {
			req, err := sc.client.Recv()
			if utils.CheckStreamErrCode(err) {
				sc.close()
				return
			}
			if err != nil {
				continue
			}
			go func(req *npool.FromGrpcProxy) {
				var respInfo *npool.GrpcInfo = nil
				var errMsg string = ""

				cc, err := grpc.Dial(
					req.Info.TargetServer,
					grpc.WithTransportCredentials(
						insecure.NewCredentials(),
					),
					grpc.WithDefaultCallOptions(
						grpc.ForceCodec(utils.RawCodec{}),
					),
				)
				if err != nil {
					errMsg = err.Error()
				} else {
					reply := &[]byte{}
					cc.Invoke(context.Background(), req.Info.Method, req.Info.RawData, reply)
					respInfo = &npool.GrpcInfo{TargetServer: req.Info.TargetServer, Method: req.Info.Method, RawData: *reply}
				}

				sc.toProxy <- &npool.ToGrpcProxy{MsgID: req.MsgID, Info: respInfo, ErrMsg: errMsg}
			}(req)
		}
	}()
	<-sc.closeRecv
}

func (sc *streamClient) send(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	go func() {
		for {
			toProxy := <-sc.toProxy
			err := sc.client.Send(toProxy)
			if utils.CheckStreamErrCode(err) {
				sc.close()
				return
			}
			if err != nil {
				continue
			}
		}
	}()
	<-sc.closeSend
}

func (sc *streamClient) close() {
	sc.closeRecv <- struct{}{}
	sc.closeSend <- struct{}{}
	logger.Sugar().Warn("a stream client close")
}
