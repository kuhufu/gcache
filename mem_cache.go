package gcache

import (
	"encoding/json"
	"github.com/coocood/freecache"
)

// 应用内存缓存
type memCache struct {
	inner *freecache.Cache
}

var _ CacheStore = (*memCache)(nil)

func (c *memCache) Set(key string, val []byte, expireSeconds int) error {
	return c.inner.Set([]byte(key), val, expireSeconds)
}

func (c *memCache) Get(key string) (value []byte, err error) {
	return c.inner.Get([]byte(key))
}

func (c *memCache) Del(key string) (affected bool) {
	return c.inner.Del([]byte(key))
}

func (c *memCache) GetUnmarshal(key string) (value interface{}, err error) {
	data, err := c.inner.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &value)
	return
}
