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
	UserRoles []*UserRole
}

func NewRole(name string, slug string) (Role, error) {
	u := Role{
		Name: name,
		Slug: slug,
	}

	return u, nil
}
