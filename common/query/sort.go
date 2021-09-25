package query

import (
	"errors"
	"regexp"
	"strings"
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

func (s *Sort) Validate(entity interface{}) error {
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
