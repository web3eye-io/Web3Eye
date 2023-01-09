package ctredis

import (
	"errors"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/web3eye-io/cyber-tracer/config"
)

var (
	redisClient *redis.ClusterClient
	poolSize    = 50
	lk          sync.RWMutex

	ErrRedisClientNotInit = errors.New("redis client not init")
)

func GetClient() *redis.ClusterClient {
	lk.Lock()
	defer lk.Unlock()

	// double read
	if redisClient != nil {
		return redisClient
	}

	service := config.GetConfig().Redis.Address
	password := config.GetConfig().Redis.Password

	redisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{service},
		Password: password,
		PoolSize: poolSize,
	})
	// check wheather is cluster mode

	// TODO: should check wheather is cluster,and auto start with cluster clent
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
