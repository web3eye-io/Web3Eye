package snapshot

import (
	"context"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

func (h *Handler) GetSnapshots(ctx context.Context) ([]*dealerpb.Snapshot, uint64, error) {
	indexes := h.Indexes
	total := orbit.Snapshot().WaitSnapshotIndex()

	if len(indexes) == 0 {
		for i := uint64(0); i < total; i++ {
			indexes = append(indexes, i)
		}
	}

	snapshots := []*dealerpb.Snapshot{}
	for _, index := range indexes {
		snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
		if err != nil {
			return nil, 0, err
		}
		snapshots = append(snapshots, snapshot)
	}

	return snapshots, total, nil
}
