package http

import (
	"rin-echo/admin/app/command"
	"rin-echo/admin/router/middleware"
	echox "rin-echo/common/echo"

	"github.com/labstack/echo/v4"
)

func (h *HttpServer) RegisterMenuRouter(g *echo.Group) {

	userRouter := g.Group("/menus")
	{
		middleware.AddJWTMiddleware(userRouter)
		middleware.AddCasbinMiddleware(userRouter)

		userRouter.POST("", echox.WrapHandler(h.CreateMenu))
	}
}

func (h HttpServer) CreateMenu(c echox.Context) error {
	var cmd command.CreateMenu
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}
	err := h.app.Commands.CreateMenu.Handle(c, &cmd)

	if err != nil {
		return err
	}

	echox.OKWithData(c, cmd.ID)
	return nil
}
