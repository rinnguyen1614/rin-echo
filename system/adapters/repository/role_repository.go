package repository

import (
	"rin-echo/common/uow"
	"rin-echo/system/domain"

	"gorm.io/gorm"
)

type roleRepository struct {
	*uow.RepositoryOfEntity
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &roleRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.Role{}),
	}
}

// func (repo roleRepository) Update(roleUpdate *domain.Role) error {
// 	tx := repo.DB()
// 	tx.Save(roleUpdate)

// 	if tx.Error != nil {
// 		return tx.Error
// 	}

// 	return nil
// }

func (r roleRepository) FirstBySlug(slug string, preloads map[string][]interface{}) (*domain.Role, error) {
	var des domain.Role

	if err := r.First(&des, map[string][]interface{}{"slug": {slug}}, preloads); err != nil {
		return nil, err
	}

	return &des, nil
}

func (r roleRepository) FindBySlug(slugs []string, preloads map[string][]interface{}) ([]*domain.Role, error) {
	var des []*domain.Role

	if err := r.Find(&des, map[string][]interface{}{"slug": {slugs}}, preloads); err != nil {
		return nil, err
	}

	return des, nil
}

func (repo *roleRepository) WithTransaction(db *gorm.DB) domain.RoleRepository {
	return &roleRepository{
		RepositoryOfEntity: repo.RepositoryOfEntity.WithTransaction(db),
	}
}
