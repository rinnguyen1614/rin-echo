package repository

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/uow"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"gorm.io/gorm"
)

type (
	securityLogRepository struct {
		*uow.RepositoryOfEntity
	}
)

func NewSecurityLogRepository(db *gorm.DB) domain.SecurityLogRepository {
	return &securityLogRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.SecurityLog{}),
	}
}

func (repo *securityLogRepository) WithTransaction(db *gorm.DB) domain.SecurityLogRepository {
	return &securityLogRepository{
		RepositoryOfEntity: repo.RepositoryOfEntity.WithTransaction(db),
	}
}
