package response

import (
	"rin-echo/common/model"
	"rin-echo/system/domain"
)

type User struct {
	model.FullAuditedEntityModel

	Username   string `json:"username"`
	FullName   string `json:"full_name" `
	Email      string `json:"email" `
	AvatarPath string `json:"avatar_path"`

	UserRoles []UserRole `json:"user_roles"`
}

type Users []*User

type UserRole struct {
	model.FullAuditedEntityModel
	Role struct {
		model.Model
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"role"`
}

func NewUser(e domain.User) User {
	return User{
		FullAuditedEntityModel: model.NewFullAuditedModelWithEntity(e.FullAuditedEntity),
		Username:               e.Username,
		FullName:               e.FullName,
		Email:                  e.Email,
		AvatarPath:             e.AvatarPath,
	}
}
