//nolint:nolintlint,dupl
package snapshot

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
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

func CreateSnapshot(ctx context.Context, in *npool.CreateSnapshotRequest) (resp *npool.CreateSnapshotResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateSnapshot(ctx, in)
		return resp, err
	})
	return resp, err
}

func CreateSnapshots(ctx context.Context, in *npool.CreateSnapshotsRequest) (resp *npool.CreateSnapshotsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateSnapshots(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateSnapshot(ctx context.Context, in *npool.UpdateSnapshotRequest) (resp *npool.UpdateSnapshotResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateSnapshot(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetSnapshot(ctx context.Context, in *npool.GetSnapshotRequest) (resp *npool.GetSnapshotResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetSnapshot(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetSnapshotOnly(ctx context.Context, in *npool.GetSnapshotOnlyRequest) (resp *npool.GetSnapshotOnlyResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetSnapshotOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetSnapshots(ctx context.Context, in *npool.GetSnapshotsRequest) (resp *npool.GetSnapshotsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetSnapshots(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistSnapshot(ctx context.Context, in *npool.ExistSnapshotRequest) (resp *npool.ExistSnapshotResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistSnapshot(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistSnapshotConds(ctx context.Context, in *npool.ExistSnapshotCondsRequest) (resp *npool.ExistSnapshotCondsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistSnapshotConds(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteSnapshot(ctx context.Context, in *npool.DeleteSnapshotRequest) (resp *npool.DeleteSnapshotResponse, err error) {
	ret, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteSnapshot(ctx, in)
		return resp, err
	})
	return ret.(*npool.DeleteSnapshotResponse), err
}
