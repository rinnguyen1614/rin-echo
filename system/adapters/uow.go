package adapters

import (
	"github.com/rinnguyen1614/rin-echo-core/uow"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"

	"gorm.io/gorm"
)

func NewUnitOfWork(store *gorm.DB) iuow.UnitOfWork {
	return uow.NewUnitOfWork(store)
}
