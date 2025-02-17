package cache

import (
	"sync"
	"time"
)

type valueWithTimeout[V any] struct {
	value   V
	expires time.Time
}

type Cache[K comparable, V any] struct {
	ttl  time.Duration
	mu   sync.RWMutex
	data map[K]valueWithTimeout[V]
}

func New[K comparable, V any](ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:  ttl,
		data: make(map[K]valueWithTimeout[V]),
	}
}

func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroV V

	value, found := c.data[key]
	switch {
	case !found:
		return zeroV, false
	case value.expires.Before(time.Now()):
		delete(c.data, key)
		return zeroV, false
	default:
		return value.value, found
	}
}

func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = valueWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	return nil
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}
