package response

import (
	"rin-echo/common/model"
	"rin-echo/common/utils"
	"rin-echo/system/domain"
)

type Resource struct {
	model.Model

	Name               string               `json:"name"`
	Slug               string               `json:"slug"`
	Object             string               `json:"object"`
	Action             string               `json:"action"`
	Description        string               `json:"description"`
	ParentID           uint                 `json:"parent_id"`
	ResourcePermission []ResourcePermission `json:"permissions"`
}

func NewResource(e domain.Resource) Resource {
	return Resource{
		Model:       model.NewModel(e.ID),
		Name:        e.Name,
		Slug:        e.Slug,
		Object:      e.Object,
		Action:      e.Action,
		Description: e.Description,
		ParentID:    utils.DefaultValue(e.ParentID, uint(0)).(uint),
	}
}

type Resources []*Resource

type ResourceTree struct {
	Resource
	Children []ResourceTree `json:"children"`
}

type ResourceTrees []*ResourceTree

type ResourcePermission struct {
	model.FullAuditedEntityModel
	Role struct {
		model.Model
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"role"`
}
