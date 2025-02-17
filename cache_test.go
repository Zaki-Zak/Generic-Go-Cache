package cache_test

import (
	"testing"
	"time"

	cache "github.com/Zaki-Zak/Generic-Go-Cache"
	"github.com/stretchr/testify/assert"
)

func TestCache_TTL(t *testing.T) {
	t.Parallel()
	c := cache.New[string, string](time.Millisecond * 100)
	c.Upsert("Norwegian", "White")

	got, found := c.Read("Norwegian")
	assert.True(t, found)
	assert.Equal(t, "White", got)

	time.Sleep(time.Millisecond * 200)

	got, found = c.Read("Norwegian")
	assert.False(t, found)
	assert.Equal(t, "", got)
}

func TestCache_Goroutine(t *testing.T) {
	c := cache.New[int, string](time.Millisecond * 100)
	t.Run("write 9", func(t *testing.T) {
		t.Parallel()
		c.Upsert(9, "nine")
	})
	t.Run("write something", func(t *testing.T) {
		t.Parallel()
		c.Upsert(9, "something")
	})
}
