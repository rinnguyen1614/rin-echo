package response

import (
	"rin-echo/common/model"
	"rin-echo/system/domain"
)

type Resource struct {
	model.Model

	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
	ParentID    uint   `json:"parent_id"`
	Menus       []struct {
		model.Model

		Name string `json:"name" `
		Slug string `json:"slug" `
		Path string `json:"path" `
	} `json:"menus"`
}

func NewResource(e domain.Resource) Resource {
	return Resource{
		Model:       model.NewModel(e.ID),
		Name:        e.Name,
		Slug:        e.Slug,
		Path:        e.Path,
		Method:      e.Method,
		Description: e.Description,
		ParentID:    *e.ParentID,
	}
}

type Resources []*Resource

type ResourceTree struct {
	Resource
	Children []ResourceTree `json:"children"`
}

type ResourceTrees []*ResourceTree
