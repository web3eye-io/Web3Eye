package ctredis

import (
	"context"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

const (
	ConnectTimeout = 5 * time.Minute
)

func Set(key string, value interface{}, expire time.Duration) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	err := cli.Set(ctx, key, value, expire).Err()
	return errFilter(err)
}

func Get(key string) (interface{}, error) {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	v, err := cli.Get(ctx, key).Result()
	err = errFilter(err)
	if err != nil {
		return nil, xerrors.Errorf("fail get key %v: %v", key, err)
	}

	return v, nil
}

func Del(key string) error {
	cli := GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	err := cli.Del(ctx, key).Err()
	return errFilter(err)
}

func errFilter(err error) error {
	if strings.Contains(err.Error(), "MOVED") {
		return nil
	}
	return err
}
