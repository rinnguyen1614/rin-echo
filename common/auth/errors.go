package auth

import (
	"rin-echo/common"
	"rin-echo/common/log"
)

var (
	ERR_TOKEN_INVALID   = NewAuthenticationError("token_invalid", "Token is invalid")
	ERR_TOKEN_EXPIRED   = NewAuthenticationError("token_expired", "Token is expired")
	ERR_TOKEN_MALFORMED = NewAuthenticationError("token_malformed", "Token is malformed")
	ERR_TOKEN_MISSING   = NewAuthenticationError("token_missing", "Token is missing")
	ERR_TOKEN_ISSUER    = NewAuthenticationError("token_issuer_invalid", "Token's issuer isn't valid")
)

type AuthenticationError struct {
	level log.Level

	*common.RinError
}

var DefaultLogLevel = log.WarnLevel

func NewAuthenticationError(id, message string) *AuthenticationError {
	return &AuthenticationError{
		DefaultLogLevel,
		common.NewRinError(id, message),
	}
}

func NewAuthenticationErrorWithInner(inner error, id, message string) *AuthenticationError {
	return &AuthenticationError{
		DefaultLogLevel,
		common.NewRinErrorWithInner(inner, id, message),
	}
}

func (a *AuthenticationError) Level() log.Level {
	return a.level
}

func (a *AuthenticationError) SetLevel(l log.Level) {
	a.level = l
}
