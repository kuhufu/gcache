package gcache

import (
	"fmt"
	"testing"
)

var memCache = NewMemCache(100 * 1024 * 1024)
var memKey = "gcache_mem_cache_test_key"

func TestMemCache_Set(t *testing.T) {
	err := memCache.Set(memKey, []byte("test_data"), -1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMemCache_Get(t *testing.T) {
	data, err := memCache.Get(memKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestMemCache_Del(t *testing.T) {
	if !memCache.Del(memKey) {
		t.Fatal("not affected")
	}
}
