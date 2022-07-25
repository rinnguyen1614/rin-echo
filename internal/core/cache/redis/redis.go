package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/rinnguyen1614/rin-echo/internal/core/cache"

	"github.com/go-redis/redis/v8"
)

// DefaultKey defines the collection name of redis for the cache adapter.
var DefaultKey = "rin-echo-redis"

type (
	RedisCache struct {
		rdb    *redis.Client
		prefix string
	}
	Config struct {
		Key string
	}
)

//
// Examples:
//		redis://user:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2
func NewRedisCache(conninfo string) (cache.Cache, error) {
	return NewRedisCacheWithPrefix(DefaultKey, conninfo)
}

func NewRedisCacheWithPrefix(prefix, conninfo string) (cache.Cache, error) {
	opt, err := redis.ParseURL(conninfo)
	if err != nil {
		return nil, err
	}
	return NewRedisCacheWithOption(prefix, opt)
}

func NewRedisCacheWithOption(prefix string, opt *redis.Options) (cache.Cache, error) {
	rc := &RedisCache{
		prefix: prefix,
		rdb:    redis.NewClient(opt),
	}
	_, err := rc.rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rc, nil
}

// associate with config prefix.
func (c *RedisCache) associate(key interface{}) string {
	return fmt.Sprintf("%s:%s", c.prefix, key)
}

func (c *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := c.rdb.Get(ctx, c.associate(key)).Result()
	if err != nil {
		return nil, redisError(err)
	}

	return val, nil
}

func (c *RedisCache) GetMulti(ctx context.Context, keys ...string) ([]interface{}, error) {
	var args []string
	for _, key := range keys {
		args = append(args, c.associate(key))
	}

	vals, err := c.rdb.MGet(ctx, args...).Result()
	if err != nil {
		return nil, redisError(err)
	}

	return vals, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := c.rdb.Set(ctx, c.associate(key), value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	_, err := c.rdb.Del(ctx, c.associate(key)).Result()
	if err != nil {
		return redisError(err)
	}

	return nil
}

func (c *RedisCache) Incr(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.Incr(ctx, c.associate(key)).Result()
	if err != nil {
		return 0, redisError(err)
	}

	return val, nil
}

func (c *RedisCache) Decr(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.Decr(ctx, c.associate(key)).Result()
	if err != nil {
		return 0, redisError(err)
	}

	return val, nil
}

func (c *RedisCache) IsExist(ctx context.Context, key string) (bool, error) {
	val, err := c.rdb.Exists(ctx, c.associate(key)).Result()
	if err != nil {
		return false, redisError(err)
	}

	return val != 0, nil
}

func (c *RedisCache) ClearAll(ctx context.Context) error {
	iter := c.rdb.Scan(ctx, 0, c.prefix+":*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		if err := c.Delete(ctx, key); err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}

func redisError(err error) error {
	if err == redis.Nil {
		return cache.ErrKeyDoNotExists
	}
	return nil
}
