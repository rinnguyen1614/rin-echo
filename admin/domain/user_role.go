package domain

import "rin-echo/common/domain"

type UserRole struct {
	domain.CreationAuditedEntity

	UserID uint  `gorm:"user_id"`
	RoleID uint  `gorm:"role_id"`
	User   *User `gorm:"foreignKey:UserID;references:id"`
	Role   *Role `gorm:"foreignKey:RoleID;references:id"`
}

func NewUserRole(userID uint, roleID uint) (UserRole, error) {
	return UserRole{
		UserID: userID,
		RoleID: roleID,
	}, nil
}
