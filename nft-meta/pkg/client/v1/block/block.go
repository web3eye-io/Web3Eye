//nolint:nolintlint,dupl
package block

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/block"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
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

func CreateBlock(ctx context.Context, in *npool.CreateBlockRequest) (resp *npool.CreateBlockResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateBlock(ctx, in)
		return resp, err
	})
	return resp, err
}

func CreateBlocks(ctx context.Context, in *npool.CreateBlocksRequest) (resp *npool.CreateBlocksResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateBlocks(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateBlock(ctx context.Context, in *npool.UpdateBlockRequest) (resp *npool.UpdateBlockResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateBlock(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetBlock(ctx context.Context, in *npool.GetBlockRequest) (resp *npool.GetBlockResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetBlock(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetBlockOnly(ctx context.Context, in *npool.GetBlockOnlyRequest) (resp *npool.GetBlockOnlyResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetBlockOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetBlocks(ctx context.Context, in *npool.GetBlocksRequest) (resp *npool.GetBlocksResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetBlocks(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistBlock(ctx context.Context, in *npool.ExistBlockRequest) (resp *npool.ExistBlockResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistBlock(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistBlockConds(ctx context.Context, in *npool.ExistBlockCondsRequest) (resp *npool.ExistBlockCondsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistBlockConds(ctx, in)
		return resp, err
	})
	return resp, err
}

func CountBlocks(ctx context.Context, in *npool.CountBlocksRequest) (resp *npool.CountBlocksResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CountBlocks(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteBlock(ctx context.Context, in *npool.DeleteBlockRequest) (resp *npool.DeleteBlockResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteBlock(ctx, in)
		return resp, err
	})
	return resp, err
}
