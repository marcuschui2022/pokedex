package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data  map[string]cacheEntry
	mutex sync.Mutex
	done  chan struct{}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]cacheEntry),
		done: make(chan struct{}),
	}

	cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	data, exists := c.data[key]
	if exists {
		return data.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-c.done:
				ticker.Stop()
				return
			case <-ticker.C:
				c.mutex.Lock()
				for key, data := range c.data {
					if data.createdAt.Add(interval).Before(time.Now()) {
						delete(c.data, key)
					}
				}
				c.mutex.Unlock()
			}
		}
	}()
}

func (c *Cache) Stop() {
	close(c.done) // Signal the goroutine to stop

}
