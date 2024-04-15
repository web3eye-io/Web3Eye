package endpoint

import (
	"context"
	"fmt"
	"time"

	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type handler func(context.Context, rankerproto.ManagerClient) (cruder.Any, error)

var (
	cc      grpc.ClientConnInterface = nil
	timeout                          = 6 * time.Second
)

func WithCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if cc == nil {
		conn, err := grpc.NewClient(
			fmt.Sprintf("%v:%v",
				config.GetConfig().Ranker.Domain,
				config.GetConfig().Ranker.GrpcPort),
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
	cli := rankerproto.NewManagerClient(cc)
	return handler(_ctx, cli)
}

func UseCloudProxyCC() {
	cc = &cloudproxy.CloudProxyCC{
		TargetServer: fmt.Sprintf("%v:%v",
			config.GetConfig().Ranker.Domain,
			config.GetConfig().Ranker.GrpcPort,
		)}
}

func CreateEndpoint(ctx context.Context, in *rankerproto.CreateEndpointRequest) (resp *rankerproto.CreateEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateEndpoint(ctx context.Context, in *rankerproto.UpdateEndpointRequest) (resp *rankerproto.UpdateEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpoint(ctx context.Context, in *rankerproto.GetEndpointRequest) (resp *rankerproto.GetEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpoints(ctx context.Context, in *rankerproto.GetEndpointsRequest) (resp *rankerproto.GetEndpointsResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteEndpoint(ctx context.Context, in *rankerproto.DeleteEndpointRequest) (resp *rankerproto.DeleteEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}
