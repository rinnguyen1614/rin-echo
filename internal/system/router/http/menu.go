package http

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/router/middleware"

	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterMenuRouter(g *echo.Group) {

	router := g.Group("/menus")
	{
		middleware.AddRequestLoggerMiddleware(router)
		middleware.AddJWTMiddleware(router)
		middleware.AddCasbinMiddleware(router)

		router.POST("", echox.WrapHandler(h.app.MenuHandler.Create))
		router.PUT("/:id", echox.WrapHandler(h.app.MenuHandler.Update))
		router.DELETE("/:id", echox.WrapHandler(h.app.MenuHandler.Delete))
		router.GET("/:id", echox.WrapHandler(h.app.MenuHandler.Get))
		router.GET("/trees", echox.WrapHandler(h.app.MenuHandler.TreeQuery))
	}
}
