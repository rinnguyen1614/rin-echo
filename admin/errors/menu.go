package errors

import "rin-echo/common"

var (
	ERR_MENU_NOT_FOUND   = common.NewRinError("menu_not_found", "Menu not found")
	ERR_MENU_SLUG_EXISTS = common.NewRinError("menu_slug_exists", "Slug already exists.")
)
