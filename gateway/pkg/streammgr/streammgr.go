package streammgr

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	v1 "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
	"github.com/web3eye-io/Web3Eye/common/utils"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	reqTimeout    = time.Second * 3
	retryInterval = time.Second * 3
)

type StreamClient struct {
	client    npool.Manager_GrpcProxyChannelClient
	toProxy   chan *npool.ToGrpcProxy
	closeChan chan struct{}
	id        string
}

func (sc *StreamClient) Start(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	if sc.id == "" {
		sc.id = uuid.NewString()
	}

	sc.closeChan = make(chan struct{})
	go func() {
		for {
			sc.toProxy = make(chan *npool.ToGrpcProxy)

			ctx, cancel := context.WithCancel(ctx)

			client, err := v1.GrpcProxyChannel(context.Background())

			if err == nil {
				sc.client = client
				go sc.recv(ctx, cancel)
				go sc.send(ctx, cancel)
				logger.Sugar().Infof("client %v successfully connected to proxy", sc.id)
				<-ctx.Done()
			}
			logger.Sugar().Errorf("client %v failed to connect to proxy, will retry", sc.id)
			time.Sleep(retryInterval)
		}
	}()

	select {
	case <-sc.closeChan:
		if sc.client != nil {
			sc.client.CloseSend()
		}
	case <-ctx.Done():
		if sc.client != nil {
			sc.client.CloseSend()
		}
		time.Sleep(time.Second * 5)
	}

}

func (sc *StreamClient) Close() {
	if sc.closeChan != nil {
		sc.closeChan <- struct{}{}
	}
}

func (sc *StreamClient) recv(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	go func() {
		for {
			req, err := sc.client.Recv()
			if utils.CheckStreamErrCode(err) {
				cancel()
				return
			}
			if err != nil {
				continue
			}
			logger.Sugar().Infof("client %v stream revc msg %v", sc.id, req.MsgID)

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
					err = cc.Invoke(context.Background(), req.Info.Method, req.Info.RawData, reply)
					if err != nil {
						errMsg = err.Error()
					}
					respInfo = &npool.GrpcInfo{TargetServer: req.Info.TargetServer, Method: req.Info.Method, RawData: *reply}
				}

				sc.toProxy <- &npool.ToGrpcProxy{MsgID: req.MsgID, Info: respInfo, ErrMsg: errMsg}
			}(req)
		}
	}()

	<-ctx.Done()
	logger.Sugar().Infof("client %v stream revc exit", sc.id)
}

func (sc *StreamClient) send(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	go func() {
		for {
			toProxy := <-sc.toProxy
			err := sc.client.Send(toProxy)
			if utils.CheckStreamErrCode(err) {
				cancel()
				return
			}
			if err != nil {
				continue
			}
			logger.Sugar().Infof("client %v stream send msg %v", sc.id, toProxy.MsgID)
		}
	}()

	<-ctx.Done()
	logger.Sugar().Infof("client %v stream send exit", sc.id)
}
