package domain

import (
	"rin-echo/common/domain"
)

var (
	PROVIDER_PHONE    = "phone"
	PROVIDER_PASSWORD = "username/password"
)

type UserLogin struct {
	domain.Entity
	UserID   uint
	DeviceID string
	// external login provider. Ex: google, facebook, twitter...
	LoginProvider string
	// key provided by the login provider.
	ProviderKey string
	User        User
}
