package domain

import "rin-echo/common/utils"

var (
	UserSettingProviderName    = "U"
	GlobalSettingProviderName  = "G"
	DefaultSettingProviderName = "D"
)

type Setting struct {
	Name  string
	Value string
	// if provider key is U, key is user id
	ProviderKey string
	// U: For User, G: Global, D: Default
	ProviderName string
}

func newSetting(name, value, providerName, providerKey string) *Setting {
	return &Setting{
		Name:         name,
		Value:        value,
		ProviderKey:  providerKey,
		ProviderName: providerName,
	}
}

func NewSettingForUser(name string, value string, userID uint) *Setting {
	return newSetting(name, value, UserSettingProviderName, utils.ToString(userID))
}

func NewSettingGlobal(name string, value string) *Setting {
	return newSetting(name, value, GlobalSettingProviderName, "")
}

func NewSettingDefault(name string, value string) *Setting {
	return newSetting(name, value, DefaultSettingProviderName, "")
}

type Settings []*Setting

func (s Settings) ToMap() map[string]*Setting {
	dest := make(map[string]*Setting)
	for _, a := range s {
		dest[a.Name] = a
	}
	return dest
}
