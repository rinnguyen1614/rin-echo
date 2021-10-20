package domain

import (
	"rin-echo/common/domain"
)

var (
	MenuTypes       = []string{"M", "A"}
	MenuTypeDefault = "M"
)

type Menu struct {
	domain.FullAuditedEntity

	Name      string `gorm:"column:name;size:255;default:'';not null;"`
	Slug      string `gorm:"column:slug;size:128;index;default:'';not null;"`
	ParentID  *uint  `gorm:"column:parent_id;index;"`
	Path      string `gorm:"column:path;size:255;index;"`
	Hidden    bool   `gorm:"column:hidden;"`
	Component string `gorm:"column:component;size:255;"`
	Sort      int    `gorm:"column:sort;index;size:4;default:0;"`
	MenuLevel uint   `gorm:"column:menu_level;size:4;default:0;"`
	Type      string `gorm:"column:type;size:2;index;"`
	Meta

	Resources Resources `gorm:"many2many:menu_resources"`
}

type Meta struct {
	Title string `gorm:"column:title;size:255"`
	Icon  string `gorm:"column:icon;size:128;"`
}

func NewMenu(name string, slug string, path string, hidden bool, component string, sort int) (Menu, error) {
	m := Menu{
		Name:      name,
		Slug:      slug,
		Path:      path,
		Hidden:    hidden,
		Component: component,
		Sort:      sort,
	}

	m.SetTypeDefault()

	return m, nil
}

func (m *Menu) SetTypeDefault() {
	m.Type = MenuTypeDefault
}

func (m *Menu) SetType(typ string) {
	m.Type = typ
}

func (m *Menu) SetResources(resources Resources) {
	m.Resources = resources
}

func (m *Menu) SetMeta(title, icon string) {
	m.Title = title
	m.Icon = icon
}

func (m *Menu) SetParent(parent *Menu) {
	if parent == nil {
		m.ParentID = nil
		m.MenuLevel = 0
	} else {
		m.ParentID = &parent.ID
		m.MenuLevel = parent.MenuLevel + 1
	}
}

type Menus []*Menu

func (rs Menus) ToMap() map[uint]*Menu {
	result := make(map[uint]*Menu)

	for _, r := range rs {
		result[r.ID] = r
	}

	return result
}
