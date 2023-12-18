package v1

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/utils"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

type streamServer struct {
	server    npool.Manager_GrpcProxyChannelServer
	id        string
	closeChan chan struct{}
	toProxy   chan *npool.ToGrpcProxy
	fromProxy chan *npool.FromGrpcProxy
	streamMGR *streamMGR
}

func (p *streamServer) streamServerSend(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	go func() {
		for {
			info := <-p.fromProxy
			err := p.server.Send(info)
			if utils.CheckStreamErrCode(err) {
				cancel()
				return
			}

			if err != nil {
				logger.Sugar().Errorf(
					"grpc proxy %v send msg, MsgID: %v, err: %v",
					p.id,
					info.GetMsgID(),
					err,
				)
				continue
			}

			logger.Sugar().Infof(
				"grpc proxy %v send msg success, MsgID: %v",
				p.id,
				info.GetMsgID(),
			)
		}
	}()

	// blocking
	<-ctx.Done()
	logger.Sugar().Warnf("%v: grpc proxy stream send exit", p.id)
}

func (p *streamServer) steamUnitRecv(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	go func() {
		for {
			psResponse, err := p.server.Recv()
			if utils.CheckStreamErrCode(err) {
				cancel()
				return
			}

			if err != nil {
				logger.Sugar().Errorf(
					"grpc proxy stream recv error: %v",
					err,
				)
				continue
			}
			go p.streamMGR.recvProxyResp(psResponse)
			continue
		}
	}()

	// blocking
	<-ctx.Done()
	logger.Sugar().Warnf("%v: grpc proxy stream recv exit", p.id)
}

type streamMGR struct {
	streamArr    []string
	streamMap    sync.Map
	proxyTaskMap sync.Map
	balancer     int
}

var _streamMGR *streamMGR

func GetSteamMGR() *streamMGR {
	if _streamMGR == nil {
		_streamMGR = &streamMGR{
			streamArr:    []string{},
			streamMap:    sync.Map{},
			proxyTaskMap: sync.Map{},
		}
	}
	return _streamMGR
}

func (p *streamMGR) AddProxySteam(stream npool.Manager_GrpcProxyChannelServer) {
	ctx, cancel := context.WithCancel(context.Background())

	id := uuid.NewString()
	lp := &streamServer{
		server:    stream,
		id:        id,
		closeChan: make(chan struct{}),
		toProxy:   make(chan *npool.ToGrpcProxy),
		fromProxy: make(chan *npool.FromGrpcProxy),
		streamMGR: p,
	}

	go lp.streamServerSend(ctx, cancel)
	go lp.steamUnitRecv(ctx, cancel)

	// add to list
	p.streamMap.Store(id, lp)
	p.streamArr = append(p.streamArr, id)

	logger.Sugar().Infof(" grpc proxy stream (%v) client connect successfully", lp.id)
	<-ctx.Done()

	// add to list
	for i := 0; i < len(p.streamArr); i++ {
		if p.streamArr[i] == id {
			p.streamArr = append(p.streamArr[:i], p.streamArr[i+1:]...)
		}
	}

	logger.Sugar().Infof("some grpc proxy stream (%v) client down, close it", lp.id)
}

func (p *streamMGR) recvProxyResp(psResponse *npool.ToGrpcProxy) {
	// check task and give it response
	if _recvResp, ok := p.proxyTaskMap.Load(psResponse.MsgID); ok {
		recvResp, ok := _recvResp.(chan *npool.ToGrpcProxy)
		if !ok {
			logger.Sugar().Errorf(
				"failed to parse map value to chan, msgID: %v",
				psResponse.MsgID,
			)
			return
		}

		recvDone := make(chan struct{})
		go func() {
			recvResp <- psResponse
			recvDone <- struct{}{}
		}()
		select {
		case <-time.NewTicker(time.Second).C:
			logger.Sugar().Warnf(
				"grpc proxy recv the msg %v but give it out timeout",
				psResponse.MsgID,
			)
		case <-recvDone:
		}
	} else {
		logger.Sugar().Warnf(
			"grpc proxy recv the msg %v but nobody want recv it",
			psResponse.MsgID,
		)
	}
}

func (p *streamMGR) InvokeMSG(ctx context.Context, req *npool.FromGrpcProxy) (*npool.ToGrpcProxy, error) {
	var sServer *streamServer
	for {
		if len(p.streamArr) == 0 {
			return nil, errors.New("have no stream connnected")
		}
		p.balancer %= len(p.streamArr)

		if id := p.streamArr[p.balancer]; len(id) > 0 {
			_sServer, ok := p.streamMap.Load(id)
			if !ok {
				continue
			}
			sServer, ok = _sServer.(*streamServer)
			if !ok {
				sServer = nil
				continue
			}
		}

		p.balancer++

		if sServer != nil {
			break
		}
	}

	if sServer == nil {
		return nil, errors.New("have no stream connected")
	}

	if req.MsgID == "" {
		req.MsgID = uuid.NewString()
	}

	// add task to map
	recvResp := make(chan *npool.ToGrpcProxy)
	p.proxyTaskMap.Store(req.MsgID, recvResp)
	defer func() {
		p.proxyTaskMap.Delete(req.MsgID)
	}()

	// send req to stream send
	go func() {
		sServer.fromProxy <- req
	}()

	// wait response
	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	case resp := <-recvResp:
		return resp, nil
	}
}
