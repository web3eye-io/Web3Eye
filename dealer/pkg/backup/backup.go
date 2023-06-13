package backup

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"

	"github.com/filecoin-project/go-fil-markets/storagemarket/network"
	"github.com/libp2p/go-libp2p/core/host"
)

type backup struct {
	host   host.Host
	stream network.StorageDealStream
}

func (b *backup) backupOne(ctx context.Context, index uint64) error {
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"backupOne",
		"Snapshot", snapshot,
		"Index", index,
	)

	switch snapshot.BackupState {
	case deaperpb.BackupState_BackupStateSuccess:
		fallthrough // nolint
	case deaperpb.BackupState_BackupStateFail:
		return nil
	default:
	}

	if snapshot.SnapshotCID == "" {
		return fmt.Errorf("invalid snapshot cid")
	}
	if snapshot.SnapshotURI == "" {
		return fmt.Errorf("invalid snapshot uri")
	}

	// TODO: backup items to IPFS
	// TODO: backup car to Filecoin

	return nil
}

func (b *backup) backupAll(ctx context.Context) {
	waits, err := orbit.Backup().Waits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"backupAll",
			"Error", err,
		)
		return
	}
	for _, wait := range waits {
		if err := b.backupOne(ctx, wait); err != nil {
			logger.Sugar().Errorw(
				"backupAll",
				"Wait", wait,
				"Error", err,
			)
		}
	}
}

var newSnapshot chan struct{}

func Watch(ctx context.Context) (err error) {
	backup := &backup{}

	if err := backup.buildHost(ctx); err != nil {
		return err
	}
	if err := backup.connectMiner(ctx); err != nil {
		return err
	}

	newSnapshot = make(chan struct{})
	backup.backupAll(ctx)

	for {
		select {
		case <-newSnapshot:
			backup.backupAll(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func NewSnapshot() {
	go func() {
		newSnapshot <- struct{}{}
	}()
}
