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

func GetTransfer(ctx context.Context, id string) (*nftmetaproto.Transfer, error) {
	info, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTransfer(ctx, &nftmetaproto.GetTransferRequest{
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
	return info.(*nftmetaproto.Transfer), nil
}

func GetTransferOnly(ctx context.Context, conds *nftmetaproto.Conds) (*nftmetaproto.Transfer, error) {
	info, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTransferOnly(ctx, &nftmetaproto.GetTransferOnlyRequest{
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
	return info.(*nftmetaproto.Transfer), nil
}

func GetTransfers(ctx context.Context, conds *nftmetaproto.Conds, offset, limit int32) ([]*nftmetaproto.Transfer, uint32, error) {
	var total uint32
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTransfers(ctx, &nftmetaproto.GetTransfersRequest{
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
	return infos.([]*nftmetaproto.Transfer), total, nil
}

func CountTransfers(ctx context.Context, conds *nftmetaproto.Conds) (uint32, error) {
	infos, err := WithCRUD(ctx, func(ctx context.Context, cli rankerproto.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountTransfers(ctx, &nftmetaproto.CountTransfersRequest{
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
