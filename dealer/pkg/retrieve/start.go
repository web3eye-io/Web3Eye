package retrieve

import (
	"context"
	"fmt"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	retrieverpb "github.com/web3eye-io/Web3Eye/proto/web3eye/retriever/v1"
)

const mock = true

func (h *Handler) StartRetrieve(ctx context.Context) (*retrieverpb.Retrieve, error) {
	uid := fmt.Sprintf("%v:%v", h.Contract, h.TokenID)
	index, err := orbit.FileState().GetFileSnapshot(ctx, h.ChainType, uid, h.ChainID)
	if err != nil {
		return nil, err
	}
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return nil, err
	}
	if snapshot.DealID == 0 {
		if !mock {
			return nil, fmt.Errorf("file backup not exist")
		}
		snapshot.DealID = 89468
	}
	return nil, nil
}
