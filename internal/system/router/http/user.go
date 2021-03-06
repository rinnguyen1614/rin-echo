package http

import (
	"github.com/rinnguyen1614/rin-echo/internal/system/router/middleware"

	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterUserRouter(g *echo.Group) {

	router := g.Group("/users")
	{
		operationName := "UserHandler"
		middleware.AddRequestLoggerMiddleware(router)
		middleware.AddJWTMiddleware(router)
		middleware.AddCasbinMiddleware(router)

		router.POST("", echox.WrapHandler(h.app.UserHandler.Create))
		router.PUT("/:id", echox.WrapHandler(h.app.UserHandler.Update))
		router.DELETE("/:id", echox.WrapHandler(h.app.UserHandler.Delete))
		router.GET("/:id", echox.WrapHandler(h.app.UserHandler.Get))
		router.GET("", echox.WrapHandlerWithOperation(h.app.UserHandler.Query, operationName, operationName+".Query"))
	}
}
