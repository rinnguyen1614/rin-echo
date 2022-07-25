package http

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/router/middleware"

	"github.com/labstack/echo/v4"

	"github.com/rinnguyen1614/rin-echo/internal/system/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app: app}
}

func (h HttpServer) RegisterRouter(g *echo.Group) {
	middleware.AddDefaultMiddleware(g)
	h.RegisterResourceRouter(g)
	h.RegisterSettingRouter(g)
	h.RegisterMenuRouter(g)
	h.RegisterAccountRouter(g)
	h.RegisterRoleRouter(g)
	h.RegisterUserRouter(g)
	h.RegisterAuditLogRouter(g)

	h.RegisterFileRouter(g)
}
