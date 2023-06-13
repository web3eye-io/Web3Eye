package snapshot

import (
	"context"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

func (h *Handler) CreateSnapshot(ctx context.Context) (*dealerpb.Snapshot, error) {
	if err := orbit.Snapshot().SetWaitSnapshot(ctx, h.SnapshotCID, h.SnapshotURI, h.Items); err != nil {
		return nil, err
	}
	if err := orbit.Snapshot().NextWaitSnapshot(ctx); err != nil {
		return nil, err
	}
	return &dealerpb.Snapshot{
		SnapshotURI: h.SnapshotURI,
		Items:       h.Items,
	}, nil
}
