package domain

import (
	"github.com/rinnguyen1614/rin-echo-core/domain"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"

	"gorm.io/gorm"
)

type Resource struct {
	domain.FullAuditedEntity

	Name        string `gorm:"column:name;size:100;default:'';not null;"`
	Slug        string `gorm:"column:slug;size:100;uniqueIndex;default:'';not null;"`
	Object      string `gorm:"column:object;size:100;default:'';index:idx_resources_object_action,unique,where: object <> '' and action <> '';"`
	Action      string `gorm:"column:action;size:100;default:'';index:idx_resources_object_action,unique,where: object <> '' and action <> '';"`
	Description string `gorm:"column:description;"`
	ParentID    *uint  `gorm:"column:parent_id;index;"`

	Permissions Permissions
	Children    []*Resource `gorm:"-"`
}

func NewResource(name, slug, object, action, description string, parent *Resource) (*Resource, error) {
	r := Resource{
		Name:        name,
		Slug:        slug,
		Object:      object,
		Action:      action,
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

func (r *Resource) IsEmptyObjectOrAction() bool {
	return r.Action == "" || r.Object == ""
}

type Resources []*Resource

func (rs Resources) IDs() []uint {
	var ids []uint
	for _, a := range rs {
		ids = append(ids, a.ID)
	}
	return ids
}

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
}
