package main

import (
	"rin-echo/admin/inject"
	"rin-echo/admin/router"
	"rin-echo/admin/router/http"
	"rin-echo/admin/service"

	"github.com/labstack/echo/v4"
)

func main() {
	router.RunHTTPServer(inject.GetEcho(), func(g *echo.Group) {
		http := http.NewHttpServer(service.NewApplication(inject.GetDB(), inject.GetCasbin()))
		http.RegisterRouter(g)
	})
}
