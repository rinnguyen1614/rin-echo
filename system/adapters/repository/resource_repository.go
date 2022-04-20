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

func (r resourceRepository) QueryByMenus(menuIDs []uint, preloads map[string][]interface{}) *gorm.DB {
	return r.Query(nil, preloads).
		Joins("INNER JOIN menu_resources ON menu_resources.resource_id = resources.id AND menu_resources.menu_id IN (?)", menuIDs).
		Select("resources.*").
		Distinct()
}
