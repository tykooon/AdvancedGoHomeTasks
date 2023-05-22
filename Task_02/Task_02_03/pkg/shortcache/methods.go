package shortcache

import (
	"time"
)

func (cache *ShortCache) Set(key string, object any) {
	cache.mutex.Lock()
	if obj, ok := cache.storage[key]; ok {
		obj.Obj = object
		obj.LastAccess = time.Now()
	} else {
		if len(cache.storage) == cache.capacity {
			var maxAge time.Duration
			var oldestKey string
			for k, v := range cache.storage {
				if time.Since(v.LastAccess) > maxAge {
					maxAge, oldestKey = time.Since(v.LastAccess), k
				}
			}
			delete(cache.storage, oldestKey)
		}
		obj := TimedObject{
			Obj:        object,
			LastAccess: time.Now(),
		}
		cache.storage[key] = obj
	}
	cache.mutex.Unlock()
}

func (cache *ShortCache) Get(key string) (res any) {
	cache.mutex.Lock()
	obj, ok := cache.storage[key]
	if ok {
		obj.LastAccess = time.Now()
		res = obj.Obj
	}
	cache.mutex.Unlock()
	return res
}

func (cache *ShortCache) Remove(key string) bool {
	cache.mutex.Lock()
	_, ok := cache.storage[key]
	if ok {
		delete(cache.storage, key)
	}
	cache.mutex.Unlock()
	return ok
}
