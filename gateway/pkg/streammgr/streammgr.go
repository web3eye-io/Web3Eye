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
	client  npool.Manager_GrpcProxyChannelClient
	toProxy chan *npool.ToGrpcProxy
	cg      *utils.CloseGroup
	id      string
}

func (sc *StreamClient) Start(ctx context.Context) {
	if sc.id == "" {
		sc.id = uuid.NewString()
	}
	for {
		sc.toProxy = make(chan *npool.ToGrpcProxy)
		sc.cg = &utils.CloseGroup{}

		client, err := v1.GrpcProxyChannel(context.Background())
		if err == nil {
			sc.client = client
			go sc.recv()
			go sc.send()
			logger.Sugar().Infof("client %v successfully connected to proxy", sc.id)
			sc.cg.Wait()
		}
		logger.Sugar().Errorf("client %v failed to connect to proxy, will retry", sc.id)
		time.Sleep(retryInterval)
	}
}

func (sc *StreamClient) Close() {
	sc.cg.Close()
}

func (sc *StreamClient) recv() {
	go func() {
		for {
			req, err := sc.client.Recv()
			if utils.CheckStreamErrCode(err) {
				sc.cg.Close()
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

	sc.cg.Start()
	logger.Sugar().Infof("client %v stream revc exit", sc.id)
}

func (sc *StreamClient) send() {
	go func() {
		for {
			toProxy := <-sc.toProxy
			err := sc.client.Send(toProxy)
			if utils.CheckStreamErrCode(err) {
				sc.cg.Close()
				return
			}
			if err != nil {
				continue
			}
		}
	}()

	sc.cg.Start()
	logger.Sugar().Infof("client %v stream send exit", sc.id)
}
