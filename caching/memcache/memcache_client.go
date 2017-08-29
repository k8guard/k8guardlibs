package memcache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/k8guard/k8guardlibs/caching/types"
	"github.com/k8guard/k8guardlibs/config"
)

type memcacheCache struct {
	client memcache.Client
}

func NewCache(cfg config.Config) (types.Cache, error) {
	client := memcache.New(fmt.Sprintf("%s:11211",
		cfg.MemCachedHostname))

	return &memcacheCache{client: *client}, nil
}

func (cache *memcacheCache) Set(key string, value interface{}, expiration time.Duration) error {
	exp := int32(expiration) / 1000
	err := cache.client.Set(&memcache.Item{Key: key, Value: getBytes(value), Expiration: exp})
	return err
}

func (cache *memcacheCache) Get(key string) (interface{}, error) {
	return cache.client.Get(key)
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
