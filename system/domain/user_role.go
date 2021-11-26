package domain

import (
	"rin-echo/common/domain"

	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type UserRole struct {
	domain.CreationAuditedEntity

	UserID uint  `gorm:"user_id"`
	RoleID uint  `gorm:"role_id"`
	User   *User `gorm:"foreignKey:UserID;references:id"`
	Role   *Role `gorm:"foreignKey:RoleID;references:id"`
}

func NewUserRole(userID uint, roleID uint) (*UserRole, error) {
	return &UserRole{
		UserID: userID,
		RoleID: roleID,
	}, nil
}

type UserRoles []*UserRole

func (urs UserRoles) IDs() []uint {
	var ids []uint
	for _, a := range urs {
		ids = append(ids, a.ID)
	}
	return ids
}

func (urs UserRoles) RoleIDs() []uint {
	var ids []uint
	for _, a := range urs {
		ids = append(ids, a.RoleID)
	}
	return ids
}

func (urs UserRoles) UserIDs() []uint {
	var ids []uint
	for _, a := range urs {
		ids = append(ids, a.UserID)
	}
	return ids
}

func (p UserRoles) ToMap() map[uint]*UserRole {
	dest := make(map[uint]*UserRole)
	for _, a := range p {
		dest[a.ID] = a
	}
	return dest
}

func (p UserRoles) ToMapByUserID() map[uint]*UserRole {
	dest := make(map[uint]*UserRole)
	for _, a := range p {
		dest[a.UserID] = a
	}
	return dest
}

func (p UserRoles) ToMapByRoleID() map[uint]*UserRole {
	dest := make(map[uint]*UserRole)
	for _, a := range p {
		dest[a.RoleID] = a
	}
	return dest
}

type UserRoleRepository interface {
	iuow.RepositoryOfEntity
	FindByRoles(roleIDs []uint, preloads map[string][]interface{}) (UserRoles, error)
	FindByUsers(userIDs []uint, preloads map[string][]interface{}) (UserRoles, error)
	FindByUsersAndRoles(userIDs []uint, roleIDs []uint, preloads map[string][]interface{}) (UserRoles, error)

	WithTransaction(db *gorm.DB) UserRoleRepository
}
