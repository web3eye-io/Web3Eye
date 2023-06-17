package retrieve

import (
	"context"
	_ "encoding/base64"
	"fmt"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	retrieverpb "github.com/web3eye-io/Web3Eye/proto/web3eye/retriever/v1"
)

func (h *Handler) StatRetrieve(ctx context.Context) (*retrieverpb.Retrieve, error) {
	uid := fmt.Sprintf("%v:%v", h.Contract, h.TokenID)
	index, err := orbit.FileState().GetFileSnapshot(ctx, h.ChainType, uid, h.ChainID)
	if err != nil {
		return nil, err
	}
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return nil, err
	}
	state, err := orbit.FileState().GetFileRetrieve(ctx, h.ChainType, uid, h.ChainID)
	if err != nil {
		return nil, err
	}

	return &retrieverpb.Retrieve{
		ChainType:        h.ChainType,
		ChainID:          h.ChainID,
		Contract:         h.Contract,
		TokenID:          h.TokenID,
		RetrieveState:    state,
		ProposalCID:      snapshot.ProposalCID,
		DealID:           snapshot.DealID,
		BackupPayloadCID: snapshot.SnapshotRoot,
	}, nil
}
