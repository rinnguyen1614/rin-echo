package uow

import (
	"errors"
	"reflect"

	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type (
	Repository struct {
		model  interface{}
		store  *gorm.DB
		schema *schema.Schema
	}
)

func NewRepository(store *gorm.DB, model interface{}) *Repository {
	if store == nil {
		panic("NewRepository requires store")
	}

	if model == nil {
		panic("NewRepository requires model")
	}

	var (
		re  = Repository{store: store}
		smt = gorm.Statement{DB: store}
	)

	err := smt.Parse(model)
	if err != nil {
		panic(err)
	}

	reflectValueType := reflect.ValueOf(model).Type().Elem()
	if reflectValueType.Kind() == reflect.Ptr {
		reflectValueType = reflectValueType.Elem()
	}
	re.model = reflect.New(reflectValueType).Interface()
	re.schema = smt.Schema
	return &re
}

func (r *Repository) DB() *gorm.DB {
	return r.store
}

func (r *Repository) Model() *gorm.DB {
	return r.store.Model(r.model)
}

func (r *Repository) WithTransaction(tx *gorm.DB) *Repository {
	if tx == nil {
		panic("nil tx")
	}
	r2 := new(Repository)
	*r2 = *r

	if r.schema != nil {
		r2.schema = new(schema.Schema)
		*r2.schema = *r.schema
	}

	r2.store = tx

	return r2
}

// func (r *Repository) Transaction(fc func(tx *gorm.DB) error) error {
// 	return transaction(r.store, fc)
// }

func (r Repository) Query(conds map[string][]interface{}, preloads map[string][]interface{}) *gorm.DB {
	tx := r.Model()

	for cQ, cArgs := range conds {
		tx = tx.Where(cQ, cArgs...)
	}

	for pQ, pArgs := range preloads {
		tx = tx.Preload(pQ, pArgs...)
	}

	return tx
}

func (r Repository) Create(v interface{}) error {
	return r.store.Create(v).Error
}

// CreateInBatches insert the value in batches into database
func (r Repository) CreateInBatches(v interface{}, batchSize int) error {
	return r.store.CreateInBatches(v, batchSize).Error
}

func (r Repository) Update(v interface{}) error {
	tx := r.DB()
	tx.Save(v)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r Repository) UpdateValues(conds map[string][]interface{}, values map[string]interface{}) error {
	return r.update(conds, values)
}

// If you want to skip Hooks methods and don’t track the update time when updating
func (r Repository) UpdateWithoutHooks(conds map[string][]interface{}, values map[string]interface{}) error {
	return r.updateWithoutHooks(conds, values)
}

func (r Repository) UpdateWithPrimaryKey(id uint, values map[string]interface{}) error {
	return r.update(map[string][]interface{}{"id": {id}}, values)
}

func (r Repository) UpdateWithoutHooksWithPrimaryKey(id uint, values map[string]interface{}) error {
	return r.UpdateWithoutHooks(map[string][]interface{}{"id": {id}}, values)
}

func (r Repository) update(conds map[string][]interface{}, values map[string]interface{}) error {
	tx := r.Model()
	for cQ, cArgs := range conds {
		tx = tx.Where(cQ, cArgs...)
	}
	return tx.Updates(values).Error
}

// If you want to skip Hooks methods and don’t track the update time when updating
func (r Repository) updateWithoutHooks(conds map[string][]interface{}, values map[string]interface{}) error {
	tx := r.Model()
	for cQ, cArgs := range conds {
		tx = tx.Where(cQ, cArgs...)
	}
	return tx.UpdateColumns(values).Error
}

func (r Repository) Delete(id uint) error {
	return r.store.Delete(r.model, id).Error
}

func (r Repository) DeleteMany(ids []uint) error {
	return r.store.Delete(r.model, ids).Error
}

func (r Repository) Find(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error {
	return Find(r.Query(conds, preloads), dest)
}

func (r Repository) First(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error {
	return First(r.Query(conds, preloads), dest)
}

func (r Repository) Get(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error {
	err := First(r.Query(conds, preloads), dest)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		name := r.schema.Name
		return core.NewRinError(name+"_not_found", name+" not found")
	}
	return err
}

func (r Repository) Count(conds map[string][]interface{}) (int64, error) {
	return Count(r.Query(conds, nil))
}

func (r Repository) Contains(conds map[string][]interface{}) (bool, error) {
	return Contains(r.Query(conds, nil))
}

func (r Repository) QueryBuilder(queryBuilder iuow.QueryBuilder) *gorm.DB {
	return BuildQuery(r.Model(), queryBuilder)
}

func (r Repository) QueryBuilderFind(dest interface{}, queryBuilder iuow.QueryBuilder) error {
	return Find(r.QueryBuilder(queryBuilder), dest)
}

func (r Repository) QueryBuilderFirst(dest interface{}, queryBuilder iuow.QueryBuilder) error {
	return First(r.QueryBuilder(queryBuilder), dest)
}

func (r Repository) QueryBuilderGet(dest interface{}, queryBuilder iuow.QueryBuilder) error {
	err := First(r.QueryBuilder(queryBuilder), dest)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		name := r.schema.Name
		return core.NewRinError(name+"_not_found", name+" not found")
	}
	return err
}

func (r Repository) QueryBuilderCount(queryBuilder iuow.QueryBuilder) (int64, error) {
	return Count(r.QueryBuilder(queryBuilder))
}

func (r Repository) QueryBuilderContains(queryBuilder iuow.QueryBuilder) (bool, error) {
	return Contains(r.QueryBuilder(queryBuilder))
}
