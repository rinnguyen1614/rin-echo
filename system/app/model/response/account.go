package response

import (
	"rin-echo/common/model"
	"rin-echo/common/utils"
	"rin-echo/system/domain"
)

type Profile struct {
	model.Model

	UUID     utils.UUID
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Avatar   string `json:"avatar"`
}

func NewProfile(user domain.User) Profile {
	return Profile{
		Model:    model.NewModel(user.ID),
		UUID:     user.UUID,
		Email:    user.Email,
		FullName: user.FullName,
	}
}

type UserMenu struct {
	model.Model

	Name      string `json:"name" `
	Slug      string `json:"slug" `
	ParentID  uint   `json:"parent_id" `
	Path      string `json:"path" `
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Type      string `json:"type"`
	Meta      struct {
		Title string `json:"title"`
		Icon  string `json:"icon"`
	}
	Children UserMenus
}

func NewUserMenu(e domain.Menu) UserMenu {
	return UserMenu{
		Model:     model.NewModel(e.ID),
		Name:      e.Name,
		Slug:      e.Slug,
		Path:      e.Path,
		ParentID:  *e.ParentID,
		Component: e.Component,
		Sort:      e.Sort,
		Type:      e.Type,
		Meta: struct {
			Title string `json:"title"`
			Icon  string `json:"icon"`
		}(e.Meta),
	}
}

type UserMenus []UserMenu
