package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
}

type casheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(t time.Duration) *Cache {
	var C Cache {
		interval: t,
	}
	go reapLoop()

	return C

}

func (C *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer C.mu.Unlock()
	c.entry[key] = val
	fmt.Printf("added %s -> %v\n", key, val)
	
}

func (c *cache) Get(key string, val []byte) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_,ok := c.etnry[key]
	if !ok {
		return []byte, false
	} else {
		return c.entry[key].val, true
	}
}

func (c *cache) reapLoop(){
	c.mu.Lock()
	defer c.mu.Unlock()
	for _,i := range c.etnry{
		dif := time.Time - createdAt
		if dif > c.interval {
			delete(i, c.entry)
		}
	}
	return
}