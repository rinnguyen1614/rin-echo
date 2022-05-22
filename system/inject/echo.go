package inject

import (
	echox "rin-echo/common/echo"
	mwx "rin-echo/common/echo/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	_ "rin-echo/system/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func GetEcho() *echo.Echo {
	if service.echo == nil {
		e := echo.New()
		e.Validator = GetValidator()
		e.Logger = echox.NewLogger(GetLogger(), "system")
		e.Logger.SetLevel(log.DEBUG)
		e.HTTPErrorHandler = echox.HTTPErrorHandlerWrapOnError(GetConfig().IsDevelopment())
		// setup static folders.
		e.Static("/public", "./static/public")
		// add swagger
		e.GET("/swagger/*", echoSwagger.WrapHandler)

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
