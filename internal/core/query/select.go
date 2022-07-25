package query

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type Select struct {
	Fields []string
}

func newSelect() Select {
	return Select{
		Fields: make([]string, 0),
	}
}

func ParseSelect(str, separateFields string) (Select, error) {
	if strings.Trim(str, " ") == "" {
		return Select{}, nil
	}

	var (
		sel = newSelect()
		re  = regexp.MustCompile(strings.Replace(EscapedCommaPattern, SeparateTemp, separateFields, 1))
	)

	for _, fs := range re.Split(str, -1) {
		if fs != "" {
			sel.Fields = append(sel.Fields, fs)
		}
	}

	return sel, nil
}

func (s *Select) Validate(mapFields map[string]reflect.StructField) error {
	// _, dif := funk.Difference(fields, s.Fields)
	// if notFounds := dif.([]string); len(notFounds) != 0 {
	// 	return fmt.Errorf("failed to found select's fields: %s", strings.Join(notFounds, ", "))
	// }

	notFounds, err := FindFieldNotExists(s.Fields, mapFields)
	if err != nil {
		return err
	}
	if len(notFounds) != 0 {
		return fmt.Errorf("failed to found select's fields: %s", strings.Join(notFounds, ", "))
	}
	return nil
}
