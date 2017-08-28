package types

import "time"

type CacheType string

const (
	// supported caches
	MEMCACHE_CACHE CacheType = "MEMCACHE"
	REDIS_CACHE    CacheType = "REDIS"
)

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (interface{}, error)
}
