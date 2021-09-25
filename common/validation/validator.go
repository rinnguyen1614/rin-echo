package validation

import (
	i18nx "rin-echo/common/validation/i18n"
	"text/template"

	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type (
	Config struct {
		FieldTag    string
		LanguageTag []language.Tag
		Bundle      *i18n.Bundle
	}

	Validator struct {
		validator *validator.Validate
	}
)

var (
	DefaultConfig = Config{
		FieldTag:    "json",
		LanguageTag: []language.Tag{language.English},
	}

	DefaultFuncMaps = template.FuncMap{
		"Field": func(fieldError FieldError) string {
			return fieldError.Field()
		},
		"Tag": func(fieldError FieldError) string {
			return fieldError.Tag()
		},
		"Param": func(fieldError FieldError) string {
			return fieldError.Param()
		},
	}
)

func NewValidator(bundle *i18n.Bundle) *Validator {
	c := DefaultConfig
	c.Bundle = bundle
	return NewValidatorWithConfig(c)
}

func NewValidatorWithConfig(config Config) *Validator {
	if config.Bundle == nil {
		panic("validator requies Bundle")
	}
	if config.FieldTag == "" {
		config.FieldTag = DefaultConfig.FieldTag
	}
	if len(config.LanguageTag) == 0 {
		config.LanguageTag = DefaultConfig.LanguageTag
	}

	validate := validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get(config.FieldTag), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	for _, tag := range config.LanguageTag {
		i18nx.RegisterDefaultTranslation(config.Bundle, tag)
	}

	return &Validator{
		validator: validate,
	}
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err == nil {
		return nil
	}

	if err, ok := err.(*validator.InvalidValidationError); ok {
		return NewValidationError(err.Type.String(), err.Error())
	}

	fieldErrors := make([]FieldError, 0)
	for _, err := range err.(validator.ValidationErrors) {
		fe := NewFieldError(err)
		fieldErrors = append(fieldErrors, fe)
	}

	return NewValidationErrorWithFields(err, fieldErrors)
}
