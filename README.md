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