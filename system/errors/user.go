package errors

import "rin-echo/common"

var (
	ErrUserNotFound             = common.NewRinError("user_not_found", "User not found")
	ErrUserNameExists           = common.NewRinError("username_exists", "Username already exists in our system.")
	ErrEmailExists              = common.NewRinError("email_exists", "Email already exists in our system.")
	ErrUserNameNotExists        = common.NewRinError("username_not_exists", "Username doesn't exist in our system.")
	ErrUserNamePasswordNotMatch = common.NewRinError("username_password_not_match", "Your username and password didn't match.")
	ErrPasswordRequired         = common.NewRinError("password_required", "Password is a required field")
	ErrUserReferencedRole       = common.NewRinError("user_referenced_role", "This user is being referenced by a role")
)
