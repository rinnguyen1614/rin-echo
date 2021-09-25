package uow

import (
	"errors"
	"reflect"
	"rin-echo/common"
	iuow "rin-echo/common/uow/interfaces"

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

func (r *Repository) WithContext(ctx common.Context) *gorm.DB {
	return r.store.WithContext(ctx)
}

func (r *Repository) Model(ctx common.Context) *gorm.DB {
	return r.WithContext(ctx).Model(r.model)
}

func (r *Repository) Transaction(ctx common.Context, fc func(tx *gorm.DB) error) error {
	return transaction(r.WithContext(ctx), fc)
}

func (r Repository) Query(ctx common.Context, conds map[string][]interface{}, preloads map[string][]interface{}) *gorm.DB {
	tx := r.Model(ctx)

	for cQ, cArgs := range conds {
		tx = tx.Where(cQ, cArgs...)
	}

	for pQ, pArgs := range preloads {
		tx = tx.Preload(pQ, pArgs...)
	}

	return tx
}

func (r Repository) Create(ctx common.Context, v interface{}) error {
	return r.store.Create(v).Error
}

// CreateInBatches insert the value in batches into database
func (r Repository) CreateInBatches(ctx common.Context, v interface{}, batchSize int) error {
	return r.WithContext(ctx).CreateInBatches(v, batchSize).Error
}

func (r Repository) Update(ctx common.Context, conds map[string][]interface{}, values map[string]interface{}) error {
	return r.update(ctx, conds, values)
}

// If you want to skip Hooks methods and don’t track the update time when updating
func (r Repository) UpdateWithoutHooks(ctx common.Context, conds map[string][]interface{}, values map[string]interface{}) error {
	return r.updateWithoutHooks(ctx, conds, values)
}

func (r Repository) UpdateWithPrimaryKey(ctx common.Context, id uint, values map[string]interface{}) error {
	return r.Update(ctx, map[string][]interface{}{"id = ?": {id}}, values)
}

func (r Repository) UpdateWithoutHooksWithPrimaryKey(ctx common.Context, id uint, values map[string]interface{}) error {
	return r.UpdateWithoutHooks(ctx, map[string][]interface{}{"id = ?": {id}}, values)
}

func (r Repository) update(ctx common.Context, conds map[string][]interface{}, values map[string]interface{}) error {
	return r.Model(ctx).Where(conds).Updates(values).Error
}

// If you want to skip Hooks methods and don’t track the update time when updating
func (r Repository) updateWithoutHooks(ctx common.Context, conds map[string][]interface{}, values map[string]interface{}) error {
	return r.Model(ctx).Where(conds).UpdateColumns(values).Error
}

func (r Repository) Find(ctx common.Context, dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error {
	if err := r.Query(ctx, conds, preloads).Find(dest).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		dest = reflect.Zero(reflect.TypeOf(dest))
	}

	return nil
}

func (r Repository) First(ctx common.Context, dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error {
	if err := r.Query(ctx, conds, preloads).First(dest).Error; err != nil {
		return err
	}
	return nil
}
