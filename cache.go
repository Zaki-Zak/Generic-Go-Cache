package cache

import (
	"slices"
	"sync"
	"time"
)

type valueWithTimeout[V any] struct {
	value   V
	expires time.Time
}

type Cache[K comparable, V any] struct {
	ttl time.Duration

	mu   sync.Mutex
	data map[K]valueWithTimeout[V]

	maxSize    int
	chronoKeys []K
}

func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:        ttl,
		data:       make(map[K]valueWithTimeout[V]),
		maxSize:    maxSize,
		chronoKeys: make([]K, 0, maxSize),
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
		c.deleteKeyValue(key)
		return zeroV, false
	default:
		return value.value, found
	}
}

func (c *Cache[K, V]) Upsert(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, alreadyExists := c.data[key]
	switch {
	case alreadyExists:
		c.deleteKeyValue(key)
	case len(c.data) == c.maxSize:
		c.deleteKeyValue(c.chronoKeys[0])
	}
	c.addKeyValue(key, value)
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.deleteKeyValue(key)
}

func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = valueWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	c.chronoKeys = append(c.chronoKeys, key)
}

func (c *Cache[K, V]) deleteKeyValue(key K) {
	c.chronoKeys = slices.DeleteFunc(c.chronoKeys, func(k K) bool { return k == key })
	delete(c.data, key)
}
