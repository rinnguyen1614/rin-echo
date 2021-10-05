package query

import (
	"database/sql"
	"fmt"
	"reflect"
	"rin-echo/common/utils"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type (
	Config struct {
		MaxPageSize       int
		MinPageSize       int
		SeparateFields    string
		SeparateSortField string
	}
)

var (
	SeparateTemp        = "_separate_"
	EscapedCommaPattern = `_separate_\s*`
)

type Interface interface {
	Validate() error
}

type Query struct {
	paging   Paging
	sort     Sort
	sel      Select
	filter   Filter
	preloads map[string]Query
	Error    error
}

func newQuery() Query {
	return Query{
		sel:      newSelect(),
		sort:     newSort(),
		filter:   newFilter(),
		preloads: make(map[string]Query),
	}
}

func (q *Query) Paging() Paging {
	return q.paging
}

func (q *Query) Select() Select {
	return q.sel
}

func (q *Query) Sort() Sort {
	return q.sort
}

func (q *Query) Filter() Filter {
	return q.filter
}

func Parse(sorts, selects, filters string, page, pageSize int, config Config) (*Query, error) {
	q := newQuery()

	q.paging = ParsePaging(pageSize, page, config.MaxPageSize, config.MinPageSize)

	sel, err := ParseSelect(selects, config.SeparateFields)
	if err != nil {
		return nil, err
	}

	for _, sField := range sel.Fields {
		iterFunc(&q, sField, sField, func(q *Query, fieldLastDot string) {
			q.sel.Fields = append(q.sel.Fields, fieldLastDot)
		})
	}

	sort, err := ParseSort(sorts, config.SeparateFields, config.SeparateSortField)
	if err != nil {
		return nil, err
	}

	for sName, sField := range sort.FieldsByName {
		iterFunc(&q, sName, sField, func(q *Query, fieldLastDot string) {
			sortField := SortField{
				Field: fieldLastDot,
				Order: sField.Order,
			}
			q.sort.Fields = append(q.sort.Fields, sortField)
			q.sort.FieldsByName[fieldLastDot] = sortField
		})
	}

	fils, err := ParseFilter(filters)
	if err != nil {
		return nil, err
	}
	q.filter = fils

	return &q, nil
}

func iterFunc(query *Query, name string, field interface{}, setFunc func(*Query, string)) {
	if iDot := strings.IndexRune(name, '.'); iDot != -1 {
		preStr := strings.Split(name, ".")
		preNext := query.preloads

		for i, len := 0, len(preStr); i < len-1; i++ {
			preName := preStr[i]
			curPre, ok := preNext[preName]
			if !ok {
				curPre = newQuery()
			}

			if i+1 == len-1 && setFunc != nil {
				setFunc(&curPre, preStr[i+1])
			}

			preNext[preName] = curPre
			preNext = curPre.preloads
		}
	} else if setFunc != nil {
		setFunc(query, name)
	}
}

func (q *Query) Validate(entity interface{}) error {
	q.AddError(q.sel.Validate(entity))

	q.AddError(q.sort.Validate(entity))

	return q.Error
}

func (q *Query) AddError(err error) error {
	if q.Error == nil {
		q.Error = err
	} else if err != nil {
		q.Error = fmt.Errorf("%v; %w", q.Error, err)
	}
	return q.Error
}

func (q Query) Bind(queryBuilder QueryBuilder, preloadBuilders map[string]QueryBuilder, modelRes interface{}) error {

	err := q.bindSelectAndOrder(queryBuilder, preloadBuilders, modelRes)
	if err != nil {
		return err
	}

	err = q.bindCondition(queryBuilder, modelRes)
	if err != nil {
		return err
	}

	return nil
}

func (q Query) bindSelectAndOrder(queryBuilder QueryBuilder, preloadBuilders map[string]QueryBuilder, modelRes interface{}) error {
	fields, _, err := utils.GetFieldsByJsonTag(modelRes)
	if err != nil {
		return err
	}

	queryBuilder.SetPagination(q.Paging().Limit, q.Paging().Offset)

	for _, sort := range q.Sort().Fields {
		queryBuilder.SetOrder(sort.Field, sort.Order)
	}

	if sels := q.Select().Fields; len(sels) != 0 {
		queryBuilder.SetSelect(sels...)
	}

	for pre, preQuery := range q.preloads {
		field, ok := fields[pre]
		if !ok {
			return fmt.Errorf("Not found '%s' field", pre)
		}

		if preBuilder, ok := preloadBuilders[field.Name]; ok {
			queryBuilder.SetPreload(field.Name, preBuilder)
			if err := preQuery.bindSelectAndOrder(preBuilder, preloadBuilders, reflect.New(field.Type).Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (q Query) bindCondition(queryBuilder QueryBuilder, modelRes interface{}) error {
	var (
		tx     *gorm.DB
		filter = q.Filter()
	)

	tx, hasCondition, fieldNames, err := filter.BuildQuery(queryBuilder, modelRes)
	if err != nil {
		return err
	}

	if hasCondition {
		var funcSetCondition func(*gorm.DB, QueryBuilder) error

		funcSetCondition = func(tx *gorm.DB, qB QueryBuilder) error {
			err := setCondition(tx, qB)
			if err != nil {
				return err
			}

			for preloadName, preloadBuilder := range qB.Preloads() {
				if _, ok := fieldNames[preloadName]; ok {
					err = funcSetCondition(tx, preloadBuilder)
					if err != nil {
						return err
					}
				}
			}
			return nil
		}

		err = funcSetCondition(tx, queryBuilder)
		if err != nil {
			return err
		}

	}
	return nil
}

func setCondition(db *gorm.DB, queryBuilder QueryBuilder) error {
	var (
		tx          = db
		columns     []string
		columnTypes []reflect.Type
		modelSchema *schema.Schema
	)

	modelSchema = queryBuilder.Schema()

	for _, field := range modelSchema.PrimaryFieldDBNames {
		col := modelSchema.Table + "." + field
		columns = append(columns, col)
		columnTypes = append(columnTypes, modelSchema.FieldsByDBName[field].FieldType)
	}

	var (
		valueMaps = make(map[string][]interface{}, 0)
		values    = make([]interface{}, len(columns))
		rows      *sql.Rows
	)

	rows, err := tx.Distinct(columns).Rows()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		prepareValues(values, columnTypes)

		rows.Scan(values...)

		for i := range columns {
			valueMaps[columns[i]] = append(valueMaps[columns[i]], values[i])
		}
	}

	for k, v := range valueMaps {
		queryBuilder.SetCondition(k, v)
	}

	return nil
}

func prepareValues(values []interface{}, columnTypes []reflect.Type) {
	for idx := range columnTypes {
		values[idx] = reflect.New(reflect.PtrTo(columnTypes[idx])).Interface()
	}
}
