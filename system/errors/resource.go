package errors

import core "github.com/rinnguyen1614/rin-echo-core"

var (
	ErrResourceNotFound              = core.NewRinError("resource_not_found", "Resource not found")
	ErrResourceParentNotFound        = core.NewRinError("resource_parent_not_found", "Parent of resource not found")
	ErrResourceSlugExists            = core.NewRinError("resource_slug_exists", "Slug already exists in our system")
	ErrResourceObjectAndActionExists = core.NewRinError("resource_object_action_exists", "Object and action already exists in our system")
	ErrResourceReferencedRole        = core.NewRinError("resource_referenced_role", "This resource is being referenced by a role")
)
