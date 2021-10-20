package uow

import (
	"reflect"
	iuow "rin-echo/common/uow/interfaces"

	gormx "rin-echo/common/gorm"

	"gorm.io/gorm"
)

type (
	Repository struct {
		model interface{}
		store *gorm.DB
	}
)

func NewRepository(store *gorm.DB, model interface{}) iuow.Repository {
	if store == nil {
		panic("NewRepository requires store")
	}

	if model == nil {
		panic("NewRepository requires model")
	}

	re := Repository{store: store}

	reflectValueType := reflect.ValueOf(model).Type().Elem()
	if reflectValueType.Kind() == reflect.Ptr {
		reflectValueType = reflectValueType.Elem()
	}
	re.model = reflect.New(reflectValueType).Interface()
	return &re
}

func (r *Repository) Model() *gorm.DB {
	return r.store.Model(r.model)
}

func (r *Repository) Transaction(fc func(tx *gorm.DB) error) error {
	return transaction(r.store, fc)
}

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

func (r Repository) Update(conds map[string][]interface{}, values map[string]interface{}) error {
	return r.update(conds, values)
}

// If you want to skip Hooks methods and don’t track the update time when updating
func (r Repository) UpdateWithoutHooks(conds map[string][]interface{}, values map[string]interface{}) error {
	return r.updateWithoutHooks(conds, values)
}

func (r Repository) UpdateWithPrimaryKey(id uint, values map[string]interface{}) error {
	return r.Update(map[string][]interface{}{"id = ?": {id}}, values)
}

func (r Repository) UpdateWithoutHooksWithPrimaryKey(id uint, values map[string]interface{}) error {
	return r.UpdateWithoutHooks(map[string][]interface{}{"id = ?": {id}}, values)
}

func (r Repository) update(conds map[string][]interface{}, values map[string]interface{}) error {
	return r.Model().Where(conds).Updates(values).Error
}

// If you want to skip Hooks methods and don’t track the update time when updating
func (r Repository) updateWithoutHooks(conds map[string][]interface{}, values map[string]interface{}) error {
	return r.Model().Where(conds).UpdateColumns(values).Error
}

func (r Repository) Find(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error {
	return gormx.FindWrapError(r.Query(conds, preloads), dest)
}

func (r Repository) Get(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error {
	tx := r.Query(conds, preloads)
	return tx.Find(&dest).Error
}

func (r Repository) First(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error {
	return r.Query(conds, preloads).First(dest).Error
}

func (r Repository) Count(conds map[string][]interface{}) int64 {
	var count int64
	r.Query(conds, nil).Count(&count)
	return count
}

func (r Repository) Contains(conds map[string][]interface{}) bool {
	count := r.Count(conds)
	return count > 0
}
