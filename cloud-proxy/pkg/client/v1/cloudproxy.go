//nolint:nolintlint,dupl
package v1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

var timeout = 60 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

type CloudProxyCC struct {
	TargetServer string
}

func (p *CloudProxyCC) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	reqRaw, err := proto.Marshal(args.(proto.Message))
	if err != nil {
		return err
	}

	msgID := uuid.NewString()
	proxyResp, err := GrpcProxy(ctx,
		&npool.GrpcProxyRequest{
			MsgID: msgID,
			Info: &npool.GrpcInfo{
				TargetServer: p.TargetServer,
				Method:       method,
				RawData:      reqRaw,
			},
		})

	if err != nil {
		return err
	}

	if proxyResp.MsgID != msgID {
		return fmt.Errorf("msg_id wrong, expect %v but get %v", msgID, proxyResp.MsgID)
	}

	if proxyResp.ErrMsg != "" {
		return errors.New(proxyResp.ErrMsg)
	}

	return proto.Unmarshal(proxyResp.Info.RawData, reply.(proto.Message))
}

func (p *CloudProxyCC) NewStream(
	ctx context.Context,
	desc *grpc.StreamDesc,
	method string,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {
	// not impl
	return nil, errors.New("CloudProxyCC.NewStream not implementation")
}

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v",
			config.GetConfig().CloudProxy.Domain,
			config.GetConfig().CloudProxy.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func withNoConnClose(ctx context.Context, handler handler) (cruder.Any, error) {
	addr := fmt.Sprintf("%v:%v",
		config.GetConfig().CloudProxy.Domain,
		config.GetConfig().CloudProxy.GrpcPort)
	logger.Sugar().Infow("withNoConnClose", "action", "prepare to connect to cloudproxy", "address", addr)

	conn, err := grpc.Dial(
		addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Sugar().Errorw("withNoConnClose", "action", "failed to connect to cloudproxy", "address", addr)
		return nil, err
	}
	logger.Sugar().Infow("withNoConnClose", "action", "success to connect to cloudproxy", "address", addr)
	cli := npool.NewManagerClient(conn)

	return handler(ctx, cli)
}

func GrpcProxyChannel(ctx context.Context) (resp npool.Manager_GrpcProxyChannelClient, err error) {
	_, err = withNoConnClose(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GrpcProxyChannel(_ctx)
		return resp, err
	})
	return
}

func GrpcProxy(ctx context.Context, in *npool.GrpcProxyRequest) (*npool.GrpcProxyResponse, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GrpcProxy(_ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.GrpcProxyResponse), nil
}
