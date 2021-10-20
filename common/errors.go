package common

import (
	"errors"
)

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

type RinErrors struct {
	errors map[int][]error
	*RinError
}

// func (r *RinErrors) Error() string {
// 	var buf bytes.Buffer
// 	for _, err := range r.errors {
// 		buf.WriteString(err.Error())
// 	}

// 	return buf.String()
// }

func NewRinErrors(errors map[int][]error, id, message string) *RinErrors {
	return &RinErrors{
		errors:   errors,
		RinError: NewRinError(id, message),
	}
}
