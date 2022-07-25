package request

import "github.com/rinnguyen1614/rin-echo/internal/system/domain"

type RoleCommon struct {
	Name        string `validate:"required,min=3"`
	Slug        string `validate:"required,min=3"`
	IsDefault   bool   `json:"is_default"`
	ResourceIDs []uint `json:"resource_ids"`
	MenuIDs     []uint `json:"menu_ids"`
}

type CreateRole struct {
	RoleCommon
}

type UpdateRole struct {
	RoleCommon
}

func (cmd UpdateRole) Populate(role *domain.Role) {
	role.Name = cmd.Name
	role.Slug = cmd.Slug
	role.IsDefault = cmd.IsDefault
}
