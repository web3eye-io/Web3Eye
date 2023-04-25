package token

import (
	"context"

	"github.com/web3eye-io/Web3Eye/nft-meta/api/v1/token"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
	token.Server
}

func (s *Server) GetToken(ctx context.Context, in *nftmetanpool.GetTokenRequest) (*nftmetanpool.GetTokenResponse, error) {
	return s.Server.GetToken(ctx, in)
}

func (s *Server) GetTokenOnly(ctx context.Context, in *nftmetanpool.GetTokenOnlyRequest) (*nftmetanpool.GetTokenOnlyResponse, error) {
	return s.Server.GetTokenOnly(ctx, in)
}

func (s *Server) GetTokens(ctx context.Context, in *nftmetanpool.GetTokensRequest) (*nftmetanpool.GetTokensResponse, error) {
	return s.Server.GetTokens(ctx, in)
}

func (s *Server) CountTokens(ctx context.Context, in *nftmetanpool.CountTokensRequest) (*nftmetanpool.CountTokensResponse, error) {
	return s.Server.CountTokens(ctx, in)
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}
