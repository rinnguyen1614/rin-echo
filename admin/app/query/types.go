package query

import (
	"rin-echo/common/domain"
	"rin-echo/common/echo/models"
)

type User struct {
	models.FullAuditedEntityModel

	UUID      domain.UUID `json:"uuid"`
	Username  string      `json:"username"`
	FullName  string      `json:"full_name"`
	Email     *string     `json:"email"`
	UserRoles []UserRole  `json:"user_roles"`
}

type Users []*User
type UserRole struct {
	models.CreationAuditedModel

	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
	Role   Role `json:"role"`
}

type UserRoles []*UserRole

type Role struct {
	models.FullAuditedEntityModel

	Slug string `json:"slug"`
}
