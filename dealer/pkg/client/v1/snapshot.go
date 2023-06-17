//nolint:nolintlint,dupl
package v1

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

func CreateSnapshot(ctx context.Context, in *npool.CreateSnapshotRequest) (*npool.CreateSnapshotResponse, error) {
	info, err := WithCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
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
	info, err := WithCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
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
