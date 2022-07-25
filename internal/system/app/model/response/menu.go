package response

import (
	"github.com/rinnguyen1614/rin-echo/internal/core/model"
	"github.com/rinnguyen1614/rin-echo/internal/core/utils"
	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
)

type Menu struct {
	model.FullAuditedEntityModel

	Name      string `json:"name" `
	Slug      string `json:"slug" `
	ParentID  uint   `json:"parent_id" `
	Path      string `json:"path" `
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
}

func NewMenu(e domain.Menu) Menu {
	return Menu{
		FullAuditedEntityModel: model.NewFullAuditedModelWithEntity(e.FullAuditedEntity),
		Name:                   e.Name,
		Slug:                   e.Slug,
		Path:                   e.Path,
		ParentID:               utils.DefaultValue(e.ParentID, uint(0)).(uint),
		Hidden:                 e.Hidden,
		Component:              e.Component,
		Sort:                   e.Sort,
		Type:                   e.Type,
		Icon:                   e.Icon,
		Title:                  e.Title,
	}
}

type Menus []*Menu

type MenuTree struct {
	Menu
	Children []MenuTree `json:"children"`
}

type MenuTrees []*MenuTree
