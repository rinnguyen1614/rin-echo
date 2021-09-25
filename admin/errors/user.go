package errors

import "rin-echo/common"

var (
	ERR_USERNAME_EXISTS             = common.NewRinError("username_exists", "Username already exists in our system.")
	ERR_USERNAME_PASSWORD_NOT_MATCH = common.NewRinError("username_password_not_match", "Your username and password didn't match.")
)
