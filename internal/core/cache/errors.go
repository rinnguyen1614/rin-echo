package cache

import "errors"

var (
	ErrKeyDoNotExists    = errors.New("key does not exist")
	ErrKeyExpired        = errors.New("key expired")
	ErrIncrementOverflow = errors.New("this incr invocation will overflow.")
	ErrDecrementOverflow = errors.New("this decr invocation will overflow.")
	ErrNotIntegerType    = errors.New("item val is not (u)int (u)int32 (u)int64")
)
