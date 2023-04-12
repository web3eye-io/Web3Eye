//nolint:nolintlint,dupl
package v1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

type CloudProxyCC struct{}

func (p *CloudProxyCC) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	reqRaw, err := proto.Marshal(args.(proto.Message))
	if err != nil {
		return err
	}

	msgID := uuid.NewString()
	proxyResp, err := Proxy(ctx, &npool.ProxyRequest{MsgID: msgID, Msg: reqRaw})
	if err != nil {
		return err
	}

	if proxyResp.MsgID != msgID {
		return fmt.Errorf("msg_id wrong, expect %v but get %v", msgID, proxyResp.MsgID)
	}
	fmt.Println(proxyResp.MsgID)
	fmt.Println(msgID)
	fmt.Println(proxyResp.Msg)

	return proto.Unmarshal(proxyResp.Msg, reply.(proto.Message))
}

func (p *CloudProxyCC) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	// not impl
	return nil, errors.New("CloudProxyCC.NewStream not implementation")
}

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v",
			"127.0.0.1",
			// config.GetConfig().CloudProxy.IP,
			config.GetConfig().CloudProxy.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func ProxyChannel(ctx context.Context) (npool.Manager_ProxyChannelClient, error) {
	var resp npool.Manager_ProxyChannelClient = nil
	var err error = nil
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ProxyChannel(ctx)
		return resp, err
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Proxy(ctx context.Context, in *npool.ProxyRequest) (*npool.ProxyResponse, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.Proxy(_ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.ProxyResponse), nil
}
