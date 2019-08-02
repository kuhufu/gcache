package gcache

import (
	"fmt"
	"testing"
	"time"
)

var cacheStore = NewRedisCache("tcp", "127.0.0.1:6379", RedisOption{
	MaxIdle:     10,
	MaxActive:   30,
	IdleTimeout: time.Second * 180,
	Password:    "",
})

var redisKey = "gcache_redis_cache_test_key"

func TestRedisCache_Set(t *testing.T) {
	_ = cacheStore.Set(redisKey, []byte("test_data"), -1)
}

func TestRedisCache_Get(t *testing.T) {
	data, err := cacheStore.Get(redisKey).Bytes()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestRedisCache_Del(t *testing.T) {
	if err := cacheStore.Del(redisKey); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkRedisCache_Get1KB(b *testing.B) {
	data := make([]byte, 1024)
	cacheStore.Set(redisKey, data, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cacheStore.Get(redisKey).Bytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkRedisCache_Get10KB(b *testing.B) {
	data := make([]byte, 10240)
	cacheStore.Set(redisKey, data, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cacheStore.Get(redisKey).Bytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRedisCache_Get100KB(b *testing.B) {
	data := make([]byte, 102400)
	cacheStore.Set(redisKey, data, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cacheStore.Get(redisKey).Bytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}
