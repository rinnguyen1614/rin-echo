package domain

import (
	"rin-echo/common/domain"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

const (
	RoleAdministrator = "administrator"
	RoleUser          = "user"
)

type Role struct {
	domain.FullAuditedEntity

	Name        string `gorm:"column:name;size:100;default:'';not null;"`
	Slug        string `gorm:"column:slug;size:100;uniqueIndex;default:'';not null;"`
	IsStatic    bool   `gorm:"column:is_static;"`
	IsDefault   bool   `gorm:"column:is_default;"`
	UserRoles   []*UserRole
	Permissions Permissions
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

func (r *Role) SetPermissions(permissions Permissions) {
	r.Permissions = permissions
}

func (r *Role) ComparePermissions(newPermissions Permissions) (permissionsNews, permissionsDels Permissions) {
	var (
		oldByMenuID = r.Permissions.ToMapByMenuID()
		newByMenuID = newPermissions.ToMapByMenuID()
	)

	if len(newPermissions) != 0 {
		for rID, ur := range newByMenuID {
			_, ok := oldByMenuID[rID]
			if ok {
				delete(oldByMenuID, rID)
			} else {
				permissionsNews = append(permissionsNews, ur)
			}
		}

		for _, ur := range oldByMenuID {
			permissionsDels = append(permissionsDels, ur)
		}
	} else {
		permissionsDels = r.Permissions
	}

	return
}

type RoleRepository interface {
	iuow.RepositoryOfEntity

	WithTransaction(db *gorm.DB) RoleRepository

	FirstBySlug(slug string, preloads map[string][]interface{}) (*Role, error)

	FindBySlug(slugs []string, preloads map[string][]interface{}) ([]*Role, error)
}
