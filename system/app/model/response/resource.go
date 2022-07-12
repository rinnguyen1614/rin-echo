package response

import (
	"rin-echo/system/domain"

	"github.com/rinnguyen1614/rin-echo-core/model"
	"github.com/rinnguyen1614/rin-echo-core/utils"
)

type Resource struct {
	model.FullAuditedEntityModel

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
		FullAuditedEntityModel: model.NewFullAuditedModelWithEntity(e.FullAuditedEntity),
		Name:                   e.Name,
		Slug:                   e.Slug,
		Object:                 e.Object,
		Action:                 e.Action,
		Description:            e.Description,
		ParentID:               utils.DefaultValue(e.ParentID, uint(0)).(uint),
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
