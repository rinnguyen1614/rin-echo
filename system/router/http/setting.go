package http

import (
	echox "rin-echo/common/echo"
	"rin-echo/system/router/middleware"

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
