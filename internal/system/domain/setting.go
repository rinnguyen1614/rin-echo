package domain

import (
	setting "github.com/rinnguyen1614/rin-echo/internal/core/setting/model"
)

type Setting struct {
	setting.Setting

	Name         string `gorm:"column:name;size:128;default:'';not null;index:idx_settings_name_providerName_providerKey,unique"`
	Value        string `gorm:"column:value;type:text;"`
	ProviderKey  string `gorm:"column:provider_key;size:4;default:'';not null;index:idx_settings_name_providerName_providerKey,unique"`
	ProviderName string `gorm:"column:provider_name;size:128;default:'';not null;index:idx_settings_name_providerName_providerKey,unique"`
}
