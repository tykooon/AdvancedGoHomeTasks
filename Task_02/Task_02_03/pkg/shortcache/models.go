package shortcache

import (
	"sync"
	"time"
)

const ObjectLifeTime time.Duration = 10 * time.Second

type TimedObject struct {
	Obj        any
	LastAccess time.Time
}

type ShortCache struct {
	storage  map[string]TimedObject
	capacity int
	ticker   time.Ticker
	mutex    sync.Mutex
}

func New(cap int) *ShortCache {
	res := &ShortCache{
		storage:  make(map[string]TimedObject, cap),
		capacity: cap,
		ticker:   *time.NewTicker(time.Second),
		mutex:    sync.Mutex{},
	}
	go TickHandler(res)
	return res
}

func TickHandler(cache *ShortCache) {
	for range cache.ticker.C {
		cache.mutex.Lock()
		for key, obj := range cache.storage {
			if time.Since(obj.LastAccess) > 10*time.Second {
				delete(cache.storage, key)
			}
		}
		cache.mutex.Unlock()
	}
}

func (c *ShortCache) Count() int {
	return len(c.storage)
}
