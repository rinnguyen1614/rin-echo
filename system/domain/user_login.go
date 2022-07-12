package domain

import (
	"github.com/rinnguyen1614/rin-echo-core/domain"
)

var (
	ProviderPhone    = "phone"
	ProviderPassword = "username/password"
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
