package query

import (
	"database/sql"
	"fmt"
	"reflect"
	"rin-echo/common"
	"rin-echo/common/model"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/common/utils"
	"strings"

	"github.com/jinzhu/copier"
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
	EscapedCommaPattern = `\s*_separate_\s*`
)

type Interface interface {
	Validate() error
}

type Query struct {
	// all select fields
	allSelect Select
	// all sort fields
	allSort Sort

	paging Paging
	// sort fields by query
	sort Sort
	// select fields by query
	sel      Select
	filter   Filter
	preloads map[string]Query
	Error    error
}

func newQuery() Query {
	return Query{
		allSelect: newSelect(),
		allSort:   newSort(),
		sel:       newSelect(),
		sort:      newSort(),
		filter:    newFilter(),
		preloads:  make(map[string]Query),
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

func (q *Query) FlatSelect() []string {
	return q.allSelect.Fields
}

func (q *Query) FlatSort() []SortField {
	return q.allSort.Fields
}

func Parse(sorts, selects, filters string, page, pageSize int, config Config) (*Query, error) {
	q := newQuery()

	q.paging = ParsePaging(pageSize, page, config.MaxPageSize, config.MinPageSize)

	sel, err := ParseSelect(selects, config.SeparateFields)
	if err != nil {
		return nil, err
	}

	q.allSelect = sel

	for _, sField := range q.allSelect.Fields {
		iterFunc(&q, sField, sField, func(q *Query, fieldLastDot string) {
			q.sel.Fields = append(q.sel.Fields, fieldLastDot)
		})
	}

	sort, err := ParseSort(sorts, config.SeparateFields, config.SeparateSortField)
	if err != nil {
		return nil, err
	}

	q.allSort = sort

	for sName, sField := range q.allSort.FieldsByName {
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

func (q *Query) Validate(fields map[string]reflect.StructField) error {
	q.AddError(q.allSelect.Validate(fields))
	q.AddError(q.allSort.Validate(fields))
	q.AddError(q.filter.Validate(fields))
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

// Get queryResult from query.
//
// - repo: the repository executes queryBuilder (find & count).
//
// - queryBuilder: it is filled by query's select, order and filter
//
// - preloadBuilders: Each of preloadBuilders is filled by select, order and filter of preload's query.
// It is preload of queryBuilder
//
// - srcModel: should be a struct type/ pointer of struct type, it is used to querying in DB
//
// - desModel: should be a struct type/ pointer of struct type, it is used to responsed to the client
func (q *Query) QueryResult(repo iuow.Repository, queryBuilder iuow.QueryBuilder, preloadBuilders map[string]iuow.QueryBuilder, srcModel interface{}, desModel interface{}) (*model.QueryResult, error) {
	var (
		typs = map[string]reflect.Type{
			"src":  reflect.ValueOf(srcModel).Type(),
			"dest": reflect.ValueOf(desModel).Type()}
	)

	for k, typ := range typs {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}

		if typ.Kind() != reflect.Struct {
			return nil, fmt.Errorf("%s must is underlying struct type", k)
		}
	}

	var (
		srcModels      = reflect.New(reflect.SliceOf(typs["src"])).Interface()
		fields, _, err = utils.GetFullFieldsByJsonTag(desModel)
	)
	if err != nil {
		return nil, err
	}

	totalRecords, err := q.BindQueryBuilder(queryBuilder, preloadBuilders, repo.DB(), fields)
	if err != nil {
		return nil, err
	}

	if err = repo.QueryBuilderFind(srcModels, queryBuilder); err != nil {
		return nil, err
	}

	// new slice of desModel with fields that get from query' selects
	prune, err := utils.NewSliceOfStructsByTag(desModel, q.FlatSelect(), "json")
	if err != nil {
		return nil, err
	}

	err = copier.CopyWithOption(prune, srcModels, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}

	return model.NewQueryResult(prune, totalRecords, q.Paging().Limit, q.Paging().Offset), nil
}

// bind query's select, order, filter to queryBuilder and preloadBuilders
//
// - db: use to build query for filter's query
//
// - fields: are allowed to query
func (q Query) BindQueryBuilder(queryBuilder iuow.QueryBuilder, preloadBuilders map[string]iuow.QueryBuilder, db *gorm.DB, fields map[string]reflect.StructField) (totalRecords int64, err error) {
	if err := q.Validate(fields); err != nil {
		return 0, common.NewRinErrorWithInner(err, "query_parameters_errors", q.Error.Error())
	}

	queryBuilder.SetPagination(q.Paging().Limit, q.Paging().Offset)

	if err := q.bindSelectAndOrder(queryBuilder, preloadBuilders, fields); err != nil {
		return 0, err
	}

	totalRecords, err = q.bindCondition(queryBuilder, preloadBuilders, db, fields)
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}

func (q Query) bindSelectAndOrder(queryBuilder iuow.QueryBuilder, preloadBuilders map[string]iuow.QueryBuilder, fields map[string]reflect.StructField) error {
	for _, sort := range q.Sort().Fields {
		queryBuilder.SetOrder(sort.Field, sort.Order)
	}

	if sels := q.Select().Fields; len(sels) != 0 {
		queryBuilder.SetSelect(sels...)
	}

	for pre, preQuery := range q.preloads {
		field, ok := fields[pre]
		if !ok {
			return fmt.Errorf("failed to found '%s' field", pre)
		}

		if preBuilder, ok := preloadBuilders[field.Name]; ok {
			queryBuilder.SetPreload(field.Name, preBuilder)
			fields, _, err := utils.GetFieldsByJsonTag(reflect.New(field.Type).Interface())
			if err != nil {
				return err
			}

			if err := preQuery.bindSelectAndOrder(preBuilder, preloadBuilders, fields); err != nil {
				return err
			}
		}
	}

	return nil
}

// if query has condition, it
//
// - executes WHERE CLAUSE with limit & offset
// to get the values of PrimaryFieldDBNames and set them in queryBuilder's condition
//
// - counts records with filter (without limit & offset)
//
// - resets querybuilder's offset, because the values of queryBuiler's condition had limited & offseted.
func (q Query) bindCondition(queryBuilder iuow.QueryBuilder, preloadBuilders map[string]iuow.QueryBuilder, db *gorm.DB, fields map[string]reflect.StructField) (totalRecords int64, err error) {
	var (
		tx             = db.WithContext(db.Statement.Context)
		stmt           = gorm.Statement{DB: db}
		filter         = q.Filter()
		preloadSchemas = make(map[string]*schema.Schema)
		primarySchema  *schema.Schema
	)

	if err := stmt.Parse(queryBuilder.Model()); err != nil {
		panic(err)
	}
	primarySchema = stmt.Schema

	for pre, preBuilder := range preloadBuilders {
		if err := stmt.Parse(preBuilder.Model()); err != nil {
			panic(err)
		}
		preloadSchemas[pre] = stmt.Schema
	}

	tx, hasCondition, fieldNames, err := filter.BuildQuery(tx, primarySchema, preloadSchemas, fields)
	if err != nil {
		return 0, err
	}

	if !hasCondition {
		return uow.Count(stmt.DB), err
	}

	if err = setCondition(tx, queryBuilder); err != nil {
		return 0, err
	}

	for preloadName, preloadBuilder := range preloadBuilders {
		if _, ok := fieldNames[preloadName]; ok {
			if err = setCondition(tx, preloadBuilder); err != nil {
				return 0, err
			}
		}
	}

	// - resets querybuilder's offset, because the values of queryBuiler's condition had limited & offseted.
	limit, _ := queryBuilder.Pagination()
	queryBuilder.SetPagination(limit, 0)

	var columns []string
	for _, field := range primarySchema.PrimaryFieldDBNames {
		col := primarySchema.Table + "." + field
		columns = append(columns, col)
	}
	return uow.Count(tx.Table(primarySchema.Table).Distinct(columns)), nil
}

func setCondition(db *gorm.DB, queryBuilder iuow.QueryBuilder) error {
	var (
		tx          = db.WithContext(db.Statement.Context)
		stmt        = gorm.Statement{DB: tx}
		columns     []string
		columnTypes []reflect.Type
		modelSchema *schema.Schema
	)

	err := stmt.Parse(queryBuilder.Model())
	if err != nil {
		panic(err)
	}
	modelSchema = stmt.Schema

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

	if limit, offset := queryBuilder.Pagination(); offset >= 0 && limit > 0 {
		tx = tx.Offset(offset).Limit(limit)
	}

	rows, err = tx.Distinct(columns).Rows()
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

	if len(valueMaps) != 0 {
		for k, v := range valueMaps {
			queryBuilder.SetCondition(k, v)
		}
	} else {
		queryBuilder.SetCondition("1 = 0")
	}

	return nil
}

func prepareValues(values []interface{}, columnTypes []reflect.Type) {
	for idx := range columnTypes {
		values[idx] = reflect.New(reflect.PtrTo(columnTypes[idx])).Interface()
	}
}
