package errors

import core "github.com/rinnguyen1614/rin-echo/internal/core"

var (
	ErrMenuTypeNotFound   = core.NewRinError("menu_type_not_found", "Menu type not found.")
	ErrMenuParentNotFound = core.NewRinError("menu_parent_not_found", "Parent of menu not found")
	ErrMenuSlugExists     = core.NewRinError("menu_slug_exists", "Slug already exists in our system")
	ErrMenuReferencedRole = core.NewRinError("menu_referenced_role", "This menu is being referenced by a role")
)
