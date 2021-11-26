package repository

import (
	"rin-echo/common/uow"
	"rin-echo/system/domain"

	"gorm.io/gorm"
)

type (
	userLoginLogRepository struct {
		*uow.RepositoryOfEntity
	}
)

func NewUserLoginLogRepository(db *gorm.DB) domain.UserLoginLogRepository {
	return &userLoginLogRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.LoginLog{}),
	}
}

func (repo *userLoginLogRepository) WithTransaction(db *gorm.DB) domain.UserLoginLogRepository {
	return &userLoginLogRepository{
		RepositoryOfEntity: repo.RepositoryOfEntity.WithTransaction(db),
	}
}
