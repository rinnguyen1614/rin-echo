package domain

import (
	"rin-echo/common/domain"
	iuow "rin-echo/common/uow/interfaces"
	"rin-echo/system/errors"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

var (
	MenuTypes       = []string{"M", "A"} // M: Menu, A: Action
	MenuTypeDefault = "M"
)

type Menu struct {
	domain.FullAuditedEntity

	Name      string `gorm:"column:name;size:255;default:'';not null;"`
	Slug      string `gorm:"column:slug;size:128;uniqueIndex;default:'';not null;"`
	ParentID  *uint  `gorm:"column:parent_id;index;"`
	Path      string `gorm:"column:path;size:255;index;"`
	Hidden    bool   `gorm:"column:hidden;"`
	Component string `gorm:"column:component;size:255;"`
	Sort      int    `gorm:"column:sort;index;size:4;default:0;"`
	MenuLevel uint   `gorm:"column:menu_level;size:4;default:0;"`
	Type      string `gorm:"column:type;size:2;index;"`
	Title     string `gorm:"column:title;size:255"`
	Icon      string `gorm:"column:icon;size:128;"`

	Resources   Resources `gorm:"many2many:menu_resources"`
	Permissions Permissions
	Children    Menus `gorm:"-"`
}

func NewMenu(name string, slug string, path string, hidden bool, component string, sort int, typ string, icon, title string, parent *Menu, resourceIDs []uint) (*Menu, error) {
	m := &Menu{
		Name:      name,
		Slug:      slug,
		Path:      path,
		Hidden:    hidden,
		Component: component,
		Sort:      sort,
		Icon:      icon,
		Title:     title,
	}

	if typ != "" {
		if !funk.Contains(MenuTypes, typ) {
			return nil, errors.ErrMenuTypeNotFound
		}
		m.SetType(typ)
	} else {
		m.SetTypeDefault()
	}

	m.SetParent(parent)
	m.AssignToResources(resourceIDs)

	return m, nil
}

func (m *Menu) SetTypeDefault() {
	m.Type = MenuTypeDefault
}

func (m *Menu) SetType(typ string) {
	m.Type = typ
}

func (m *Menu) AssignToResource(resourceID uint) {
	re := Resource{Entity: domain.Entity{ID: resourceID}}
	m.Resources = append(m.Resources, &re)
}

func (m *Menu) AssignToResources(resourceIDs []uint) {
	for _, rID := range resourceIDs {
		m.AssignToResource(rID)
	}
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

func (m Menu) CompareResources(newResources Resources) (resourceNews, resourceDels Resources) {
	var (
		oldByID = m.Resources.ToMap()
		newByID = newResources.ToMap()
	)

	if len(newResources) != 0 {
		for rID, ur := range newByID {
			_, ok := oldByID[rID]
			if ok {
				delete(oldByID, rID)
			} else {
				resourceNews = append(resourceNews, ur)
			}
		}

		for _, ur := range oldByID {
			resourceDels = append(resourceDels, ur)
		}
	} else {
		resourceDels = m.Resources
	}

	return
}

type Menus []*Menu

func (rs Menus) ToMap() map[uint]*Menu {
	result := make(map[uint]*Menu)

	for _, r := range rs {
		result[r.ID] = r
	}

	return result
}

func (ms Menus) ToTree() Menus {
	var (
		trees     = make(Menus, 0)
		treesByID = make(map[uint]*Menu)
	)

	for _, m := range ms {
		treesByID[m.ID] = m
		if parentID := m.ParentID; parentID != nil {
			if parent, ok := treesByID[*parentID]; ok {
				parent.Children = append(parent.Children, m)
			} else {
				trees = append(trees, m)
			}
		} else {
			trees = append(trees, m)
		}
	}

	return trees
}

type MenuRepository interface {
	iuow.RepositoryOfEntity

	WithTransaction(db *gorm.DB) MenuRepository

	QueryByUser(userID uint, preloads map[string][]interface{}) *gorm.DB
}
