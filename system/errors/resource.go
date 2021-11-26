package errors

import "rin-echo/common"

var (
	ErrResourceNotFound            = common.NewRinError("resource_not_found", "Resource not found")
	ErrResourceParentNotFound      = common.NewRinError("resource_parent_not_found", "Parent of resource not found")
	ErrResourceSlugExists          = common.NewRinError("resource_slug_exists", "Slug already exists in our system")
	ErrResourcePathAndMethodExists = common.NewRinError("resource_path_and_method_exists", "Path and method already exists in our system")
	ErrResourceReferencedMenu      = common.NewRinError("resource_referenced_menu", "Resource it is being referenced by a menu")
)
