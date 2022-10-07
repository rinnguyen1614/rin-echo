package redis

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	ctx        = context.Background()
	expiration = 5 * time.Second
	rc         *RedisCache
)

func init() {
	c, err := NewRedisCache("redis://localhost/")
	if err != nil {
		panic(err)
	}
	rc = c.(*RedisCache)
}

func TestRedisCache(t *testing.T) {

	assert.Nil(t, rc.Set(ctx, "key", 1, expiration))

	res, _ := rc.IsExist(ctx, "key")
	assert.True(t, res)

	time.Sleep(5 * time.Second)

	res, _ = rc.IsExist(ctx, "key-2")
	assert.False(t, res)

	assert.Nil(t, rc.Set(ctx, "key", 1, expiration))

	val, _ := rc.Get(ctx, "key")
	assert.Equal(t, float64(1), val)

	// test incr
	incr, err := rc.Incr(ctx, "key")
	assert.Nil(t, err)
	assert.Equal(t, int64(2), incr)

	// test decr
	decr, err := rc.Decr(ctx, "key")
	assert.Nil(t, err)
	val, _ = rc.Get(ctx, "key")
	assert.Equal(t, int64(1), decr)
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

func TestRedisCache_ExplicitType(t *testing.T) {

	// struct type
	type testType struct {
		Field1 string
		Field2 int
	}
	structExpected := testType{
		Field1: "value1",
		Field2: 2,
	}

	var valStruct testType
	assert.Nil(t, rc.Set(ctx, "key", structExpected, expiration))
	assert.Nil(t, rc.GetParse(ctx, "key", &valStruct))
	assert.Equal(t, structExpected, valStruct)

	var valInt int
	assert.Nil(t, rc.Set(ctx, "key", 1, expiration))
	assert.Nil(t, rc.GetParse(ctx, "key", &valInt))
	assert.Equal(t, 1, valInt)

	var valSliceInt []int
	assert.Nil(t, rc.Set(ctx, "key", []int{1, 2, 3}, expiration))
	assert.Nil(t, rc.GetParse(ctx, "key", &valSliceInt))
	assert.Equal(t, []int{1, 2, 3}, valSliceInt)
}
