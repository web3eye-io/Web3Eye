package backup

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
)

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
		logger.Sugar().Infow(
			"backupAll",
			"Wait", wait,
		)
	}
}

var newSnapshot chan struct{}

func Watch(ctx context.Context) {
	newSnapshot = make(chan struct{})
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
