package domain

import (
	"rin-echo/common/domain"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type Role struct {
	domain.FullAuditedEntity

	Name      string `gorm:"column:name;size:100;default:'';not null;"`
	Slug      string `gorm:"column:slug;size:100;uniqueIndex;default:'';not null;"`
	IsStatic  bool   `gorm:"column:is_static;"`
	IsDefault bool   `gorm:"column:is_default;"`

	UserRoles   []*UserRole
	Permissions Permissions
	Menus       Menus `gorm:"many2many:menu_roles"`
}

func NewRole(name string, slug string, isStatic, isDefault bool) (*Role, error) {
	u := Role{
		Name:      name,
		Slug:      slug,
		IsStatic:  isStatic,
		IsDefault: isDefault,
	}

	return &u, nil
}

func NewNotRoleStatic(name string, slug string, isDefault bool) (*Role, error) {
	return NewRole(name, slug, false, true)
}

func (r *Role) SetPermissions(permissions Permissions) {
	r.Permissions = permissions
}

func (r *Role) ComparePermissions(newPermissions Permissions) (permissionsNews, permissionsDels Permissions) {
	var (
		oldByResourceID = r.Permissions.ToMapByResourceID()
		newByResourceID = newPermissions.ToMapByResourceID()
	)

	if len(newPermissions) != 0 {
		for rID, ur := range newByResourceID {
			_, ok := oldByResourceID[rID]
			if ok {
				delete(oldByResourceID, rID)
			} else {
				permissionsNews = append(permissionsNews, ur)
			}
		}

		for _, ur := range oldByResourceID {
			permissionsDels = append(permissionsDels, ur)
		}
	} else {
		permissionsDels = r.Permissions
	}

	return
}

func (r *Role) CompareMenus(newMenus Menus) (menuNews, menuDels Menus) {
	var (
		oldByID = r.Menus.ToMap()
		newByID = newMenus.ToMap()
	)

	if len(newMenus) != 0 {
		for rID, ur := range newByID {
			_, ok := oldByID[rID]
			if ok {
				delete(oldByID, rID)
			} else {
				menuNews = append(menuNews, ur)
			}
		}

		for _, ur := range oldByID {
			menuDels = append(menuDels, ur)
		}
	} else {
		menuDels = r.Menus
	}

	return
}

type RoleRepository interface {
	iuow.RepositoryOfEntity

	WithTransaction(db *gorm.DB) RoleRepository

	FirstBySlug(slug string, preloads map[string][]interface{}) (*Role, error)

	FindBySlug(slugs []string, preloads map[string][]interface{}) ([]*Role, error)
}
