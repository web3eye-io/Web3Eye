//nolint:nolintlint,dupl
package token

import (
	"context"
	"fmt"
	"time"

	cloudproxy "github.com/web3eye-io/Web3Eye/cloud-proxy/pkg/client/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	nftmetaproto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	rankerproto "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type handler func(context.Context, rankerproto.ManagerClient) (cruder.Any, error)

var (
	cc      grpc.ClientConnInterface = nil
	timeout                          = 12 * time.Second
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

func GetToken(ctx context.Context, in *nftmetaproto.GetTokenRequest) (resp *nftmetaproto.GetTokenResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetToken(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTokenOnly(ctx context.Context, in *nftmetaproto.GetTokenOnlyRequest) (resp *nftmetaproto.GetTokenOnlyResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTokenOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTokens(ctx context.Context, in *nftmetaproto.GetTokensRequest) (resp *nftmetaproto.GetTokensResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTokens(ctx, in)
		return resp, err
	})
	return resp, err
}

func Search(ctx context.Context, in *rankerproto.SearchTokenRequest) (resp *rankerproto.SearchResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.Search(ctx, in)
		return resp, err
	})
	return resp, err
}

func SearchPage(ctx context.Context, in *rankerproto.SearchPageRequest) (resp *rankerproto.SearchResponse, err error) {
	_, err = WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.SearchPage(ctx, in)
		return resp, err
	})
	return resp, err
}
