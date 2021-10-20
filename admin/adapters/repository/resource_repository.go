package repository

import (
	"rin-echo/admin/domain"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type ResourceRepository struct {
	iuow.RepositoryOfEntity
}

func NewResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.Resource{}),
	}
}

func (r ResourceRepository) FirstBySlug(slug string, preloads map[string][]interface{}) (*domain.Resource, error) {
	var des domain.Resource

	if err := r.First(&des, map[string][]interface{}{"slug": {slug}}, preloads); err != nil {
		return nil, err
	}

	return &des, nil
}

func (r ResourceRepository) FindBySlug(slugs []string, preloads map[string][]interface{}) (*domain.Resource, error) {
	var des domain.Resource

	if err := r.Find(&des, map[string][]interface{}{"slug": {slugs}}, preloads); err != nil {
		return nil, err
	}

	return &des, nil
}
