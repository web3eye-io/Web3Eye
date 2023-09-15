package transfer

import (
	"context"

	entrancepool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/transfer"
	rankerpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/transfer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	entrancepool.UnimplementedManagerServer
}

func (s *Server) GetTransfers(ctx context.Context, in *rankerpool.GetTransfersRequest) (*rankerpool.GetTransfersResponse, error) {
	// ret := `{
	// 	"Infos": [
	// 		{
	// 			"ID": "de812157-2175-49d0-8171-5ad38597e9a6",
	// 			"ChainType": "Ethereum",
	// 			"ChainID": "1",
	// 			"Contract": "0x5Af0D9827E0c53E4799BB226655A1de152A425a5",
	// 			"TokenType": "\n",
	// 			"TokenID": "1687",
	// 			"From": "0x0000000000000000000000007df70b612040c682d1cb2e32017446e230fcd747",
	// 			"To": "0x00000000000000000000000029469395eaf6f95920e59f858042f0e28d98a20b",
	// 			"Amount": "1",
	// 			"BlockNumber": "18068781",
	// 			"TxHash": "0x98b9f7ccb9d8f68e0a8f1e65b6a894daa581dd3c73a3c998c9b718e04125740a",
	// 			"BlockHash": "0xf97544bedf905e7d74ddd6dc891a71dd983c04f19995133396e7b7526e790db1",
	// 			"TxTime": 0,
	// 			"Remark": ""
	// 		},
	// 		{
	// 			"ID": "ec0ac03f-b110-470e-a9f1-db8cbe851975",
	// 			"ChainType": "Ethereum",
	// 			"ChainID": "1",
	// 			"Contract": "0x5Af0D9827E0c53E4799BB226655A1de152A425a5",
	// 			"TokenType": "\n",
	// 			"TokenID": "1687",
	// 			"From": "0x000000000000000000000000bc2dc51e35873d09725bfb10e6ce6ac6d0033a01",
	// 			"To": "0x0000000000000000000000007df70b612040c682d1cb2e32017446e230fcd747",
	// 			"Amount": "1",
	// 			"BlockNumber": "18068779",
	// 			"TxHash": "0x95d9fb43760dc1540c0357466f5a39ad66f62467610fd94f0588d76ad6eb0611",
	// 			"BlockHash": "0xae780ebf5d9a7ec77f9f9e3d77757f0fbb65b034464494fa0757271698593295",
	// 			"TxTime": 0,
	// 			"Remark": ""
	// 		}
	// 	],
	// 	"Total": 2
	// }`
	// resp := &rankerpool.GetTransfersResponse{}
	// json.Unmarshal([]byte(ret), resp)
	// return resp, nil
	client.UseCloudProxyCC()
	return client.GetTransfers(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancepool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancepool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
