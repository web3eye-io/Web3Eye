package snapshot

import (
	"context"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

func (h *Handler) GetSnapshots(ctx context.Context) ([]*dealerpb.Snapshot, uint64, error) {
	indexes := h.Indexes
	total := uint64(0)

	switch h.SnapshotType {
	case dealerpb.SnapshotType_SnapshotWait:
		total = orbit.Snapshot().WaitSnapshotIndex()
	case dealerpb.SnapshotType_SnapshotBackup:
		total = orbit.Snapshot().BackupSnapshotIndex()
	}

	if len(indexes) == 0 {
		for i := uint64(0); i < total; i++ {
			indexes = append(indexes, i)
		}
	}

	snapshots := []*dealerpb.Snapshot{}

	switch h.SnapshotType {
	case dealerpb.SnapshotType_SnapshotWait:
		for _, index := range indexes {
			uri, items, err := orbit.Snapshot().GetWaitSnapshot(ctx, index)
			if err != nil {
				return nil, 0, err
			}
			snapshots = append(snapshots, &dealerpb.Snapshot{
				SnapshotURI: uri,
				Items:       items,
			})
		}
	case dealerpb.SnapshotType_SnapshotBackup:
		for _, index := range indexes {
			uri, items, err := orbit.Snapshot().GetBackupSnapshot(ctx, index)
			if err != nil {
				return nil, 0, err
			}
			snapshots = append(snapshots, &dealerpb.Snapshot{
				SnapshotURI: uri,
				Items:       items,
			})
		}
	}

	return snapshots, total, nil
}
