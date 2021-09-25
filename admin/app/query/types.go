package query

import "rin-echo/common/echo/models"

type User struct {
	models.FullAuditedEntityModel

	Username  string      `json:"username"`
	FullName  string      `json:"full_name"`
	Email     *string     `json:"email"`
	UserRoles []*UserRole `json:"user_roles"`
}

type UserRole struct {
	models.CreationAuditedModel

	UserID uint  `json:"user_id"`
	RoleID uint  `json:"role_id"`
	User   *User `json:"user"`
	Role   *Role `json:"role"`
}

type Role struct {
	models.FullAuditedEntityModel

	Slug string `json:"slug"`
}
