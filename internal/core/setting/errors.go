package setting

import "errors"

var (
	ErrProviderNotFound = errors.New("provider not found")
	ErrSettingNotFound  = errors.New("setting not found")
	ErrSettingExists    = errors.New("setting not found")
)
