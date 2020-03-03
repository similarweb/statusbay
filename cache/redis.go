package cache

import (
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

// RedisConfig describe redis configuration
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type RedisCache struct {
	client *redis.Client
}

// NewRedisClient create new redis client
func NewRedisClient(config *RedisConfig) (cacheManager *CacheManager) {

	cacheManager = &CacheManager{
		Client: &NoOpCache{},
	}

	if config == nil {
		return
	}

	var client *redis.Client
	c := make(chan int, 1)
	go func() {
		for {
			client = redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%d", config.Addr, config.Port),
				Password: config.Password,
				DB:       config.DB,
			})

			_, err := client.Ping().Result()
			if err == nil {
				break
			}
			log.Warn("could not initialize connection to redis, retrying for 5 seconds")
			time.Sleep(5 * time.Second)
		}
		c <- 1
	}()

	select {
	case <-c:
	case <-time.After(60 * time.Second):
		log.Error("could not connect redis, timed out after 1 minute")
		return
	}

	cacheManager = &CacheManager{
		Client: &RedisCache{
			client: client,
		},
	}

	return

}

// Set key value [expiration]` command.
func (cm *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return cm.client.Set(key, fmt.Sprintf("%v", value), expiration).Err()
}

// Get key command. It returns redis.Nil error when key does not exist.
func (cm *RedisCache) Get(key string) (string, error) {
	return cm.client.Get(key).Result()

}

// Ping to redis server
func (cm *RedisCache) Ping() (string, error) {
	return cm.client.Ping().Result()
}
