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
	retryInterval = time.Second * 3
)

type streamClient struct {
	client    npool.Manager_GrpcProxyChannelClient
	toProxy   chan *npool.ToGrpcProxy
	closeChan chan struct{}
	id        string
}

func NewStreamClient() *streamClient {
	return &streamClient{
		closeChan: make(chan struct{}),
		id:        uuid.NewString(),
		toProxy:   make(chan *npool.ToGrpcProxy),
	}
}

func (sc *streamClient) Start(ctx context.Context) {
	go func() {
		for {
			ctx, cancel := context.WithCancel(ctx)
			client, err := v1.GrpcProxyChannel(context.Background())

			if err == nil {
				sc.client = client
				go sc.recv(ctx, cancel)
				go sc.send(ctx, cancel)
				logger.Sugar().Infof("client %v successfully connected to proxy", sc.id)
				<-ctx.Done()
			}

			time.Sleep(retryInterval)
			logger.Sugar().Errorf("client %v failed to connect to cloud proxy, will retry", sc.id)
		}
	}()

	<-sc.closeChan
	if sc.client != nil {
		_ = sc.client.CloseSend()
	}
	<-ctx.Done()
}

func (sc *streamClient) Close() {
	if sc.closeChan != nil {
		sc.closeChan <- struct{}{}
	}
}

func (sc *streamClient) recv(ctx context.Context, cancel context.CancelFunc) {
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
				var errMsg string

				cc, err := grpc.NewClient(
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

func (sc *streamClient) send(ctx context.Context, cancel context.CancelFunc) {
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
