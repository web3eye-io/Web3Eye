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

func GetToken(ctx context.Context, in *nftmetaproto.GetTokenRequest) (*nftmetaproto.GetTokenResponse, error) {
	info, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetToken(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*nftmetaproto.GetTokenResponse), nil
}

func GetTokenOnly(ctx context.Context, in *nftmetaproto.GetTokenOnlyRequest) (*nftmetaproto.GetTokenOnlyResponse, error) {
	info, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTokenOnly(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*nftmetaproto.GetTokenOnlyResponse), nil
}

func GetTokens(ctx context.Context, in *nftmetaproto.GetTokensRequest) (*nftmetaproto.GetTokensResponse, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTokens(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.(*nftmetaproto.GetTokensResponse), nil
}

func Search(ctx context.Context, in *rankerproto.SearchTokenRequest) (*rankerproto.SearchResponse, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.Search(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.(*rankerproto.SearchResponse), nil
}

func SearchPage(ctx context.Context, in *rankerproto.SearchPageRequest) (*rankerproto.SearchResponse, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.SearchPage(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.(*rankerproto.SearchResponse), nil
}

func CountTokens(ctx context.Context, in *nftmetaproto.CountTokensRequest) (*nftmetaproto.CountTokensResponse, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountTokens(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.(*nftmetaproto.CountTokensResponse), nil
}
