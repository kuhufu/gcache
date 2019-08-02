# 缓存工具库
提供两种缓存方式

### 1. 应用内存缓存
```go
cache := gcache.NewMemCache()
```

### 2. Redis缓存
```go
var cacheStore = NewRedisCache("tcp", "127.0.0.1:6379", RedisOption{
	MaxIdle:     10,
	MaxActive:   30,
	IdleTimeout: time.Second * 180,
	Password:    "",
})
```

### 使用

Set
```go
expireSeconds := 10 //10s后过期，expireSeconds <= 0时，永不过期
cache.Set("key", "value", expireSeconds)
```

Get
```go
//"key"不存在，err != nil
value, err := cache.Get("key").String()
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
```go
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
cache := gcache.NewMemCache()
updater := gcache.NewUpdaterOf(cache)

//每5s更新一次，true表示立刻更新一次
updater.Interval("key", 5, func() interface{} {
	//获取新值，返回新值
}, true)

//5s后更新一次，true同上
updater.Timeout("key", 5, func() interface{} {
	//获取新值，返回新值
}, true)
```

### Benchmark
```
```