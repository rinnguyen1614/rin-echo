package request

import "time"

type Login struct {
	Username string `validate:"required,min=5"`
	Password string `validate:"required,min=6"`
}

type Register struct {
	Username string `validate:"username_validate"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password_validate"`
}

func (r Register) ToCreateUser() CreateUser {
	return CreateUser{
		UserCommon: UserCommon{
			Username:                  r.Username,
			FullName:                  r.FullName,
			Email:                     r.Email,
			Password:                  r.Password,
			SendActivationEmail:       true,
			RandomPassword:            false,
			ChangePasswordOnNextLogin: false,
			Active:                    false,
		},
	}
}

type ChangePassword struct {
	Username        string `json:"username" validate:"required"`
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"password_validate"`
}

type UpdateProfile struct {
	FullName    string     `json:"full_name" validate:"required"`
	Email       string     `json:"email" validate:"required,email"`
	DateOfBirth *time.Time `json:"date_of_birth" validate:"datetime"`
	Phone       string     `json:"phone"`
	Username    string     `json:"username"`
	AvatarPath  string     `json:"avatar_path"`
	Gender      uint       `json:"gender"`
}

type ChangeEmail struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyEmail struct {
	Code string `json:"code" validate:"required,len=20"`
}

type ChangePhone struct {
	Phone string `json:"phone" validate:"required,min=10"`
}

type VerifyPhone struct {
	Code string `json:"code" validate:"required,len=5"`
}
