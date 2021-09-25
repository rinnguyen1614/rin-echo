package inject

import (
	echox "rin-echo/common/echo"
	mwx "rin-echo/common/echo/middleware"
	"rin-echo/common/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func GetEcho() *echo.Echo {
	if service.echo == nil {
		e := echo.New()
		e.Validator = validation.NewValidator(GetI18n())
		e.Logger = echox.NewLogger(GetLogger(), "admin")
		e.Logger.SetLevel(log.ERROR)
		e.HTTPErrorHandler = echox.HTTPErrorHandlerWrapOnError(GetConfig().IsDevelopment())

		e.Use(mwx.Contextx())
		e.Use(mwx.Logger())
		e.Use(mwx.Recover())
		e.Use(middleware.RemoveTrailingSlash())
		e.Use(mwx.Localizer(GetI18n()))
		e.Use(mwx.RequestID())

		service.echo = e
	}
	return service.echo
}
