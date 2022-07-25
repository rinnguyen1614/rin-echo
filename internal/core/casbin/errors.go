package casbin

import (
	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"github.com/rinnguyen1614/rin-echo/internal/core/log"
)

var (
	ERR_NOT_PERMISSION = NewAuthorizationError("not_permission", "You don't have permission for this resource.")
)

type AuthorizationError struct {
	level log.Level

	*core.RinError
}

var DefaultLogLevel = log.WarnLevel

func NewAuthorizationError(id, message string) *AuthorizationError {
	return &AuthorizationError{
		DefaultLogLevel,
		core.NewRinError(id, message),
	}
}

func NewAuthorizationErrorWithInner(inner error, id, message string) *AuthorizationError {
	return &AuthorizationError{
		DefaultLogLevel,
		core.NewRinErrorWithInner(inner, id, message),
	}
}

func (a *AuthorizationError) Level() log.Level {
	return a.level
}

func (a *AuthorizationError) SetLevel(l log.Level) {
	a.level = l
}
