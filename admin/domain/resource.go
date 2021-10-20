package domain

import "rin-echo/common/domain"

type Resource struct {
	domain.Entity

	Name        string      `gorm:"column:name;size:100;default:'';not null;"`
	Slug        string      `gorm:"column:slug;size:100;uniqueIndex;default:'';not null;"`
	Path        string      `gorm:"column:path;size:100;default:'';"`
	Method      string      `gorm:"column:method;size:100;default:'';"`
	Description string      `gorm:"column:description;"`
	ParentID    *uint       `gorm:"column:parent_id;index;"`
	Children    []*Resource `gorm:"-"`
}

func NewResource(name, slug, path, method, description string) (Resource, error) {
	return Resource{
		Name:        name,
		Slug:        slug,
		Path:        path,
		Method:      method,
		Description: description,
	}, nil
}

func (r *Resource) SetParent(parent *Resource) {
	if parent == nil {
		r.ParentID = nil
	} else {
		r.ParentID = &parent.ID
	}
}

type Resources []*Resource

func (rs Resources) ToMap() map[uint]*Resource {
	result := make(map[uint]*Resource)

	for _, r := range rs {
		result[r.ID] = r
	}

	return result
}
