package ctredis

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func TryLock(key string, expire time.Duration) (string, error) {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	lockID := uuid.New().String()
	resp := cli.SetNX(ctx, key, lockID, expire)
	locked, err := resp.Result()
	err = ErrFilter(err)
	if err != nil {
		return "", xerrors.Errorf("fail lock: %v", err)
	}

	if !locked {
		return "", xerrors.Errorf("fail lock")
	}

	return lockID, nil
}

func Unlock(lockKey, lockID string) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	_lockID, err := cli.Get(ctx, lockKey).Result()
	err = ErrFilter(err)
	if err != nil {
		return err
	}

	if _lockID != lockID {
		return errors.New("lockID not match")
	}

	err = cli.Del(ctx, lockKey).Err()
	return ErrFilter(err)
}
