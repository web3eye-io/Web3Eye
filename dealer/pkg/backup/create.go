package backup

import (
	"context"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

func (h *Handler) CreateBackup(ctx context.Context) (*dealerpb.Snapshot, error) {
	if err := orbit.Backup().Create(ctx, h.Index); err != nil {
		return nil, err
	}
	NewSnapshot()
	return orbit.Snapshot().UpdateSnapshotState(ctx, h.Index, dealerpb.BackupState_BackupStateCreated)
}
