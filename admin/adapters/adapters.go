package adapters

import (
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

func NewUnitOfWork(store *gorm.DB) iuow.UnitOfWork {
	uow := uow.NewUnitOfWork(store)
	registerRepositories(uow)
	return uow
}

func registerRepositories(uow iuow.UnitOfWork) {
	uow.SetRepository("UserRepository", NewUserRepository(uow.DB()))
}
