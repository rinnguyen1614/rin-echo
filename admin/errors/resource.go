package errors

import "rin-echo/common"

var (
	ERR_RESOURCE_SLUG_EXISTS = common.NewRinError("resource_slug_exists", "Slug already exists.")
)
