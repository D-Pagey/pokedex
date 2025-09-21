// Package pokecache caches api responses from the Pokemon API
package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]CacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache:    make(map[string]CacheEntry),
		interval: interval,
	}

	go cache.reapLoop()
	return &cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()

		for key, entry := range c.cache {
			isOld := entry.createdAt.Add(c.interval).Before(time.Now())
			if isOld {
				delete(c.cache, key)
			}
		}

		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cache[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, ok := c.cache[key]; ok {
		return val.val, true
	}
	return nil, false
}
