package backup

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
)

func backOne(ctx context.Context, index uint64) error {
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"backOne",
		"Snapshot", snapshot,
		"Index", index,
	)

	return nil
}

func backupAll(ctx context.Context) {
	waits, err := orbit.Backup().Waits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"backupAll",
			"Error", err,
		)
		return
	}
	for _, wait := range waits {
		if err := backOne(ctx, wait); err != nil {
			logger.Sugar().Errorw(
				"backupAll",
				"Wait", wait,
				"Error", err,
			)
		}
	}
}

var newSnapshot chan struct{}

func Watch(ctx context.Context) {
	newSnapshot = make(chan struct{})

	backupAll(ctx)

	for {
		select {
		case <-newSnapshot:
			backupAll(ctx)
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
