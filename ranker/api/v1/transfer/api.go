package transfer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/transfer"
	"github.com/web3eye-io/Web3Eye/ranker/pkg/crud/v1/transfer"
	"google.golang.org/grpc"
)

type Server struct {
	rankernpool.UnimplementedManagerServer
}

func (s *Server) GetTransfers(ctx context.Context, in *rankernpool.GetTransfersRequest) (*rankernpool.GetTransfersResponse, error) {
	infos, total, err := transfer.Rows(ctx, in)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		for _, item := range info.OfferItems {
			FillAmountStr(item)
		}
		for _, item := range info.TargetItems {
			FillAmountStr(item)
		}
	}
	return &rankernpool.GetTransfersResponse{Infos: infos, Total: uint32(total)}, nil
}

func FillAmountStr(item *rankernpool.OrderItem) {
	if item.TokenType == v1.TokenType_ERC1155 ||
		item.TokenType == v1.TokenType_ERC1155_WITH_CRITERIA ||
		item.TokenType == v1.TokenType_ERC721 ||
		item.TokenType == v1.TokenType_ERC721_WITH_CRITERIA {
		item.AmountStr = fmt.Sprintf("%v #%v", item.Symbol, item.TokenID)
	} else {
		maxAmountStrFixed := int32(4)
		amount := big.NewInt(0).SetUint64(item.Amount)
		amountStr := decimal.NewFromBigInt(amount, -int32(item.Decimals)).Round(maxAmountStrFixed)
		item.AmountStr = fmt.Sprintf("%v %v", amountStr, item.Symbol)
	}
}

func Register(server grpc.ServiceRegistrar) {
	rankernpool.RegisterManagerServer(server, &Server{})
}
