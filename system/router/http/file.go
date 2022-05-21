package http

import (
	echox "rin-echo/common/echo"
	"rin-echo/system/router/middleware"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterFileRouter(g *echo.Group) {

	router := g.Group("/files")
	{
		middleware.AddRequestLoggerMiddleware(router)
		middleware.AddJWTMiddleware(router)
		middleware.AddCasbinMiddleware(router)

		router.POST("/upload", echox.WrapHandler(h.app.FileHandler.Upload))
		router.POST("/download", echox.WrapHandler(h.app.FileHandler.Download))
	}
}
