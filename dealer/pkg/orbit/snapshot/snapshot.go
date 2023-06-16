package snapshot

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"

	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	orbitdb "berty.tech/go-orbit-db"
	orbitiface "berty.tech/go-orbit-db/iface"

	"github.com/google/uuid"
)

type SnapshotKV struct {
	odb                 orbitiface.OrbitDB
	kvSnapshotIndex     orbitdb.KeyValueStore
	waitSnapshotIndex   uint64
	backupSnapshotIndex uint64
}

const (
	CurrentWaitSnapshotIndex   = "current-wait-index"
	CurrentBackupSnapshotIndex = "current-wait-index"
	SnapshotID                 = "snapshot-id"
	SnapshotURI                = "snapshot-uri"
	ContentItems               = "content-items"
	SnapshotRoot               = "snapshot-root"
	SnapshotCommP              = "snapshot-commp"
	SnapshotProposalCID        = "snapshot-proposal-cid"
	SnapshotBackupState        = "snapshot-backup-state"
	KVStoreSnapshotIndex       = "snapshot-index"
	KVStoreSnapshot            = "snapshot-"
)

func NewSnapshotKV(ctx context.Context, odb orbitiface.OrbitDB) (*SnapshotKV, error) {
	kv := &SnapshotKV{
		odb: odb,
	}
	var err error

	replicate := true
	kv.kvSnapshotIndex, err = odb.KeyValue(ctx, KVStoreSnapshotIndex, &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return nil, err
	}

	if err := kv.kvSnapshotIndex.Load(ctx, -1); err != nil { //nolint
		return nil, err
	}

	b, err := kv.kvSnapshotIndex.Get(ctx, CurrentWaitSnapshotIndex)
	if err != nil {
		return nil, err
	}
	kv.waitSnapshotIndex, _ = binary.Uvarint(b)

	b, err = kv.kvSnapshotIndex.Get(ctx, CurrentBackupSnapshotIndex)
	if err != nil {
		return nil, err
	}
	kv.backupSnapshotIndex, _ = binary.Uvarint(b)

	return kv, nil
}

func (kv *SnapshotKV) NextWaitSnapshot(ctx context.Context) error {
	if kv.kvSnapshotIndex == nil {
		return fmt.Errorf("invalid kvstore")
	}

	b := make([]byte, 8)
	kv.waitSnapshotIndex += 1
	binary.PutUvarint(b, kv.waitSnapshotIndex)
	if _, err := kv.kvSnapshotIndex.Put(ctx, CurrentWaitSnapshotIndex, b); err != nil {
		return err
	}

	return nil
}

func (kv *SnapshotKV) setSnapshot(ctx context.Context, kvStoreName, snapshotCommP, snapshotRoot, snapshotURI string, items []*dealerpb.ContentItem) error {
	replicate := true
	_kv, err := kv.odb.KeyValue(ctx, kvStoreName, &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return err
	}
	defer _kv.Close()

	if err := _kv.Load(ctx, -1); err != nil {
		return err
	}

	uid := uuid.NewString()
	if _, err := _kv.Put(ctx, SnapshotID, []byte(uid)); err != nil {
		return err
	}
	if _, err := _kv.Put(ctx, SnapshotCommP, []byte(snapshotCommP)); err != nil {
		return err
	}
	if _, err := _kv.Put(ctx, SnapshotRoot, []byte(snapshotRoot)); err != nil {
		return err
	}
	if _, err := _kv.Put(ctx, SnapshotURI, []byte(snapshotURI)); err != nil {
		return err
	}

	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	if _, err := _kv.Put(ctx, ContentItems, b); err != nil {
		return err
	}

	if _, err := _kv.Put(ctx, SnapshotBackupState, []byte(dealerpb.BackupState_BackupStateNone.String())); err != nil {
		return err
	}

	return nil
}

func (kv *SnapshotKV) SetWaitSnapshot(ctx context.Context, snapshotCommP, snapshotRoot, snapshotURI string, items []*dealerpb.ContentItem) error {
	return kv.setSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, kv.waitSnapshotIndex), snapshotCommP, snapshotRoot, snapshotURI, items)
}

func (kv *SnapshotKV) getSnapshot(ctx context.Context, kvStoreName string) (*dealerpb.Snapshot, error) {
	replicate := true
	_kv, err := kv.odb.KeyValue(ctx, kvStoreName, &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return nil, err
	}
	defer _kv.Close()

	if err := _kv.Load(ctx, -1); err != nil {
		return nil, err
	}

	_id, err := _kv.Get(ctx, SnapshotID)
	if err != nil {
		return nil, err
	}

	_snapshotURI, err := _kv.Get(ctx, SnapshotURI)
	if err != nil {
		return nil, err
	}

	_commP, err := _kv.Get(ctx, SnapshotCommP)
	if err != nil {
		return nil, err
	}

	_cid, err := _kv.Get(ctx, SnapshotRoot)
	if err != nil {
		return nil, err
	}

	_items, err := _kv.Get(ctx, ContentItems)
	if err != nil {
		return nil, err
	}
	items := []*dealerpb.ContentItem{}
	_ = json.Unmarshal(_items, &items)

	_state, err := _kv.Get(ctx, SnapshotBackupState)
	if err != nil {
		return nil, err
	}
	state := dealerpb.BackupState(dealerpb.BackupState_value[string(_state)])

	_proposal, err := _kv.Get(ctx, SnapshotProposalCID)
	if err != nil {
		return nil, err
	}

	return &dealerpb.Snapshot{
		ID:            string(_id),
		SnapshotCommP: string(_commP),
		SnapshotRoot:  string(_cid),
		SnapshotURI:   string(_snapshotURI),
		ProposalCID:   string(_proposal),
		Items:         items,
		BackupState:   state,
	}, nil
}

func (kv *SnapshotKV) GetCurrentWaitSnapshot(ctx context.Context) (*dealerpb.Snapshot, error) {
	return kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, kv.waitSnapshotIndex))
}

func (kv *SnapshotKV) NextBackupSnapshot(ctx context.Context) error {
	if kv.kvSnapshotIndex == nil {
		return fmt.Errorf("invalid kvstore")
	}

	b := make([]byte, 8)
	kv.backupSnapshotIndex += 1
	binary.PutUvarint(b, kv.backupSnapshotIndex)
	if _, err := kv.kvSnapshotIndex.Put(ctx, CurrentBackupSnapshotIndex, b); err != nil {
		return err
	}

	return nil
}

func (kv *SnapshotKV) GetCurrentBackupSnapshot(ctx context.Context) (*dealerpb.Snapshot, error) {
	snapshot, err := kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, kv.backupSnapshotIndex))
	if err != nil {
		return nil, err
	}
	snapshot.Index = kv.backupSnapshotIndex
	return snapshot, nil
}

func (kv *SnapshotKV) GetSnapshot(ctx context.Context, index uint64) (*dealerpb.Snapshot, error) {
	if index >= kv.waitSnapshotIndex {
		return nil, fmt.Errorf("invalid snapshot index")
	}
	snapshot, err := kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, index))
	if err != nil {
		return nil, err
	}
	snapshot.Index = index
	return snapshot, nil
}

func (kv *SnapshotKV) UpdateSnapshotState(ctx context.Context, index uint64, state dealerpb.BackupState) (*dealerpb.Snapshot, error) {
	if index >= kv.waitSnapshotIndex {
		return nil, fmt.Errorf("invalid snapshot index")
	}

	replicate := true
	_kv, err := kv.odb.KeyValue(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, index), &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return nil, err
	}
	defer _kv.Close()

	if err := _kv.Load(ctx, -1); err != nil {
		return nil, err
	}

	if _, err := _kv.Put(ctx, SnapshotBackupState, []byte(state.String())); err != nil {
		return nil, err
	}

	snapshot, err := kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, index))
	if err != nil {
		return nil, err
	}
	snapshot.Index = index
	return snapshot, nil
}

func (kv *SnapshotKV) UpdateSnapshotProposalCID(ctx context.Context, index uint64, proposal string) (*dealerpb.Snapshot, error) {
	if index >= kv.waitSnapshotIndex {
		return nil, fmt.Errorf("invalid snapshot index")
	}

	replicate := true
	_kv, err := kv.odb.KeyValue(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, index), &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return nil, err
	}
	defer _kv.Close()

	if err := _kv.Load(ctx, -1); err != nil {
		return nil, err
	}

	if _, err := _kv.Put(ctx, SnapshotProposalCID, []byte(proposal)); err != nil {
		return nil, err
	}

	snapshot, err := kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, index))
	if err != nil {
		return nil, err
	}
	snapshot.Index = index
	return snapshot, nil
}

func (kv *SnapshotKV) WaitSnapshotIndex() uint64 {
	return kv.waitSnapshotIndex
}

func (kv *SnapshotKV) BackupSnapshotIndex() uint64 {
	return kv.backupSnapshotIndex
}

func (kv *SnapshotKV) Close() {
	if kv.kvSnapshotIndex != nil {
		kv.kvSnapshotIndex.Close()
	}
}
