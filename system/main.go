package main

import (
	"rin-echo/system/app"
	"rin-echo/system/inject"
	"rin-echo/system/router"
	"rin-echo/system/router/http"

	initdata "rin-echo/system/init_data"

	"github.com/labstack/echo/v4"
)

//go:generate swag init --parseDependency --parseInternal --parseDepth 1

// @title rin-echo API
// @version 1.0.0
// @description Application using Go Echo framework
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:1809
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
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
