//nolint:nolintlint,dupl
package v1

import (
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
