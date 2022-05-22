package errors

import "rin-echo/common"

var (
	ErrRoleNotFound           = common.NewRinError("role_not_found", "Role not found")
	ErrRoleSlugExists         = common.NewRinError("role_slug_exists", "Slug has already in roles")
	ErrRoleReferencedResource = common.NewRinError("role_referenced_resource", "This role is being referenced by a resource")
	ErrRoleReferencedMenu     = common.NewRinError("role_referenced_menu", "This role is being referenced by a menu")
	ErrRoleReferencedUser     = common.NewRinError("role_referenced_user", "This role is being referenced by a user")
)
