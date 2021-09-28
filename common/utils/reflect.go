package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func GetTypeName(i interface{}) string {
	return reflect.TypeOf(i).String()
}

// func RealValue(v reflect.Value) (interface{}, error) {
// 	val := reflect.ValueOf(v)
// 	ptr := val
// 	if ptr.Kind() != reflect.Ptr {
// 		return nil, fmt.Errorf("type not a pointer: " + val.Type().String())
// 	}
// 	switch v := ptr.Elem(); v.Kind() {
// 	case reflect.Bool:
// 		v.Bool()
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		v.Int()
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
// 		v.Uint()
// 	case reflect.String:
// 		v.String()
// 	case reflect.Slice:
// 		// For now, can only handle (renamed) []byte.
// 		typ := v.Type()
// 		if typ.Elem().Kind() != reflect.Uint8 {
// 			s.errorString("can't scan type: " + val.Type().String())
// 		}
// 		str := s.convertString(verb)
// 		v.Set(reflect.MakeSlice(typ, len(str), len(str)))
// 		for i := 0; i < len(str); i++ {
// 			v.Index(i).SetUint(uint64(str[i]))
// 		}
// 		v.Interface()
// 	case reflect.Float32, reflect.Float64:
// 		v.Float()
// 	case reflect.Complex64, reflect.Complex128:
// 		v.Complex()
// 	default:
// 		s.errorString("can't scan type: " + val.Type().String())
// 	}
// }

// reflectValue := reflect.ValueOf(dest)
// reflectValueType := reflectValue.Type().Elem()
// if reflectValueType.Kind() == reflect.Ptr {
// 	reflectValueType = reflectValueType.Elem()
// }
// dest = reflect.New(reflectValueType).Interface()

// key is tag's name, value is field of tag
func GetFieldsByTag(model interface{}, tag string) (fieldsByKey map[string]reflect.StructField, keyOptions map[string][]string, err error) {
	var (
		retype = reflect.ValueOf(model).Type().Elem()
	)

	fieldsByKey = make(map[string]reflect.StructField)
	keyOptions = make(map[string][]string)

	for retype.Kind() == reflect.Ptr || retype.Kind() == reflect.Array || retype.Kind() == reflect.Slice {
		retype = retype.Elem()
	}

	if retype.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("%s must is underlying struct type", retype.Name())
	}

	for i := 0; i < retype.NumField(); i++ {
		field := retype.Field(i)
		if value, ok := field.Tag.Lookup(tag); ok {
			key, options := parseTag(value)
			fieldsByKey[key] = field
			keyOptions[key] = options
		}
	}

	return fieldsByKey, keyOptions, nil
}

// key is json tag's name, value is field of tag
func GetFieldsByJsonTag(model interface{}) (map[string]reflect.StructField, map[string][]string, error) {
	return GetFieldsByTag(model, "json")
}

// tag is one of followings:
// ""
// "name"
// "name,opt"
// "name,opt,opt2"
// ",opt"
func parseTag(tag string) (key string, options []string) {
	res := strings.Split(tag, ",")
	return res[0], res[1:]
}
