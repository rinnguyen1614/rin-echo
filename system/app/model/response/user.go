package response

import "rin-echo/common/model"

type User struct {
	model.FullAuditedEntityModel

	Username   string `json:"username"`
	FullName   string `json:"full_name" `
	Email      string `json:"email" `
	AvatarPath string `json:"avatar_path"`

	UserRoles []UserRole `json:"user_roles"`
}

type UserRole struct {
	model.FullAuditedEntityModel
	Role struct {
		model.Model
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"role"`
}
