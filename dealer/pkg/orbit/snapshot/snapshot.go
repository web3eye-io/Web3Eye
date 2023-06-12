package snapshot

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"

	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	orbitdb "berty.tech/go-orbit-db"
	orbitiface "berty.tech/go-orbit-db/iface"
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
	SnapshotURI                = "snapshot-url"
	ContentItems               = "content-items"
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

func (kv *SnapshotKV) setSnapshot(ctx context.Context, kvStoreName, snapshotURI string, items []*dealerpb.ContentItem) error {
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

func (kv *SnapshotKV) SetWaitSnapshot(ctx context.Context, snapshotURI string, items []*dealerpb.ContentItem) error {
	return kv.setSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, kv.waitSnapshotIndex), snapshotURI, items)
}

func (kv *SnapshotKV) getSnapshot(ctx context.Context, kvStoreName string) (snapshotURI string, items []*dealerpb.ContentItem, state dealerpb.BackupState, err error) {
	replicate := true
	_kv, err := kv.odb.KeyValue(ctx, kvStoreName, &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}
	defer _kv.Close()

	if err := _kv.Load(ctx, -1); err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}

	_snapshotURI, err := _kv.Get(ctx, SnapshotURI)
	if err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}

	_items, err := _kv.Get(ctx, ContentItems)
	if err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}
	if err := json.Unmarshal(_items, &items); err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}

	_state, err := _kv.Get(ctx, SnapshotBackupState)
	if err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}
	state = dealerpb.BackupState(dealerpb.BackupState_value[string(_state)])

	return string(_snapshotURI), items, state, nil
}

func (kv *SnapshotKV) GetCurrentWaitSnapshot(ctx context.Context) (snapshotURI string, items []*dealerpb.ContentItem, state dealerpb.BackupState, err error) {
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

func (kv *SnapshotKV) GetCurrentBackupSnapshot(ctx context.Context) (snapshotURI string, items []*dealerpb.ContentItem, state dealerpb.BackupState, err error) {
	return kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, kv.backupSnapshotIndex))
}

func (kv *SnapshotKV) GetSnapshot(ctx context.Context, index uint64) (snapshotURI string, items []*dealerpb.ContentItem, state dealerpb.BackupState, err error) {
	if index >= kv.backupSnapshotIndex {
		return "", nil, dealerpb.BackupState_DefaultBackupState, fmt.Errorf("invalid snapshot index")
	}
	return kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, index))
}

func (kv *SnapshotKV) UpdateSnapshot(ctx context.Context, index uint64, state dealerpb.BackupState) (string, []*dealerpb.ContentItem, dealerpb.BackupState, error) {
	if index >= kv.backupSnapshotIndex {
		return "", nil, dealerpb.BackupState_DefaultBackupState, fmt.Errorf("invalid snapshot index")
	}

	replicate := true
	_kv, err := kv.odb.KeyValue(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, index), &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}
	defer _kv.Close()

	if err := _kv.Load(ctx, -1); err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}

	if _, err := _kv.Put(ctx, SnapshotBackupState, []byte(state.String())); err != nil {
		return "", nil, dealerpb.BackupState_DefaultBackupState, err
	}

	return kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreSnapshot, index))
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
