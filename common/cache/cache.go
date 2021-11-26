package cache

import (
	"context"
	"time"
)

type Cache interface {
	// Get a cached value by key.
	Get(ctx context.Context, key string) (interface{}, error)
	// GetMulti is a batch version of Get.
	GetMulti(ctx context.Context, keys ...string) ([]interface{}, error)
	// Set a cached value with key and expire time.
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	// Delete cached value by key.
	// Should not return error if key not found
	Delete(ctx context.Context, key string) error
	// Increment a cached int value by key, as a counter.
	Incr(ctx context.Context, key string) (int64, error)
	// Decrement a cached int value by key, as a counter.
	Decr(ctx context.Context, key string) (int64, error)
	// Check if a cached value exists or not.
	// if key is expired, return (false, nil)
	IsExist(ctx context.Context, key string) (bool, error)
	// Clear all cache.
	ClearAll(ctx context.Context) error
}
