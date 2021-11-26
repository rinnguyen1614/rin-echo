package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func GetTypeName(i interface{}) string {
	return reflect.TypeOf(i).String()
}

// It includes embbed struct's fields.
// key is field name, value is field
func GetFullFields(model interface{}) (map[string]reflect.StructField, error) {
	fields, _, err := getFieldsByTag(model, nil, true)
	return fields, err
}

// key is field name, value is field
func GetFields(model interface{}) (map[string]reflect.StructField, error) {
	fields, _, err := getFieldsByTag(model, nil, false)
	return fields, err
}

// It includes embbed struct's fields.
// key is tag's name, value is field of tag
func GetFullFieldsByTag(model interface{}, tag string) (fieldsByKey map[string]reflect.StructField, keyOptions map[string][]string, err error) {
	return getFieldsByTag(model, &tag, true)
}

// key is tag's name, value is field of tag
func GetFieldsByTag(model interface{}, tag string) (fieldsByKey map[string]reflect.StructField, keyOptions map[string][]string, err error) {
	return getFieldsByTag(model, &tag, false)
}

// It includes embbed struct's fields.
// key is json tag's name, value is field of tag
func GetFullFieldsByJsonTag(model interface{}) (fieldsByKey map[string]reflect.StructField, keyOptions map[string][]string, err error) {
	return GetFullFieldsByTag(model, "json")
}

// key is json tag's name, value is field of tag
func GetFieldsByJsonTag(model interface{}) (fieldsByKey map[string]reflect.StructField, keyOptions map[string][]string, err error) {
	return GetFieldsByTag(model, "json")
}

func getFieldsByTag(model interface{}, tag *string, includeEmbbed bool) (fieldsByKey map[string]reflect.StructField, keyOptions map[string][]string, err error) {
	var (
		retype = reflect.ValueOf(model).Type()
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
		if field.Anonymous && includeEmbbed {
			fieldsEmbbed, optionsEmbbed, _ := getFieldsByTag(reflect.New(field.Type).Interface(), tag, includeEmbbed)
			for k, v := range fieldsEmbbed {
				fieldsByKey[k] = v
			}

			for k, v := range optionsEmbbed {
				keyOptions[k] = v
			}
		} else {
			if tag == nil {
				fieldsByKey[field.Name] = field
			} else if value, ok := field.Tag.Lookup(*tag); ok {
				key, options := parseTag(value)
				fieldsByKey[key] = field
				keyOptions[key] = options
			}
		}
	}

	return fieldsByKey, keyOptions, nil
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
