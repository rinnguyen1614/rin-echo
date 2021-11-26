package http

import (
	echox "rin-echo/common/echo"
	"rin-echo/system/router/middleware"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterRoleRouter(g *echo.Group) {

	router := g.Group("/roles")
	{
		middleware.AddRequestLoggerMiddleware(router)
		middleware.AddJWTMiddleware(router)
		middleware.AddCasbinMiddleware(router)

		router.POST("", echox.WrapHandler(h.app.RoleHandler.Create))
		router.PUT("/:id", echox.WrapHandler(h.app.RoleHandler.Update))
		router.DELETE("/:id", echox.WrapHandler(h.app.RoleHandler.Delete))
		router.GET("/:id", echox.WrapHandler(h.app.RoleHandler.Get))
	}
}
