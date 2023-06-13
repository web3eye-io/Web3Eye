package filestate

import (
	"context"
	"fmt"

	orbitdb "berty.tech/go-orbit-db"
	orbitiface "berty.tech/go-orbit-db/iface"
)

const (
	Invalid           = "Invalid"
	Waiting           = "Waiting"
	IpfsUploading     = "IpfsUploading"
	FilecoinUploading = "FilecoinUploading"
	FilecoinSealing   = "FilecoinSealing"
	BackupSuccess     = "BackupSuccess"
	BackupFail        = "BackupFail"
)

type FileState string

type FileStateKV struct {
	kvFileState orbitdb.KeyValueStore
}

const (
	KVStoreFileState = "file-state"
)

func NewFileStateKV(ctx context.Context, odb orbitiface.OrbitDB) (*FileStateKV, error) {
	kv := &FileStateKV{}
	var err error

	replicate := true
	kv.kvFileState, err = odb.KeyValue(ctx, KVStoreFileState, &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return nil, err
	}

	if err := kv.kvFileState.Load(ctx, -1); err != nil {
		return nil, err
	}

	return kv, nil
}

func (kv *FileStateKV) SetFileState(ctx context.Context, chainType, uid, chainId, state string) error {
	if kv.kvFileState == nil {
		return fmt.Errorf("invalid kvstore")
	}

	key := fmt.Sprintf("%v:%v:%v", chainType, uid, chainId)
	if _, err := kv.kvFileState.Put(ctx, key, []byte(state)); err != nil {
		return err
	}

	return nil
}

func (kv *FileStateKV) GetFileState(ctx context.Context, chainType, uid, chainId string) (FileState, error) {
	if kv.kvFileState == nil {
		return Invalid, fmt.Errorf("invalid kvstore")
	}

	key := fmt.Sprintf("%v:%v:%v", chainType, uid, chainId)
	b, err := kv.kvFileState.Get(ctx, key)
	if err != nil {
		return Invalid, err
	}

	return FileState(string(b)), nil
}

func (kv *FileStateKV) Close() {
	if kv.kvFileState != nil {
		kv.kvFileState.Close()
	}
}
