package uow

import (
	"fmt"
	"rin-echo/common"
	iuow "rin-echo/common/uow/interfaces"
	"sync"

	"gorm.io/gorm"
)

type (
	UnitOfWork struct {
		store        *gorm.DB
		repositories map[string]iuow.Repository
		lock         sync.RWMutex
	}
)

func NewUnitOfWork(store *gorm.DB) iuow.UnitOfWork {
	return &UnitOfWork{
		store: store,
	}
}

func (uow *UnitOfWork) DB() *gorm.DB {
	return uow.store
}

func (uow *UnitOfWork) WithContext(ctx common.Context) *gorm.DB {
	return uow.DB().WithContext(ctx)
}

func (uow *UnitOfWork) Transaction(fc func(*gorm.DB) error) (err error) {
	return transaction(uow.store, fc)
}

func (uow *UnitOfWork) TransactionUnitOfWork(fc func(iuow.UnitOfWork) error) (err error) {
	tx := uow.store.Begin()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		err = finishTransaction(err, tx)
	}()
	err = fc(NewUnitOfWork(tx))
	return
}

func (uow *UnitOfWork) Rollback(tx *gorm.DB) (err error) {
	return tx.Rollback().Error
}

func (uow *UnitOfWork) RollbackUnitOfWork(ux iuow.UnitOfWork) (err error) {
	return ux.DB().Rollback().Error
}

func (uow *UnitOfWork) GetRepository(key string) iuow.Repository {
	uow.lock.RLock()
	defer uow.lock.RUnlock()
	return uow.repositories[key]
}

func (uow *UnitOfWork) SetRepository(key string, val iuow.Repository) {
	uow.lock.Lock()
	defer uow.lock.Unlock()

	if uow.repositories == nil {
		uow.repositories = make(map[string]iuow.Repository)
	}
	uow.repositories[key] = val
}

func transaction(db *gorm.DB, fc func(*gorm.DB) error) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		err = finishTransaction(err, tx)
	}()
	err = fc(tx)
	return
}

func finishTransaction(err error, tx *gorm.DB) error {
	if err != nil {
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(); commitErr != nil {
			return err
		}

		return nil
	}
}
