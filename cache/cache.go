package cache

import (
	"time"

	redis "github.com/go-redis/redis/v7"
)

// CacheDescriptor describe the cache interface
type CacheDescriptor interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Ping() (string, error)
}

// CacheManager defined cache manager struct
type CacheManager struct {
	Client CacheDescriptor
}

type NoOpCache struct {
	client *redis.Client
}

func (cm *NoOpCache) Set(key string, value interface{}, expiration time.Duration) (err error) {
	return
}

func (cm *NoOpCache) Get(key string) (value string, err error) {
	return

}

func (cm *NoOpCache) Ping() (value string, err error) {
	return
}
