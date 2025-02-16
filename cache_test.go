package cache_test

import (
	"testing"

	cache "github.com/Zaki-Zak/Generic-Go-Cache"
)

func TestCache_Goroutine(t *testing.T) {
	c := cache.New[int, string]()
	t.Run("write 9", func(t *testing.T) {
		t.Parallel()
		c.Upsert(9, "nine")
	})
	t.Run("write something", func(t *testing.T) {
		t.Parallel()
		c.Upsert(9, "something")
	})
}
