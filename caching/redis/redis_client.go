package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/k8guard/k8guardlibs/caching/types"
	"github.com/k8guard/k8guardlibs/config"
)

type redisCache struct {
	client redis.Client
}

func NewCache(cfg config.Config) (types.Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RmqBroker,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &redisCache{client: *client}, nil
}

func (cache *redisCache) Set(key string, value interface{}, expiration time.Duration) error {
	err := cache.client.Set(key, value, expiration).Err()
	return err
}

func (cache *redisCache) Get(key string) (interface{}, error) {
	val, err := cache.client.Get(key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return val, nil
	}
}
