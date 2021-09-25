package domain

import (
	"rin-echo/common"
	"rin-echo/common/log"
)

type EntityNotFoundError struct {
	EntityError
}

type EntityError struct {
	level log.Level

	*common.RinError
}

var DefaultLogLevel = log.WarnLevel

func NewEntityError(id, message string) *EntityError {
	return &EntityError{
		DefaultLogLevel,
		common.NewRinError(id, message),
	}
}

func NewEntityErrorWithInner(inner error, id, message string) *EntityError {
	return &EntityError{
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
