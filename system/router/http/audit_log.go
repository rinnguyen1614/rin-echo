package http

import (
	echox "rin-echo/common/echo"
	"rin-echo/system/router/middleware"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterAuditLogRouter(g *echo.Group) {
	router := g.Group("/audit_logs")
	{
		middleware.AddJWTMiddleware(router)
		middleware.AddCasbinMiddleware(router)

		router.GET("/:id", echox.WrapHandler(h.app.AuditLogHandler.Get))
		router.GET("", echox.WrapHandler(h.app.AuditLogHandler.Query))
	}
}
