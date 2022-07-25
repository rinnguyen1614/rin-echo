package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

type (
	CORSConfig echomw.CORSConfig
)

func CORS() echo.MiddlewareFunc {
	return echomw.CORS()
}

func CORSWithConfig(config CORSConfig) echo.MiddlewareFunc {
	return echomw.CORSWithConfig(echomw.CORSConfig(config))
}
