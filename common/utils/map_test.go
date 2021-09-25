package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_ConvertMapToSliceOfValues_NotPtrType(t *testing.T) {
	src := make(map[string]string)
	v1 := "nguyen"
	v2 := "anh"
	src[v1] = v1
	src[v2] = v2

	var dest []string
	err := ConvertMapToSliceOfValues(&dest, src)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	assert.Equal(t, len(src), len(dest))

}

func TestMap_ConvertMapToSliceOfValues_SlicePtrType(t *testing.T) {
	src := make(map[string]string)
	v1 := "nguyen"
	v2 := "anh"
	src[v1] = v1
	src[v2] = v2

	var dest []*string
	err := ConvertMapToSliceOfValues(&dest, src)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	assert.Equal(t, len(src), len(dest))

}

func TestMap_ConvertMapToSliceOfValues_MapPtrType(t *testing.T) {
	src := make(map[string]*string)
	v1 := "nguyen"
	v2 := "anh"
	src[v1] = &v1
	src[v2] = &v2

	var dest []*string
	err := ConvertMapToSliceOfValues(&dest, src)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	assert.Equal(t, len(src), len(dest))

}
