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
