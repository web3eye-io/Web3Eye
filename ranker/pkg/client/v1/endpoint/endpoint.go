package endpoint

import (
	"context"
	"fmt"
	"time"

	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
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
		conn, err := grpc.Dial(
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

func CreateEndpoint(ctx context.Context, in *nftmetaproto.CreateEndpointRequest) (resp *nftmetaproto.CreateEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func CreateEndpoints(ctx context.Context, in *nftmetaproto.CreateEndpointsRequest) (resp *nftmetaproto.CreateEndpointsResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateEndpoint(ctx context.Context, in *nftmetaproto.UpdateEndpointRequest) (resp *nftmetaproto.UpdateEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateEndpoints(ctx context.Context, in *nftmetaproto.UpdateEndpointsRequest) (resp *nftmetaproto.UpdateEndpointsResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpoint(ctx context.Context, in *nftmetaproto.GetEndpointRequest) (resp *nftmetaproto.GetEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpointOnly(ctx context.Context, in *nftmetaproto.GetEndpointOnlyRequest) (resp *nftmetaproto.GetEndpointOnlyResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpointOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpoints(ctx context.Context, in *nftmetaproto.GetEndpointsRequest) (resp *nftmetaproto.GetEndpointsResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistEndpoint(ctx context.Context, in *nftmetaproto.ExistEndpointRequest) (resp *nftmetaproto.ExistEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistEndpointConds(ctx context.Context, in *nftmetaproto.ExistEndpointCondsRequest) (resp *nftmetaproto.ExistEndpointCondsResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistEndpointConds(ctx, in)
		return resp, err
	})
	return resp, err
}

func CountEndpoints(ctx context.Context, in *nftmetaproto.CountEndpointsRequest) (resp *nftmetaproto.CountEndpointsResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CountEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteEndpoint(ctx context.Context, in *nftmetaproto.DeleteEndpointRequest) (resp *nftmetaproto.DeleteEndpointResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}
