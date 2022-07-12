package repository

import (
	"rin-echo/system/domain"

	"github.com/rinnguyen1614/rin-echo-core/uow"

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

func (repo *menuRepository) QueryByUser(userID uint, conds map[string][]interface{}) *gorm.DB {
	return repo.Query(conds, nil).
		Joins("inner join menu_roles on menu_roles.menu_id = menus.id ").Joins("inner join user_roles on user_roles.role_id = menu_roles.role_id and user_roles.user_id = ?", userID)
}

func (repo *menuRepository) FindByUser(userID uint, conds map[string][]interface{}) (domain.Menus, error) {
	var dest domain.Menus
	if err := uow.Find(repo.QueryByUser(userID, conds), &dest); err != nil {
		return nil, err
	}
	return dest, nil
}
