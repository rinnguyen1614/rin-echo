package utils

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func NewStruct(src interface{}, lookupFields []string) (interface{}, error) {
	return newStruct(src, lookupFields, nil)
}

func NewStructByTag(src interface{}, lookupFields []string, tag string) (interface{}, error) {
	return newStruct(src, lookupFields, &tag)
}

func newStruct(src interface{}, lookupFields []string, tag *string) (interface{}, error) {
	typ, err := newStructType(src, lookupFields, tag)
	if err != nil {
		return nil, err
	}

	return reflect.New(typ).Interface(), nil
}

func NewSliceOfStructs(src interface{}, lookupFields []string) (interface{}, error) {
	return newSliceOfStructs(src, lookupFields, nil)
}

func NewSliceOfStructsByTag(src interface{}, lookupFields []string, tag string) (interface{}, error) {
	return newSliceOfStructs(src, lookupFields, &tag)
}

func newSliceOfStructs(src interface{}, lookupFields []string, tag *string) (interface{}, error) {
	typ, err := newStructType(src, lookupFields, tag)
	if err != nil {
		return nil, err
	}
	return reflect.New(reflect.SliceOf(typ)).Interface(), nil
}

func newStructType(src interface{}, lookupFields []string, tag *string) (reflect.Type, error) {
	var (
		retype = reflect.ValueOf(src).Type()
	)

	if retype.Kind() == reflect.Ptr {
		retype = retype.Elem()
	}

	if retype.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s must is underlying struct type", retype.Name())
	}

	var (
		retStructFields    []reflect.StructField
		structFields, _, _ = getFieldsByTag(src, tag, true)
	)

	sort.Strings(lookupFields)

	for i := 0; i < len(lookupFields); i++ {
		var (
			field = lookupFields[i]
			name  = field
			iDot  = strings.IndexByte(field, '.')
			typ   reflect.Type
		)

		if iDot != -1 {
			name = field[:iDot]
		}

		if structField, ok := structFields[name]; ok {
			typ = structField.Type

			if iDot != -1 {
				sfTypes := make([]reflect.Type, 0)
				sfType := structField.Type
				for sfType.Kind() == reflect.Ptr || sfType.Kind() == reflect.Array || sfType.Kind() == reflect.Slice {
					sfTypes = append(sfTypes, sfType)
					sfType = sfType.Elem()
				}

				if sfType.Kind() == reflect.Struct {
					partfields := make([]string, 0)
					partfields = append(partfields, field[iDot+1:])

					// group by name
					// ex: a.b.c, a.d.e => [b.c, d.e]
					for ; i < len(lookupFields)-1; i++ {
						part := lookupFields[i+1]
						iPartDot := strings.IndexByte(part, '.')
						if iPartDot != -1 && part[:iPartDot] == name {
							partfields = append(partfields, part[iPartDot+1:])
						} else {
							// because all fields in lookupFields that are sorted, so fields after the index will not be equal
							break
						}
					}

					typeEle, err := newStructType(reflect.New(sfType).Interface(), partfields, tag)
					if err != nil {
						return nil, err
					}

					for i := len(sfTypes) - 1; i >= 0; i-- {
						switch typ := sfTypes[i]; typ.Kind() {
						case reflect.Slice:
							typeEle = reflect.SliceOf(typeEle)
						case reflect.Array:
							typeEle = reflect.ArrayOf(typ.Len(), typeEle)
						case reflect.Ptr:
							typeEle = reflect.PtrTo(typ)
						}
					}

					typ = typeEle
				}
			}

			retStructFields = append(retStructFields, reflect.StructField{
				Name: structField.Name,
				Type: typ,
				Tag:  reflect.StructTag(structField.Tag),
			})

		} else {
			return nil, fmt.Errorf("field name '%s' is not found in struct '%s'", name, retype.String())
		}

	}

	return reflect.StructOf(retStructFields), nil
}
