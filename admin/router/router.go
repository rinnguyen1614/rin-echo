package router

import (
	"fmt"
	"rin-echo/admin/inject"

	"github.com/labstack/echo/v4"
)

func RunHTTPServer(e *echo.Echo, registerHandler func(g *echo.Group)) {
	c := inject.GetConfig().Server
	RunHTTPServerOnAddr(fmt.Sprintf("%s:%s", c.Host, c.Port), e, registerHandler)
}

func RunHTTPServerOnAddr(addr string, e *echo.Echo, registerHandler func(g *echo.Group)) {

	api := e.Group("api")
	{
		registerHandler(api)
	}

	e.Logger.Fatal(e.Start(addr))
}
