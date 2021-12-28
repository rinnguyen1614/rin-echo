package http

import (
	"rin-echo/system/app"
	"rin-echo/system/router/middleware"

	"github.com/labstack/echo/v4"
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
}
