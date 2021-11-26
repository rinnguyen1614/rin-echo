package memory

import (
	"context"
	"fmt"
	"math"
	"rin-echo/common/cache"
	"strings"
	"sync"
	"time"
)

type MemoryItem struct {
	val         interface{}
	createdTime time.Time
	lifespan    time.Duration
}

func (i *MemoryItem) isExpire() bool {
	// 0 means forever
	if i.lifespan == 0 {
		return false
	}
	return time.Since(i.createdTime) > i.lifespan
}

type MemoryCache struct {
	mu       sync.RWMutex
	interval time.Duration
	items    map[string]*MemoryItem
	close    chan struct{}
}

// interval is a timer for how often to recycle the expired cache items in memory (in seconds)
// interval = 0, it means "don't clean up"
func NewMemoryCache(interval time.Duration) cache.Cache {
	var c = &MemoryCache{
		interval: interval,
		items:    make(map[string]*MemoryItem),
		close:    make(chan struct{}),
	}

	if interval > 0 {
		go func() {
			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					c.ClearAllExpired(context.Background())
				case <-c.close:
					return
				}
			}
		}()
	}

	return c
}

func (c *MemoryCache) Close() {
	close(c.close)
}

func (c *MemoryCache) Get(ctx context.Context, key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if item, ok := c.items[key]; ok {
		if item.isExpire() {
			return nil, cache.ErrKeyExpired
		}
		return item.val, nil
	}

	return nil, cache.ErrKeyDoNotExists
}

func (c *MemoryCache) GetMulti(ctx context.Context, keys ...string) ([]interface{}, error) {
	var (
		values  = make([]interface{}, len(keys))
		keysErr = make([]string, 0)
	)

	for i, key := range keys {
		val, err := c.Get(context.Background(), key)
		if err != nil {
			keysErr = append(keysErr, fmt.Sprintf("key [%s] error: %v", key, err.Error()))
			continue
		}
		values[i] = val
	}

	if len(keysErr) != 0 {
		return nil, fmt.Errorf(strings.Join(keysErr, "; "))
	}

	return values, nil
}

func (c *MemoryCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &MemoryItem{
		val:         value,
		createdTime: time.Now(),
		lifespan:    expiration,
	}
	return nil
}

func (c *MemoryCache) Delete(ctx context.Context, key string) error {
	return c.delete(key)
}

func (c *MemoryCache) Incr(ctx context.Context, key string) (int64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.items[key]
	if !ok {
		return 0, cache.ErrKeyDoNotExists
	}

	val, err := incr(item.val)
	if err != nil {
		return 0, err
	}
	item.val = val
	return val, nil
}

func (c *MemoryCache) Decr(ctx context.Context, key string) (int64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, ok := c.items[key]
	if !ok {
		return 0, cache.ErrKeyDoNotExists
	}

	val, err := decr(item.val)
	if err != nil {
		return 0, err
	}
	item.val = val
	return val, nil
}

func (c *MemoryCache) IsExist(ctx context.Context, key string) (bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if item, ok := c.items[key]; ok {
		return !item.isExpire(), nil
	}
	return false, nil
}

//
func (c *MemoryCache) ClearAllExpired(ctx context.Context) error {
	keys := c.getExpriedKeys()
	if len(keys) != 0 {
		c.delete(keys...)
	}
	return nil
}

func (c *MemoryCache) ClearAll(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*MemoryItem)
	return nil
}

func (c *MemoryCache) getExpriedKeys() (keys []string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for key, itm := range c.items {
		if itm.isExpire() {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c *MemoryCache) delete(keys ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, key := range keys {
		delete(c.items, key)
	}
	return nil
}

func incr(val interface{}) (int64, error) {
	switch v := val.(type) {
	case int:
		if v == math.MaxInt {
			return 0, cache.ErrIncrementOverflow
		}
		return int64(v + 1), nil
	case int32:
		if v == math.MaxInt32 {
			return 0, cache.ErrIncrementOverflow
		}
		return int64(v + 1), nil
	case int64:
		if v == math.MaxInt64 {
			return 0, cache.ErrIncrementOverflow
		}
		return int64(v + 1), nil
	case uint:
		if v == math.MaxUint {
			return 0, cache.ErrIncrementOverflow
		}
		return int64(v + 1), nil
	case uint32:
		if v == math.MaxUint32 {
			return 0, cache.ErrIncrementOverflow
		}
		return int64(v + 1), nil
	case uint64:
		if v == math.MaxUint64 {
			return 0, cache.ErrIncrementOverflow
		}
		return int64(v + 1), nil
	default:
		return 0, cache.ErrNotIntegerType
	}
}

func decr(val interface{}) (int64, error) {
	switch v := val.(type) {
	case int:
		if v == math.MinInt {
			return 0, cache.ErrDecrementOverflow
		}
		return int64(v - 1), nil
	case int32:
		if v == math.MinInt32 {
			return 0, cache.ErrDecrementOverflow
		}
		return int64(v - 1), nil
	case int64:
		if v == math.MinInt64 {
			return 0, cache.ErrDecrementOverflow
		}
		return int64(v - 1), nil
	case uint:
		if v == 0 {
			return 0, cache.ErrDecrementOverflow
		}
		return int64(v - 1), nil
	case uint32:
		if v == 0 {
			return 0, cache.ErrDecrementOverflow
		}
		return int64(v - 1), nil
	case uint64:
		if v == 0 {
			return 0, cache.ErrDecrementOverflow
		}
		return int64(v - 1), nil
	default:
		return 0, cache.ErrNotIntegerType
	}
}
