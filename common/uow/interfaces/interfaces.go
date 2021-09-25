package interfaces

import (
	"rin-echo/common"
	"rin-echo/common/utils"

	"gorm.io/gorm"
)

type (
	UnitOfWork interface {
		DB() *gorm.DB

		WithContext(common.Context) *gorm.DB

		Transaction(func(*gorm.DB) error) error

		TransactionUnitOfWork(func(UnitOfWork) error) error

		Rollback(tx *gorm.DB) error

		RollbackUnitOfWork(ux UnitOfWork) error

		GetRepository(key string) Repository

		SetRepository(key string, val Repository)
	}

	Repository interface {
		Model(common.Context) *gorm.DB

		Transaction(ctx common.Context, fc func(tx *gorm.DB) error) (err error)

		/*
			- conds : key is query string, value is the array of the query's arguments
				Ex1: conds= map[string][]interface{}{ "id" : 1 } // is equal `id = 1`
				Ex2: conds= map[string][]interface{}{ "id NOT IN (?) AND id <> ?" : { []uint{1,2,3}, 6} } // is equal `id NOT IN (1,2,3) AND id <> 6`
			- preloads : key is table name & value (options) is function (same type as the one used for Scopes) or query:
				If it is query string, [0]: condition for preload (option), [1:] arguments for condition.
				Ex: preloads = map[string][]interface{}{ "users" : {"username = ? AND role_id = ? ", "rin-echo", 1}} // is equal `username = "rin-echo" AND role_id = 1`
		*/
		Query(ctx common.Context, conds map[string][]interface{}, preloads map[string][]interface{}) *gorm.DB

		Create(ctx common.Context, v interface{}) error

		CreateInBatches(ctx common.Context, v interface{}, batchSize int) error

		Update(ctx common.Context, conds map[string][]interface{}, values map[string]interface{}) error

		UpdateWithoutHooks(ctx common.Context, conds map[string][]interface{}, values map[string]interface{}) error

		UpdateWithPrimaryKey(ctx common.Context, id uint, values map[string]interface{}) error

		UpdateWithoutHooksWithPrimaryKey(ctx common.Context, id uint, values map[string]interface{}) error

		Find(ctx common.Context, dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error

		First(ctx common.Context, dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error
	}

	RepositoryOfEntity interface {
		Repository

		// find by id
		FindID(ctx common.Context, dest interface{}, ids []uint, preloads map[string][]interface{}) error

		FirstID(ctx common.Context, dest interface{}, id uint, preloads map[string][]interface{}) error

		FindUUID(ctx common.Context, dest interface{}, uuids []utils.UUID, preloads map[string][]interface{}) error

		FirstUUID(ctx common.Context, dest interface{}, uuid utils.UUID, preloads map[string][]interface{}) error
	}
)
