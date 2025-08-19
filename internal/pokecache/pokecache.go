package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	locs          map[string]cacheEntry
	mux           *sync.Mutex
	cleanInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		locs:          make(map[string]cacheEntry),
		mux:           &sync.Mutex{},
		cleanInterval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, value []byte) {
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.mux.Lock()
	c.locs[key] = newEntry
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, ok := c.locs[key]
	c.mux.Unlock()

	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.cleanInterval)

	for range ticker.C {
		c.mux.Lock()
		for k, v := range c.locs {
			if time.Since(v.createdAt) >= c.cleanInterval {
				delete(c.locs, k)
			}
		}
		c.mux.Unlock()
	}
}
