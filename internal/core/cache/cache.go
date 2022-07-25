package cache

import (
	"context"
	"errors"
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

type (
	Factory func(key string) (value interface{})
)

func GetOrCreate(cache Cache, ctx context.Context, key string, factory Factory) (interface{}, error) {
	value, err := cache.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, ErrKeyDoNotExists) {
			return nil, err
		}
		value = factory(key)
		if err = cache.Set(ctx, key, value, 0); err != nil {
			return nil, err
		}
	}

	return value, nil
}

func GetOrCreateMulti(cache Cache, ctx context.Context, keys []string, factory Factory) ([]interface{}, error) {
	values, err := cache.GetMulti(ctx, keys...)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		values = make([]interface{}, len(keys))
	}

	for i, value := range values {
		if value != nil {
			continue
		}

		var (
			key      = keys[i]
			newValue = factory(key)
		)

		if err = cache.Set(ctx, key, newValue, 0); err != nil {
			return nil, err
		}

		values[i] = newValue
	}
	return values, nil
}
