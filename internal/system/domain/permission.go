package domain

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/domain"
	iuow "github.com/rinnguyen1614/rin-echo/internal/core/uow/interfaces"

	"gorm.io/gorm"
)

type Permission struct {
	domain.FullAuditedEntity

	ResourceID uint  `gorm:"column:resource_id;index;not null;"`
	RoleID     *uint `gorm:"column:role_id;index;"`
	UserID     *uint `gorm:"column:user_id;index;"`
	IsGranted  bool  `gorm:"column:is_granted;"`

	Resource *Resource `gorm:"foreignKey:ResourceID;references:id"`
	User     *User     `gorm:"foreignKey:UserID;references:id"`
	Role     *Role     `gorm:"foreignKey:RoleID;references:id"`
}

func NewPermission(resourceID uint, roleID *uint, userID *uint) (*Permission, error) {
	return &Permission{
		ResourceID: resourceID,
		RoleID:     roleID,
		UserID:     userID,
		IsGranted:  true,
	}, nil
}

func NewPermissionForRole(roleID uint, resourceID uint) (*Permission, error) {
	return NewPermission(resourceID, &roleID, nil)
}

func NewPermissionsForRole(roleID uint, resourceIDs []uint) (Permissions, error) {
	var pers Permissions
	for _, mID := range resourceIDs {
		per, _ := NewPermissionForRole(roleID, mID)
		pers = append(pers, per)
	}

	return pers, nil
}

func NewPermissionForUser(userID uint, resourceID uint) (*Permission, error) {
	return NewPermission(resourceID, nil, &userID)
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

func (p Permissions) ToMapByResourceID() map[uint]*Permission {
	dest := make(map[uint]*Permission)
	for _, a := range p {
		dest[a.ResourceID] = a
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

func (p Permissions) ResourceIDs() []uint {
	var ids []uint
	for _, a := range p {
		ids = append(ids, a.ResourceID)
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

	FindByUser(userID uint) ([]map[string]interface{}, error)
}
