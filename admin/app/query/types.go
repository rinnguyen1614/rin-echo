package query

import (
	"rin-echo/admin/domain"
	"rin-echo/common/echo/models"
)

type User struct {
	models.FullAuditedEntityModel

	Username  string      `json:"username,omitempty"`
	FullName  string      `json:"full_name,omitempty"`
	Email     *string     `json:"email,omitempty"`
	UserRoles []*UserRole `json:"user_roles,omitempty"`
}

func newUser(e *domain.User) *User {
	if e == nil {
		return nil
	}

	return &User{
		FullAuditedEntityModel: models.NewFullAuditedModelWithEntity(&e.FullAuditedEntity),
		Username:               e.Username,
		FullName:               e.FullName,
		Email:                  e.Email,
		UserRoles:              newUserRoles(e.UserRoles),
	}
}

type Users []*User

func newUsers(es []*domain.User) []*User {
	var m []*User
	for i := 0; i < len(es); i++ {
		m = append(m, newUser(es[i]))
	}
	return m
}

type UserRole struct {
	models.CreationAuditedModel

	UserID uint  `json:"user_id,omitempty"`
	RoleID uint  `json:"role_id,omitempty"`
	Role   *Role `json:"role,omitempty"`
}

func newUserRole(e *domain.UserRole) *UserRole {
	if e == nil {
		return nil
	}

	return &UserRole{
		CreationAuditedModel: models.NewCreationAuditedModelWithEntity(&e.CreationAuditedEntity),
		UserID:               e.UserID,
		RoleID:               e.RoleID,
		Role:                 newRole(e.Role),
	}
}

type UserRoles []*UserRole

func newUserRoles(es []*domain.UserRole) []*UserRole {
	var m []*UserRole
	for i := 0; i < len(es); i++ {
		m = append(m, newUserRole(es[i]))
	}
	return m
}

type Role struct {
	models.FullAuditedEntityModel

	Slug string `json:"slug,omitempty"`
}

func newRole(e *domain.Role) *Role {
	if e == nil {
		return nil
	}

	return &Role{
		FullAuditedEntityModel: models.NewFullAuditedModelWithEntity(&e.FullAuditedEntity),
		Slug:                   e.Slug,
	}
}
