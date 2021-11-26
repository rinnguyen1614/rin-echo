package domain

import (
	"rin-echo/common/domain"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type Permission struct {
	domain.FullAuditedEntity

	MenuID    uint  `gorm:"column:menu_id;index;not null;"`
	RoleID    *uint `gorm:"column:role_id;index;"`
	UserID    *uint `gorm:"column:user_id;index;"`
	IsGranted bool  `gorm:"column:is_granted;"`

	Menu *Menu `gorm:"foreignKey:MenuID;references:id"`
	User *User `gorm:"foreignKey:UserID;references:id"`
	Role *Role `gorm:"foreignKey:RoleID;references:id"`
}

func NewPermission(menuID uint, roleID *uint, userID *uint) (*Permission, error) {
	return &Permission{
		MenuID:    menuID,
		RoleID:    roleID,
		UserID:    userID,
		IsGranted: true,
	}, nil
}

func NewPermissionForRole(roleID uint, menuID uint) (*Permission, error) {
	return NewPermission(menuID, &roleID, nil)
}

func NewPermissionsForRole(roleID uint, menuIDs []uint) (Permissions, error) {
	var pers Permissions
	for _, mID := range menuIDs {
		per, _ := NewPermissionForRole(roleID, mID)
		pers = append(pers, per)
	}

	return pers, nil
}

func NewPermissionForUser(userID uint, menuID uint) (*Permission, error) {
	return NewPermission(menuID, nil, &userID)
}

func (p *Permission) Grant(allow bool) {
	p.IsGranted = allow
}

type Permissions []*Permission

func (p Permissions) ToMap() map[uint]*Permission {
	dest := make(map[uint]*Permission)
	for _, a := range p {
		dest[a.ID] = a
	}
	return dest
}

func (p Permissions) ToMapByMenuID() map[uint]*Permission {
	dest := make(map[uint]*Permission)
	for _, a := range p {
		dest[a.MenuID] = a
	}
	return dest
}

func (p Permissions) IDs() []uint {
	var ids []uint
	for _, a := range p {
		ids = append(ids, a.ID)
	}
	return ids
}

func (p Permissions) MenuIDs() []uint {
	var ids []uint
	for _, a := range p {
		ids = append(ids, a.MenuID)
	}
	return ids
}

func (p Permissions) RoleIDs() []uint {
	var ids []uint
	for _, a := range p {
		ids = append(ids, *a.RoleID)
	}
	return ids
}

func (p Permissions) UserIDs() []uint {
	var ids []uint
	for _, a := range p {
		ids = append(ids, *a.UserID)
	}
	return ids
}

type PermissionRepository interface {
	iuow.RepositoryOfEntity

	WithTransaction(db *gorm.DB) PermissionRepository

	QueryByResources(resourceIDs []uint, preloads map[string][]interface{}) *gorm.DB
}
