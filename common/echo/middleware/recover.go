package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

type (
	RecoverConfig echomw.RecoverConfig
)

func Recover() echo.MiddlewareFunc {
	return echomw.Recover()
}

func RecoverWithConfig(config RecoverConfig) echo.MiddlewareFunc {
	return echomw.RecoverWithConfig(echomw.RecoverConfig(config))
}
