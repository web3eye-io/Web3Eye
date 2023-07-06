//nolint:nolintlint,dupl
package snapshot

import (
	"context"
	"fmt"
	"time"

	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/snapshot"
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

func GetSnapshot(ctx context.Context, in *nftmetaproto.GetSnapshotRequest) (*nftmetaproto.GetSnapshotResponse, error) {
	info, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSnapshot(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*nftmetaproto.GetSnapshotResponse), nil
}

func GetSnapshotOnly(ctx context.Context, in *nftmetaproto.GetSnapshotOnlyRequest) (*nftmetaproto.GetSnapshotOnlyResponse, error) {
	info, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSnapshotOnly(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*nftmetaproto.GetSnapshotOnlyResponse), nil
}

func GetSnapshots(ctx context.Context, in *nftmetaproto.GetSnapshotsRequest) (*nftmetaproto.GetSnapshotsResponse, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSnapshots(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.(*nftmetaproto.GetSnapshotsResponse), nil
}

func CountSnapshots(ctx context.Context, in *nftmetaproto.CountSnapshotsRequest) (*nftmetaproto.CountSnapshotsResponse, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountSnapshots(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.(*nftmetaproto.CountSnapshotsResponse), nil
}
