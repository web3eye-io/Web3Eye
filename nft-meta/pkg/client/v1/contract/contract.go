//nolint:nolintlint,dupl
package contract

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/web3eye-io/Web3Eye/config"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"

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

func CreateContract(ctx context.Context, in *npool.CreateContractRequest) (resp *npool.CreateContractResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateContract(_ctx, in)
		return resp, err
	})
	return resp, err
}

func CreateContracts(ctx context.Context, in *npool.CreateContractsRequest) (resp *npool.CreateContractsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateContracts(_ctx, in)
		return resp, err
	})
	return resp, err
}

func UpdateContract(ctx context.Context, in *npool.UpdateContractRequest) (resp *npool.UpdateContractResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.UpdateContract(_ctx, in)
		return resp, err
	})
	return resp, err
}

func GetContract(ctx context.Context, in *npool.GetContractRequest) (resp *npool.GetContractResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetContract(_ctx, in)
		return resp, err
	})
	return resp, err
}

func GetContractOnly(ctx context.Context, in *npool.GetContractOnlyRequest) (resp *npool.GetContractOnlyResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetContractOnly(_ctx, in)
		return resp, err
	})
	return resp, err
}

func GetContracts(ctx context.Context, in *npool.GetContractsRequest) (resp *npool.GetContractsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetContracts(_ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistContract(ctx context.Context, in *npool.ExistContractRequest) (resp *npool.ExistContractResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistContract(_ctx, in)
		return resp, err
	})
	return resp, err
}

func ExistContractConds(ctx context.Context, in *npool.ExistContractCondsRequest) (resp *npool.ExistContractCondsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ExistContractConds(_ctx, in)
		return resp, err
	})
	return resp, err
}

func CountContracts(ctx context.Context, in *npool.CountContractsRequest) (resp *npool.CountContractsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CountContracts(_ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteContract(ctx context.Context, in *npool.DeleteContractRequest) (resp *npool.DeleteContractResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteContract(_ctx, in)
		return resp, err
	})
	return resp, err
}
