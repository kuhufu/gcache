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