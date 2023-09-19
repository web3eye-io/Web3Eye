//nolint:nolintlint,dupl
package endpoint

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var timeout = 10 * time.Second

type handlerFunc func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handlerFunc) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v",
			config.GetConfig().NFTMeta.Domain,
			config.GetConfig().NFTMeta.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateEndpoint(ctx context.Context, in *npool.CreateEndpointRequest) (resp *npool.CreateEndpointResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func CreateEndpoints(ctx context.Context, in *npool.CreateEndpointsRequest) (resp *npool.CreateEndpointsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateEndpoint(ctx context.Context, in *npool.UpdateEndpointRequest) (resp *npool.UpdateEndpointResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateEndpoints(ctx context.Context, in *npool.UpdateEndpointsRequest) (resp *npool.UpdateEndpointsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpoint(ctx context.Context, in *npool.GetEndpointRequest) (resp *npool.GetEndpointResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpointOnly(ctx context.Context, in *npool.GetEndpointOnlyRequest) (resp *npool.GetEndpointOnlyResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpointOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpoints(ctx context.Context, in *npool.GetEndpointsRequest) (resp *npool.GetEndpointsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistEndpoint(ctx context.Context, in *npool.ExistEndpointRequest) (resp *npool.ExistEndpointResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistEndpointConds(ctx context.Context, in *npool.ExistEndpointCondsRequest) (resp *npool.ExistEndpointCondsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistEndpointConds(ctx, in)
		return resp, err
	})
	return resp, err
}

func CountEndpoints(ctx context.Context, in *npool.CountEndpointsRequest) (resp *npool.CountEndpointsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CountEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteEndpoint(ctx context.Context, in *npool.DeleteEndpointRequest) (resp *npool.DeleteEndpointResponse, err error) {
	ret, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteEndpoint(ctx, in)
		return resp, err
	})
	return ret.(*npool.DeleteEndpointResponse), err
}
