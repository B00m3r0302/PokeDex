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
	return &Cache{
		cacheEntry: make(map[string]CacheEntry),
		interval:   interval,
		mutex:      sync.Mutex{},
	}
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

func (c *Cache) Get() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
}

func (c *Cache) reapLoop() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

}
