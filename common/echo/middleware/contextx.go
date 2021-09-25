package middleware

import (
	echox "rin-echo/common/echo"

	"github.com/labstack/echo/v4"
)

func Contextx() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(echox.NewContextx(c))
		}
	}
}
