package utils

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testConvert[T any](t *testing.T, s string, want interface{}, hasErr bool) {
	got, err := Parse[T](s)
	if err != nil {
		if hasErr {
			assert.Error(t, err)
		}
		fmt.Println(err)
	} else {
		if !assert.Equal(t, want, got) {
			fmt.Printf("%s want:%v , got: %v\n", s, want, got)
		}
	}
}

func TestConvert_ToBool(t *testing.T) {
	testConvert[bool](t, "true", true, false)
	testConvert[bool](t, "false", false, false)
	testConvert[bool](t, "1", true, false)
	testConvert[bool](t, "0", false, false)
	testConvert[bool](t, "re", true, true)
}

func TestConvert_ToInt(t *testing.T) {
	want := math.MaxInt8
	str := ToString(want)
	testConvert[int](t, str, want, false)
	testConvert[int64](t, str, int64(want), false)
	testConvert[int8](t, str, int8(want), false)
	testConvert[int16](t, str, int16(want), false)
	testConvert[int32](t, str, int32(want), false)
}

func TestConvert_ToUInt(t *testing.T) {
	want := math.MaxUint8
	str := ToString(want)
	testConvert[uint](t, str, uint(want), false)
	testConvert[uint64](t, str, uint64(want), false)
	testConvert[uint8](t, str, uint8(want), false)
	testConvert[uint16](t, str, uint16(want), false)
	testConvert[uint32](t, str, uint32(want), false)
	testConvert[uint32](t, str+"invalid", 0, true)
}

func TestConvert_ToFloat(t *testing.T) {
	want := math.MaxFloat32
	str := ToString(want)
	testConvert[float32](t, str, float32(want), false)
	testConvert[float64](t, str, float64(want), false)
}

func TestConvert_ToComplex(t *testing.T) {
	want := complex(5, 6)
	str := "5+6i"
	testConvert[complex64](t, str, complex64(want), false)
	testConvert[complex128](t, str, complex128(want), false)
}

func TestConvert_ToByte(t *testing.T) {
	want := []byte("test")
	str := string(want)
	testConvert[[]byte](t, str, want, false)
}

func TestConvert_ToStruct(t *testing.T) {
	type tempT struct {
		Number int `json:"name"`
	}
	want := tempT{
		Number: 1,
	}
	str := ToString(want)
	testConvert[tempT](t, str, want, false)
}

func TestConvert_ToSlice(t *testing.T) {
	want := []int{1, 2, 3}
	str := ToString(want)
	testConvert[[]int](t, str, want, false)
}
