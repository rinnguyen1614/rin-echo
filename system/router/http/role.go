package http

import (
	"rin-echo/system/router/middleware"

	echox "github.com/rinnguyen1614/rin-echo-core/echo"

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
		router.GET("", echox.WrapHandler(h.app.RoleHandler.Query))
	}
}
