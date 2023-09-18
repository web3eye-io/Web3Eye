package backup

import (
	"context"
	"fmt"
	"strconv"

	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	orbitdb "berty.tech/go-orbit-db"
	orbitiface "berty.tech/go-orbit-db/iface"
)

type BackupKV struct {
	kvBackup orbitdb.KeyValueStore
	creates  map[uint64]struct{}
	waits    map[uint64]struct{}
	dones    map[uint64]struct{}
}

const (
	KVStoreBackup = "backup"
)

func NewBackupKV(ctx context.Context, odb orbitiface.OrbitDB) (*BackupKV, error) {
	kv := &BackupKV{
		creates: map[uint64]struct{}{},
		waits:   map[uint64]struct{}{},
		dones:   map[uint64]struct{}{},
	}
	var err error

	replicate := true
	kv.kvBackup, err = odb.KeyValue(ctx, KVStoreBackup, &orbitdb.CreateDBOptions{
		Replicate: &replicate,
	})
	if err != nil {
		return nil, err
	}

	if err := kv.kvBackup.Load(ctx, -1); err != nil {
		return nil, err
	}

	backups := kv.kvBackup.All()
	for index, state := range backups {
		//nolint:gomnd
		_index, err := strconv.ParseUint(index, 10, 64)
		if err != nil {
			return nil, err
		}
		switch string(state) {
		case dealerpb.BackupState_BackupStateCreated.String():
			kv.creates[_index] = struct{}{}
		case dealerpb.BackupState_BackupStateProposed.String():
			fallthrough //nolint
		case dealerpb.BackupState_BackupStateAccepted.String():
			kv.waits[_index] = struct{}{}
		case dealerpb.BackupState_BackupStateSuccess.String():
			kv.dones[_index] = struct{}{}
		}
	}

	return kv, nil
}

func (kv *BackupKV) Create(ctx context.Context, index uint64) error {
	if kv.kvBackup == nil {
		return fmt.Errorf("invalid keyvalue")
	}

	val, err := kv.kvBackup.Get(ctx, fmt.Sprintf("%v", index))
	if err != nil {
		return err
	}

	switch string(val) {
	case dealerpb.BackupState_BackupStateProposed.String():
		fallthrough //nolint
	case dealerpb.BackupState_BackupStateAccepted.String():
		fallthrough //nolint
	case dealerpb.BackupState_BackupStateSuccess.String():
		return fmt.Errorf("already created")
	}

	if _, err := kv.kvBackup.Put(ctx, fmt.Sprintf("%v", index), []byte(dealerpb.BackupState_BackupStateCreated.String())); err != nil {
		return err
	}
	kv.creates[index] = struct{}{}

	return nil
}

func (kv *BackupKV) Creates(ctx context.Context) ([]uint64, error) {
	creates := []uint64{}
	for index := range kv.creates {
		creates = append(creates, index)
	}
	return creates, nil
}

func (kv *BackupKV) Wait(ctx context.Context, index uint64) error {
	if kv.kvBackup == nil {
		return fmt.Errorf("invalid keyvalue")
	}

	val, err := kv.kvBackup.Get(ctx, fmt.Sprintf("%v", index))
	if err != nil {
		return err
	}

	switch string(val) {
	case dealerpb.BackupState_BackupStateCreated.String():
	default:
		return fmt.Errorf("invalid state")
	}

	if _, err := kv.kvBackup.Put(ctx, fmt.Sprintf("%v", index), []byte(dealerpb.BackupState_BackupStateProposed.String())); err != nil {
		return err
	}
	delete(kv.creates, index)
	kv.waits[index] = struct{}{}

	return nil
}

func (kv *BackupKV) Waits(ctx context.Context) ([]uint64, error) {
	waits := []uint64{}
	for index := range kv.waits {
		waits = append(waits, index)
	}
	return waits, nil
}

func (kv *BackupKV) Done(ctx context.Context, index uint64, fail bool) error {
	if kv.kvBackup == nil {
		return fmt.Errorf("invalid keyvalue")
	}

	val, err := kv.kvBackup.Get(ctx, fmt.Sprintf("%v", index))
	if err != nil {
		return err
	}

	switch string(val) {
	case dealerpb.BackupState_BackupStateProposed.String():
	default:
		return fmt.Errorf("invalid state")
	}

	state := dealerpb.BackupState_BackupStateSuccess
	if fail {
		state = dealerpb.BackupState_BackupStateFail
	}

	if _, err := kv.kvBackup.Put(ctx, fmt.Sprintf("%v", index), []byte(state.String())); err != nil {
		return err
	}
	delete(kv.waits, index)

	if !fail {
		kv.dones[index] = struct{}{}
	}

	return nil
}

func (kv *BackupKV) Dones(ctx context.Context) ([]uint64, error) {
	dones := []uint64{}
	for index, _ := range kv.dones {
		dones = append(dones, index)
	}
	return dones, nil
}

func (kv *BackupKV) Close() {
	if kv.kvBackup != nil {
		kv.kvBackup.Close()
	}
}
