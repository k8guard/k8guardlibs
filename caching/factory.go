package caching

import (
	"errors"

	"github.com/k8guard/k8guardlibs/caching/memcache"
	"github.com/k8guard/k8guardlibs/caching/redis"
	"github.com/k8guard/k8guardlibs/caching/types"
	"github.com/k8guard/k8guardlibs/config"
)

func CreateCache(p types.CacheType, cfg config.Config) (types.Cache, error) {
	switch p {
	case types.MEMCACHE_CACHE:
		return memcache.NewCache(cfg)
	case types.REDIS_CACHE:
		return redis.NewCache(cfg)
	default:
		return nil, errors.New("Invalid Cache Type")
	}
}
