package snapshot

import (
	"context"
	"encoding/binary"
	"fmt"

	orbitdb "berty.tech/go-orbit-db"
	orbitiface "berty.tech/go-orbit-db/iface"
)

type SnapshotKV struct {
	kvSnapshotIndex orbitdb.KeyValueStore
	snapshotIndex   uint64
}

const (
	CurrentSnapshotIndex = "current-index"
	KVStoreSnapshotIndex = "snapshot-index"
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

	b, err := kv.kvSnapshotIndex.Get(ctx, CurrentSnapshotIndex)
	if err != nil {
		return nil, err
	}
	kv.snapshotIndex, _ = binary.Uvarint(b)

	if kv.snapshotIndex == 0 {
		if err := kv.NextSnapshot(ctx); err != nil {
			return nil, err
		}
	}

	return kv, nil
}

func (kv *SnapshotKV) NextSnapshot(ctx context.Context) error {
	if kv.kvSnapshotIndex == nil {
		return fmt.Errorf("invalid kvstore")
	}

	b := make([]byte, 8)
	kv.snapshotIndex += 1
	binary.PutUvarint(b, kv.snapshotIndex)
	if _, err := kv.kvSnapshotIndex.Put(ctx, CurrentSnapshotIndex, b); err != nil {
		return err
	}

	return nil
}

func (kv *SnapshotKV) Close() {
	if kv.kvSnapshotIndex != nil {
		kv.kvSnapshotIndex.Close()
	}
}
