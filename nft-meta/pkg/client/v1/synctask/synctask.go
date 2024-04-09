//nolint:nolintlint,dupl
package synctask

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.NewClient(
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

func CreateSyncTask(ctx context.Context, in *npool.CreateSyncTaskRequest) (resp *npool.CreateSyncTaskResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}

func TriggerSyncTask(ctx context.Context, in *npool.TriggerSyncTaskRequest) (resp *npool.TriggerSyncTaskResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.TriggerSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateSyncTask(ctx context.Context, in *npool.UpdateSyncTaskRequest) (resp *npool.UpdateSyncTaskResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetSyncTask(ctx context.Context, in *npool.GetSyncTaskRequest) (resp *npool.GetSyncTaskResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetSyncTaskOnly(ctx context.Context, in *npool.GetSyncTaskOnlyRequest) (resp *npool.GetSyncTaskOnlyResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetSyncTaskOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetSyncTasks(ctx context.Context, in *npool.GetSyncTasksRequest) (resp *npool.GetSyncTasksResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetSyncTasks(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistSyncTask(ctx context.Context, in *npool.ExistSyncTaskRequest) (resp *npool.ExistSyncTaskResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistSyncTask(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistSyncTaskConds(ctx context.Context, in *npool.ExistSyncTaskCondsRequest) (resp *npool.ExistSyncTaskCondsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistSyncTaskConds(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteSyncTask(ctx context.Context, in *npool.DeleteSyncTaskRequest) (resp *npool.DeleteSyncTaskResponse, err error) {
	ret, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteSyncTask(ctx, in)
		return resp, err
	})
	return ret.(*npool.DeleteSyncTaskResponse), err
}
