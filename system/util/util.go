package util

import (
	"reflect"
)

func DefaultValue(v interface{}, defaultValue interface{}) interface{} {
	rv := reflect.ValueOf(v)
	rDefaultValue := reflect.ValueOf(defaultValue)
	if rv.IsZero() {
		return defaultValue
	}

	if rv.Kind() == rDefaultValue.Kind() {
		return v
	}

	if rDefaultValue.Kind() == reflect.Ptr {
		ptr := reflect.New(rDefaultValue.Elem().Type())
		temp := ptr.Elem()
		temp.Set(rv)
		return ptr.Interface()
	}

	if rv.Kind() == reflect.Ptr {
		return rv.Elem().Interface()
	}

	return rv.Interface()

}
