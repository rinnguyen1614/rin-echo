package request

import "rin-echo/system/domain"

type ResourceCommon struct {
	Name        string `json:"name" validate:"required,min=5"`
	Slug        string `json:"slug" validate:"required,min=6"`
	ParentID    uint   `json:"parent_id"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

func (cmd ResourceCommon) IsEmptyPathAndMethod() bool {
	return cmd.Path == "" || cmd.Method == ""
}

type CreateResource struct {
	ResourceCommon

	Children CreateResources
}

type CreateResources []*CreateResource

type UpdateResource struct {
	ResourceCommon
}

func (cmd UpdateResource) Populate(resource *domain.Resource) {
	resource.Name = cmd.Name
	resource.Slug = cmd.Slug
	resource.Description = cmd.Description
	resource.Method = cmd.Method
	resource.Path = cmd.Path
}
