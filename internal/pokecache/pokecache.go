package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	storage  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		storage:  make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Lock()
	c.storage[key] = newEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.storage[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.storage {
			if time.Since(v.createdAt) >= c.interval {
				delete(c.storage, k)
			}
		}
		c.mu.Unlock()
	}
}
