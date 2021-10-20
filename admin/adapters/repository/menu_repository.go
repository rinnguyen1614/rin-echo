package repository

import (
	"rin-echo/admin/domain"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type MenuRepository struct {
	iuow.RepositoryOfEntity
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.Menu{}),
	}
}
