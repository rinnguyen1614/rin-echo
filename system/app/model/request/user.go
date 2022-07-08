package request

type CreateUser struct {
	Username                  string `validate:"required,min=5,alphanum"`
	FullName                  string `json:"full_name" validate:"required"`
	Password                  string `json:"password" validate:"min=6"`
	Email                     string `json:"email" validate:"required,email"`
	RoleIDs                   []uint `json:"role_ids"`
	RandomPassword            bool
	ChangePasswordOnNextLogin bool
	SendActivationEmail       bool
	Active                    bool
	LockoutEnabled            bool
}

func (u CreateUser) IsRandomPassword() bool {
	return u.RandomPassword || u.Password == ""
}

type UpdateUser struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	RoleIDs  []uint `json:"role_ids"`
}
