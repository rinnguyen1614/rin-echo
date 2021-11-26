package adapters

import (
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

func NewUnitOfWork(store *gorm.DB) iuow.UnitOfWork {
	return uow.NewUnitOfWork(store)
}
