package http

import (
	echox "rin-echo/common/echo"
	"rin-echo/system/router/middleware"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterResourceRouter(g *echo.Group) {

	router := g.Group("/resources")
	{
		middleware.AddRequestLoggerMiddleware(router)
		middleware.AddJWTMiddleware(router)
		middleware.AddCasbinMiddleware(router)

		router.POST("", echox.WrapHandler(h.app.ResourceHandler.Create))
		router.PUT("/:id", echox.WrapHandler(h.app.ResourceHandler.Update))
		router.DELETE("/:id", echox.WrapHandler(h.app.ResourceHandler.Delete))
		router.GET("/:id", echox.WrapHandler(h.app.ResourceHandler.Get))
		router.GET("/trees", echox.WrapHandler(h.app.ResourceHandler.TreeQuery))
	}
}
