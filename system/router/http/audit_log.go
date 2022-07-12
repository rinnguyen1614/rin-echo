package http

import (
	"rin-echo/system/router/middleware"

	echox "github.com/rinnguyen1614/rin-echo-core/echo"

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
