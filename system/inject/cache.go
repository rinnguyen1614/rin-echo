package inject

import (
	"rin-echo/common/cache"
	"rin-echo/common/cache/memory"
)

var (
	SettingCacheName = "setting"
)

func GetCache() *cache.CacheManager {
	if service.cache == nil {
		cacheManager := cache.NewCacheManager()

		cacheManager.Register(SettingCacheName, memory.NewMemoryCache(0))

		service.cache = cacheManager
	}
	return service.cache
}
