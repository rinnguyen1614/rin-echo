package http

import (
	"rin-echo/system/router/middleware"

	echox "github.com/rinnguyen1614/rin-echo-core/echo"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterSettingRouter(g *echo.Group) {
	router := g.Group("/settings")
	{
		middleware.AddRequestLoggerMiddleware(router)
		middleware.AddJWTMiddleware(router)
		middleware.AddCasbinMiddleware(router)

		router.PUT("", echox.WrapHandler(h.app.SettingHandler.Set))
		router.GET("", echox.WrapHandler(h.app.SettingHandler.Get))
	}
}
