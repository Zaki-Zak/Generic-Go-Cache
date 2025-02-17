# Generic-Go-Cache

Generic-Go-Cache is a generic, thread-safe in-memory cache library written in Go. It supports time-to-live (TTL) expiration for cache entries and limits the number of items stored in the cache (max size). The cache is designed to ensure efficient concurrent access and automatic cleanup of expired or evicted items.

## Features

- **Thread Safety**: Uses a `sync.Mutex` to ensure safe concurrent access.
- **TTL Support**: Cache entries automatically expire after the specified TTL.
- **Max Size**: The cache is limited to a maximum number of items. When the cache reaches this size, the oldest entry is evicted.
- **FIFO Eviction**: When the cache exceeds the maximum size, the oldest item (based on insertion order) is removed.
- **Simple API**: Provides basic operations such as `Read`, `Upsert`, and `Delete`.

## Installation

To install GoCache, run the following command:

```bash
$ go get github.com/Zaki-Zak/Generic-Go-Cache
```

## Usage

### Create a new Cache

```go
cache := gocache.New(100, time.Minute*10) // Max size 100, TTL 10 minutes
```

### Upsert an item
```go
cache.Upsert("key", "value")
```
### Read an item
```go
value, found := cache.Read("key")
```
### Delete an item
```go
cache.Delete("key")
```
### Example

```go
package main

import (
    "fmt"
    "time"
   cache "github.com/Zaki-Zak/Generic-Go-Cache"
)

func main() {
    cache := cache.New(5, time.Minute*1)

    cache.Upsert("key", "value")
    if value, found := cache.Read("key"); found {
        fmt.Println("Cache hit:", value)
    } else {
        fmt.Println("Cache miss")
    }
}
```

## License

This is just a tiny side project I used to learn more about generics and it's not really intended to practical use. This project still misses a lot of functionalities that are expected for a cache memory so maybe it will be added later.
