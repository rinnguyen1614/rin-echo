package http

import (
	"rin-echo/admin/app/command"
	"rin-echo/admin/router/middleware"
	echox "rin-echo/common/echo"

	"github.com/labstack/echo/v4"
)

func (h *HttpServer) RegisterResourceRouter(g *echo.Group) {

	userRouter := g.Group("/resources")
	{
		middleware.AddJWTMiddleware(userRouter)
		middleware.AddCasbinMiddleware(userRouter)

		userRouter.POST("", echox.WrapHandler(h.CreateResource))
	}
}

func (h HttpServer) CreateResource(c echox.Context) error {
	var cmd command.CreateResource
	if err := c.Bind(&cmd); err != nil {
		return err
	}
	if err := c.Validate(cmd); err != nil {
		return err
	}
	err := h.app.Commands.CreateResource.Handle(c, &cmd)

	if err != nil {
		return err
	}

	echox.OKWithData(c, cmd.ID)
	return nil
}
