# 缓存工具库
提供两种缓存方式

### 1. 应用内存缓存
```go
cacheSize := 100 * 1024 * 1024 //100MB
cache := gcache.NewMemCache(cacheSize)
```

### 2. Redis缓存
```go
var cache = gcache.NewRedisCache(10, "tcp", "127.0.0.1:6379", "password")
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

### 更新器
```go
cacheSize := 100 * 1024 * 1024 //100MB
cache := gcache.NewMemCache(cacheSize)
updater := gcache.NewUpdaterOf(cache)

//每5s更新一次，true表示立刻更新一次
updater.Interval("key", 5, func()[]byte{
	//获取新值，返回新值
}, true)

//5s后更新一次，true同上
updater.Timeout("key", 5, func()[]byte{
	//获取新值，返回新值
}, true)
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