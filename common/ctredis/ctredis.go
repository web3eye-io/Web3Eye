package ctredis

import (
	"errors"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/web3eye-io/cyber-tracer/config"
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

	service := config.GetConfig().Redis.Address
	password := config.GetConfig().Redis.Password

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
