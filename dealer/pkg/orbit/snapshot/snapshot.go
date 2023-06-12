package snapshot

import (
	"context"
	"encoding/binary"
	"fmt"

	orbitdb "berty.tech/go-orbit-db"
	orbitiface "berty.tech/go-orbit-db/iface"
)

type SnapshotKV struct {
	kvSnapshotIndex     orbitdb.KeyValueStore
	waitSnapshotIndex   uint64
	backupSnapshotIndex uint64
}

const (
	CurrentWaitSnapshotIndex   = "current-wait-index"
	CurrentBackupSnapshotIndex = "current-wait-index"
	KVStoreSnapshotIndex       = "snapshot-index"
)

func NewSnapshotKV(ctx context.Context, odb orbitiface.OrbitDB) (*SnapshotKV, error) {
	kv := &SnapshotKV{}
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

func (kv *SnapshotKV) Close() {
	if kv.kvSnapshotIndex != nil {
		kv.kvSnapshotIndex.Close()
	}
}
