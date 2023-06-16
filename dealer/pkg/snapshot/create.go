package snapshot

import (
	"context"
	"fmt"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	metacli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/snapshot"
	metapb "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

func (h *Handler) CreateSnapshot(ctx context.Context) (*dealerpb.Snapshot, error) {
	if err := orbit.Snapshot().SetWaitSnapshot(ctx, h.SnapshotCommP, h.SnapshotRoot, h.SnapshotURI, h.Items); err != nil {
		return nil, err
	}
	if err := orbit.Snapshot().NextWaitSnapshot(ctx); err != nil {
		return nil, err
	}
	index := orbit.Snapshot().WaitSnapshotIndex()
	for _, item := range h.Items {
		uid := fmt.Sprintf("%v:%v", item.Contract, item.TokenID)
		if err := orbit.FileState().SetFileState(ctx, item.ChainType, uid, item.ChainID, dealerpb.BackupState_BackupStateNone); err != nil {
			return nil, err
		}
		if err := orbit.FileState().SetFileSnapshot(ctx, item.ChainType, uid, item.ChainID, index); err != nil {
			return nil, err
		}
	}
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return nil, err
	}
	state := snapshot.BackupState.String()
	if _, err := metacli.CreateSnapshot(ctx, &metapb.SnapshotReq{
		ID:            &snapshot.ID,
		Index:         &snapshot.Index,
		SnapshotCommP: &snapshot.SnapshotCommP,
		SnapshotRoot:  &snapshot.SnapshotRoot,
		SnapshotURI:   &snapshot.SnapshotURI,
		BackupState:   &state,
	}); err != nil {
		return nil, err
	}
	return snapshot, nil
}
