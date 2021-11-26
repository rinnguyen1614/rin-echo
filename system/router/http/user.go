package http

import (
	echox "rin-echo/common/echo"
	"rin-echo/system/router/middleware"

	"github.com/labstack/echo/v4"
)

func (h HttpServer) RegisterUserRouter(g *echo.Group) {

	userRouter := g.Group("/users")
	{
		operationName := "UserHandler"
		middleware.AddRequestLoggerMiddleware(userRouter)
		middleware.AddJWTMiddleware(userRouter)
		middleware.AddCasbinMiddleware(userRouter)

		userRouter.POST("", echox.WrapHandler(h.app.UserHandler.Create))
		userRouter.PUT("/:id", echox.WrapHandler(h.app.UserHandler.Update))
		userRouter.GET("", echox.WrapHandlerWithOperation(h.app.UserHandler.Query, operationName, operationName+".Query"))
	}
}
