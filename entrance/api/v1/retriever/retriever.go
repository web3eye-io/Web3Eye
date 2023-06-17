package retriever

import (
	"context"
	"math/rand"

	// client "github.com/web3eye-io/Web3Eye/retriever/pkg/client/v1/retriever"

	entrancernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/retriever"
)

type Server struct {
	entrancernpool.UnimplementedManagerServer
}

func (s *Server) StartRetrieve(ctx context.Context, in *entrancernpool.StartRetrieveRequest) (*entrancernpool.StatRetrieveResponse, error) {
	// client.UseCloudProxyCC()
	// return client.StartRetrieve(ctx, in)
	return &entrancernpool.StatRetrieveResponse{
		Info: &entrancernpool.Retrieve{RetrieveState: "Start"},
	}, nil
}

func (s *Server) StatRetrieve(ctx context.Context, in *entrancernpool.StatRetrieveRequest) (*entrancernpool.StatRetrieveResponse, error) {
	// client.UseCloudProxyCC()
	// return client.StatRetrieve(ctx, in)
	if rand.Int31n(2) == 0 {
		return &entrancernpool.StatRetrieveResponse{}, nil
	}
	return &entrancernpool.StatRetrieveResponse{Info: &entrancernpool.Retrieve{RetrieveState: "Start"}}, nil

}
