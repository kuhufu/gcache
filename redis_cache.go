package gcache

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"time"
)

// redis缓存
type redisCache struct {
	inner *flyredis.Pool
}

var _ CacheStore = (*redisCache)(nil)

type RedisOption struct {
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	Wait            bool
	MaxConnLifetime time.Duration
	Password        string
	TestOnBorrow    func(c redis.Conn, t time.Time) error
	DialOptions     []redis.DialOption
}

func (c *redisCache) Incr(key string) (result Result) {
	return c.inner.Do("INCR", key)
}

func (c *redisCache) IncrBy(key string, v int) (result Result) {
	return c.inner.Do("INCRBY", key, v)
}

func (c *redisCache) Expire(key string, sec int) error {
	return c.inner.Do("EXPIRE", key, sec).Error()
}

func (c *redisCache) Set(key string, val interface{}, expireSeconds int) error {
	var err error
	if expireSeconds <= 0 {
		err = c.inner.Do("SET", key, val).Error()
		return err
	}
	err = c.inner.Do("SET", key, val, "EX", expireSeconds).Error()
	return err
}

func (c *redisCache) Get(key string) (result Result) {
	return c.inner.Do("GET", key)
}

func (c *redisCache) Del(key string) (err error) {
	return c.inner.Do("DEL", key).Error()
}

func (c *redisCache) GetUnmarshal(key string) (value interface{}, err error) {
	reply, err := c.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reply, &value)
	return
}
