package domain

import "rin-echo/common/domain"

var (
	PROVIDER_PHONE    = "phone"
	PROVIDER_PASSWORD = "username/password"
)

type UserLogin struct {
	domain.Entity
	UserID        uint
	LoginProvider string
	ProviderKey   string
	User          User
}

type LocationInfo struct {
	Location  string
	IPAddress string
	UserAgent string
	DevideID  string
}
