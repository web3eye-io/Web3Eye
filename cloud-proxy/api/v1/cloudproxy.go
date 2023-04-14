//nolint:nolintlint,dupl
package v1

import (
	"context"

	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/cloudproxy/v1"
)

func (s *Server) GrpcProxyChannel(stream npool.Manager_GrpcProxyChannelServer) error {
	streamMGR := GetSteamMGR()
	streamMGR.AddProxySteam(stream)
	return nil
}

func (s *Server) GrpcProxy(ctx context.Context, in *npool.GrpcProxyRequest) (*npool.GrpcProxyResponse, error) {
	streamMGR := GetSteamMGR()
	resp, err := streamMGR.InvokeMSG(ctx, &npool.FromGrpcProxy{MsgID: in.MsgID, Method: in.Method, ReqRaw: in.ReqRaw})
	if err != nil {
		return nil, err
	}
	return &npool.GrpcProxyResponse{MsgID: resp.MsgID, Method: resp.Method, RespRaw: resp.RespRaw}, nil
}
