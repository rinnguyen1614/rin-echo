package request

type CreateUser struct {
	Username                        string `validate:"required,min=5,alphanum"`
	FullName                        string `json:"full_name" validate:"required"`
	Password                        string `json:"password" validate:"min=6"`
	Email                           string `json:"email" validate:"required,email"`
	RoleIDs                         []uint `json:"role_ids"`
	SetRandomPassword               bool
	ShouldChangePasswordOnNextLogin bool
	SendActivationEmail             bool
	IsActive                        bool
	IsLockoutEnabled                bool
}

type UpdateUser struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	RoleIDs  []uint `json:"role_ids"`
}
