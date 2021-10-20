package errors

import "rin-echo/common"

var (
	ERR_ROLE_SLUG_EXISTS = common.NewRinError("role_slug_exists", "Slug already exists.")
)
