package gcache

import (
	"github.com/coocood/freecache"
	"github.com/gomodule/redigo/redis"
	"github.com/kuhufu/flyredis"
	"github.com/kuhufu/scheduler"
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

var s = scheduler.New()

func StartSchedule() {
	s.Start()
}

func StopSchedule() {
	s.Stop()
}

//@param immediately 是否立刻执行一次fetch函数
func Interval(store CacheStore, key string, seconds int, fetch func() []byte, immediately bool) {
	if immediately {
		store.Set(key, fetch(), -1)
	}
	s.AddIntervalFunc(time.Duration(seconds)*time.Second, func() {
		store.Set(key, fetch(), -1)
	})
}
