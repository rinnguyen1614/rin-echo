package validation

import (
	"fmt"
	"reflect"
	"rin-echo/common"
	"rin-echo/common/log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	ValidationError struct {
		fieldErrors []FieldError
		level       log.Level
		*common.RinError
	}
)

var DefaultLogLevel = log.WarnLevel

func NewValidationError(id, message string) *ValidationError {
	return &ValidationError{
		level:       DefaultLogLevel,
		fieldErrors: make([]FieldError, 0),
		RinError:    common.NewRinError(id, message),
	}
}

func NewValidationErrorWithFields(err error, fieldErrors []FieldError) *ValidationError {
	return &ValidationError{
		level:       DefaultLogLevel,
		fieldErrors: fieldErrors,
		RinError:    common.NewRinErrorWithInner(err, "validation_error", "You have some errors for validation."),
	}
}

func (a *ValidationError) Level() log.Level {
	return a.level
}

func (a *ValidationError) SetLevel(l log.Level) {
	a.level = l
}

type (
	FieldError interface {
		validator.FieldError
		common.Error
		TranslateWithI18n(*i18n.Localizer) string
	}

	fieldError struct {
		validator.FieldError
		*common.RinError
	}
)

func (v *ValidationError) FieldErrors() []FieldError {
	return v.fieldErrors
}

func NewFieldError(err validator.FieldError) FieldError {
	message := fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag())
	f := fieldError{
		FieldError: err,
		RinError:   common.NewRinErrorWithInner(err, err.Field()+"_invalid", message),
	}

	return &f
}

func (f *fieldError) ID() string {
	if len(f.Param()) == 0 {
		return f.Tag()
	}
	kind := f.Kind()
	if kind == reflect.Ptr {
		kind = f.Type().Elem().Kind()
	}

	var sub string
	switch kind {
	case reflect.String:
		sub = "string"
	case reflect.Slice, reflect.Map, reflect.Array:
		if f.Type() != reflect.TypeOf(time.Time{}) {
			sub = "datetime"
			goto END
		}
		sub = "items"
	default:
		sub = "number"
	}

END:
	return f.Tag() + "_" + sub
}

func (f *fieldError) Error() string {
	return f.Cause().Error()
}

func (f *fieldError) TranslateWithI18n(localizer *i18n.Localizer) string {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    f.ID(),
		Funcs:        DefaultFuncMaps,
		TemplateData: f,
	})
	if err != nil {
		return f.Message()
	}

	return msg
}
