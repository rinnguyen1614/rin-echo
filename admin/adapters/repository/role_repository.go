package repository

import (
	"rin-echo/admin/domain"
	"rin-echo/common/uow"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type RoleRepository struct {
	iuow.RepositoryOfEntity
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.Role{}),
	}
}

func (r RoleRepository) FirstBySlug(slug string, preloads map[string][]interface{}) (*domain.Role, error) {
	var des domain.Role

	if err := r.First(&des, map[string][]interface{}{"slug": {slug}}, preloads); err != nil {
		return nil, err
	}

	return &des, nil
}
