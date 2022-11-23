package redis

import (
	"errors"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	poolSize    = 50
	lk          sync.RWMutex

	ErrRedisClientNotInit = errors.New("redis client not init")
)

func GetClient() *redis.Client {
	lk.Lock()
	defer lk.Unlock()

	// double read
	if redisClient != nil {
		return redisClient
	}

	// TODO: should be managed by config
	service := "127.0.0.1:6379"
	password := ""

	redisClient = redis.NewClient(&redis.Options{
		Addr:     service,
		Password: password,
		DB:       0,
		PoolSize: poolSize,
	})

	return redisClient
}

func Close() error {
	lk.Lock()
	defer lk.Unlock()

	if redisClient != nil {
		redisClient.Close()
		redisClient = nil
	}

	return nil
}
