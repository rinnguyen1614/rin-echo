package query

import (
	"errors"
	"reflect"
	"rin-echo/common"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type (
	QueryBuilder interface {
		Model() interface{}

		Select() []string

		SetSelect(columns ...string)

		Conditions() map[string][]interface{}

		SetCondition(query string, args ...interface{})

		Pagination() (limit, offset int)

		SetPagination(limit, offset int)

		Orders() []string

		SetOrder(field, order string)

		Preloads() map[string]QueryBuilder

		SetPreload(tableName string, query QueryBuilder)

		DB() *gorm.DB

		WithContext(ctx common.Context) QueryBuilder

		Schema() *schema.Schema

		Query() *gorm.DB

		Find(dest interface{}) error

		First(dest interface{}) error

		Count() int64
	}

	queryBuilder struct {
		sel        []string
		conditions map[string][]interface{}
		orders     []string
		limit      int
		offset     int
		preloads   map[string]QueryBuilder
		model      interface{}
		tableName  string
		db         *gorm.DB
		schema     *schema.Schema
	}
)

func NewQueryBuilder(db *gorm.DB, model interface{}) (QueryBuilder, error) {
	if db == nil {
		panic("NewQueryBuilder requires db")
	}

	if model == nil {
		panic("NewQueryBuilder requires model")
	}

	var (
		tx  = db
		smt = tx.Statement
		q   = &queryBuilder{
			sel:        make([]string, 0),
			conditions: make(map[string][]interface{}),
			orders:     make([]string, 0),
			preloads:   make(map[string]QueryBuilder),
		}
	)

	err := smt.Parse(model)
	if err != nil {
		return nil, err
	}

	reflectType := reflect.ValueOf(model).Type().Elem()
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	q.model = reflect.New(reflectType).Interface()
	q.db = db
	q.schema = smt.Schema
	return q, nil
}

func (q *queryBuilder) Model() interface{} {
	return q.model
}

func (q *queryBuilder) Select() []string {
	return q.sel
}

func (q *queryBuilder) SetSelect(columns ...string) {
	if len(q.sel) == 0 {
		q.sel = make([]string, 0)
	}

	q.sel = append(q.sel, columns...)
}

func (q *queryBuilder) Conditions() map[string][]interface{} {
	return q.conditions
}

func (q *queryBuilder) SetCondition(query string, args ...interface{}) {
	if q.conditions == nil {
		q.conditions = make(map[string][]interface{})
	}
	q.conditions[query] = args
}

func (q *queryBuilder) Pagination() (limit, offset int) {
	return q.limit, q.offset
}

func (q *queryBuilder) SetPagination(limit, offset int) {
	q.limit = limit
	q.offset = offset
}

func (q *queryBuilder) Orders() []string {
	return q.orders
}

func (q *queryBuilder) SetOrder(field, order string) {
	if len(q.orders) == 0 {
		q.orders = make([]string, 0)
	}

	q.orders = append(q.orders, field+" "+order)
}

func (q *queryBuilder) Preloads() map[string]QueryBuilder {
	if q.preloads == nil {
		q.preloads = make(map[string]QueryBuilder)
	}
	return q.preloads
}

func (q *queryBuilder) SetPreload(tableName string, query QueryBuilder) {
	q.preloads[tableName] = query
}

func (q *queryBuilder) DB() *gorm.DB {
	return q.db
}

func (q *queryBuilder) Schema() *schema.Schema {
	return q.schema
}

func (q queryBuilder) WithContext(ctx common.Context) QueryBuilder {
	q.db.WithContext(ctx)
	return &q
}

func (q *queryBuilder) Query() *gorm.DB {
	return buildQuery(q.db.Model(q.model), q)
}

func buildQuery(db *gorm.DB, queryBuilder QueryBuilder) *gorm.DB {
	var (
		tx = db
	)

	for pQ, pBuilder := range queryBuilder.Preloads() {
		if sels := pBuilder.Select(); len(sels) != 0 {
			refs := queryBuilder.Schema().Relationships.Relations[pQ].References
			for _, ref := range refs {
				if ref.OwnPrimaryKey {
					pBuilder.SetSelect(ref.ForeignKey.DBName)
					queryBuilder.SetSelect(ref.PrimaryKey.DBName)
				} else {
					pBuilder.SetSelect(ref.PrimaryKey.DBName)
					queryBuilder.SetSelect(ref.ForeignKey.DBName)
				}
			}
		}

		tx = tx.Preload(pQ, func(db *gorm.DB) *gorm.DB {
			//if the select of preload doesn't exist foreign key, we should add the field that is foreign key in this case.
			return buildQuery(db, pBuilder)
		})
	}

	for cQ, cArgs := range queryBuilder.Conditions() {
		tx = tx.Where(cQ, cArgs...)
	}

	if sels := queryBuilder.Select(); len(sels) != 0 {
		tx = tx.Select(sels)
	}

	for _, or := range queryBuilder.Orders() {
		tx = tx.Order(or)
	}

	if limit, offset := queryBuilder.Pagination(); offset != 0 && limit != 0 {
		tx = tx.Offset(offset).Limit(limit)
	}

	return tx
}

func (q *queryBuilder) Find(dest interface{}) error {
	query := q.Query()
	if err := query.Find(dest).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		dest = reflect.Zero(reflect.TypeOf(dest))
	}

	return nil
}

func (q *queryBuilder) First(dest interface{}) error {
	return q.Query().First(dest).Error
}

func (q *queryBuilder) Count() (total int64) {
	q.Query().Count(&total)
	return
}
