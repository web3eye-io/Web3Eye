package synctask

import (
	"context"
	"fmt"
	"time"

	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/synctask"
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

func CreateSyncTask(ctx context.Context, in *rankerproto.CreateSyncTaskRequest) (resp *rankerproto.CreateSyncTaskResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateSyncTask(ctx context.Context, in *rankerproto.UpdateSyncTaskRequest) (resp *rankerproto.UpdateSyncTaskResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetSyncTask(ctx context.Context, in *rankerproto.GetSyncTaskRequest) (resp *rankerproto.GetSyncTaskResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetSyncTasks(ctx context.Context, in *rankerproto.GetSyncTasksRequest) (resp *rankerproto.GetSyncTasksResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetSyncTasks(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteSyncTask(ctx context.Context, in *rankerproto.DeleteSyncTaskRequest) (resp *rankerproto.DeleteSyncTaskResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}
