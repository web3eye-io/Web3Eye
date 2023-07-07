//nolint:nolintlint,dupl
package transfer

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
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

func CreateTransfer(ctx context.Context, in *npool.CreateTransferRequest) (resp *npool.CreateTransferResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateTransfer(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpsertTransfer(ctx context.Context, in *npool.UpsertTransferRequest) (resp *npool.UpsertTransferResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpsertTransfer(ctx, in)
		return resp, err
	})
	return resp, err
}

func CreateTransfers(ctx context.Context, in *npool.CreateTransfersRequest) (resp *npool.CreateTransfersResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateTransfers(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpsertTransfers(ctx context.Context, in *npool.UpsertTransfersRequest) (resp *npool.UpsertTransfersResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpsertTransfers(ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateTransfer(ctx context.Context, in *npool.UpdateTransferRequest) (resp *npool.UpdateTransferResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateTransfer(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTransfer(ctx context.Context, in *npool.GetTransferRequest) (resp *npool.GetTransferResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTransfer(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTransferOnly(ctx context.Context, in *npool.GetTransferOnlyRequest) (resp *npool.GetTransferOnlyResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTransferOnly(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetTransfers(ctx context.Context, in *npool.GetTransfersRequest) (resp *npool.GetTransfersResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetTransfers(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistTransfer(ctx context.Context, in *npool.ExistTransferRequest) (resp *npool.ExistTransferResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistTransfer(ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistTransferConds(ctx context.Context, in *npool.ExistTransferCondsRequest) (resp *npool.ExistTransferCondsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistTransferConds(ctx, in)
		return resp, err
	})
	return resp, err
}

func CountTransfers(ctx context.Context, in *npool.CountTransfersRequest) (resp *npool.CountTransfersResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CountTransfers(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteTransfer(ctx context.Context, in *npool.DeleteTransferRequest) (resp *npool.DeleteTransferResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteTransfer(ctx, in)
		return resp, err
	})
	return resp, err
}
