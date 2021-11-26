package uow

import (
	"fmt"
	"rin-echo/common"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type (
	FuncNewInstance func(*gorm.DB) iuow.UnitOfWork
	UnitOfWork      struct {
		store           *gorm.DB
		funcNewInstance FuncNewInstance
	}
)

func NewUnitOfWork(store *gorm.DB) iuow.UnitOfWork {
	return &UnitOfWork{
		store:           store,
		funcNewInstance: NewUnitOfWork,
	}
}

func NewUnitOfWorkByFunc(store *gorm.DB, funcNewInstance FuncNewInstance) iuow.UnitOfWork {
	return funcNewInstance(store)
}

func (uow *UnitOfWork) DB() *gorm.DB {
	return uow.store
}

func (uow *UnitOfWork) Model(v interface{}) *gorm.DB {
	return uow.store.Model(v)
}

func (uow *UnitOfWork) Association(v interface{}, name string) *gorm.Association {
	return uow.store.Model(v).Association(name)
}

func (uow UnitOfWork) WithContext(ctx common.Context) iuow.UnitOfWork {
	return uow.funcNewInstance(uow.store.WithContext(ctx))
}

func (uow *UnitOfWork) Transaction(fc func(*gorm.DB) error) (err error) {
	return transaction(uow.store, fc)
}

func (uow *UnitOfWork) TransactionUnitOfWork(fc func(iuow.UnitOfWork) error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	err = uow.store.Transaction(func(tx *gorm.DB) error {
		return fc(uow.funcNewInstance(tx))
	})
	return
}

func transaction(db *gorm.DB, fc func(*gorm.DB) error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	err = db.Transaction(fc)
	return
}

func finishTransaction(err error, tx *gorm.DB) error {
	if err != nil {
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit().Error; commitErr != nil {
			return commitErr
		}

		return nil
	}
}
