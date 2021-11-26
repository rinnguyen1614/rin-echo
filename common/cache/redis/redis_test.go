package redis

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
	rc  *RedisCache
)

func init() {
	c, err := NewRedisCache("redis-14954.c289.us-west-1-2.ec2.cloud.redislabs.com:14954")
	if err != nil {
		panic(err)
	}
	rc = c.(*RedisCache)
}

func TestRedisCache(t *testing.T) {
	expiration := 3 * time.Second

	assert.Nil(t, rc.Set(ctx, "key", 1, expiration))

	res, _ := rc.IsExist(ctx, "key")
	assert.True(t, res)

	time.Sleep(5 * time.Second)

	res, _ = rc.IsExist(ctx, "key")
	assert.False(t, res)

	assert.Nil(t, rc.Set(ctx, "key", 1, expiration))

	val, _ := rc.Get(ctx, "key")
	assert.Equal(t, 1, val)

	// test incr
	incr, err := rc.Incr(ctx, "key")
	assert.Nil(t, err)
	assert.Equal(t, 2, incr)

	// test decr
	decr, err := rc.Decr(ctx, "key")
	assert.Nil(t, err)
	val, _ = rc.Get(ctx, "key")
	assert.Equal(t, 1, decr)
	rc.Delete(ctx, "key")

	res, _ = rc.IsExist(ctx, "key")
	assert.False(t, res)

	assert.Nil(t, rc.Set(ctx, "key", "value", expiration))
	// test string

	res, _ = rc.IsExist(ctx, "key")
	assert.True(t, res)

	val, _ = rc.Get(ctx, "key")
	assert.Equal(t, "value", val)

	// test GetMulti
	assert.Nil(t, rc.Set(ctx, "key1", "value1", expiration))

	res, _ = rc.IsExist(ctx, "key1")
	assert.True(t, res)

	vv, _ := rc.GetMulti(ctx, "key", "key1")
	assert.Equal(t, 2, len(vv))
	assert.Equal(t, "value", vv[0].(string))
	assert.Equal(t, "value1", vv[1].(string))

	vv, _ = rc.GetMulti(ctx, "key0", "key1")
	assert.Nil(t, vv[0])
	assert.Equal(t, "value1", vv[1].(string))

	// test clear all
	assert.Nil(t, rc.ClearAll(ctx))
}
