package filestate

import (
	"context"
	"encoding/binary"
	"fmt"

	orbitdb "berty.tech/go-orbit-db"
	orbitiface "berty.tech/go-orbit-db/iface"

	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
)

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

func (kv *FileStateKV) SetFileState(ctx context.Context, chainType, uid, chainId string, state dealerpb.BackupState) error {
	if kv.kvFileState == nil {
		return fmt.Errorf("invalid kvstore")
	}

	key := fmt.Sprintf("%v:%v:%v:state", chainType, uid, chainId)
	if _, err := kv.kvFileState.Put(ctx, key, []byte(state.String())); err != nil {
		return err
	}

	return nil
}

func (kv *FileStateKV) GetFileState(ctx context.Context, chainType, uid, chainId string) (dealerpb.BackupState, error) {
	if kv.kvFileState == nil {
		return dealerpb.BackupState_DefaultBackupState, fmt.Errorf("invalid kvstore")
	}

	key := fmt.Sprintf("%v:%v:%v:state", chainType, uid, chainId)
	b, err := kv.kvFileState.Get(ctx, key)
	if err != nil {
		return dealerpb.BackupState_DefaultBackupState, err
	}

	return dealerpb.BackupState(dealerpb.BackupState_value[string(b)]), nil
}

func (kv *FileStateKV) SetFileSnapshot(ctx context.Context, chainType, uid, chainId string, index uint64) error {
	if kv.kvFileState == nil {
		return fmt.Errorf("invalid kvstore")
	}

	key := fmt.Sprintf("%v:%v:%v:snapshot", chainType, uid, chainId)
	b := make([]byte, 8)
	binary.PutUvarint(b, index)

	if _, err := kv.kvFileState.Put(ctx, key, b); err != nil {
		return err
	}

	return nil
}

func (kv *FileStateKV) GetFileSnapshot(ctx context.Context, chainType, uid, chainId string) (uint64, error) {
	if kv.kvFileState == nil {
		return 0, fmt.Errorf("invalid kvstore")
	}

	key := fmt.Sprintf("%v:%v:%v:snapshot", chainType, uid, chainId)
	b, err := kv.kvFileState.Get(ctx, key)
	if err != nil {
		return 0, err
	}
	index, _ := binary.Uvarint(b)

	return index, nil
}

func (kv *FileStateKV) SetFileRetrieve(ctx context.Context, chainType, uid, chainId string, state string) error {
	if kv.kvFileState == nil {
		return fmt.Errorf("invalid kvstore")
	}

	key := fmt.Sprintf("%v:%v:%v:retrieve", chainType, uid, chainId)
	if _, err := kv.kvFileState.Put(ctx, key, []byte(state)); err != nil {
		return err
	}

	return nil
}

func (kv *FileStateKV) GetFileRetrieve(ctx context.Context, chainType, uid, chainId string) (string, error) {
	if kv.kvFileState == nil {
		return "", fmt.Errorf("invalid kvstore")
	}

	key := fmt.Sprintf("%v:%v:%v:retrieve", chainType, uid, chainId)
	b, err := kv.kvFileState.Get(ctx, key)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (kv *FileStateKV) Close() {
	if kv.kvFileState != nil {
		kv.kvFileState.Close()
	}
}
