package request

type RoleCommon struct {
	Name      string `validate:"required,min=5"`
	Slug      string `validate:"min=6"`
	IsStatic  bool   `json:"is_static"`
	IsDefault bool   `json:"is_default"`
	MenuIDs   []uint `json:"menu_ids"`
}

type CreateRole struct {
	RoleCommon
}

type UpdateRole struct {
	RoleCommon
}
