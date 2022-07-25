package domain

import (
	core "github.com/rinnguyen1614/rin-echo/internal/core"
	"github.com/rinnguyen1614/rin-echo/internal/core/log"
)

var DefaultLogLevel = log.WarnLevel

type EntityNotFoundError struct {
	EntityError
}

type EntityError struct {
	field string

	level log.Level
	*core.RinError
}

func NewEntityError(id, field, message string) *EntityError {
	return &EntityError{
		field,
		DefaultLogLevel,
		core.NewRinError(id, message),
	}
}

func NewEntityErrorWithInner(inner error, id, field, message string) *EntityError {
	return &EntityError{
		field,
		DefaultLogLevel,
		core.NewRinErrorWithInner(inner, id, message),
	}
}

func (a *EntityError) Level() log.Level {
	return a.level
}

func (a *EntityError) SetLevel(l log.Level) {
	a.level = l
}

func (a *EntityError) Field() string {
	return a.field
}
