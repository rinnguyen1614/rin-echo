package query

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

var (
	OrderDESC = "desc"
	OrderASC  = ""
)

type Sort struct {
	Fields       []SortField
	FieldsByName map[string]SortField
}

func newSort() Sort {
	return Sort{
		Fields:       make([]SortField, 0),
		FieldsByName: make(map[string]SortField),
	}
}

func ParseSort(str, separateFields, separateSortField string) (Sort, error) {
	if strings.Trim(str, " ") == "" {
		return Sort{}, nil
	}

	var (
		sort = newSort()
		re   = regexp.MustCompile(strings.Replace(EscapedCommaPattern, SeparateTemp, separateFields, 1))
	)

	for _, fs := range re.Split(str, -1) {
		sortField, err := sortFieldString(strings.Trim(fs, " ")).Parse(separateSortField)
		if err == nil {
			sort.Fields = append(sort.Fields, sortField)
			sort.FieldsByName[sortField.Field] = sortField
		}
	}

	return sort, nil
}

func (s *Sort) Validate(mapFields map[string]reflect.StructField) error {
	// var (
	// 	notFounds []string
	// 	mapFields = funk.Map(fields, func(x string) (string, string) { return x, x }).(map[string]string)
	// )
	// for sField := range s.FieldsByName {
	// 	if _, ok := mapFields[sField]; !ok {
	// 		notFounds = append(notFounds, sField)
	// 	}
	// }
	// if len(notFounds) != 0 {
	// 	return fmt.Errorf("failed to found sort's fields: %s", strings.Join(notFounds, ", "))
	// }

	notFounds, err := FindFieldNotExists(funk.Keys(s.FieldsByName).([]string), mapFields)
	if err != nil {
		return err
	}
	if len(notFounds) != 0 {
		return fmt.Errorf("failed to found sort's fields: %s", strings.Join(notFounds, ", "))
	}
	return nil
}

type SortField struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

func (f *SortField) String() string {
	return f.Field + " " + f.Order
}

type sortFieldString string

func (s sortFieldString) Parse(sep string) (SortField, error) {
	if s == "" {
		return SortField{}, errors.New("empty sort field")
	}

	if sep == "" {
		return SortField{}, errors.New("sep empty")
	}

	var field, order string
	splited := strings.Split(string(s), sep)
	switch len := len(splited); {
	case len >= 2:
		field = splited[0]
		order = splited[1]
		if order != OrderDESC && order != OrderASC {
			order = OrderASC
		}
	default:
		field = splited[0]
	}

	return SortField{
		Field: field,
		Order: order,
	}, nil
}
