package gcache

import (
	"fmt"
	"testing"
	"time"
)

var memc = NewMemCache()
var memKey = "gcache_mem_cache_test_key"

func TestMemCache_Exist(t *testing.T) {
	exist, err := memc.Exist(memKey)
	if exist {
		t.Error("err")
	}

	if err != nil {
		t.Error("err")
	}

	memc.Set(memKey, memKey, 0)
	exist, err = memc.Exist(memKey)
	if !exist {
		t.Error("err")
	}

	if err != nil {
		t.Error("err")
	}
}

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
	if err := memc.Del(memKey); err != nil {
		t.Fatal(err)
	}
}

func TestMemCache_IncrBy(t *testing.T) {
	memc.IncrBy("test", 1)
	memc.IncrBy("test", 1)
	reply, err := memc.Get("test").Int()
	if err != nil {
		t.Error(err)
	}

	if reply != 2 {
		t.Error("error incr 1", reply)
	}
}

func TestMemCache_Expire(t *testing.T) {
	sec := 2
	memc.Set("test", 3, sec)
	if v, err := memc.Get("test").Int(); err != nil || v != 3 {
		t.Error(err)
	}

	time.Sleep(time.Duration(sec) * time.Second)

	if memc.Get("test").Error() == nil {
		t.Error("expire failure")
	}

	fmt.Println(memc.Get("test").Error())
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
