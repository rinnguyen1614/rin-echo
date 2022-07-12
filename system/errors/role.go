package errors

import core "github.com/rinnguyen1614/rin-echo-core"

var (
	ErrRoleNotFound           = core.NewRinError("role_not_found", "Role not found")
	ErrRoleSlugExists         = core.NewRinError("role_slug_exists", "Slug has already in roles")
	ErrRoleReferencedResource = core.NewRinError("role_referenced_resource", "This role is being referenced by a resource")
	ErrRoleReferencedMenu     = core.NewRinError("role_referenced_menu", "This role is being referenced by a menu")
	ErrRoleReferencedUser     = core.NewRinError("role_referenced_user", "This role is being referenced by a user")
)
