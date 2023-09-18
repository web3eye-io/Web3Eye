//nolint:nolintlint,dupl
package token

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var timeout = 10 * time.Second

type handlerFunc func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handlerFunc) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v",
			config.GetConfig().NFTMeta.Domain,
			config.GetConfig().NFTMeta.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func LookupHost() (addrs []string, err error) {
	return net.LookupHost(config.GetConfig().NFTMeta.Domain)
}

func CreateToken(ctx context.Context, in *npool.CreateTokenRequest) (resp *npool.CreateTokenResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateToken(ctx, in)
		return resp, err
	})
	return resp, err
}

func CreateTokens(ctx context.Context, in *npool.CreateTokensRequest) (resp *npool.CreateTokensResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateTokens(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpsertToken(ctx context.Context, in *npool.UpsertTokenRequest) (resp *npool.UpsertTokenResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpsertToken(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateToken(ctx context.Context, in *npool.UpdateTokenRequest) (resp *npool.UpdateTokenResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateToken(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateImageVector(ctx context.Context, in *npool.UpdateImageVectorRequest) (resp *npool.UpdateImageVectorResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateImageVector(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetToken(ctx context.Context, in *npool.GetTokenRequest) (resp *npool.GetTokenResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetToken(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTokenOnly(ctx context.Context, in *npool.GetTokenOnlyRequest) (resp *npool.GetTokenOnlyResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTokenOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTokens(ctx context.Context, in *npool.GetTokensRequest) (resp *npool.GetTokensResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTokens(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistToken(ctx context.Context, in *npool.ExistTokenRequest) (resp *npool.ExistTokenResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistToken(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistTokenConds(ctx context.Context, in *npool.ExistTokenCondsRequest) (resp *npool.ExistTokenCondsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistTokenConds(ctx, in)
		return resp, err
	})
	return resp, err
}

func CountTokens(ctx context.Context, in *npool.CountTokensRequest) (resp *npool.CountTokensResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CountTokens(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteToken(ctx context.Context, in *npool.DeleteTokenRequest) (resp *npool.DeleteTokenResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteToken(ctx, in)
		return resp, err
	})
	return resp, err
}
