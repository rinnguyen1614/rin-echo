package domain

import (
	"rin-echo/common/domain"
	iuow "rin-echo/common/uow/interfaces"

	"gorm.io/gorm"
)

type Resource struct {
	domain.Entity

	Name        string `gorm:"column:name;size:100;default:'';not null;"`
	Slug        string `gorm:"column:slug;size:100;uniqueIndex;default:'';not null;"`
	Path        string `gorm:"column:path;size:100;default:'';index:idx_resources_path_method,unique,where: path <> '' and method <> '';"`
	Method      string `gorm:"column:method;size:100;default:'';index:idx_resources_path_method,unique,where: path <> '' and method <> '';"`
	Description string `gorm:"column:description;"`
	ParentID    *uint  `gorm:"column:parent_id;index;"`

	Menus    Menus       `gorm:"many2many:menu_resources"`
	Children []*Resource `gorm:"-"`
}

func NewResource(name, slug, path, method, description string, parent *Resource) (*Resource, error) {
	r := Resource{
		Name:        name,
		Slug:        slug,
		Path:        path,
		Method:      method,
		Description: description,
	}

	r.SetParent(parent)
	return &r, nil
}

func (r *Resource) SetParent(parent *Resource) {
	if parent == nil {
		r.ParentID = nil
	} else {
		r.ParentID = &parent.ID
	}
}

func (r *Resource) IsEmptyPathOrMethod() bool {
	return r.Method == "" || r.Path == ""
}

type Resources []*Resource

func (rs Resources) ToMap() map[uint]*Resource {
	result := make(map[uint]*Resource)

	for _, r := range rs {
		result[r.ID] = r
	}

	return result
}

func (rs Resources) ToTree() Resources {
	var (
		trees     = make(Resources, 0)
		treesByID = make(map[uint]*Resource)
	)

	for _, r := range rs {
		treesByID[r.ID] = r
		if parentID := r.ParentID; parentID != nil {
			if parent, ok := treesByID[*parentID]; ok {
				parent.Children = append(parent.Children, r)
			} else {
				trees = append(trees, r)
			}
		} else {
			trees = append(trees, r)
		}
	}

	return trees
}

type ResourceRepository interface {
	iuow.RepositoryOfEntity

	WithTransaction(db *gorm.DB) ResourceRepository

	QueryByMenus(menuIDs []uint, preloads map[string][]interface{}) *gorm.DB
}
