package gcache

import (
	"fmt"
	"testing"
)

var redisCache = NewRedisCache(10, "tcp", "127.0.0.1:6379", "")
var redisKey = "gcache_redis_cache_test_key"

func TestRedisCache_Set(t *testing.T) {
	_ = redisCache.Set(redisKey, []byte("test_data"), -1)
}

func TestRedisCache_Get(t *testing.T) {
	data, err := redisCache.Get(redisKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestRedisCache_Del(t *testing.T) {
	if !redisCache.Del(redisKey) {
		t.Fatal("not affected")
	}
}
