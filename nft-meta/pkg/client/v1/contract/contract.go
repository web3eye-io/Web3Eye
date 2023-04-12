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
	"google.golang.org/protobuf/proto"
)

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

var (
	cc      grpc.ClientConnInterface = nil
	timeout                          = 10 * time.Second
)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if cc == nil {
		conn, err := GetGrpcConn()
		if err != nil {
			return nil, err
		}
		defer conn.Close()
		cc = conn
	}
	cli := npool.NewManagerClient(cc)
	return handler(_ctx, cli)
}

type Po struct{}
type ReqItem struct {
	M string
	A proto.Message
	R proto.Message
	G []grpc.CallOption
}

type rawCodec struct{}

func (cb rawCodec) Marshal(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}

func (cb rawCodec) Unmarshal(data []byte, v interface{}) error {
	ba, ok := v.(*[]byte)
	if !ok {
		panic("ss")
	}
	*ba = append(*ba, data...)

	return nil
}

func (cb rawCodec) Name() string { return "dtm_raw" }

func (p *Po) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	reqRaw, err := proto.Marshal(args.(proto.Message))
	if err != nil {
		fmt.Println("1", err)
		panic("1sss")
	}

	conn, err := grpc.Dial("nft-meta:30101",
		grpc.WithTransportCredentials(
			insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.ForceCodec(rawCodec{}),
		),
	)
	if err != nil {
		fmt.Println("2", err)
		panic("2sss")
	}

	out := &[]byte{}
	err = conn.Invoke(ctx, method, reqRaw, out, opts...)
	if err != nil {
		fmt.Println("3", err)
		panic("3sss")
	}

	err = proto.Unmarshal(*out, reply.(proto.Message))
	if err != nil {
		fmt.Println("3", err)
		panic("3sss")
	}
	return nil
}

func (p *Po) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	// not impl
	return nil, nil
}

func GetGrpcConn() (*grpc.ClientConn, error) {
	return grpc.Dial(
		fmt.Sprintf("%v:%v",
			config.GetConfig().NFTMeta.IP,
			config.GetConfig().NFTMeta.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func SetClientConnInterface(c grpc.ClientConnInterface) {
	cc = c
}

func CreateContract(ctx context.Context, in *npool.ContractReq) (*npool.Contract, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateContract(_ctx, &npool.CreateContractRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Contract), nil
}

func CreateContracts(ctx context.Context, in []*npool.ContractReq) ([]*npool.Contract, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateContracts(_ctx, &npool.CreateContractsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.Contract), nil
}

func UpdateContract(ctx context.Context, in *npool.ContractReq) (*npool.Contract, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateContract(_ctx, &npool.UpdateContractRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Contract), nil
}

func GetContract(ctx context.Context, id string) (*npool.Contract, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetContract(_ctx, &npool.GetContractRequest{
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
	return info.(*npool.Contract), nil
}

func GetContractOnly(ctx context.Context, conds *npool.Conds) (*npool.Contract, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetContractOnly(_ctx, &npool.GetContractOnlyRequest{
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
	return info.(*npool.Contract), nil
}

func GetContracts(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Contract, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetContracts(_ctx, &npool.GetContractsRequest{
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
	return infos.([]*npool.Contract), total, nil
}

func ExistContract(ctx context.Context, id string) (bool, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistContract(_ctx, &npool.ExistContractRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func ExistContractConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistContractConds(_ctx, &npool.ExistContractCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func CountContracts(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountContracts(_ctx, &npool.CountContractsRequest{
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

func DeleteContract(ctx context.Context, id string) (*npool.Contract, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteContract(_ctx, &npool.DeleteContractRequest{
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
	return info.(*npool.Contract), nil
}
