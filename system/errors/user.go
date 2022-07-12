package errors

import core "github.com/rinnguyen1614/rin-echo-core"

var (
	ErrUserNotFound             = core.NewRinError("user_not_found", "User not found")
	ErrUserNameExists           = core.NewRinError("username_exists", "Username already exists in our system.")
	ErrEmailExists              = core.NewRinError("email_exists", "Email already exists in our system.")
	ErrUserNameNotExists        = core.NewRinError("username_not_exists", "Username doesn't exist in our system.")
	ErrUserNamePasswordNotMatch = core.NewRinError("username_password_not_match", "Your username and password didn't match.")
	ErrPasswordRequired         = core.NewRinError("password_required", "Password is a required field")
	ErrUserReferencedRole       = core.NewRinError("user_referenced_role", "This user is being referenced by a role")
	ErrGenderNotFound           = core.NewRinError("gender_not_found", "Gender not found.")
	ErrVericationEmail          = core.NewRinError("verification_email", "Verification email in error")
	ErrVericationPhone          = core.NewRinError("verification_phone", "Verification phone in error")
)
