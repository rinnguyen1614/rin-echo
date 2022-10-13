package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rinnguyen1614/rin-echo/internal/core/auth/jwt"
	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	"github.com/rinnguyen1614/rin-echo/internal/core/validation"
	"golang.org/x/text/language"
)

var autherTest = jwt.NewJWT("secret")

func newEchoTest() *echo.Echo {
	e := echo.New()
	e.Validator = validation.NewValidator(i18n.NewBundle(language.English))
	e.HTTPErrorHandler = echox.HTTPErrorHandlerWrapOnError(false)
	return e
}
