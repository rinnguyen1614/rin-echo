package query

import (
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
		sel.Fields = append(sel.Fields, fs)
	}

	return sel, nil
}

func (s *Select) Validate(entity interface{}) error {
	return nil
}
