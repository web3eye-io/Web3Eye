//nolint:nolintlint,dupl
package v1

import (
	"context"
	"fmt"
	"time"

	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

var (
	cc      grpc.ClientConnInterface = nil
	timeout                          = 6 * time.Second
)

func WithCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if cc == nil {
		conn, err := grpc.NewClient(
			fmt.Sprintf("%v:%v",
				config.GetConfig().Dealer.Domain,
				config.GetConfig().Dealer.GrpcPort),
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
	cli := npool.NewManagerClient(cc)
	return handler(_ctx, cli)
}

func UseCloudProxyCC() {
	cc = &cloudproxy.CloudProxyCC{
		TargetServer: fmt.Sprintf("%v:%v",
			config.GetConfig().Dealer.Domain,
			config.GetConfig().Dealer.GrpcPort,
		)}
}

func CreateBackup(ctx context.Context, in *npool.CreateBackupRequest) (*npool.CreateBackupResponse, error) {
	info, err := WithCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
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
