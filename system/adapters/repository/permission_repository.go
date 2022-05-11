package repository

import (
	"rin-echo/common/uow"
	"rin-echo/system/domain"

	"gorm.io/gorm"
)

type permissionRepository struct {
	*uow.RepositoryOfEntity
}

func NewPermissionRepository(db *gorm.DB) domain.PermissionRepository {
	return &permissionRepository{
		RepositoryOfEntity: uow.NewRepositoryOfEntity(db, &domain.Permission{}),
	}
}

func (repo *permissionRepository) WithTransaction(db *gorm.DB) domain.PermissionRepository {
	return &permissionRepository{
		RepositoryOfEntity: repo.RepositoryOfEntity.WithTransaction(db),
	}
}

func (r permissionRepository) QueryByResources(resourceIDs []uint, preloads map[string][]interface{}) *gorm.DB {
	return r.Query(nil, preloads).
		Joins("inner join menus on menus.id = permissions.menu_id").
		Joins("inner join menu_resources on menus.id = menu_resources.menu_id and menu_resources.resource_id IN ?", resourceIDs)

}

func (r permissionRepository) FindByRole(roleID uint, preloads map[string][]interface{}) (domain.Permissions, error) {
	var des domain.Permissions

	if err := r.Find(&des, map[string][]interface{}{"role_id": {roleID}}, preloads); err != nil {
		return nil, err
	}

	return des, nil
}

func (r permissionRepository) FindByMenu(menuID uint, preloads map[string][]interface{}) (domain.Permissions, error) {
	var des domain.Permissions

	if err := r.Find(&des, map[string][]interface{}{"menu_id": {menuID}}, preloads); err != nil {
		return nil, err
	}

	return des, nil
}

func (r permissionRepository) FindByMenus(menuIDs []uint, preloads map[string][]interface{}) (domain.Permissions, error) {
	var des domain.Permissions

	if err := r.Find(&des, map[string][]interface{}{"menu_id": {menuIDs}}, preloads); err != nil {
		return nil, err
	}

	return des, nil
}

func (r permissionRepository) FindByUser(userID uint) ([]map[string]interface{}, error) {
	var des = make([]map[string]interface{}, 0)

	if err := uow.Find(r.Query(nil, nil).
		Joins("inner join resources on resources.id = permissions.resource_id").
		Joins("inner join resources as parent_resources on parent_resources.id = resources.parent_id").
		Joins("inner join user_roles on user_roles.role_id = permissions.role_id and user_roles.user_id = ?", userID).
		Select("resources.slug, parent_resources.slug as parent_slug"), &des); err != nil {
		return nil, err
	}

	return des, nil
}
