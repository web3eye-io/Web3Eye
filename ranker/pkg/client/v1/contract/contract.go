//nolint:nolintlint,dupl
package contract

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"
	"github.com/web3eye-io/Web3Eye/config"
	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/contract"
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

func GetContract(ctx context.Context, in *nftmetaproto.GetContractRequest) (resp *nftmetaproto.GetContractResponse, err error) {
	_, err = WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetContract(_ctx, in)
		return resp, err
	})
	return resp, err
}

func GetContractOnly(ctx context.Context, in *nftmetaproto.GetContractOnlyRequest) (resp *nftmetaproto.GetContractOnlyResponse, err error) {
	_, err = WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetContractOnly(_ctx, in)
		return resp, err
	})
	return resp, err
}

func GetContracts(ctx context.Context, in *nftmetaproto.GetContractsRequest) (resp *nftmetaproto.GetContractsResponse, err error) {
	_, err = WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetContracts(_ctx, in)
		return resp, err
	})
	return resp, err
}

func CountContracts(ctx context.Context, in *nftmetaproto.CountContractsRequest) (resp *nftmetaproto.CountContractsResponse, err error) {
	_, err = WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CountContracts(_ctx, in)
		return resp, err
	})
	return resp, err
}

func GetContractAndTokens(ctx context.Context, in *rankerproto.GetContractAndTokensReq) (resp *rankerproto.GetContractAndTokensResp, err error) {
	_, err = WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetContractAndTokens(_ctx, in)
		return resp, err
	})
	return resp, err
}
