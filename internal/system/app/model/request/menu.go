package request

import "github.com/rinnguyen1614/rin-echo/internal/system/domain"

type MenuCommon struct {
	Name      string `json:"name" validate:"required,min=5"`
	Slug      string `json:"slug" validate:"required,min=6"`
	ParentID  uint   `json:"parent_id" `
	Path      string `json:"path" validate:"required,min=6"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
}

type CreateMenu struct {
	MenuCommon

	Children CreateMenus
}

type CreateMenus []*CreateMenu

type UpdateMenu struct {
	MenuCommon
}

func (cmd UpdateMenu) Populate(menu *domain.Menu) {
	menu.Name = cmd.Name
	menu.Slug = cmd.Slug
	menu.Hidden = cmd.Hidden
	menu.Component = cmd.Component
	menu.Path = cmd.Path
	menu.Sort = cmd.Sort
	menu.Icon = cmd.Icon
	menu.Title = cmd.Title
	menu.SetType(cmd.Type)
}
