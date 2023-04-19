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

func GetContract(ctx context.Context, id string) (*nftmetaproto.Contract, error) {
	info, err := WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetContract(_ctx, &nftmetaproto.GetContractRequest{
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
	return info.(*nftmetaproto.Contract), nil
}

func GetContractOnly(ctx context.Context, conds *nftmetaproto.Conds) (*nftmetaproto.Contract, error) {
	info, err := WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetContractOnly(_ctx, &nftmetaproto.GetContractOnlyRequest{
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
	return info.(*nftmetaproto.Contract), nil
}

func GetContracts(ctx context.Context, conds *nftmetaproto.Conds, offset, limit int32) ([]*nftmetaproto.Contract, uint32, error) {
	var total uint32
	infos, err := WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetContracts(_ctx, &nftmetaproto.GetContractsRequest{
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
	return infos.([]*nftmetaproto.Contract), total, nil
}

func CountContracts(ctx context.Context, conds *nftmetaproto.Conds) (uint32, error) {
	infos, err := WithCRUD(ctx, func(_ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountContracts(_ctx, &nftmetaproto.CountContractsRequest{
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
