package model

import "github.com/rinnguyen1614/rin-echo/internal/core/domain"

type Setting struct {
	domain.FullAuditedEntity

	Name  string
	Value string
	// if provider key is U, key is user id
	ProviderKey string
	// U: For User, G: Global, D: Default
	ProviderName string
}

func (s *Setting) TableName() string {
	return "settings"
}

type Settings []*Setting

func (s Settings) ToMapByName() map[string]*Setting {
	dest := make(map[string]*Setting)
	for _, item := range s {
		dest[item.Name] = item
	}
	return dest
}

func (s Settings) ToValuesByName() map[string]string {
	dest := make(map[string]string)
	for _, item := range s {
		dest[item.Name] = item.Value
	}
	return dest
}
