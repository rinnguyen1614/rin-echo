package repository

import (
	"rin-echo/common/uow"
	"rin-echo/system/domain"

	"gorm.io/gorm"
)

type menuRepository struct {
	*uow.RepositoryOfEntity
}

func NewMenuRepository(db *gorm.DB) domain.MenuRepository {
	return &menuRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.Menu{}),
	}
}

func (repo *menuRepository) WithTransaction(db *gorm.DB) domain.MenuRepository {
	return &menuRepository{
		RepositoryOfEntity: repo.RepositoryOfEntity.WithTransaction(db),
	}
}

func (r menuRepository) QueryByUser(userID uint, preloads map[string][]interface{}) *gorm.DB {
	return r.Query(nil, preloads).
		Joins("inner join permissions on menus.id = permissions.menu_id and permissions.is_granted = 'true'").
		Joins("inner join user_roles on permissions.role_id = user_roles.role_id and user_roles.user_id = ?", userID)
}
