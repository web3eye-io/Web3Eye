package contract

import (
	"context"
	"encoding/json"

	entrancepool "github.com/web3eye-io/Web3Eye/proto/web3eye/entrance/v1/contract"
	rankerpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/contract"
	client "github.com/web3eye-io/Web3Eye/ranker/pkg/client/v1/contract"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	entrancepool.UnimplementedManagerServer
}

func (s *Server) GetContractAndTokens(ctx context.Context, in *rankerpool.GetContractAndTokensReq) (*rankerpool.GetContractAndTokensResp, error) {
	ret := `{
		"Contract": {
			"ID": "ae0c722f-2427-4649-9c3e-d0bebd0a345d",
			"ChainType": "Ethereum",
			"ChainID": "1",
			"Address": "0xCC845392C20a5836b8f5d2D3D88EF7b5B1820644",
			"Name": "BoredPunk",
			"Symbol": "BP",
			"Decimals": 0,
			"Creator": "0x0000000000000000000000000000000000000000",
			"BlockNum": "0",
			"TxHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"TxTime": 0,
			"ProfileURL": "",
			"BaseURL": "",
			"BannerURL": "",
			"Description": "",
			"Remark": ""
		},
		"Tokens": [
			{
				"ID": "63de8c6d-ae1a-46ba-a135-4c3f55c68449",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3707",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3707.png",
				"Name": "BoredPunks #3707",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "ce24e328-a297-4f74-8801-b72ef41d0352",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3708",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3708.png",
				"Name": "BoredPunks #3708",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "ba07e380-d8e2-4401-b84d-f7eb947e9a3a",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3710",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3710.png",
				"Name": "BoredPunks #3710",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "00e616a1-5bc2-470b-a924-bd2ee8b9e638",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3711",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3711.png",
				"Name": "BoredPunks #3711",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "aa08aee6-4596-41d6-a23a-a82509bf9dc0",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3717",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3717.png",
				"Name": "BoredPunks #3717",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "9e026ab9-b957-4218-9434-f6e93444d42b",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3718",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3718.png",
				"Name": "BoredPunks #3718",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "10699b15-8404-4d48-aff9-0c9977e34021",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3709",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3709.png",
				"Name": "BoredPunks #3709",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "f1da83c3-b019-4bcb-a466-23addfb603e5",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3720",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3720.png",
				"Name": "BoredPunks #3720",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "471f2f2e-4890-48f0-a7d1-920ce4e88fe9",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3719",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3719.png",
				"Name": "BoredPunks #3719",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "70020b45-d647-4ebb-8e7a-b00370ded7ee",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3721",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3721.png",
				"Name": "BoredPunks #3721",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "833fcb28-a384-4c82-8d0a-fca5ac357953",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3712",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3712.png",
				"Name": "BoredPunks #3712",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "ca1eb8d6-89f5-47f0-b268-46c8e4b95c7c",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3713",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3713.png",
				"Name": "BoredPunks #3713",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "f77dcc04-c728-4cb5-8bcd-86c0c11c04dc",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3715",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3715.png",
				"Name": "BoredPunks #3715",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "6ad5fb56-6299-45b2-b545-ad2844064d5a",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3714",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3714.png",
				"Name": "BoredPunks #3714",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "f11385a3-9bac-42e3-8f31-397d862e302a",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3716",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3716.png",
				"Name": "BoredPunks #3716",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "15556859-bacb-4077-92ad-097c8220c5ae",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3724",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3724.png",
				"Name": "BoredPunks #3724",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "1b261267-1f12-4a07-b853-a7d0e3d12b47",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3722",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3722.png",
				"Name": "BoredPunks #3722",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "cea02cc9-f34d-4387-b5fa-eb693f762ff3",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3723",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3723.png",
				"Name": "BoredPunks #3723",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "ef13af49-a822-488f-9134-faccf21ce56c",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3726",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3726.png",
				"Name": "BoredPunks #3726",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "8ad4fcdb-2d5c-4c8e-901d-0cc2ee761f5d",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3727",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3727.png",
				"Name": "BoredPunks #3727",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "c67cd091-25dc-4422-a281-247190d4dd7e",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3729",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3729.png",
				"Name": "BoredPunks #3729",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "892514e2-b394-4bc8-9e2e-28af353424ff",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3728",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3728.png",
				"Name": "BoredPunks #3728",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "99f2dbae-7f8f-47c7-87a0-73106d58ed63",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3730",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3730.png",
				"Name": "BoredPunks #3730",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			},
			{
				"ID": "84118b23-afea-4a49-9ba1-8f2166667cdc",
				"ChainType": "Ethereum",
				"TokenType": "ERC721",
				"TokenID": "3725",
				"Owner": "",
				"ImageURL": "ipfs://QmaZkDroZm9bnqtUrtuwGV1QDp7yB5kyPUvNhAvKy4dj5d/3725.png",
				"Name": "BoredPunks #3725",
				"IPFSImageURL": "",
				"ImageSnapshotID": "",
				"TransfersNum": 1
			}
		],
		"TotalTokens": 24
	}`
	resp := &rankerpool.GetContractAndTokensResp{}
	json.Unmarshal([]byte(ret), resp)
	return resp, nil
	client.UseCloudProxyCC()
	return client.GetContractAndTokens(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	entrancepool.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return entrancepool.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
