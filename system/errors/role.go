package errors

import "rin-echo/common"

var (
	ErrRoleNotFound   = common.NewRinError("role_not_found", "Role not found")
	ErrRoleSlugExists = common.NewRinError("role_slug_exists", "Slug has already in roles")
)
