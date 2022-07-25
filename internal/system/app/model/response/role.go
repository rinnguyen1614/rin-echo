package response

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/model"

	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
)

type Role struct {
	model.FullAuditedEntityModel

	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	IsStatic    bool             `json:"is_static"`
	IsDefault   bool             `json:"is_default"`
	Permissions []RolePermission `json:"permissions"`
	Menus       []RoleMenu       `json:"menus"`
}

type Roles []*Role

type RolePermission struct {
	model.Model
	Resource struct {
		model.Model
		Name   string `json:"name"`
		Slug   string `json:"slug"`
		Object string `json:"object"`
		Action string `json:"action"`
	} `json:"resource"`
}

type RoleMenu struct {
	model.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
	Path string `json:"path"`
}

func NewRole(e domain.Role) Role {
	var (
		permissions []RolePermission
		menus       []RoleMenu
	)

	return Role{
		FullAuditedEntityModel: model.NewFullAuditedModelWithEntity(e.FullAuditedEntity),
		Name:                   e.Name,
		Slug:                   e.Slug,
		IsDefault:              e.IsDefault,
		IsStatic:               e.IsStatic,
		Permissions:            permissions,
		Menus:                  menus,
	}
}
