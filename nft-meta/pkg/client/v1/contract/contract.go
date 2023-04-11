//nolint:nolintlint,dupl
package contract

import (
	"context"
	"encoding/json"
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

func (p *Po) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {

	_args, ok := args.(proto.Message)

	fmt.Println(_args.ProtoReflect())

	if !ok {
		fmt.Println("stop 1")
		return nil
	}

	_reply, ok := reply.(proto.Message)
	if !ok {
		fmt.Println("stop 2")
		return nil
	}

	req := ReqItem{M: method, A: _args, R: _reply, G: opts}

	reqBytes, err := json.Marshal(req)
	fmt.Println(err)
	fmt.Println(string(reqBytes))

	_req := &ReqItem{}
	err = json.Unmarshal(reqBytes, _req)
	fmt.Println(err)

	conn, err := GetGrpcConn()
	fmt.Println(err)
	if err != nil {
		return err
	}

	err = conn.Invoke(ctx, method, _args, reply, opts...)
	fmt.Println(err)
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
