package errors

import "rin-echo/common"

var (
	ErrResourceNotFound              = common.NewRinError("resource_not_found", "Resource not found")
	ErrResourceParentNotFound        = common.NewRinError("resource_parent_not_found", "Parent of resource not found")
	ErrResourceSlugExists            = common.NewRinError("resource_slug_exists", "Slug already exists in our system")
	ErrResourceObjectAndActionExists = common.NewRinError("resource_object_action_exists", "Object and action already exists in our system")
	ErrResourceReferencedRole        = common.NewRinError("resource_referenced_role", "This resource is being referenced by a role")
)
