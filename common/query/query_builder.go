package query

import (
	"errors"
	"reflect"

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

		Schema(db *gorm.DB) (*schema.Schema, error)

		Execute(db *gorm.DB) *gorm.DB

		Find(db *gorm.DB, dest interface{}) error

		Count(db *gorm.DB) int64
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
	}
)

func NewQueryBuilder(model interface{}) (QueryBuilder, error) {
	if model == nil {
		panic("NewQueryBuilder requires model")
	}

	var (
		q = &queryBuilder{
			sel:        make([]string, 0),
			conditions: make(map[string][]interface{}),
			orders:     make([]string, 0),
			preloads:   make(map[string]QueryBuilder),
		}
	)

	reflectType := reflect.ValueOf(model).Type().Elem()
	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	q.model = reflect.New(reflectType).Interface()

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

func (q *queryBuilder) Schema(db *gorm.DB) (*schema.Schema, error) {
	var (
		tx  = db
		smt = tx.Statement
	)

	err := smt.Parse(q.Model())
	if err != nil {
		return nil, err
	}

	return smt.Schema, nil
}

func (q *queryBuilder) Execute(db *gorm.DB) *gorm.DB {
	return fQueryBuilder(db.Model(q.model), q)
}

func fQueryBuilder(db *gorm.DB, queryBuilder QueryBuilder) *gorm.DB {
	var (
		tx = db
	)

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

	for pQ, pBuilder := range queryBuilder.Preloads() {
		tx = tx.Preload(pQ, func(db *gorm.DB) *gorm.DB {
			//if the select of preload doesn't exist foreign key, we should add the field that is foreign key in this case.
			if sels := pBuilder.Select(); len(sels) != 0 {
				refs := tx.Statement.Schema.Relationships.Relations[pQ].References
				for _, ref := range refs {
					if ref.OwnPrimaryKey {
						pBuilder.SetSelect(ref.ForeignKey.DBName)
					} else {
						pBuilder.SetSelect(ref.PrimaryKey.DBName)
					}
				}
			}

			return fQueryBuilder(db, pBuilder)
		})
	}

	return tx
}

func (q *queryBuilder) Find(db *gorm.DB, dest interface{}) error {
	query := q.Execute(db)
	if err := query.Find(dest).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		dest = reflect.Zero(reflect.TypeOf(dest))
	}

	return nil
}

func (q *queryBuilder) Count(db *gorm.DB) (total int64) {
	q.Execute(db).Count(&total)
	return
}
