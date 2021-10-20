package interfaces

import (
	"rin-echo/common"

	"gorm.io/gorm"
)

type (
	UnitOfWork interface {
		DB() *gorm.DB

		WithContext(common.Context) UnitOfWork

		Transaction(func(*gorm.DB) error) error

		TransactionUnitOfWork(func(UnitOfWork) error) error

		Rollback(tx *gorm.DB) error

		RollbackUnitOfWork(ux UnitOfWork) error
	}

	Repository interface {
		Model() *gorm.DB

		Transaction(fc func(tx *gorm.DB) error) (err error)

		/*
			- conds : key is query string, value is the array of the query's arguments
				Ex1: conds= map[string][]interface{}{ "id" : 1 } // is equal `id = 1`
				Ex2: conds= map[string][]interface{}{ "id NOT IN (?) AND id <> ?" : { []uint{1,2,3}, 6} } // is equal `id NOT IN (1,2,3) AND id <> 6`
			- preloads : key is table name & value (options) is function (same type as the one used for Scopes) or query:
				If it is query string, [0]: condition for preload (option), [1:] arguments for condition.
				Ex: preloads = map[string][]interface{}{ "users" : {"username = ? AND role_id = ? ", "rin-echo", 1}} // is equal `username = "rin-echo" AND role_id = 1`
		*/
		Query(conds map[string][]interface{}, preloads map[string][]interface{}) *gorm.DB

		Create(v interface{}) error

		CreateInBatches(v interface{}, batchSize int) error

		Update(conds map[string][]interface{}, values map[string]interface{}) error

		UpdateWithoutHooks(conds map[string][]interface{}, values map[string]interface{}) error

		UpdateWithPrimaryKey(id uint, values map[string]interface{}) error

		UpdateWithoutHooksWithPrimaryKey(id uint, values map[string]interface{}) error

		Find(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error

		Get(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error

		First(dest interface{}, conds map[string][]interface{}, preloads map[string][]interface{}) error

		Count(conds map[string][]interface{}) int64

		Contains(conds map[string][]interface{}) bool
	}

	RepositoryOfEntity interface {
		Repository

		// find by id
		FindID(dest interface{}, ids []uint, preloads map[string][]interface{}) error

		GetID(dest interface{}, ids []uint, preloads map[string][]interface{}) error

		FirstID(dest interface{}, id uint, preloads map[string][]interface{}) error
	}
)
