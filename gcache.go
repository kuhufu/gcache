package gcache

import (
	"github.com/coocood/freecache"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"time"
)

type CacheStore interface {
	//expireSeconds <= 0 表示永不过期
	Set(key string, val []byte, expireSeconds int) error

	Get(key string) (value []byte, err error)

	Del(key string) (affected bool)

	GetUnmarshal(key string) (value interface{}, err error)
}

func NewMemCache(size int) CacheStore {
	return &memCache{
		inner: freecache.NewCache(size),
	}
}

func NewRedisCache(size int, network, address, password string) CacheStore {
	return &redisCache{
		inner: flyredis.NewPool(&redis.Pool{
			MaxIdle:     size,
			IdleTimeout: 240 * time.Second,
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
			Dial: func() (redis.Conn, error) {
				return dial(network, address, password)
			},
		}),
	}
}
