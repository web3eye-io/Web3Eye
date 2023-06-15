//nolint:nolintlint,dupl
package v1

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

func CreateBackup(ctx context.Context, in *npool.CreateBackupRequest) (*npool.CreateBackupResponse, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateBackup(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.CreateBackupResponse), nil
}
