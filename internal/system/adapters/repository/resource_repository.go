package repository

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/uow"

	"gorm.io/gorm"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
)

type resourceRepository struct {
	*uow.RepositoryOfEntity
}

func NewResourceRepository(db *gorm.DB) domain.ResourceRepository {
	return &resourceRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.Resource{}),
	}
}

func (repo *resourceRepository) WithTransaction(db *gorm.DB) domain.ResourceRepository {
	return &resourceRepository{
		RepositoryOfEntity: repo.RepositoryOfEntity.WithTransaction(db),
	}
}
