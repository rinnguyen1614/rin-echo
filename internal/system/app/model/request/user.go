package request

type UserCommon struct {
	Username                  string `validate:"required,min=5,alphanum"`
	FullName                  string `json:"full_name" validate:"required"`
	Password                  string
	Email                     string `json:"email" validate:"required,email"`
	RoleIDs                   []uint `json:"role_ids"`
	RandomPassword            bool   `json:"random_password"`
	ChangePasswordOnNextLogin bool   `json:"change_password_on_next_login"`
	SendActivationEmail       bool   `json:"send_activation_email"`
	Active                    bool
	LockoutEnabled            bool `json:"lockout_enabled"`
	Gender                    uint
}

type CreateUser struct {
	UserCommon
	Password string `json:"password" validate:"required_if=RandomPassword false"`
}

func (u CreateUser) IsRandomPassword() bool {
	return u.RandomPassword || u.Password == ""
}

type UpdateUser struct {
	UserCommon
}
