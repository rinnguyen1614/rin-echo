package casbin

import (
	"rin-echo/common"
	"rin-echo/common/log"
)

var (
	ERR_NOT_PERMISSION = NewAuthorizationError("not_permission", "You don't have permission for this resource.")
)

type AuthorizationError struct {
	level log.Level

	*common.RinError
}

var DefaultLogLevel = log.WarnLevel

func NewAuthorizationError(id, message string) *AuthorizationError {
	return &AuthorizationError{
		DefaultLogLevel,
		common.NewRinError(id, message),
	}
}

func NewAuthorizationErrorWithInner(inner error, id, message string) *AuthorizationError {
	return &AuthorizationError{
		DefaultLogLevel,
		common.NewRinErrorWithInner(inner, id, message),
	}
}

func (a *AuthorizationError) Level() log.Level {
	return a.level
}

func (a *AuthorizationError) SetLevel(l log.Level) {
	a.level = l
}
