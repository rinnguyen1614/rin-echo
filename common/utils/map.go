package utils

import (
	"errors"
	"reflect"
)

// Convert map to slice of values.
func ConvertMapToSliceOfValues(dest interface{}, src interface{}) error {
	srcReflectValue := reflect.ValueOf(src)
	if srcReflectValue.Kind() == reflect.Ptr {
		srcReflectValue = srcReflectValue.Elem()
	}
	if srcReflectValue.Kind() != reflect.Map {
		return errors.New("The src's not a map type")
	}
	if srcReflectValue.Len() == 0 {
		return nil
	}

	destReflectValue := reflect.ValueOf(dest)
	if destReflectValue.Kind() == reflect.Ptr {
		destReflectValue = destReflectValue.Elem()
	}

	if destReflectValue.Kind() != reflect.Slice {
		return errors.New("The dest's not a slice type")
	}

	isSrcPtr := srcReflectValue.Type().Elem().Kind() == reflect.Ptr
	isDestPtr := destReflectValue.Type().Elem().Kind() == reflect.Ptr
	iter := srcReflectValue.MapRange()
	for iter.Next() {
		v := iter.Value()
		if isSrcPtr != isDestPtr {
			if isSrcPtr {
				destReflectValue.Set(reflect.Append(destReflectValue, v.Elem()))
			} else {
				ptr := reflect.New(v.Type())
				temp := ptr.Elem()
				temp.Set(v)
				destReflectValue.Set(reflect.Append(destReflectValue, ptr))
			}
		} else {
			destReflectValue.Set(reflect.Append(destReflectValue, v))
		}
	}
	return nil
}

// Convert map to slice of keys.

func ConvertMapToStruct(dest interface{}, src map[string]interface{}) error {
	destReflectValue := reflect.ValueOf(dest)

	if destReflectValue.Kind() != reflect.Ptr || destReflectValue.IsNil() {
		return errors.New("dest requires non-nil pointer")
	}

	destReflectValue = destReflectValue.Elem()

	for k, v := range src {
		// Set value
		_ = setField(&destReflectValue, k, v)
	}

	return nil
}

func setField(reflectValue *reflect.Value, name string, value interface{}) error {
	fieldValue := reflectValue.FieldByName(name)
	if !fieldValue.IsValid() {
		return errors.New("No such field " + name + " in object")
	}

	if !fieldValue.CanSet() {
		return errors.New("Cannot set" + name + " field value")
	}

	fieldType := fieldValue.Type()
	if fieldType.Kind() == reflect.Ptr {
		ptr := reflect.New(fieldType)
		fieldValue = ptr.Elem()
	}
	val := reflect.ValueOf(value)
	if fieldType != val.Type() {
		return errors.New("Provided value type didn't match reflectValue field type")
	}

	fieldValue.Set(val)

	return nil
}
