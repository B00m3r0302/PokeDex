package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]CacheEntry
	interval   time.Duration
	mutex      sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntry: make(map[string]CacheEntry),
		interval:   interval,
	}
	go cache.reapLoop()
	return cache
}

type cacheFunctions interface {
	Add()
	Get()
	reapLoop()
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntry[key] = CacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.cacheEntry[key]
	if !ok {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) reapLoop() *Cache {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.cacheEntry {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cacheEntry, key)
			}
		}
		c.mutex.Unlock()
	}
	return c
}
