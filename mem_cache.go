package gcache

import (
	"encoding/json"
	"errors"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type memCache struct {
	inner *cache.Cache
	mu    *sync.Mutex
}

var _ CacheStore = (*memCache)(nil)
var ErrKeyNotExist = errors.New("key not exist")

func (c *memCache) IncrBy(key string, v int) (result Result) {
	c.mu.Lock()
	_, exist := c.inner.Get(key)
	if !exist {
		c.Set(key, v, 0)
		c.mu.Unlock()
		return memResult{
			reply: v,
			err:   nil,
		}
	}

	err := c.inner.Increment(key, int64(v))
	if err != nil {
		c.mu.Unlock()
		return memResult{
			reply: nil,
			err:   err,
		}
	}
	result = c.Get(key)
	c.mu.Unlock()
	return result
}

func (c *memCache) Incr(key string) (result Result) {
	return c.IncrBy(key, 1)
}

func (c *memCache) Expire(key string, sec int) error {
	c.mu.Lock()
	res := c.Get(key)
	if res.Error() != nil {
		return res.Error()
	}
	err := c.Set(key, res.Reply(), sec)
	c.mu.Unlock()
	return err
}

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
		return memResult{reply: nil, err: ErrKeyNotExist}
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
