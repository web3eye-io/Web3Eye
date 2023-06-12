package backup

import (
	"context"
)

func backupAll(ctx context.Context) {

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
