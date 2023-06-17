package retrieve

import (
	"context"
	"fmt"

	retrieverpb "github.com/web3eye-io/Web3Eye/proto/web3eye/retriever/v1"
)

func (h *Handler) StatRetrieve(ctx context.Context) (*retrieverpb.Retrieve, error) {
	fmt.Printf("h = %v\n", h)
	return nil, nil
}
