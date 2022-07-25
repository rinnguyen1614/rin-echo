package memory

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	ctx        = context.Background()
	mc         *MemoryCache
	expiration = time.Second * 10
)

func init() {
	mc = NewMemoryCache(0).(*MemoryCache)
}

func TestMemoryCache_WithInterval(t *testing.T) {
	var interval = 5 * time.Second
	mc = NewMemoryCache(interval).(*MemoryCache)

	// init values
	for i := 0; i < 10; i++ {
		var (
			key = fmt.Sprintf("key-%d", i)
			val = fmt.Sprintf("val-%d", i)
		)
		_ = mc.Set(ctx, key, val, time.Duration(i)*time.Second)
	}

	fmt.Printf("Waiting for %s\n", interval.String())
	time.Sleep(interval)

	re, _ := mc.IsExist(ctx, "key-0")
	assert.True(t, re)

	re, _ = mc.IsExist(ctx, "key-1")
	assert.True(t, !re)

	re, _ = mc.IsExist(ctx, "key-9")
	assert.True(t, re)

	// stop timer
	mc.Close()
	fmt.Printf("Stopping timer and waiting for %s\n", interval.String())
	time.Sleep(interval)
	assert.ElementsMatch(t, mc.getExpriedKeys(), []string{"key-6", "key-7", "key-8", "key-9"})
}

func TestMemoryCache_Get(t *testing.T) {
	key := "key"
	value := "value"

	val, err := mc.Get(ctx, key)
	assert.NotNil(t, err)

	err = mc.Set(ctx, key, value, 0)
	assert.Nil(t, err)

	val, err = mc.Get(ctx, key)
	assert.Equal(t, value, val)
}
func TestMemoryCache_GetMulti(t *testing.T) {
	var (
		n      = 10
		keys   = make([]string, 10)
		values = make([]string, 10)
	)
	for i := 0; i < n; i++ {
		var (
			key = fmt.Sprintf("key-%d", i)
			val = fmt.Sprintf("val-%d", i)
			err = mc.Set(ctx, key, val, expiration)
		)

		if assert.Nil(t, err) {
			keys[i] = key
			values[i] = val
		}
	}

	actual, err := mc.GetMulti(ctx, keys...)
	if assert.Nil(t, err) && assert.Equal(t, len(values), len(actual)) {
		for i, v := range actual {
			assert.Equal(t, values[i], v)
		}
	}
}

func TestMemoryCache_Set(t *testing.T) {
	key := "key"
	value := "value"

	err := mc.Set(ctx, key, value, 0)
	assert.Nil(t, err)

	val, err := mc.Get(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, value, val)
}

func TestMemoryCache_Incr_And_Decr(t *testing.T) {
	v := 1
	testIncrAndDecr(t, v, expiration)
	testIncrAndDecr(t, int32(v), expiration)
	testIncrAndDecr(t, int64(v), expiration)
	testIncrAndDecr(t, uint(v), expiration)
	testIncrAndDecr(t, uint32(v), expiration)
	testIncrAndDecr(t, uint64(v), expiration)
}

func TestMemoryCache_Incr_OverFlow(t *testing.T) {
	testIncrOverFlow(t, int(math.MaxInt), expiration)
	testIncrOverFlow(t, int32(math.MaxInt32), expiration)
	testIncrOverFlow(t, int64(math.MaxInt64), expiration)
	testIncrOverFlow(t, uint(math.MaxUint), expiration)
	testIncrOverFlow(t, uint32(math.MaxUint32), expiration)
	testIncrOverFlow(t, uint64(math.MaxUint64), expiration)
}

func TestMemoryCache_Decr_OverFlow(t *testing.T) {
	testDecrOverFlow(t, int(math.MinInt), expiration)
	testDecrOverFlow(t, int32(math.MinInt32), expiration)
	testDecrOverFlow(t, int64(math.MinInt64), expiration)
	testDecrOverFlow(t, uint(0), expiration)
	testDecrOverFlow(t, uint32(0), expiration)
	testDecrOverFlow(t, uint64(0), expiration)
}

func testIncrAndDecr(t *testing.T, val interface{}, expiration time.Duration) {
	key := "counter"
	err := mc.Set(ctx, key, val, expiration)
	assert.Nil(t, err)

	// increment
	incr, err := mc.Incr(ctx, key)
	assert.Nil(t, err)
	assert.EqualValues(t, val, incr-1)

	// decrement
	decr, err := mc.Decr(ctx, key)
	assert.Nil(t, err)
	assert.EqualValues(t, val, decr)

	assert.Nil(t, mc.Delete(ctx, key))
}

func testIncrOverFlow(t *testing.T, maxVal interface{}, expiration time.Duration) {
	key := "counter"
	err := mc.Set(ctx, key, maxVal, expiration)
	assert.Nil(t, err)

	defer func() {
		assert.Nil(t, mc.Delete(ctx, key))
	}()

	_, err = mc.Incr(ctx, key)
	assert.NotNil(t, err)
}

func testDecrOverFlow(t *testing.T, minVal interface{}, expiration time.Duration) {
	key := "counter"
	err := mc.Set(ctx, key, minVal, expiration)
	assert.Nil(t, err)

	defer func() {
		assert.Nil(t, mc.Delete(ctx, key))
	}()

	_, err = mc.Decr(ctx, key)
	assert.NotNil(t, err)
}
