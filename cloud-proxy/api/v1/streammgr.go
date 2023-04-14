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

type streamUnit struct {
	steamServer   npool.Manager_GrpcProxyChannelServer
	id            string
	recvConnClose chan struct{}
	sendConnClose chan struct{}
	toProxy       chan *npool.ToGrpcProxy
	fromProxy     chan *npool.FromGrpcProxy
	streamMGR     *streamMGR
}

func (p *streamUnit) streamUnitSend(wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		for {
			select {
			case info := <-p.fromProxy:
				err := p.steamServer.Send(info)
				if utils.CheckCode(err) {
					p.close()
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
				continue
			}
		}
	}()

	select {
	case <-p.sendConnClose:
		logger.Sugar().Warnf("%v: grpc proxy stream send exit", p.id)
	}
}

func (p *streamUnit) steamUnitRecv(wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		for {
			psResponse, err := p.steamServer.Recv()
			if utils.CheckCode(err) {
				p.close()
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

	select {
	case <-p.recvConnClose:
		defer logger.Sugar().Warnf("%v: grpc proxy stream recv exit", p.id)
	}
}

func (p *streamUnit) close() {
	// delete streamUnit from streamMGR
	for i, v := range p.streamMGR.streamArr {
		if v == p.id {
			p.streamMGR.streamArr = append(p.streamMGR.streamArr[:i], p.streamMGR.streamArr[i+1:]...)
			break
		}
	}

	p.streamMGR.lock.Lock()
	if _, ok := p.streamMGR.proxyTaskMap[p.id]; ok {
		delete(p.streamMGR.proxyTaskMap, p.id)
	}
	p.streamMGR.lock.Unlock()

	// close send and recv
	p.sendConnClose <- struct{}{}
	p.recvConnClose <- struct{}{}
}

type streamMGR struct {
	streamArr    []string
	streamMap    map[string]*streamUnit
	proxyTaskMap map[string]chan *npool.ToGrpcProxy
	balancer     int
	lock         sync.Mutex
}

var _streamMGR *streamMGR

func GetSteamMGR() *streamMGR {
	if _streamMGR == nil {
		_streamMGR = &streamMGR{
			streamArr:    []string{},
			streamMap:    make(map[string]*streamUnit),
			proxyTaskMap: make(map[string]chan *npool.ToGrpcProxy),
			lock:         sync.Mutex{},
		}
	}
	return _streamMGR
}

func (p *streamMGR) AddProxySteam(stream npool.Manager_GrpcProxyChannelServer) {
	id := uuid.NewString()
	lp := &streamUnit{
		steamServer:   stream,
		id:            id,
		recvConnClose: make(chan struct{}),
		sendConnClose: make(chan struct{}),
		toProxy:       make(chan *npool.ToGrpcProxy),
		fromProxy:     make(chan *npool.FromGrpcProxy),
		streamMGR:     p,
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go lp.streamUnitSend(wg)
	go lp.steamUnitRecv(wg)

	// add to list
	p.lock.Lock()
	p.streamMap[id] = lp
	p.streamArr = append(p.streamArr, id)
	p.lock.Unlock()

	wg.Wait()
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
	var streamUnit *streamUnit
	for {
		if len(p.streamArr) == 0 {
			return nil, errors.New("have no stream connnected")
		}
		p.balancer = p.balancer % len(p.streamArr)

		if id := p.streamArr[p.balancer]; len(id) > 0 {
			streamUnit = p.streamMap[id]
		}

		p.balancer++

		if streamUnit != nil {
			break
		}
	}

	if streamUnit == nil {
		return nil, errors.New("have no stream connected")
	}

	if len(req.MsgID) == 0 {
		req.MsgID = uuid.NewString()
	}

	// add task to map
	p.proxyTaskMap[req.MsgID] = make(chan *npool.ToGrpcProxy)
	defer func() {
		delete(p.proxyTaskMap, req.MsgID)
	}()

	// send req to stream send
	go func() {
		streamUnit.fromProxy <- req
	}()

	// wait response
	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	case resp := <-p.proxyTaskMap[req.MsgID]:
		return resp, nil
	}
}
