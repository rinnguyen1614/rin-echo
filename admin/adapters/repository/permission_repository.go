package repository

import (
	"rin-echo/admin/domain"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	iuow.RepositoryOfEntity
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.Permission{}),
	}
}
