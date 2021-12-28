package main

import (
	"rin-echo/system/app"
	"rin-echo/system/inject"
	"rin-echo/system/router"
	"rin-echo/system/router/http"

	initdata "rin-echo/system/init_data"

	"github.com/labstack/echo/v4"
)

func main() {
	initdata.Init()
	router.RunHTTPServer(inject.GetEcho(), func(g *echo.Group) {
		http := http.NewHttpServer(
			app.NewApplication(
				inject.GetConfig(),
				inject.GetDB(),
				inject.GetCache(),
				inject.GetCasbin(),
				inject.GetLogger(),
				inject.GetRestQuery(),
				inject.GetAuther(),
				inject.GetValidator()))
		http.RegisterRouter(g)
	})
}
