package repository

import (
	"rin-echo/common/uow"
	"rin-echo/system/domain"

	"gorm.io/gorm"
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
