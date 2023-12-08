//nolint:nolintlint,dupl
package order

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
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

func CreateOrder(ctx context.Context, in *npool.CreateOrderRequest) (resp *npool.CreateOrderResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateOrder(ctx, in)
		return resp, err
	})
	return resp, err
}

func CreateOrders(ctx context.Context, in *npool.CreateOrdersRequest) (resp *npool.CreateOrdersResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateOrders(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateOrder(ctx context.Context, in *npool.UpdateOrderRequest) (resp *npool.UpdateOrderResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateOrder(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetOrder(ctx context.Context, in *npool.GetOrderRequest) (resp *npool.GetOrderResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetOrder(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetOrderOnly(ctx context.Context, in *npool.GetOrderOnlyRequest) (resp *npool.GetOrderOnlyResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetOrderOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetOrders(ctx context.Context, in *npool.GetOrdersRequest) (resp *npool.GetOrdersResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetOrders(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistOrder(ctx context.Context, in *npool.ExistOrderRequest) (resp *npool.ExistOrderResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistOrder(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistOrderConds(ctx context.Context, in *npool.ExistOrderCondsRequest) (resp *npool.ExistOrderCondsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistOrderConds(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteOrder(ctx context.Context, in *npool.DeleteOrderRequest) (resp *npool.DeleteOrderResponse, err error) {
	ret, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteOrder(ctx, in)
		return resp, err
	})
	return ret.(*npool.DeleteOrderResponse), err
}
