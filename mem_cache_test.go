package gcache

import (
	"fmt"
	"testing"
)

var memc = NewMemCache()
var memKey = "gcache_mem_cache_test_key"

func TestMemCache_Set(t *testing.T) {
	err := memc.Set(memKey, "data", -1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMemCache_Get(t *testing.T) {
	memc.Set(memKey, []byte{1}, -1)
	data, err := memc.Get(memKey).Bytes()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestMemCache_Del(t *testing.T) {
	if !memc.Del(memKey) {
		t.Fatal("not affected")
	}
}

func BenchmarkMemCache_Get1KB(b *testing.B) {
	data := make([]byte, 1024)
	memc.Set(memKey, data, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := memc.Get(memKey).Bytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMemCache_Get10KB(b *testing.B) {
	data := make([]byte, 10240)
	memc.Set(memKey, data, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := memc.Get(memKey).Bytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMemCache_Get100KB(b *testing.B) {
	data := make([]byte, 102400)
	err := memc.Set(memKey, data, -1)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := memc.Get(memKey).Bytes()
		if err != nil {
			b.Fatal(err)
		}
	}
}
