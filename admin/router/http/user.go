package http

import (
	"rin-echo/admin/app/command"
	"rin-echo/admin/inject"
	"rin-echo/admin/router/middleware"
	echox "rin-echo/common/echo"

	"github.com/labstack/echo/v4"
)

func (h *HttpServer) RegisterUserRouter(g *echo.Group) {

	userRouter := g.Group("/users")
	{
		middleware.AddJWTMiddleware(userRouter)
		middleware.AddCasbinMiddleware(userRouter)

		userRouter.GET("", echox.WrapHandler(h.FindUser))
		userRouter.POST("", echox.WrapHandler(h.CreateUser))
	}
}

func (h HttpServer) FindUser(c echox.Context) error {
	query, err := inject.GetRestQuery().Query(c.Request())
	if err != nil {
		return err
	}

	result, err := h.app.Queries.FindUsers.Handle(c, query)
	if err != nil {
		return err
	}

	echox.OKWithData(c, result)
	return nil
}

func (h HttpServer) CreateUser(c echox.Context) error {
	var cmd command.CreateUser
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}
	err := h.app.Commands.CreateUser.Handle(c, cmd)

	if err != nil {
		return err
	}

	token, err := h.app.Queries.TokenUser.Handle(c, cmd.Username)
	if err != nil {
		return err
	}

	echox.OKWithData(c, token)
	return nil
}
