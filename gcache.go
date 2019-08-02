package gcache

import (
	"github.com/kuhufu/flyredis"
	"github.com/patrickmn/go-cache"
	"time"
)

type CacheStore interface {
	//expireSeconds <= 0 表示永不过期
	Set(key string, val interface{}, expireSeconds int) error
	Get(key string) (result Result)
	Del(key string) (err error)
	GetUnmarshal(key string) (value interface{}, err error)
}

type Result interface {
	Value() (interface{}, error)
	Bool() (reply bool, err error)
	Int() (reply int, err error)
	Int64() (reply int64, err error)
	Float64() (reply float64, err error)
	String() (reply string, err error)
	Bytes() (reply []byte, err error)
	Reply() (reply interface{})
	Error() (err error)
}

func NewMemCache() CacheStore {
	return &memCache{
		inner: cache.New(0, time.Second*60),
	}
}

func NewRedisCache(network, address string, option RedisOption) CacheStore {
	return &redisCache{
		inner: flyredis.NewPool(network, address, flyredis.Option{
			MaxIdle:      option.MaxIdle,
			MaxActive:    option.MaxActive,
			IdleTimeout:  option.IdleTimeout,
			Wait:         option.Wait,
			Password:     option.Password,
			TestOnBorrow: option.TestOnBorrow,
			DialOptions:  option.DialOptions,
		}),
	}
}
