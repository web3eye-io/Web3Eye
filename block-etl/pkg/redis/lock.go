package redis

import (
	"context"
	"time"

	"golang.org/x/xerrors"
)

func TryLock(key string, expire time.Duration) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	resp := cli.SetNX(ctx, key, 1, expire)
	locked, err := resp.Result()
	if err != nil {
		return xerrors.Errorf("fail lock: %v", err)
	}

	if !locked {
		return xerrors.Errorf("fail lock")
	}

	return nil
}

func Unlock(key string) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	return cli.Del(ctx, key).Err()
}
