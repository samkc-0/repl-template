package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	created_at time.Time
	val        []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) (Cache, error) {
	c := Cache{}
	go c.reapLoop(interval)
	return c, nil
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{created_at: time.Now(), val: val}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.cache {
				if time.Since(entry.created_at) > interval {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
