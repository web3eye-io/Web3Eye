package retriever

import (
	"context"
	"math/rand"

	// client "github.com/web3eye-io/Web3Eye/retriever/pkg/client/v1/retriever"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	entrancernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/retriever"
	"google.golang.org/grpc"
)

type Server struct {
	entrancernpool.UnimplementedManagerServer
}

func (s *Server) StartRetrieve(ctx context.Context, in *entrancernpool.StartRetrieveRequest) (*entrancernpool.StartRetrieveResponse, error) {
	// client.UseCloudProxyCC()
	// return client.StartRetrieve(ctx, in)
	return &entrancernpool.StartRetrieveResponse{
		Info: &entrancernpool.Retrieve{
			ChainType:     in.ChainType,
			Contract:      in.Contract,
			ChainID:       in.ChainID,
			TokenID:       in.TokenID,
			RetrieveState: "Start",
		},
	}, nil
}

func (s *Server) StatRetrieve(ctx context.Context, in *entrancernpool.StatRetrieveRequest) (*entrancernpool.StatRetrieveResponse, error) {
	// client.UseCloudProxyCC()
	// return client.StatRetrieve(ctx, in)
	retrieveState := "Start"
	if rand.Int31n(2) == 0 {
		retrieveState = ""
	}
	return &entrancernpool.StatRetrieveResponse{Info: &entrancernpool.Retrieve{
		ChainType:     in.ChainType,
		Contract:      in.Contract,
		ChainID:       in.ChainID,
		TokenID:       in.TokenID,
		RetrieveState: retrieveState,
	}}, nil

}

func Register(server grpc.ServiceRegistrar) {
	entrancernpool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancernpool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
