package ctredis

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const (
	ConnectTimeout  = 3 * time.Second
	DefaultLockTime = 10 * time.Minute
)

func Set(key string, value interface{}, expire time.Duration) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	err := cli.Set(ctx, key, value, expire).Err()
	return ErrFilter(err)
}

func Get(key string, v any) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	err := cli.Get(ctx, key).Scan(v)
	err = ErrFilter(err)
	if err != nil {
		return fmt.Errorf("fail get key %v: %v", key, err)
	}

	return nil
}

func Del(key string) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	err := cli.Del(ctx, key).Err()
	return ErrFilter(err)
}

func ErrFilter(err error) error {
	if err != nil && strings.Contains(err.Error(), "MOVED") {
		return nil
	}
	return err
}
