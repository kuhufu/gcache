# 缓存工具库
提供两种缓存方式

### 1. 应用内存缓存
```go
cacheSize := 100 * 1024 * 1024 //100MB
var cache = NewMemCache(cacheSize)
```

### 2. Redis缓存
```go
var cache = NewRedisCache(10, "tcp", "127.0.0.1:6379", "password")
```

### 使用

Set
```go
expireSeconds := 10 //10s后过期，expireSeconds <= 0时，永不过期
cache.Set("key", []byte("value"), expireSeconds)
```

Get
```go
//"key"不存在，err != nil
value, err := cache.Get("key")
```

Del
```go
//"key"存在，affected == true
//"key"不存在，affected == false
affected := cache.Del("key")
```

GetUnmarshal
```go
value, err := cache.GetUnmarshal("key")
```
```
//等价于
bytes, err := cache.Get("key")
if err != nil {
    return nil, err
}
var value interface{}
err = json.Unmarshal(data, &value)
return value, err
```

### Benchmark
```
goos: windows
goarch: amd64
pkg: github.com/kuhufu/gcache
BenchmarkMemCache_Get1KB-8       	 3000000	       351 ns/op
BenchmarkMemCache_Get10KB-8      	 1000000	      1838 ns/op
BenchmarkMemCache_Get100KB-8     	  100000	     18059 ns/op
BenchmarkRedisCache_Get1KB-8     	   10000	    125026 ns/op
BenchmarkRedisCache_Get10KB-8    	   10000	    147127 ns/op
BenchmarkRedisCache_Get100KB-8   	   10000	    214976 ns/op
PASS
```