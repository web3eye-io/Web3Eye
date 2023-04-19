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
				config.GetConfig().Ranker.IP,
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
			config.GetConfig().Ranker.IP,
			config.GetConfig().Ranker.GrpcPort,
		)}
}

func GetToken(ctx context.Context, id string) (*nftmetaproto.Token, error) {
	info, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetToken(ctx, &nftmetaproto.GetTokenRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*nftmetaproto.Token), nil
}

func GetTokenOnly(ctx context.Context, conds *nftmetaproto.Conds) (*nftmetaproto.Token, error) {
	info, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTokenOnly(ctx, &nftmetaproto.GetTokenOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*nftmetaproto.Token), nil
}

func GetTokens(ctx context.Context, conds *nftmetaproto.Conds, offset, limit int32) ([]*nftmetaproto.Token, uint32, error) {
	var total uint32
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTokens(ctx, &nftmetaproto.GetTokensRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*nftmetaproto.Token), total, nil
}

func Search(ctx context.Context, vector []float32, limit int32) (*rankerproto.SearchTokenResponse, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.Search(ctx, &rankerproto.SearchTokenRequest{
			Vector: vector,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.(*rankerproto.SearchTokenResponse), nil
}

func CountTokens(ctx context.Context, conds *nftmetaproto.Conds) (uint32, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountTokens(ctx, &nftmetaproto.CountTokensRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, err
	}
	return infos.(uint32), nil
}
