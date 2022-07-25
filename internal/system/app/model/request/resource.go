package request

import "github.com/rinnguyen1614/rin-echo/internal/system/domain"

type ResourceCommon struct {
	Name        string `json:"name" validate:"required,min=5"`
	Slug        string `json:"slug" validate:"required,min=6"`
	ParentID    uint   `json:"parent_id"`
	Object      string `json:"object"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

func (cmd ResourceCommon) IsEmptyObjectAndAction() bool {
	return cmd.Object == "" || cmd.Action == ""
}

func (cmd ResourceCommon) IsEqualObjectAndAction(object, action string) bool {
	return cmd.Object == object && cmd.Action == action
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
	resource.Action = cmd.Action
	resource.Object = cmd.Object
}
