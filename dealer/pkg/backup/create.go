package backup

import (
	"context"
	"fmt"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	metacli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/snapshot"
	metapb "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

func (h *Handler) CreateBackup(ctx context.Context) (*dealerpb.Snapshot, error) {
	if err := orbit.Backup().Create(ctx, h.Index); err != nil {
		return nil, err
	}
	NewSnapshot()
	snapshot, err := orbit.Snapshot().UpdateSnapshotState(ctx, h.Index, dealerpb.BackupState_BackupStateCreated)
	if err != nil {
		return nil, err
	}
	for _, item := range snapshot.Items {
		uid := fmt.Sprintf("%v:%v", item.Contract, item.TokenID)
		if err := orbit.FileState().SetFileState(ctx, item.ChainType, uid, item.ChainID, dealerpb.BackupState_BackupStateCreated); err != nil {
			return nil, err
		}
	}

	_state := snapshot.BackupState.String()
	if _, err := metacli.UpdateSnapshot(ctx, &metapb.SnapshotReq{
		ID:          &snapshot.ID,
		BackupState: &_state,
	}); err != nil {
		return nil, err
	}
	return snapshot, nil
}
