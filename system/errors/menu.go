package errors

import "rin-echo/common"

var (
	ErrMenuTypeNotFound   = common.NewRinError("menu_type_not_found", "Menu type not found.")
	ErrMenuParentNotFound = common.NewRinError("menu_parent_not_found", "Parent of menu not found")
	ErrMenuSlugExists     = common.NewRinError("menu_slug_exists", "Slug already exists in our system")
	ErrMenuReferencedRole = common.NewRinError("menu_referenced_role", "This menu is being referenced by a role")
)
