package snapshot

import (
	"context"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
)

func (h *Handler) CreateSnapshot(ctx context.Context) error {
	if err := orbit.Snapshot().SetWaitSnapshot(ctx, h.SnapshotURI, h.Items); err != nil {
		return err
	}
	if err := orbit.Snapshot().NextWaitSnapshot(ctx); err != nil {
		return err
	}
	return nil
}
