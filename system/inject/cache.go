package inject

import (
	"github.com/rinnguyen1614/rin-echo-core/cache"
	"github.com/rinnguyen1614/rin-echo-core/cache/memory"
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
