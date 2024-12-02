package pokecache

import (
	"time"
)

type Cache struct {
	cached map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{cached: make(map[string]cacheEntry)}
	go c.clearLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cached[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheVal, ok := c.cached[key]
	return cacheVal.val, ok
}

func (c *Cache) clearLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Clear(interval)
	}
}

func (c *Cache) Clear(interval time.Duration) {
	timeToDelete := time.Now().Add(interval)
	for key, value := range c.cached {
		if value.createdAt.Before(timeToDelete) {
			delete(c.cached, key)
		}
	}
}
