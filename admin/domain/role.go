package domain

import "rin-echo/common/domain"

const (
	Role_Administrator = "administrator"
	Role_User          = "user"
)

type Role struct {
	domain.FullAuditedEntity

	Name      string `gorm:"column:name;size:100;default:'';not null;"`
	Slug      string `gorm:"column:slug;size:100;uniqueIndex;default:'';not null;"`
	IsStatic  bool   `gorm:"column:is_static;"`
	IsDefault bool   `gorm:"column:is_default;"`
	UserRoles []*UserRole
}

func NewRole(name string, slug string, isStatic, isDefault bool) (Role, error) {
	u := Role{
		Name:      name,
		Slug:      slug,
		IsStatic:  isStatic,
		IsDefault: isDefault,
	}

	return u, nil
}
