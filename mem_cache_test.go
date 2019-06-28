package gcache

import (
	"fmt"
	"testing"
)

var cache = NewMemCache(1000 * 1024 * 1024)
var memKey = "gcache_mem_cache_test_key"

func TestMemCache_Set(t *testing.T) {
	err := cache.Set(memKey, []byte("test_data"), -1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMemCache_Get(t *testing.T) {
	data, err := cache.Get(memKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestMemCache_Del(t *testing.T) {
	if !cache.Del(memKey) {
		t.Fatal("not affected")
	}
}

func BenchmarkMemCache_Get1KB(b *testing.B) {
	data := make([]byte, 1024)
	cache.Set(memKey, data, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cache.Get(memKey)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMemCache_Get10KB(b *testing.B) {
	data := make([]byte, 10240)
	cache.Set(memKey, data, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cache.Get(memKey)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMemCache_Get100KB(b *testing.B) {
	data := make([]byte, 102400)
	err := cache.Set(memKey, data, -1)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cache.Get(memKey)
		if err != nil {
			b.Fatal(err)
		}
	}
}
