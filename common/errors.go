package common

import "errors"

type (
	Error interface {
		Error() string

		Cause() error

		ID() string

		Message() string
	}

	RinError struct {
		id      string
		message string
		cause   error
	}
)

func NewRinError(id, message string) *RinError {
	return NewRinErrorWithInner(errors.New(id+": "+message), id, message)
}

func NewRinErrorWithInner(err error, id, message string) *RinError {
	return &RinError{
		id,
		message,
		err,
	}
}

func (r *RinError) Error() string {
	return r.cause.Error()
}

func (r *RinError) Cause() error {
	return r.cause
}

func (r *RinError) ID() string {
	return r.id
}

func (r *RinError) Message() string {
	return r.message
}
