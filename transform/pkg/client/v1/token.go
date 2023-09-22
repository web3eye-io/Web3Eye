package v1

import (
	"context"
	"fmt"
	"time"

	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	transformproto "github.com/web3eye-io/Web3Eye/proto/web3eye/transform/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type handler func(context.Context, transformproto.ManagerClient) (cruder.Any, error)

var (
	cc      grpc.ClientConnInterface = nil
	timeout                          = 6 * time.Second
)

func WithCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if cc == nil {
		conn, err := grpc.Dial(
			fmt.Sprintf("%v:%v",
				config.GetConfig().Transform.Domain,
				config.GetConfig().Transform.GrpcPort),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		cc = conn

		defer func() {
			conn.Close()
			cc = nil
		}()
	}
	cli := transformproto.NewManagerClient(cc)
	return handler(_ctx, cli)
}

func UseCloudProxyCC() {
	cc = &cloudproxy.CloudProxyCC{
		TargetServer: fmt.Sprintf("%v:%v",
			config.GetConfig().Transform.Domain,
			config.GetConfig().Transform.GrpcPort,
		)}
}

func UrlToVector(ctx context.Context, in *transformproto.UrlToVectorReq) (resp *transformproto.UrlToVectorResp, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli transformproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UrlToVector(ctx, in)
		return resp, err
	})
	return resp, err
}
