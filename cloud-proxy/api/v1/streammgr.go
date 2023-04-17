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
	cg        *utils.CloseGroup
	toProxy   chan *npool.ToGrpcProxy
	fromProxy chan *npool.FromGrpcProxy
	streamMGR *streamMGR
}

func (p *streamServer) streamServerSend() {
	go func() {
		for {
			info := <-p.fromProxy
			err := p.server.Send(info)
			if utils.CheckStreamErrCode(err) {
				p.cg.Close()
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
	p.cg.Start()
	logger.Sugar().Warnf("%v: grpc proxy stream send exit", p.id)
}

func (p *streamServer) steamUnitRecv() {
	go func() {
		for {
			psResponse, err := p.server.Recv()
			if utils.CheckStreamErrCode(err) {
				p.cg.Close()
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
	p.cg.Start()
	logger.Sugar().Warnf("%v: grpc proxy stream recv exit", p.id)
}

type streamMGR struct {
	streamArr    []string
	streamMap    map[string]*streamServer
	proxyTaskMap map[string]chan *npool.ToGrpcProxy
	balancer     int
	lock         sync.Mutex
}

var _streamMGR *streamMGR

func GetSteamMGR() *streamMGR {
	if _streamMGR == nil {
		_streamMGR = &streamMGR{
			streamArr:    []string{},
			streamMap:    make(map[string]*streamServer),
			proxyTaskMap: make(map[string]chan *npool.ToGrpcProxy),
			lock:         sync.Mutex{},
		}
	}
	return _streamMGR
}

func (p *streamMGR) AddProxySteam(stream npool.Manager_GrpcProxyChannelServer) {
	id := uuid.NewString()
	lp := &streamServer{
		server:    stream,
		id:        id,
		cg:        &utils.CloseGroup{},
		toProxy:   make(chan *npool.ToGrpcProxy),
		fromProxy: make(chan *npool.FromGrpcProxy),
		streamMGR: p,
	}

	go lp.streamServerSend()
	go lp.steamUnitRecv()

	// add to list
	p.lock.Lock()
	p.streamMap[id] = lp
	p.streamArr = append(p.streamArr, id)
	p.lock.Unlock()

	lp.cg.Wait()
	logger.Sugar().Infof("some grpc proxy stream (%v) client down, close it", lp.id)
}

func (p *streamMGR) recvProxyResp(psResponse *npool.ToGrpcProxy) {
	// check task and give it response
	if _, ok := p.proxyTaskMap[psResponse.MsgID]; ok {
		recvDone := make(chan struct{})
		go func() {
			p.proxyTaskMap[psResponse.MsgID] <- psResponse
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
	var streamServer *streamServer
	for {
		if len(p.streamArr) == 0 {
			return nil, errors.New("have no stream connnected")
		}
		p.balancer %= len(p.streamArr)

		if id := p.streamArr[p.balancer]; len(id) > 0 {
			streamServer = p.streamMap[id]
		}

		p.balancer++

		if streamServer != nil {
			break
		}
	}

	if streamServer == nil {
		return nil, errors.New("have no stream connected")
	}

	if req.MsgID == "" {
		req.MsgID = uuid.NewString()
	}

	// add task to map
	p.proxyTaskMap[req.MsgID] = make(chan *npool.ToGrpcProxy)
	defer func() {
		delete(p.proxyTaskMap, req.MsgID)
	}()

	// send req to stream send
	go func() {
		streamServer.fromProxy <- req
	}()

	// wait response
	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	case resp := <-p.proxyTaskMap[req.MsgID]:
		return resp, nil
	}
}
