package ctredis

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// Lock with lockID
func TryLock(key string, expire time.Duration) (string, error) {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), DefaultLockTime)
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

	ctx, cancel := context.WithTimeout(context.Background(), DefaultLockTime)
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

// Lock without lockID
func TryPubLock(key string, expire time.Duration) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), DefaultLockTime)
	defer cancel()

	resp := cli.SetNX(ctx, key, true, expire)
	locked, err := resp.Result()
	err = ErrFilter(err)
	if err != nil {
		return xerrors.Errorf("fail lock: %v", err)
	}

	if !locked {
		return xerrors.Errorf("fail lock")
	}

	return nil
}

func UnPubLock(lockKey string) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), DefaultLockTime)
	defer cancel()

	_, err := cli.Get(ctx, lockKey).Result()
	err = ErrFilter(err)
	if err != nil {
		return err
	}

	err = cli.Del(ctx, lockKey).Err()
	return ErrFilter(err)
}
