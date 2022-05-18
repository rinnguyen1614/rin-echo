package request

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
		Username:                        r.Username,
		FullName:                        r.FullName,
		Email:                           r.Email,
		Password:                        r.Password,
		SendActivationEmail:             true,
		SetRandomPassword:               false,
		ShouldChangePasswordOnNextLogin: false,
		IsActive:                        false,
	}
}

type ChangePassword struct {
	Username        string `json:"username" validate:"required"`
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"password_validate"`
}

type UpdateProfile struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UpdateAvatar struct {
}
