package task

import (
	"context"

	"github.com/web3eye-io/Web3Eye/gateway/pkg/streammgr"
)

const (
	StreamNum = 2
)

func Run(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	for i := 0; i < StreamNum; i++ {
		sc := &streammgr.StreamClient{}
		go sc.Start(ctx, cancel)
	}
	<-ctx.Done()
}
