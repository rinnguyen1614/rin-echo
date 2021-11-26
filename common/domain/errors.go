package domain

import (
	"rin-echo/common"
	"rin-echo/common/log"
)

var DefaultLogLevel = log.WarnLevel

type EntityNotFoundError struct {
	EntityError
}

type EntityError struct {
	field string

	level log.Level
	*common.RinError
}

func NewEntityError(id, field, message string) *EntityError {
	return &EntityError{
		field,
		DefaultLogLevel,
		common.NewRinError(id, message),
	}
}

func NewEntityErrorWithInner(inner error, id, field, message string) *EntityError {
	return &EntityError{
		field,
		DefaultLogLevel,
		common.NewRinErrorWithInner(inner, id, message),
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
