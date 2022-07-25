package uow

import (
	"reflect"

	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type (
	queryBuilder struct {
		selects    []string
		conditions map[string][]interface{}
		orders     []string
		limit      int
		offset     int
		preloads   map[string]iuow.QueryBuilder
		model      interface{}
		tableName  string
	}
)

func NewQueryBuilder(model interface{}) (iuow.QueryBuilder, error) {
	if model == nil {
		panic("NewQueryBuilder requires model")
	}

	var (
		q = &queryBuilder{
			selects:    make([]string, 0),
			conditions: make(map[string][]interface{}),
			orders:     make([]string, 0),
			preloads:   make(map[string]iuow.QueryBuilder),
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
	return q.selects
}

func (q *queryBuilder) SetSelect(columns ...string) {
	if len(q.selects) == 0 {
		q.selects = make([]string, 0)
	}

	q.selects = append(q.selects, columns...)
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

func (q *queryBuilder) Preloads() map[string]iuow.QueryBuilder {
	if q.preloads == nil {
		q.preloads = make(map[string]iuow.QueryBuilder)
	}
	return q.preloads
}

func (q *queryBuilder) SetPreload(tableName string, query iuow.QueryBuilder) {
	q.preloads[tableName] = query
}

func BuildQuery(db *gorm.DB, queryBuilder iuow.QueryBuilder) *gorm.DB {
	var (
		tx     = db
		stmt   = gorm.Statement{DB: db}
		schema *schema.Schema
	)

	if err := stmt.Parse(queryBuilder.Model()); err != nil {
		panic(err)
	}

	schema = stmt.Schema

	for pQ, pBuilder := range queryBuilder.Preloads() {
		if pBuilder != nil {
			var (
				selsMap    = funk.Map(queryBuilder.Select(), func(x string) (string, string) { return x, x }).(map[string]string)
				selsPreMap = funk.Map(pBuilder.Select(), func(x string) (string, string) { return x, x }).(map[string]string)
				rel        = schema.Relationships.Relations[pQ]
				refs       = rel.References
			)

			if rel.JoinTable == nil {
				for _, ref := range refs {
					var (
						key    = ref.PrimaryKey.DBName
						preKey = ref.ForeignKey.DBName
					)
					if !ref.OwnPrimaryKey {
						key = ref.ForeignKey.DBName
						preKey = ref.PrimaryKey.DBName
					}
					if _, ok := selsMap[key]; !ok {
						queryBuilder.SetSelect(key)
					}
					if _, ok := selsPreMap[preKey]; !ok {
						pBuilder.SetSelect(preKey)
					}
				}
			} else {
				for _, ref := range refs {
					var (
						key = ref.PrimaryKey.DBName
					)
					if !ref.OwnPrimaryKey {
						if _, ok := selsPreMap[key]; !ok {
							pBuilder.SetSelect(key)
						}
					} else {
						if _, ok := selsMap[key]; !ok {
							queryBuilder.SetSelect(key)
						}
					}
				}
			}

			tx = tx.Preload(pQ, func(db *gorm.DB) *gorm.DB {
				return BuildQuery(db, pBuilder)
			})
		} else {
			tx = tx.Preload(pQ)
		}
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

	if limit, offset := queryBuilder.Pagination(); offset >= 0 && limit > 0 {
		tx = tx.Offset(offset).Limit(limit)
	}

	return tx
}
