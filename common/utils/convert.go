package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func MustParseUint(s string, base int, bitSize int) uint64 {
	v, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		panic(v)
	}
	return v
}

func MustParseInt(s string, base int, bitSize int) int64 {
	v, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		panic(v)
	}
	return v
}

func MustParseFloat(s string, bitSize int) float64 {
	v, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		panic(v)
	}
	return v
}

func MustParseComplex(s string, bitSize int) complex128 {
	v, err := strconv.ParseComplex(s, bitSize)
	if err != nil {
		panic(v)
	}
	return v
}

func MustParseBool(s string) bool {
	v, err := strconv.ParseBool(s)
	if err != nil {
		panic(v)
	}
	return v
}

func MustParse[T any](str string) T {
	v, err := Parse[T](str)
	if err != nil {
		panic(err)
	}
	return v
}

func Parse[T any](str string) (T, error) {
	var (
		t   T
		v   interface{}
		err error
	)
	switch (interface{})(t).(type) {
	case string:
		v = str
	case bool:
		v, err = strconv.ParseBool(str)
	case int64:
		v, err = strconv.ParseInt(str, 10, 64)
	case int:
		v, err = strconv.Atoi(str)
	case int8:
		var i int64
		i, err = strconv.ParseInt(str, 10, 8)
		if err == nil {
			v = int8(i)
		}
	case int16:
		var i int64
		i, err = strconv.ParseInt(str, 10, 16)
		if err == nil {
			v = int16(i)
		}
	case int32:
		var i int
		i, err = strconv.Atoi(str)
		if err == nil {
			v = int32(i)
		}
	case uint:
		var u uint64
		u, err = strconv.ParseUint(str, 10, 32)
		if err == nil {
			v = uint(u)
		}
	case uint8:
		var u uint64
		u, err = strconv.ParseUint(str, 10, 8)
		if err == nil {
			v = uint8(u)
		}
	case uint16:
		var u uint64
		u, err = strconv.ParseUint(str, 10, 16)
		if err == nil {
			v = uint16(u)
		}
	case uint32:
		var u uint64
		u, err = strconv.ParseUint(str, 10, 32)
		if err == nil {
			v = uint32(u)
		}
	case uint64:
		v, err = strconv.ParseUint(str, 10, 64)
	case float32:
		var f float64
		f, err = strconv.ParseFloat(str, 32)
		if err == nil {
			v = float32(f)
		}
	case float64:
		v, err = strconv.ParseFloat(str, 64)
	case complex64:
		var c complex128
		c, err = strconv.ParseComplex(str, 64)
		if err == nil {
			v = complex64(c)
		}
	case complex128:
		v, err = strconv.ParseComplex(str, 128)
	case []byte:
		v = []byte(str)
	default:
		err = json.Unmarshal([]byte(str), &t)
		v = t
	}

	if err != nil {
		return t, err
	}
	return v.(T), nil
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', 6, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', 6, 64)
	case complex64:
		return strconv.FormatComplex(complex128(v), 'e', 6, 64)
	case complex128:
		return strconv.FormatComplex(v, 'e', 6, 128)
	case fmt.Stringer:
		return v.String()
	default:
		bytes, err := json.Marshal(v)
		if err != nil {
			panic(fmt.Sprintf("Error occured during marshaling. Error: %v.", err))
		}
		return string(bytes)
	}
}
