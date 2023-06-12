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
	KVStoreSnapshotIndex       = "snapshot-index"
	KVStoreWaitSnapshot        = "wait-snapshot-"
	KVStoreBackupSnapshot      = "backup-snapshot-"
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

	return nil
}

func (kv *SnapshotKV) SetWaitSnapshot(ctx context.Context, snapshotURI string, items []*dealerpb.ContentItem) error {
	return kv.setSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreWaitSnapshot, kv.waitSnapshotIndex), snapshotURI, items)
}

func (kv *SnapshotKV) getSnapshot(ctx context.Context, kvStoreName string) (snapshotURI string, items []*dealerpb.ContentItem, err error) {
	replicate := true
	_kv, err := kv.odb.KeyValue(ctx, kvStoreName, &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return "", nil, err
	}
	defer _kv.Close()

	if err := _kv.Load(ctx, -1); err != nil {
		return "", nil, err
	}

	_snapshotURI, err := _kv.Get(ctx, SnapshotURI)
	if err != nil {
		return "", nil, err
	}

	_items, err := _kv.Get(ctx, ContentItems)
	if err != nil {
		return "", nil, err
	}

	if err := json.Unmarshal(_items, &items); err != nil {
		return "", nil, err
	}

	return string(_snapshotURI), items, nil
}

func (kv *SnapshotKV) GetWaitSnapshot(ctx context.Context) (snapshotURI string, items []*dealerpb.ContentItem, err error) {
	return kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreWaitSnapshot, kv.waitSnapshotIndex))
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

func (kv *SnapshotKV) SetBackupSnapshot(ctx context.Context, snapshotURI string, items []*dealerpb.ContentItem) error {
	return kv.setSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreBackupSnapshot, kv.backupSnapshotIndex), snapshotURI, items)
}

func (kv *SnapshotKV) GetBackupSnapshot(ctx context.Context) (snapshotURI string, items []*dealerpb.ContentItem, err error) {
	return kv.getSnapshot(ctx, fmt.Sprintf("%v%v", KVStoreBackupSnapshot, kv.backupSnapshotIndex))
}

func (kv *SnapshotKV) Close() {
	if kv.kvSnapshotIndex != nil {
		kv.kvSnapshotIndex.Close()
	}
}
