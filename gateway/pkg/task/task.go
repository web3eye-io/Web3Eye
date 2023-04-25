package task

import (
	"context"

	"github.com/web3eye-io/Web3Eye/gateway/pkg/streammgr"
)

const (
	StreamNum = 2
)

func Run(ctx context.Context) {
	for i := 0; i < StreamNum; i++ {
		sc := streammgr.NewStreamClient()
		go sc.Start(ctx)
	}
}
