package router

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/rinnguyen1614/rin-echo/internal/system/inject"
)

func RunHTTPServer(e *echo.Echo, registerHandler func(g *echo.Group)) {
	c := inject.GetConfig().Server
	RunHTTPServerOnAddr(fmt.Sprintf("%s:%s", c.Host, c.Port), e, registerHandler)
}

func RunHTTPServerOnAddr(addr string, e *echo.Echo, registerHandler func(g *echo.Group)) {

	api := e.Group("api/v1")
	{
		registerHandler(api)
	}

	e.Logger.Fatal(e.Start(addr))
}
