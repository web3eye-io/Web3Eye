//nolint:nolintlint,dupl
package v1

import (
	"context"
	"fmt"
	"io"

	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
)

func (s *Server) ProxyChannel(stream npool.Manager_ProxyChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println(args.MsgID)
		reply := &npool.ProxyChannelResponse{MsgID: args.MsgID}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (s *Server) Proxy(ctx context.Context, in *npool.ProxyRequest) (*npool.ProxyResponse, error) {
	fmt.Println(in.MsgID)
	return &npool.ProxyResponse{MsgID: in.MsgID, Msg: in.Msg}, nil
}
