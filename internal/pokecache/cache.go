package pokecache

import (
	"fmt"
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
	return Cache{}, nil
}

func (c Cache) Add(key string, val []byte) error {
	return nil
}

func (c Cache) Get(key string) (cacheEntry, error) {
	entry, ok := c.cache[key]
	if !ok {
		return cacheEntry{}, fmt.Errorf("no cache entry with key: %s", key)
	}
	return entry, nil
}
