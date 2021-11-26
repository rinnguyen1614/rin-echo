package cache

import "sync"

type CacheManager struct {
	caches map[string]Cache
	mu     sync.RWMutex
}

func NewCacheManager() *CacheManager {
	return &CacheManager{
		caches: make(map[string]Cache),
	}
}

func (m *CacheManager) Set(name string, cache Cache) {
	if cache == nil {
		panic("cachemanager: Set cache is nil")
	}
	m.mu.Lock()
	m.caches[name] = cache
	m.mu.Unlock()
}

func (m *CacheManager) Get(name string) Cache {
	m.mu.RLock()
	cache, _ := m.caches[name]
	m.mu.RUnlock()
	return cache
}

func (m *CacheManager) Register(name string, cache Cache) {
	if cache == nil {
		panic("cache: Register cache is nil")
	}

	if cache := m.Get(name); cache == nil {
		panic("cache: Register called twice for cache " + name)
	}

	m.Set(name, cache)
}
