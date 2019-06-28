package gcache

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
)

// redis缓存
type redisCache struct {
	inner *flyredis.Pool
}

var _ CacheStore = (*redisCache)(nil)

func (c *redisCache) Set(key string, val []byte, expireSeconds int) error {
	var err error
	if expireSeconds <= 0 {
		err = c.inner.SET(key, val).Error()
		return err
	}
	err = c.inner.SET(key, val, "EX", expireSeconds).Error()
	return err
}

func (c *redisCache) Get(key string) (value []byte, err error) {
	return c.inner.GET(key).Bytes()
}

func (c *redisCache) Del(key string) (affected bool) {
	reply, _ := c.inner.DEL(key).Int()
	return reply != 0
}

func (c *redisCache) GetUnmarshal(key string) (value interface{}, err error) {
	reply, err := c.inner.GET(key).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reply, &value)
	return
}

func dial(network, address, password string) (redis.Conn, error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, err := c.Do("AUTH", password); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, err
}
