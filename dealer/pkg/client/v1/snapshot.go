//nolint:nolintlint,dupl
package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v",
			config.GetConfig().Dealer.IP,
			config.GetConfig().Dealer.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateSnapshot(ctx context.Context, in *npool.CreateSnapshotRequest) (*npool.CreateSnapshotResponse, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateSnapshot(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.CreateSnapshotResponse), nil
}

func GetSnapshots(ctx context.Context, in *npool.GetSnapshotsRequest) (*npool.GetSnapshotsResponse, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSnapshots(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.GetSnapshotsResponse), nil
}
