package repository

import (
	"rin-echo/admin/domain"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type UserRoleRepository struct {
	iuow.RepositoryOfEntity
}

func NewUserRoleRepository(db *gorm.DB) *UserRoleRepository {
	return &UserRoleRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.UserRole{}),
	}
}
