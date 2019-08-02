package gcache

import (
	"encoding/json"
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

type memCache struct {
	inner *cache.Cache
}

var _ CacheStore = (*memCache)(nil)

func (c *memCache) Set(key string, val interface{}, expireSeconds int) (err error) {
	if expireSeconds <= 0 {
		c.inner.Set(key, val, -1)
	} else {
		c.inner.Set(key, val, time.Duration(expireSeconds)*time.Second)
	}
	return
}

func (c *memCache) Get(key string) Result {
	data, exist := c.inner.Get(key)
	if !exist {
		return memResult{reply: nil, err: errors.New(key + " not exist")}
	}

	return memResult{reply: data, err: nil}
}

func (c *memCache) Del(key string) (err error) {
	c.inner.Delete(key)
	return
}

func (c *memCache) GetUnmarshal(key string) (value interface{}, err error) {
	data, err := c.Get(key).Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &value)
	return
}
