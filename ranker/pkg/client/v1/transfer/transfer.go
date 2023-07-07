//nolint:nolintlint,dupl
package transfer

import (
	"context"
	"fmt"
	"time"

	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
	"github.com/web3eye-io/Web3Eye/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
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

func GetTransfer(ctx context.Context, in *nftmetaproto.GetTransferRequest) (resp *nftmetaproto.GetTransferResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTransfer(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTransferOnly(ctx context.Context, in *nftmetaproto.GetTransferOnlyRequest) (resp *nftmetaproto.GetTransferOnlyResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTransferOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTransfers(ctx context.Context, in *rankerproto.GetTransfersRequest) (resp *nftmetaproto.GetTransfersResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTransfers(ctx, in)
		return resp, err
	})
	return resp, err
}

func CountTransfers(ctx context.Context, in *rankerproto.CountTransfersRequest) (resp *nftmetaproto.CountTransfersResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CountTransfers(ctx, in)
		return resp, err
	})
	return resp, err
}
