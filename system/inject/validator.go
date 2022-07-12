package inject

import (
	"regexp"

	"github.com/rinnguyen1614/rin-echo-core/utils"
	"github.com/rinnguyen1614/rin-echo-core/validation"

	"github.com/go-playground/validator/v10"
)

func GetValidator() *validation.Validator {
	if service.validator == nil {
		vali := validation.NewValidator(GetI18n())

		vali.Instance().RegisterValidation("username_validate", func(fl validator.FieldLevel) bool {
			return regexp.MustCompile("^[a-zA-Z0-9]{6,30}$").MatchString(fl.Field().String())
		})

		vali.Instance().RegisterValidation("password_validate", func(fl validator.FieldLevel) bool {
			return utils.ValidatePassword(fl.Field().String(), 8)
		})

		service.validator = vali
	}

	return service.validator
}
